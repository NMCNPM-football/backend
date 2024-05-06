# CREATE DATABASE SE104;
USE SE104;

-- Create a table `sea_sons` with the following columns:
   INSERT INTO `sea_sons` (`sea_son_id`,`name`, `country`, `year`) VALUES
     (SHA1(UUID()), 'NightWolf', 'VietNam', '2021-2022'),
     (SHA1(UUID()), 'NightWolf', 'VietNam', '2022-2023'),
     (SHA1(UUID()), 'NightWolf', 'VietNam', '2023-2024');

SELECT * FROM sea_sons;
-- Create a table `clubs` with the following columns:


-- 2021-22
INSERT INTO `clubs` (`id`, `name_club`, `shorthand`,`domain_email`,`sea_son`,`created_at`,`created_by`) VALUES
                                          (SHA1(UUID()),'Becamex Bình Dương', 'BDFC', 'bdfc.vn', '2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'Sài Gòn', 'SGFC', 'sgfc.vn','2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'Đông Á Thanh Hóa', 'THFC', 'thfc.vn','2021-2022', NOW() , 'admin'),
                                          (SHA1(UUID()),'Hà Nội', 'HNFC', 'hnfc.vn', '2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'Hải Phòng', 'HPFC', 'hpfc.vn','2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'LPBank Hoàng Anh Gia Lai', 'LPBHA', 'lpbha.vn','2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'Hồng Lĩnh Hà Tĩnh', 'HLHT', 'hlht.vn','2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'Thể Công - Viettel', 'TCVT', 'tcvt.vn', '2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'MerryLand Quy Nhơn Bình Định', 'MQBĐ', 'mqbd.vn', '2021-2022',NOW(), 'admin'),
                                          (SHA1(UUID()),'Sông Lam Nghệ An', 'SLNA', 'slna.vn', '2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'Thép Xanh Nam Định', 'TXND', 'txnd.vn', '2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'TP Hồ Chí Minh', 'HCMC', 'hcmc.vn', '2021-2022', NOW(), 'admin'),
                                          (SHA1(UUID()),'SHB Đà Nẵng', 'SHBĐN', 'shbdn.vn', '2021-2022', NOW(), 'admin');

-- 2022-23
INSERT INTO `clubs` (`id`, `name_club`, `shorthand`,`domain_email`,`sea_son`,`created_at`,`created_by`) VALUES
                                          (SHA1(UUID()),'Becamex Bình Dương', 'BDFC', 'bdfc.vn', '2022-2023', NOW(), 'admin'),
                                          (SHA1(UUID()),'Công An Hà Nội', 'CAHN', 'cahn.vn', '2022-2023', NOW(), 'admin'),
                                          (SHA1(UUID()),'Đông Á Thanh Hóa', 'THFC', 'thfc.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'Hà Nội', 'HNFC', 'hnfc.vn','2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'Hải Phòng', 'HPFC', 'hpfc.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'LPBank Hoàng Anh Gia Lai', 'LPBHA', 'lpbha.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'Hồng Lĩnh Hà Tĩnh', 'HLHT', 'hlht.vn','2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'Khánh Hòa', 'KHFC', 'khfc.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'Thể Công - Viettel', 'TCVT', 'tcvt.vn','2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'MerryLand Quy Nhơn Bình Định', 'MQBĐ', 'mqbd.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'Sông Lam Nghệ An', 'SLNA', 'slna.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'Thép Xanh Nam Định', 'TXND', 'txnd.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'TP Hồ Chí Minh', 'HCMC', 'hcmc.vn', '2022-2023',NOW(), 'admin'),
                                          (SHA1(UUID()),'SHB Đà Nẵng', 'SHBĐN', 'shbdn.vn','2022-2023',NOW(), 'admin');
-- 2023-24
INSERT INTO `clubs` (`id`, `name_club`, `shorthand`,`domain_email`,`sea_son`,`created_at`,`created_by`) VALUES
                                        (SHA1(UUID()),'Becamex Bình Dương', 'BDFC', 'bdfc.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Công An Hà Nội', 'CAHN', 'cahn.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Đông Á Thanh Hóa', 'THFC', 'thfc.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Hà Nội', 'HNFC', 'hnfc.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Hải Phòng', 'HPFC', 'hpfc.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'LPBank Hoàng Anh Gia Lai', 'LPBHA', 'lpbha.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Hồng Lĩnh Hà Tĩnh', 'HLHT', 'hlht.vn','2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Khánh Hòa', 'KHFC', 'khfc.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Thể Công - Viettel', 'TCVT', 'tcvt.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'MerryLand Quy Nhơn Bình Định', 'MQBĐ', 'mqbd.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Sông Lam Nghệ An', 'SLNA', 'slna.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Thép Xanh Nam Định', 'TXND', 'txnd.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'TP Hồ Chí Minh', 'HCMC', 'hcmc.vn', '2023-2024',NOW(), 'admin'),
                                        (SHA1(UUID()),'Quảng Nam', 'QNFC', 'qnfc.vn', '2023-2024',NOW(), 'admin');
SELECT * FROM `clubs`
# SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-2022';