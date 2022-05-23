CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE, 
    password VARCHAR(60) NOT NULL,
    permission BIGINT NOT NULL DEFAULT 0,
    status INT NOT NULL DEFAULT 0,
    dateline TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_del BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE urls (
    id CHAR(6) PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    is_customize BOOLEAN NOT NULL DEFAULT FALSE,
    user_id BIGINT NOT NULL,
    is_protected BOOLEAN NOT NULL DEFAULT FALSE,
    password VARCHAR(60) NOT NULL DEFAULT '',
    dateline TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    visit BIGINT NOT NULL DEFAULT 0,
    is_del BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE url_visits (
    id BIGSERIAL PRIMARY KEY, 
    url_id CHAR(6) NOT NULL,
    ip VARCHAR(45) NOT NULL DEFAULT '',
    user_agent VARCHAR(255) NOT NULL DEFAULT '',
    dateline TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (url_id) REFERENCES urls (id)
);

CREATE VIEW v_user_urls AS
SELECT ur.id AS url_id, url, is_customize, is_protected, ur.password AS protected_password, ur.dateline, ur.is_del AS url_is_del, u.id AS user_id, email, u.is_del AS user_is_del, visit
FROM urls AS ur
INNER JOIN
users AS u
ON ur.user_id=u.id;

CREATE VIEW v_user_url_visits AS
SELECT 
    u.url_id, url, is_customize, is_protected, protected_password, u.dateline, url_is_del, user_id,email, user_is_del, visit,
    v.id AS visit_id, ip, user_agent, v.dateline AS visit_dateline
FROM
    v_user_urls AS u
INNER JOIN
    url_visits AS v
ON 
    v.url_id = u.url_id;

