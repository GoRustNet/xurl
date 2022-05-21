CREATE TABLE urls (
    id CHAR(6) PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    is_customize BOOLEAN NOT NULL DEFAULT FALSE,
    is_del BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE, 
    password VARCHAR(60) NOT NULL,
    permission BIGINT NOT NULL DEFAULT 0,
    status INT NOT NULL DEFAULT 0,
    dateline TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_del BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE user_urls (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    url_id CHAR(6) NOT NULL,
    is_protected BOOLEAN NOT NULL DEFAULT FALSE,
    password VARCHAR(60) NOT NULL DEFAULT '',
    dateline TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_del BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE url_visits (
    id BIGSERIAL PRIMARY KEY, 
    user_url_id BIGINT NOT NULL,
    ip VARCHAR(45) NOT NULL DEFAULT '',
    user_agent VARCHAR(255) NOT NULL DEFAULT '',
    dateline TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE VIEW v_user_urls AS

SELECT
    u.id AS user_id,
    email,
    u.is_del AS user_is_del,
    ur.id AS url_id,
    ur.url AS target_url,
    is_customize,
    ur.is_del AS url_is_del,
    uu.id AS user_url_id,
    is_protected,
    uu.password AS protected_password,
    uu.dateline
FROM
    user_urls AS uu
INNER JOIN
    urls AS ur
ON uu.url_id=ur.id
INNER JOIN
    users AS u
ON uu.user_id=u.id;

