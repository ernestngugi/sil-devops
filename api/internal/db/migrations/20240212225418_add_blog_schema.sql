-- +goose Up
CREATE TABLE blogs(
    id              BIGSERIAL       PRIMARY KEY,
    title           VARCHAR(50)     NOT NULL,
    description     TEXT            NOT NULL,
    date_created    TIMESTAMPTZ     NOT NULL DEFAULT clock_timestamp(),
    date_modified    TIMESTAMPTZ     NOT NULL DEFAULT clock_timestamp()
);

-- +goose Down
drop table if exists blogs;
