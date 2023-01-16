CREATE SCHEMA IF NOT EXISTS smart_loader;

CREATE TABLE "user"
(
    id            bigint generated by default as identity primary key,
    name          varchar(256)        not null,
    login_email   varchar(256) unique not null,
    password_hash varchar(256)        not null,
    registered_at timestamp           not null,
    visited_at    timestamp           not null
);

INSERT INTO "user" (name, login_email, password_hash, registered_at, visited_at)
VALUES ('admin', 'admin@mail.ru', '326236336132646265e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5',
        now(), now());

CREATE TABLE token
(
    id         uuid default gen_random_uuid() primary key,
    expired_at timestamp           not null,
    user_id    bigint
        constraint fk_user references "user" not null
);

CREATE TABLE job
(
    id         uuid default gen_random_uuid() primary key,
    size       bigint                        not null,
    status     varchar(32)                   not null,
    created_at timestamp                     not null,
    user_id    bigint
        constraint fk_user references "user" not null
);

CREATE TABLE job_stage
(
    id     bigint generated by default as identity primary key,
    size   bigint                        not null,
    urls   text[] not null,
    status varchar(32)                   not null,
    job_id uuid
        constraint fk_job references job not null
);

CREATE TABLE download
(
    id            uuid default gen_random_uuid() primary key,
    hash          varchar(64) unique not null,
    downloaded_at timestamp          not null,
    size          bigint             not null
);

CREATE TABLE job_stage_download
(
    job_stage_id bigint
        constraint fk_job_stage references job_stage     not null,
    download_id  uuid
        constraint fk_download_stage references download not null
);

CREATE TABLE lock
(
    type       varchar(16) not null,
    value      varchar(64) not null,
    expired_at timestamp   not null,
    PRIMARY KEY (type, value)
);