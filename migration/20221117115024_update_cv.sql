-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
-- +goose StatementBegin
    ALTER TABLE `apply_jobs`
        ADD COLUMN `letter` text NULL AFTER `cv_id`;
-- +goose StatementEnd
-- +goose StatementBegin
    ALTER TABLE `students`
        ADD COLUMN `birthday` timestamp NULL AFTER `gender`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
