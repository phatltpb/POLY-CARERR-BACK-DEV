-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
-- +goose StatementBegin
    ALTER TABLE `student_profiles` 
        ADD CONSTRAINT FK_student_profile_category FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
