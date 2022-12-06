-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
-- +goose StatementBegin
    CREATE TABLE `student_profiles`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `student_id` int,
        `position_Wish` varchar(100),
        `level_wish` varchar(100),
        `level_current` varchar(100),
        `experience` varchar(50),
        `salary_wish` decimal(10),
        `category_id` int,
        `job_type` varchar(20),
        `province_id` int,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp
     )
-- +goose StatementEnd
-- +goose StatementBegin
    CREATE TABLE `student_educations`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `student_id` int,
        `degree` varchar(255),
        `rank` varchar(50),
        `information` text,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp
     )
-- +goose StatementEnd
-- +goose StatementBegin
    ALTER TABLE `student_profiles` 
        ADD CONSTRAINT FK_student_profile FOREIGN KEY (`student_id`) REFERENCES `students` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
    ALTER TABLE `student_profiles` 
               ADD CONSTRAINT FK_student_profile_city FOREIGN KEY (`province_id`) REFERENCES `cities` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
    ALTER TABLE `student_educations` ADD CONSTRAINT FK_student_edutcation FOREIGN KEY (`student_id`) REFERENCES `students` (`id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
