-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;
DROP TYPE IF EXISTS task_status;
CREATE TYPE task_status AS ENUM ('Active', 'Completed', 'Overdue');
CREATE TABLE IF NOT EXISTS task(
    task_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid,
    task_name VARCHAR(255) NOT NULL,
    task_desc TEXT DEFAULT NULL,
    task_status task_status NOT NULL DEFAULT ('Active'),
    priority INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP DEFAULT (NOW())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS task;
-- +goose StatementEnd
