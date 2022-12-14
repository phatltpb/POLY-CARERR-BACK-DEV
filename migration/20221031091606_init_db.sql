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
INSERT INTO `cities` VALUES (1,'00','Th??nh ph??? H?? N???i','01',NULL,NULL,NULL),(2,'00','T???nh H?? Giang','02',NULL,NULL,NULL),(3,'00','T???nh Cao B???ng','04',NULL,NULL,NULL),(4,'00','T???nh B???c K???n','06',NULL,NULL,NULL),(5,'00','T???nh Tuy??n Quang','08',NULL,NULL,NULL),(6,'00','T???nh L??o Cai','10',NULL,NULL,NULL),(7,'00','T???nh ??i???n Bi??n','11',NULL,NULL,NULL),(8,'00','T???nh Lai Ch??u','12',NULL,NULL,NULL),(9,'00','T???nh S??n La','14',NULL,NULL,NULL),(10,'00','T???nh Y??n B??i','15',NULL,NULL,NULL),(11,'00','T???nh Ho?? B??nh','17',NULL,NULL,NULL),(12,'00','T???nh Th??i Nguy??n','19',NULL,NULL,NULL),(13,'00','T???nh L???ng S??n','20',NULL,NULL,NULL),(14,'00','T???nh Qu???ng Ninh','22',NULL,NULL,NULL),(15,'00','T???nh B???c Giang','24',NULL,NULL,NULL),(16,'00','T???nh Ph?? Th???','25',NULL,NULL,NULL),(17,'00','T???nh V??nh Ph??c','26',NULL,NULL,NULL),(18,'00','T???nh B???c Ninh','27',NULL,NULL,NULL),(19,'00','T???nh H???i D????ng','30',NULL,NULL,NULL),(20,'00','Th??nh ph??? H???i Ph??ng','31',NULL,NULL,NULL),(21,'00','T???nh H??ng Y??n','33',NULL,NULL,NULL),(22,'00','T???nh Th??i B??nh','34',NULL,NULL,NULL),(23,'00','T???nh H?? Nam','35',NULL,NULL,NULL),(24,'00','T???nh Nam ?????nh','36',NULL,NULL,NULL),(25,'00','T???nh Ninh B??nh','37',NULL,NULL,NULL),(26,'00','T???nh Thanh H??a','38',NULL,NULL,NULL),(27,'00','T???nh Ngh??? An','40',NULL,NULL,NULL),(28,'00','T???nh H?? T??nh','42',NULL,NULL,NULL),(29,'00','T???nh Qu???ng B??nh','44',NULL,NULL,NULL),(30,'00','T???nh Qu???ng Tr???','45',NULL,NULL,NULL),(31,'00','T???nh Th???a Thi??n Hu???','46',NULL,NULL,NULL),(32,'00','Th??nh ph??? ???? N???ng','48',NULL,NULL,NULL),(33,'00','T???nh Qu???ng Nam','49',NULL,NULL,NULL),(34,'00','T???nh Qu???ng Ng??i','51',NULL,NULL,NULL),(35,'00','T???nh B??nh ?????nh','52',NULL,NULL,NULL),(36,'00','T???nh Ph?? Y??n','54',NULL,NULL,NULL),(37,'00','T???nh Kh??nh H??a','56',NULL,NULL,NULL),(38,'00','T???nh Ninh Thu???n','58',NULL,NULL,NULL),(39,'00','T???nh B??nh Thu???n','60',NULL,NULL,NULL),(40,'00','T???nh Kon Tum','62',NULL,NULL,NULL),(41,'00','T???nh Gia Lai','64',NULL,NULL,NULL),(42,'00','T???nh ?????k L???k','66',NULL,NULL,NULL),(43,'00','T???nh ?????k N??ng','67',NULL,NULL,NULL),(44,'00','T???nh L??m ?????ng','68',NULL,NULL,NULL),(45,'00','T???nh B??nh Ph?????c','70',NULL,NULL,NULL),(46,'00','T???nh T??y Ninh','72',NULL,NULL,NULL),(47,'00','T???nh B??nh D????ng','74',NULL,NULL,NULL),(48,'00','T???nh ?????ng Nai','75',NULL,NULL,NULL),(49,'00','T???nh B?? R???a - V??ng T??u','77',NULL,NULL,NULL),(50,'00','Th??nh ph??? H??? Ch?? Minh','79',NULL,NULL,NULL),(51,'00','T???nh Long An','80',NULL,NULL,NULL),(52,'00','T???nh Ti???n Giang','82',NULL,NULL,NULL),(53,'00','T???nh B???n Tre','83',NULL,NULL,NULL),(54,'00','T???nh Tr?? Vinh','84',NULL,NULL,NULL),(55,'00','T???nh V??nh Long','86',NULL,NULL,NULL),(56,'00','T???nh ?????ng Th??p','87',NULL,NULL,NULL),(57,'00','T???nh An Giang','89',NULL,NULL,NULL),(58,'00','T???nh Ki??n Giang','91',NULL,NULL,NULL),(59,'00','Th??nh ph??? C???n Th??','92',NULL,NULL,NULL),(60,'00','T???nh H???u Giang','93',NULL,NULL,NULL),(61,'00','T???nh S??c Tr??ng','94',NULL,NULL,NULL),(62,'00','T???nh B???c Li??u','95',NULL,NULL,NULL),(63,'00','T???nh C?? Mau','96',NULL,NULL,NULL);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO `users` (`full_name`, `email`, `password`, `role`,`is_active`) VALUES ('Super Admin', 'super.admin@fpt.edu.vn', '$2a$14$b707KjTpD.di5karjs0X/O/6x8HGUM07uWWcy0FUG673VE64VIwOC', '4','1');
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO `categories`(`name`) VALUES ("L???p tr??nh m??y t??nh"),("???ng d???ng ph???n m???m"),("Thi???t k??? Web"),("Thi???t k??? ????? ho???"),("Quan h??? c??ng ch??ng"),("Qu???n tr??? doanh nghi???p"),("Th????ng m???i ??i???n t???"),("T??? ?????ng ho??"),("Qu???n tr??? du l???ch"),("Qu???n tr??? nh?? h??ng"),("Qu???n tr??? kh??ch s???n");
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