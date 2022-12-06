-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
 
-- +goose StatementBegin
    CREATE TABLE `company_activities` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp,
    `deleted_at` timestamp
    );
-- +goose StatementEnd
-- +goose StatementBegin
    ALTER TABLE `companies`
        ADD COLUMN `company_activity` INT NULL AFTER `size`,
        ADD COLUMN `phone` varchar(15) NULL AFTER `company_activity`,
        ADD COLUMN `website` varchar(100) NULL AFTER `phone`;
-- +goose StatementEnd

-- +goose StatementBegin
    ALTER TABLE `jobs`
        ADD COLUMN `address` varchar(255) NULL AFTER `location`;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `companies` ADD CONSTRAINT FK_company_company_activies FOREIGN KEY (`company_activity`) REFERENCES `company_activities`(`id`);
-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO `company_activities`(`name`) VALUES ( "THÔNG TIN VÀ TRUYỀN THÔNG"),( "HOẠT ĐỘNG CHUYÊN MÔN, KHOA HỌC VÀ CÔNG NGHỆ"),( "HOẠT ĐỘNG KINH DOANH BẤT ĐỘNG SẢN"),( "GIÁO DỤC VÀ ĐÀO TẠO"),( "HOẠT ĐỘNG TÀI CHÍNH, NGÂN HÀNG VÀ BẢO HIỂM"),( "NGHỆ THUẬT, VUI CHƠI VÀ GIẢI TRÍ"),( "Y TẾ VÀ HOẠT ĐỘNG TRỢ GIÚP XÃ HỘI"),("NÔNG NGHIỆP, LÂM NGHIỆP VÀ THUỶ SẢN"),( "CÔNG NGHIỆP CHẾ BIẾN, CHẾ TẠO"),( "SẢN XUẤT VÀ PHÂN PHỐI ĐIỆN, KHÍ ĐỐT, NƯỚC NÓNG, HƠI NƯỚC VÀ ĐIỀU HOÀ KHÔNG KHÍ"),( "XÂY DỰNG"),( "BÁN BUÔN VÀ BÁN LẺ; SỬA CHỮA Ô TÔ, MÔ TÔ, XE MÁY VÀ XE CÓ ĐỘNG CƠ KHÁC"),( "VẬN TẢI KHO BÃI"),( "DỊCH VỤ LƯU TRÚ VÀ ĂN UỐNG"),( "HOẠT ĐỘNG DỊCH VỤ KHÁC");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE `company_activities`
-- +goose StatementEnd
