-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose StatementBegin
    CREATE TABLE `student_watches` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `student_id` int,
    `job_id` int,
    `count` decimal(5)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
