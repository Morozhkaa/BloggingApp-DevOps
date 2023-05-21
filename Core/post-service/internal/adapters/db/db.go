package db

import (
	"context"
	"errors"
	"fmt"
	"post-service/internal/domain/models"
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

func ParseTime(post_time *models.Time, t time.Time) {
	post_time.Date = fmt.Sprintf("%04d/%02d/%02d", t.Year(), int(t.Month()), t.Day())
	post_time.Time = fmt.Sprintf("%02d:%02d:%02d", (t.Hour()+3)%24, t.Minute(), t.Second())
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

func (db *DBStorage) Get(ctx context.Context, id uuid.UUID) (post models.Post, err error) {
	const queryPost = `
	SELECT title, description, content, created_at, updated_at, post_id, auth_email FROM posts WHERE post_id = $1
	`
	var created_at, updated_at time.Time

	err = db.pool.QueryRow(ctx, queryPost, id).Scan(
		&post.Title,
		&post.Description,
		&post.Content,
		&created_at,
		&updated_at,
		&post.ID,
		&post.Author.Email,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Post{}, nil
		}
		return models.Post{}, models.ErrNotFound
	}
	ParseTime(&post.Created_at, created_at)
	ParseTime(&post.Updated_at, updated_at)

	const query = `
	SELECT email, login FROM authors WHERE email = $1
	`
	err = db.pool.QueryRow(ctx, query, post.Author.Email).Scan(
		&post.Author.Email,
		&post.Author.Login,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Post{}, nil
		}
		return models.Post{}, models.ErrNotFound
	}

	return post, nil
}

func (db *DBStorage) Delete(ctx context.Context, post_id uuid.UUID, email strfmt.Email, role models.UserRole) error {
	// Check that we have the right to delete
	post, err := db.Get(ctx, post_id)
	if err != nil {
		return err
	}
	if post.Author.Email != email && role != models.RoleAdmin && role != models.RoleModerator {
		return models.ErrForbidden
	}

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	const queryPost = `
	DELETE FROM posts WHERE post_id = $1;
	`
	_, err = tx.Exec(ctx, queryPost, post_id)
	if err != nil {
		return err
	}

	// Check if the author has posts, if not, delete him from the authors table
	const queryCheck = `
		SELECT COUNT(auth_email) FROM posts WHERE auth_email = $1;
		`
	var countRows int
	err = db.pool.QueryRow(ctx, queryCheck, post.Author.Email).Scan(
		&countRows,
	)
	if err != nil {
		return err
	}

	if countRows == 1 {
		const queryAuthor = `
		DELETE FROM authors WHERE email = $1;
		`
		_, err = tx.Exec(ctx, queryAuthor, email)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
}

func (db *DBStorage) Save(ctx context.Context, post models.Post) (created_at time.Time, err error) {
	const queryAuthor = `
		SELECT EXISTS (SELECT email, login FROM authors WHERE email = $1);
		`
	var isNotFirstPost bool
	err = db.pool.QueryRow(ctx, queryAuthor, post.Author.Email).Scan(
		&isNotFirstPost,
	)
	if err != nil {
		return time.Time{}, err
	}

	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return created_at, err
	}
	defer tx.Rollback(ctx)

	if !isNotFirstPost {
		const query = `
		INSERT INTO authors (email, login)
		VALUES ($1, $2)
		`
		_, err = tx.Exec(ctx, query, post.Author.Email, post.Author.Login)
		if err != nil {
			return created_at, err
		}
	}

	const postQuery = `
	INSERT INTO posts (title, description, content, created_at, updated_at, post_id, auth_email)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	created_at = time.Now()
	_, err = tx.Exec(ctx, postQuery, post.Title, post.Description, post.Content, created_at, created_at, post.ID, post.Author.Email)
	if err != nil {
		return created_at, err
	}

	err = tx.Commit(ctx)
	return created_at, err
}
func (db *DBStorage) Update(ctx context.Context, id uuid.UUID, data models.NewPostDescription) (updated_at time.Time, err error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return updated_at, err
	}
	defer tx.Rollback(ctx)

	const query = `
	UPDATE posts SET title = $2, description = $3, content = $4, updated_at = $5
	WHERE post_id = $1
	`
	updated_at = time.Now()
	_, err = tx.Exec(ctx, query, id, data.Title, data.Description, data.Content, updated_at)
	if err != nil {
		return updated_at, err
	}
	err = tx.Commit(ctx)
	return updated_at, err
}

func (db *DBStorage) GetAll(ctx context.Context) (post_list []models.Post, err error) {

	const queryPost = `
	SELECT title, description, content, created_at, updated_at, post_id, auth_email FROM posts
	`
	rows, err := db.pool.Query(ctx, queryPost)
	if err != nil {
		return nil, fmt.Errorf("can't get all posts: %v", err)
	}

	for rows.Next() {
		post := models.Post{}

		var created_at time.Time
		var updated_at time.Time

		err = rows.Scan(
			&post.Title,
			&post.Description,
			&post.Content,
			&created_at,
			&updated_at,
			&post.ID,
			&post.Author.Email,
		)
		if err != nil {
			return post_list, err
		}
		ParseTime(&post.Created_at, created_at)
		ParseTime(&post.Updated_at, updated_at)

		const queryAuthor = `
		SELECT email, login FROM authors WHERE email = $1
		`
		author := models.Author{}
		err := db.pool.QueryRow(ctx, queryAuthor, post.Author.Email).Scan(
			&author.Email,
			&author.Login,
		)
		if err != nil {
			return post_list, err
		}
		post.Author = author
		post_list = append(post_list, post)
	}

	return post_list, err
}
