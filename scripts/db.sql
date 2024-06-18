CREATE TABLE IF NOT EXISTS sys_user
(
    id         bigserial PRIMARY KEY,
    name       text                        not null UNIQUE,
    password   text                        not null default '',
    mobile     text                        not null default '',
    email      TEXT                        not null default '',
    avatar_url text                        not null default '',
    language   text                        not null default '',
    role       int                         not null default 0,
    created_at TIMESTAMP without time zone NOT NULL default current_timestamp,
    updated_at TIMESTAMP without time zone NOT NULL default current_timestamp,
    deleted_at TIMESTAMP without time zone NULL
);

INSERT INTO "sys_user"("name","password","created_at","updated_at") VALUES('frankie','',now(),now());