USE SE104;
#  Create a table `sea_sons` with the following columns:
#     INSERT INTO `sea_sons` (`sea_son_id`,`name`, `country`, `year`) VALUES
#     (SHA1(UUID()), 'NightWolf', 'VietNam', '2021-2022'),
#     (SHA1(UUID()), 'NightWolf', 'VietNam', '2021-2022'),
#     (SHA1(UUID()), 'NightWolf', 'VietNam', '2021-2022');
#
# SELECT * FROM sea_sons;
#  Create a table `clubs` with the following columns:


#2021-22
INSERT INTO `clubs` (`id`, `name_club`, `shorthand`,`domain_email`,`sea_son_id`) VALUES
                                          (SHA1(UUID()),'Becamex Bình Dương', 'BDFC', 'bdfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Sài Gòn', 'SGFC', 'sgfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Đông Á Thanh Hóa', 'THFC', 'thfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Hà Nội', 'HNFC', 'hnfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Hải Phòng', 'HPFC', 'hpfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'LPBank Hoàng Anh Gia Lai', 'LPBHA', 'lpbha.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Hồng Lĩnh Hà Tĩnh', 'HLHT', 'hlht.vn',(SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Thể Công - Viettel', 'TCVT', 'tcvt.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'MerryLand Quy Nhơn Bình Định', 'MQBĐ', 'mqbd.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Sông Lam Nghệ An', 'SLNA', 'slna.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'Thép Xanh Nam Định', 'TXND', 'txnd.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'TP Hồ Chí Minh', 'HCMC', 'hcmc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22')),
                                          (SHA1(UUID()),'SHB Đà Nẵng', 'SHBĐN', 'shbdn.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2021-22'));

#2022-23
INSERT INTO `clubs` (`id`, `name_club`, `shorthand`,`domain_email`,`sea_son_id`) VALUES
                                          (SHA1(UUID()),'Becamex Bình Dương', 'BDFC', 'bdfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Công An Hà Nội', 'CAHN', 'cahn.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Đông Á Thanh Hóa', 'THFC', 'thfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Hà Nội', 'HNFC', 'hnfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Hải Phòng', 'HPFC', 'hpfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'LPBank Hoàng Anh Gia Lai', 'LPBHA', 'lpbha.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Hồng Lĩnh Hà Tĩnh', 'HLHT', 'hlht.vn',(SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Khánh Hòa', 'KHFC', 'khfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Thể Công - Viettel', 'TCVT', 'tcvt.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'MerryLand Quy Nhơn Bình Định', 'MQBĐ', 'mqbd.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Sông Lam Nghệ An', 'SLNA', 'slna.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'Thép Xanh Nam Định', 'TXND', 'txnd.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'TP Hồ Chí Minh', 'HCMC', 'hcmc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023')),
                                          (SHA1(UUID()),'SHB Đà Nẵng', 'SHBĐN', 'shbdn.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2022-2023'));
#2023-24
INSERT INTO `clubs` (`id`, `name_club`, `shorthand`,`domain_email`,`sea_son_id`) VALUES
                                          (SHA1(UUID()),'Becamex Bình Dương', 'BDFC', 'bdfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Công An Hà Nội', 'CAHN', 'cahn.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Đông Á Thanh Hóa', 'THFC', 'thfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Hà Nội', 'HNFC', 'hnfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Hải Phòng', 'HPFC', 'hpfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'LPBank Hoàng Anh Gia Lai', 'LPBHA', 'lpbha.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Hồng Lĩnh Hà Tĩnh', 'HLHT', 'hlht.vn',(SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Khánh Hòa', 'KHFC', 'khfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Thể Công - Viettel', 'TCVT', 'tcvt.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'MerryLand Quy Nhơn Bình Định', 'MQBĐ', 'mqbd.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Sông Lam Nghệ An', 'SLNA', 'slna.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Thép Xanh Nam Định', 'TXND', 'txnd.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'TP Hồ Chí Minh', 'HCMC', 'hcmc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024')),
                                          (SHA1(UUID()),'Quảng Nam', 'QNFC', 'qnfc.vn', (SELECT `sea_son_id` FROM `sea_sons` WHERE `year` = '2023-2024'));

# SELECT * FROM `clubs`