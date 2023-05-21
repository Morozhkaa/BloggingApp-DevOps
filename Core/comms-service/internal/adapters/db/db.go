package db

import (
	"comm-service/internal/domain/models"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

type DBStorage struct {
	pool *pgxpool.Pool
}

func ParseTime(comm_time *models.Time, t time.Time) {
	comm_time.Date = fmt.Sprintf("%04d/%02d/%02d", t.Year(), int(t.Month()), t.Day())
	comm_time.Time = fmt.Sprintf("%02d:%02d:%02d", (t.Hour()+3)%24, t.Minute(), t.Second())
}

func New(ctx context.Context, conn string) (*DBStorage, error) {
	time.Sleep(4 * time.Second)
	pool, err := pgxpool.Connect(ctx, conn)
	if err != nil {
		return nil, err
	}
	return &DBStorage{
		pool: pool,
	}, nil
}

func (db *DBStorage) Get(ctx context.Context, comm_id uuid.UUID) (comment models.Comment, err error) {
	const queryComm = `
	SELECT text, created_at, updated_at, comment_id, post_id, commenter_email FROM comments WHERE comment_id = $1
	`
	var created_at, updated_at time.Time

	err = db.pool.QueryRow(ctx, queryComm, comm_id).Scan(
		&comment.Text,
		&created_at,
		&updated_at,
		&comment.ID,
		&comment.Post_id,
		&comment.Commenter.Email,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Comment{}, nil
		}
		return models.Comment{}, models.ErrNotFound
	}
	ParseTime(&comment.Created_at, created_at)
	ParseTime(&comment.Updated_at, updated_at)

	const query = `
	SELECT email, login FROM commenters WHERE email = $1
	`
	err = db.pool.QueryRow(ctx, query, comment.Commenter.Email).Scan(
		&comment.Commenter.Email,
		&comment.Commenter.Login,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Comment{}, nil
		}
		return models.Comment{}, models.ErrNotFound
	}

	return comment, nil
}

func (db *DBStorage) Delete(ctx context.Context, comm_id uuid.UUID, email strfmt.Email, role models.UserRole) error {
	// Check that we have the right to delete
	comment, err := db.Get(ctx, comm_id)
	if err != nil {
		return err
	}
	if comment.Commenter.Email != email && role != models.RoleAdmin && role != models.RoleModerator {
		return models.ErrForbidden
	}

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	const queryComm = `
	DELETE FROM comments WHERE comment_id = $1;
	`
	_, err = tx.Exec(ctx, queryComm, comm_id)
	if err != nil {
		return err
	}

	// Check if the commenter has comments, if not, delete him from the commenters table
	const queryCheck = `
		SELECT COUNT(commenter_email) FROM comments WHERE commenter_email = $1;
		`
	var countRows int
	err = db.pool.QueryRow(ctx, queryCheck, comment.Commenter.Email).Scan(
		&countRows,
	)
	if err != nil {
		return err
	}

	if countRows == 1 {
		const queryCommenter = `
		DELETE FROM commenters WHERE email = $1;
		`
		_, err = tx.Exec(ctx, queryCommenter, email)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
}

func (db *DBStorage) Save(ctx context.Context, comment models.Comment) (created_at time.Time, err error) {
	const queryCommenter = `
		SELECT EXISTS (SELECT email FROM commenters WHERE email = $1);
		`
	var isNotFirstComm bool
	err = db.pool.QueryRow(ctx, queryCommenter, comment.Commenter.Email).Scan(
		&isNotFirstComm,
	)
	if err != nil {
		return time.Time{}, err
	}

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return created_at, err
	}
	defer tx.Rollback(ctx)

	if !isNotFirstComm {
		const query = `
		INSERT INTO commenters (email, login)
		VALUES ($1, $2)
		`
		_, err = tx.Exec(ctx, query, comment.Commenter.Email, comment.Commenter.Login)
		if err != nil {
			return created_at, err
		}
	}

	const commQuery = `
	INSERT INTO comments (text, created_at, updated_at, comment_id, post_id, commenter_email)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	created_at = time.Now()
	_, err = tx.Exec(ctx, commQuery, comment.Text, created_at, created_at, comment.ID, comment.Post_id, comment.Commenter.Email)
	if err != nil {
		return created_at, err
	}

	err = tx.Commit(ctx)
	return created_at, err
}

func (db *DBStorage) Update(ctx context.Context, comm_id uuid.UUID, data models.NewCommDescription) (updated_at time.Time, err error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return updated_at, err
	}
	defer tx.Rollback(ctx)

	const query = `
	UPDATE comments SET text = $2, updated_at = $3
	WHERE comment_id = $1
	`
	updated_at = time.Now()
	_, err = tx.Exec(ctx, query, comm_id, data.Text, updated_at)
	if err != nil {
		return updated_at, err
	}
	err = tx.Commit(ctx)
	return updated_at, err
}

func (db *DBStorage) GetAll(ctx context.Context, post_id uuid.UUID) (comm_list []models.Comment, err error) {

	const queryComm = `
	SELECT text, created_at, updated_at, comment_id, post_id, commenter_email FROM comments
	WHERE post_id = $1
	`
	rows, err := db.pool.Query(ctx, queryComm, post_id)
	if err != nil {
		return nil, fmt.Errorf("can't get all comments: %v", err)
	}

	for rows.Next() {
		comment := models.Comment{}

		var created_at time.Time
		var updated_at time.Time

		err = rows.Scan(
			&comment.Text,
			&created_at,
			&updated_at,
			&comment.ID,
			&comment.Post_id,
			&comment.Commenter.Email,
		)
		if err != nil {
			return comm_list, err
		}
		ParseTime(&comment.Created_at, created_at)
		ParseTime(&comment.Updated_at, updated_at)

		const queryCommenter = `
		SELECT email, login FROM commenters WHERE email = $1
		`
		commenter := models.Commenter{}
		err := db.pool.QueryRow(ctx, queryCommenter, comment.Commenter.Email).Scan(
			&commenter.Email,
			&commenter.Login,
		)
		if err != nil {
			return comm_list, err
		}
		comment.Commenter = commenter
		comm_list = append(comm_list, comment)
	}

	return comm_list, err
}
