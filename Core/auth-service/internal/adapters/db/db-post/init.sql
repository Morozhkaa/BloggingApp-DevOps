SELECT 'CREATE DATABASE post' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'post')\gexec
\c post;

CREATE TABLE authors (
    email TEXT PRIMARY KEY,
    login TEXT NOT NULL
);

CREATE TABLE posts (
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    post_id VARCHAR(36) PRIMARY KEY,
    auth_email VARCHAR(36) NOT NULL,
    FOREIGN KEY (auth_email) REFERENCES authors(email) ON DELETE CASCADE
);

CREATE TABLE commenters (
    email TEXT PRIMARY KEY,
    login TEXT NOT NULL
);

CREATE TABLE comments (
    text TEXT NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    comment_id VARCHAR(36) PRIMARY KEY,
    post_id VARCHAR(36) NOT NULL,
    commenter_email VARCHAR(36) NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
    FOREIGN KEY (commenter_email) REFERENCES commenters(email) ON DELETE CASCADE
);

INSERT INTO authors (email, login)
VALUES (
    'Olenka@mail.ru',
    'Olenka'
);

INSERT INTO posts (title, description, content, post_id, auth_email)
VALUES (
    'Пингвины-гиганты',
    'Раньше на планете обитали пингвины-гиганты',
    'Пингвины — забавные морские птицы, лишенные способности летать, но зато отлично умеющие плавать. Наверняка многие из нас видели, как нелепо эти животные передвигаются по суше или поскальзываются на льду. Современные пингвины редко превышают метр в длину и кажутся довольно безобидными для человека зверями. Но относились бы мы к ним так же, если бы рост среднестатистического пингвина превосходил человеческий, а масса колебалась бы у отметки в 100 кг? По словам археологов, именно такие пингвины-гиганты обитали раньше на нашей планете. Окаменелости такого пингвина, которого позже отнесли к новому виду Kumimanu biceae, были обнаружены при раскопках в Новой Зеландии. Установлено, что Kumimanu обитал на Земле около 55-59 млн лет назад — уже после начала вымирания динозавров. Окаменелые останки пингвина дают основания полагать, что его тело достигало в длину 175 см. Вряд ли ученым попалась самая крупная особь данного вида, поэтому можно предположить, что пингвины-гиганты достигали длины и в 180 см — среднестатистический рост современного мужчины.',
    '3ad7e1a0-3a7e-4891-b3ea-b0c1c2ab64f5',
    'Olenka@mail.ru'
);

INSERT INTO authors (email, login)
VALUES (
    'Kira@mail.ru',
    'Kira'
);

INSERT INTO posts (title, description, content, post_id, auth_email)
VALUES (
    'Парижский синдром',
    'Парижский синдром — когда город влюбленных не оправдывает ожиданий',
    'Париж — город-мечта, на идеально чистых улицах которого пахнет свежеиспеченными круассанами, а прохожие благодушно появляются при одном лишь вашем появлении. По крайней мере, именно такую утопическую картину о столице Франции рисует нам массовая культура. Но чаще всего приезжающий турист сталкивается с совершенно иной действительностью — крупный город спешит по своим делам, по улицам летают клубы пыли, а вокруг достопримечательностей стоят толпы назойливых продавцов сувениров. Многих людей, у которых не оправдались ожидания, охватывает депрессия. Разочарованные туристы бывают настолько шокированы, что им требуется помощь психиатра. Интересно, что подобные проблемы чаще всего наблюдаются у жителей азиатских стран, которые впервые приезжают в другую культурную реальность, в своих фантазиях представляя ее совсем иначе. Специалисты, все чаще наблюдая такие расстройства, даже придумали для них термин: «парижский синдром». Впервые его ввел в употребление в 1986 году японский психиатр Хироаки Отой, осматривая своего грустного соотечественника, побывавшего.в столице Франции. Ежегодно несколько десятков туристов заболевают парижским синдромом и лишь после профессиональной терапии они вновь приходят в нормальное состояние.',
    'f0013854-ef1f-40a9-bbf0-1349a643871d',
    'Kira@mail.ru'
);

INSERT INTO commenters (email, login)
VALUES (
    'Olenka@mail.ru',
    'Olenka'
);

INSERT INTO commenters (email, login)
VALUES (
    'Kira@mail.ru',
    'Kira'
);

INSERT INTO commenters (email, login)
VALUES (
    'Polina@mail.ru',
    'Polina'
);

INSERT INTO comments (text, comment_id, post_id, commenter_email)
VALUES (
    'Мне очень понравилась данная статья, спасибо автору! Надеюсь, вы напишете еще парочку статей, было бы интересно изучить культуру других стран!! 😀',
    'e895d80c-c0c4-42c4-8730-6faae51905bf',
    'f0013854-ef1f-40a9-bbf0-1349a643871d',
    'Polina@mail.ru'
);

INSERT INTO comments (text, comment_id, post_id, commenter_email)
VALUES (
    'Ого! 🫢',
    'de8b1c21-9eef-4e82-a80f-c63ff3bc92b7',
    '3ad7e1a0-3a7e-4891-b3ea-b0c1c2ab64f5',
    'Kira@mail.ru'
);

INSERT INTO comments (text, comment_id, post_id, commenter_email)
VALUES (
    'Было бы интересно посмотреть картинки!',
    '500ca845-caab-41e7-bd75-b31cb3afb307',
    '3ad7e1a0-3a7e-4891-b3ea-b0c1c2ab64f5',
    'Polina@mail.ru'
);