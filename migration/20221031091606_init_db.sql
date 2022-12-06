-- +goose Up
-- +goose StatementBegin

    CREATE TABLE `cities` (
        `id` int PRIMARY KEY,
        `parent` varchar(11) DEFAULT NULL,
        `name` varchar(255) DEFAULT NULL,
        `code` varchar(11) DEFAULT NULL,
        `created_at` datetime(3) DEFAULT NULL,
        `updated_at` datetime(3) DEFAULT NULL,
        `deleted_at` datetime(3) DEFAULT NULL
    );
-- +goose StatementEnd
-- +goose StatementBegin
    Create table `students`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `full_name` varchar(255) NOT NULL,
        `email` varchar(50) NOT NULL UNIQUE,
        `password` varchar(255) NOT NULL,
        `phone` varchar(15),
        `location` int,
        `address` varchar(255),
        `avatar` varchar(100),
        `gender` int,
        `is_noti` TINYINT(1) DEFAULT 0,
        `is_active` TINYINT(1) DEFAULT 0,
        `status` int,
        `role` int,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp,
        `deleted_at` timestamp
    );
-- +goose StatementEnd
-- +goose StatementBegin
    Create table `student_cvs`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `student_id` int,
        `title` varchar(50),
        `link` varchar(100),
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `deleted_at` timestamp
    );
-- +goose StatementEnd
-- +goose StatementBegin
    Create table `categories`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `parent_id` int,
        `name` varchar(100),
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `deleted_at` timestamp
    );    
-- +goose StatementEnd
-- +goose StatementBegin
    Create table `users`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `full_name` varchar(255) NOT NULL,
        `email` varchar(50) NOT NULL,
        `email_noti` varchar(50),
        `password` varchar(255) NOT NULL,
        `avatar` varchar(100),
        `address` varchar(255),
        `role` int NOT NULL,
        `phone` varchar(15),
        `company_id` int,
        `is_owner` TINYINT(1) DEFAULT 1,
        `is_active` TINYINT(1) DEFAULT 0,
        `is_noti` TINYINT(1) DEFAULT 0,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp,
        `deleted_at` timestamp
    );   
-- +goose StatementEnd
-- +goose StatementBegin
    Create table `companies`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `name` varchar(255) NOT NULL,
        `location` int,
        `address` varchar(255),
        `banner` varchar(100),
        `avatar` varchar(100),
        `tax_code` varchar(20) NOT NULL,
        `information` text,
        `size` int,
        `is_hidden` TINYINT(1) DEFAULT 0,
        `is_active` TINYINT(1) DEFAULT 1,
        `is_noti` TINYINT(1) DEFAULT 0,
        `status` int NOT NULL DEFAULT 1
    );    
-- +goose StatementEnd

-- +goose StatementBegin
    Create table `jobs`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `title`  varchar(255),
        `category_id`  int,
        `user_id` int NOT NULL,
        `gender` varchar(20),
        `job_type` varchar(20),
        `count` int,
        `experience` varchar(50),
        `position` varchar(50),
        `salary` decimal(10),
        `location` int,
        `level` varchar(50),
        `description` text,
        `require` text,
        `benefit` text,
        `is_hidden` TINYINT(1) DEFAULT 0,
        `status` int,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp,
        `deleted_at` timestamp
    )    
