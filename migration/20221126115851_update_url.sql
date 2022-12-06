-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `users` 
    CHANGE COLUMN `avatar` `avatar` VARCHAR(150) NULL DEFAULT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `students` 
    CHANGE COLUMN `avatar` `avatar` VARCHAR(150) NULL DEFAULT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `student_cvs` 
    CHANGE COLUMN `link` `link` VARCHAR(200) NULL DEFAULT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `companies` 
    CHANGE COLUMN `avatar` `avatar` VARCHAR(150) NULL DEFAULT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `companies` 
    CHANGE COLUMN `banner` `banner` VARCHAR(150) NULL DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