-- +goose StatementEnd
-- +goose StatementBegin
    Create table `apply_jobs`(
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `post_id`  int,
        `cv_id`  int,
        `date_apply` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp,
        `status` int
    )    
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `categories` ADD CONSTRAINT FK_category_parent FOREIGN KEY (`parent_id`) REFERENCES `categories` (`id`);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `users` ADD CONSTRAINT FK_user_company FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `jobs` ADD CONSTRAINT FK_user FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `jobs` ADD CONSTRAINT FK_category_post FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `student_cvs` ADD CONSTRAINT FK_student FOREIGN KEY (`student_id`) REFERENCES `students` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `apply_jobs` ADD CONSTRAINT FK_post FOREIGN KEY (`post_id`) REFERENCES `jobs` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `apply_jobs` ADD CONSTRAINT FK_cv FOREIGN KEY (`cv_id`) REFERENCES `student_cvs` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `students` ADD CONSTRAINT FK_student_location FOREIGN KEY (`location`) REFERENCES `cities` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `companies` ADD CONSTRAINT FK_company_location FOREIGN KEY (`location`) REFERENCES `cities` (`id`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `jobs` ADD CONSTRAINT FK_job_location FOREIGN KEY (`location`) REFERENCES `cities` (`id`);
-- +goose StatementEnd




-- +goose StatementBegin
INSERT INTO `cities` VALUES (1,'00','Thành phố Hà Nội','01',NULL,NULL,NULL),(2,'00','Tỉnh Hà Giang','02',NULL,NULL,NULL),(3,'00','Tỉnh Cao Bằng','04',NULL,NULL,NULL),(4,'00','Tỉnh Bắc Kạn','06',NULL,NULL,NULL),(5,'00','Tỉnh Tuyên Quang','08',NULL,NULL,NULL),(6,'00','Tỉnh Lào Cai','10',NULL,NULL,NULL),(7,'00','Tỉnh Điện Biên','11',NULL,NULL,NULL),(8,'00','Tỉnh Lai Châu','12',NULL,NULL,NULL),(9,'00','Tỉnh Sơn La','14',NULL,NULL,NULL),(10,'00','Tỉnh Yên Bái','15',NULL,NULL,NULL),(11,'00','Tỉnh Hoà Bình','17',NULL,NULL,NULL),(12,'00','Tỉnh Thái Nguyên','19',NULL,NULL,NULL),(13,'00','Tỉnh Lạng Sơn','20',NULL,NULL,NULL),(14,'00','Tỉnh Quảng Ninh','22',NULL,NULL,NULL),(15,'00','Tỉnh Bắc Giang','24',NULL,NULL,NULL),(16,'00','Tỉnh Phú Thọ','25',NULL,NULL,NULL),(17,'00','Tỉnh Vĩnh Phúc','26',NULL,NULL,NULL),(18,'00','Tỉnh Bắc Ninh','27',NULL,NULL,NULL),(19,'00','Tỉnh Hải Dương','30',NULL,NULL,NULL),(20,'00','Thành phố Hải Phòng','31',NULL,NULL,NULL),(21,'00','Tỉnh Hưng Yên','33',NULL,NULL,NULL),(22,'00','Tỉnh Thái Bình','34',NULL,NULL,NULL),(23,'00','Tỉnh Hà Nam','35',NULL,NULL,NULL),(24,'00','Tỉnh Nam Định','36',NULL,NULL,NULL),(25,'00','Tỉnh Ninh Bình','37',NULL,NULL,NULL),(26,'00','Tỉnh Thanh Hóa','38',NULL,NULL,NULL),(27,'00','Tỉnh Nghệ An','40',NULL,NULL,NULL),(28,'00','Tỉnh Hà Tĩnh','42',NULL,NULL,NULL),(29,'00','Tỉnh Quảng Bình','44',NULL,NULL,NULL),(30,'00','Tỉnh Quảng Trị','45',NULL,NULL,NULL),(31,'00','Tỉnh Thừa Thiên Huế','46',NULL,NULL,NULL),(32,'00','Thành phố Đà Nẵng','48',NULL,NULL,NULL),(33,'00','Tỉnh Quảng Nam','49',NULL,NULL,NULL),(34,'00','Tỉnh Quảng Ngãi','51',NULL,NULL,NULL),(35,'00','Tỉnh Bình Định','52',NULL,NULL,NULL),(36,'00','Tỉnh Phú Yên','54',NULL,NULL,NULL),(37,'00','Tỉnh Khánh Hòa','56',NULL,NULL,NULL),(38,'00','Tỉnh Ninh Thuận','58',NULL,NULL,NULL),(39,'00','Tỉnh Bình Thuận','60',NULL,NULL,NULL),(40,'00','Tỉnh Kon Tum','62',NULL,NULL,NULL),(41,'00','Tỉnh Gia Lai','64',NULL,NULL,NULL),(42,'00','Tỉnh Đắk Lắk','66',NULL,NULL,NULL),(43,'00','Tỉnh Đắk Nông','67',NULL,NULL,NULL),(44,'00','Tỉnh Lâm Đồng','68',NULL,NULL,NULL),(45,'00','Tỉnh Bình Phước','70',NULL,NULL,NULL),(46,'00','Tỉnh Tây Ninh','72',NULL,NULL,NULL),(47,'00','Tỉnh Bình Dương','74',NULL,NULL,NULL),(48,'00','Tỉnh Đồng Nai','75',NULL,NULL,NULL),(49,'00','Tỉnh Bà Rịa - Vũng Tàu','77',NULL,NULL,NULL),(50,'00','Thành phố Hồ Chí Minh','79',NULL,NULL,NULL),(51,'00','Tỉnh Long An','80',NULL,NULL,NULL),(52,'00','Tỉnh Tiền Giang','82',NULL,NULL,NULL),(53,'00','Tỉnh Bến Tre','83',NULL,NULL,NULL),(54,'00','Tỉnh Trà Vinh','84',NULL,NULL,NULL),(55,'00','Tỉnh Vĩnh Long','86',NULL,NULL,NULL),(56,'00','Tỉnh Đồng Tháp','87',NULL,NULL,NULL),(57,'00','Tỉnh An Giang','89',NULL,NULL,NULL),(58,'00','Tỉnh Kiên Giang','91',NULL,NULL,NULL),(59,'00','Thành phố Cần Thơ','92',NULL,NULL,NULL),(60,'00','Tỉnh Hậu Giang','93',NULL,NULL,NULL),(61,'00','Tỉnh Sóc Trăng','94',NULL,NULL,NULL),(62,'00','Tỉnh Bạc Liêu','95',NULL,NULL,NULL),(63,'00','Tỉnh Cà Mau','96',NULL,NULL,NULL);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO `users` (`full_name`, `email`, `password`, `role`,`is_active`) VALUES ('Super Admin', 'super.admin@fpt.edu.vn', '$2a$14$b707KjTpD.di5karjs0X/O/6x8HGUM07uWWcy0FUG673VE64VIwOC', '4','1');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO `categories`(`name`) VALUES ("Lập trình máy tính"),("Ứng dụng phần mềm"),("Thiết kế Web"),("Thiết kế đồ hoạ"),("Quan hệ công chúng"),("Quản trị doanh nghiệp"),("Thương mại điện tử"),("Tự động hoá"),("Quản trị du lịch"),("Quản trị nhà hàng"),("Quản trị khách sạn");
-- +goose StatementEnd

-- +goose Down
    DROP TABLE IF EXISTS `apply_jobs`;
    DROP TABLE IF EXISTS `student_cvs`; 
    DROP TABLE IF EXISTS `students`; 
    DROP TABLE IF EXISTS `company_users`; 
    DROP TABLE IF EXISTS `company_users`; 
    DROP TABLE IF EXISTS `students`; 
    DROP TABLE IF EXISTS `jobs`; 
    DROP TABLE IF EXISTS `categories`; 
    DROP TABLE IF EXISTS `users`; 
    DROP TABLE IF EXISTS `companies`; 
    DROP TABLE IF EXISTS `cities`; 