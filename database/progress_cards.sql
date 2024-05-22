USE SE104;

INSERT INTO `progress_cards` (id, created_at, updated_at, deleted_at, match_id, club_name, player_name, card_type, time_in_match) VALUES
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-20 18:00' AND club_one_name = 'Hải Phòng' AND club_two_name = 'LPBank Hoàng Anh Gia Lai'), '', '', 'LT03', ''),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-21 18:00' AND club_one_name = 'Đông Á Thanh Hóa' AND club_two_name = 'Hồng Lĩnh Hà Tĩnh'), 'Đông Á Thanh Hóa', 'Popov Velizar Emilov', 'LT01', '26'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-21 18:00' AND club_one_name = 'Đông Á Thanh Hóa' AND club_two_name = 'Hồng Lĩnh Hà Tĩnh'), 'Hồng Lĩnh Hà Tĩnh', 'Vũ Viết Triều', 'LT01', '34'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-21 18:00' AND club_one_name = 'Đông Á Thanh Hóa' AND club_two_name = 'Hồng Lĩnh Hà Tĩnh'), 'Hồng Lĩnh Hà Tĩnh', 'Nguyễn Văn Hạnh', 'LT01', '55'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-21 18:00' AND club_one_name = 'Đông Á Thanh Hóa' AND club_two_name = 'Hồng Lĩnh Hà Tĩnh'), 'Đông Á Thanh Hóa', 'Nguyễn Trọng Hùng', 'LT01', '55'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 17:00' AND club_one_name = 'Sông Lam Nghệ An' AND club_two_name = 'Thể Công - Viettel'), 'Thể Công – Viettel', 'Abdumuminov Jakhongir', 'LT01', '40'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 17:00' AND club_one_name = 'Sông Lam Nghệ An' AND club_two_name = 'Thể Công - Viettel'), 'Sông Lam Nghệ An', 'Phan Bá Quyền', 'LT01', '79'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 17:00' AND club_one_name = 'Sông Lam Nghệ An' AND club_two_name = 'Thể Công - Viettel'), 'Thể Công – Viettel', 'Vũ Văn Quyết', 'LT01', '85'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 18:00' AND club_one_name = 'Thép Xanh Nam Định' AND club_two_name = 'Quảng Nam'), 'Thép Xanh Nam Định', 'Nguyễn Phong Hồng Duy', 'LT01', '21'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 18:00' AND club_one_name = 'Thép Xanh Nam Định' AND club_two_name = 'Quảng Nam'), 'Quảng Nam', 'Lê Xuân Tú', 'LT01', '75'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'Công An Hà Nội' AND club_two_name = 'MerryLand Quy Nhơn Bình Định'), 'MerryLand Quy Nhơn Bình Định', 'Marlon Rangel De Almeida', 'LT01', '30'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'Công An Hà Nội' AND club_two_name = 'MerryLand Quy Nhơn Bình Định'), 'MerryLand Quy Nhơn Bình Định', 'Đỗ Thanh Thịnh', 'LT01', '77'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'Công An Hà Nội' AND club_two_name = 'MerryLand Quy Nhơn Bình Định'), 'Công An Hà Nội', 'Hoàng Văn Toản', 'LT01', '81'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'TP Hồ Chí Minh' AND club_two_name = 'Khánh Hòa'), 'TP Hồ Chí Minh', 'Nguyễn Hạ Long', 'LT01', '5'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'TP Hồ Chí Minh' AND club_two_name = 'Khánh Hòa'), 'Khánh Hòa', 'Trần Mạnh Hùng', 'LT01', '23'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'TP Hồ Chí Minh' AND club_two_name = 'Khánh Hòa'), 'Khánh Hòa', 'Sesay Alie', 'LT01', '39'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'TP Hồ Chí Minh' AND club_two_name = 'Khánh Hòa'), 'Khánh Hòa', 'Guirassy Mamadou', 'LT01', '52'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'TP Hồ Chí Minh' AND club_two_name = 'Khánh Hòa'), 'TP Hồ Chí Minh', 'Nguyễn Minh Tùng', 'LT01', '72'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-22 19:15' AND club_one_name = 'TP Hồ Chí Minh' AND club_two_name = 'Khánh Hòa'), 'TP Hồ Chí Minh', 'Đào   Quốc Gia', 'LT01', '86'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-11-24 18:00' AND club_one_name = 'Becamex Bình Dương' AND club_two_name = 'Hà Nội'), 'Hà Nội', 'Vũ Tiến Long', 'LT02', '42'),
                          #Vong 2
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công - Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Đông Á Thanh Hóa', 'Nguyễn Thái Sơn', 'LT01', '6'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công - Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Đông Á Thanh Hóa', 'Rimario Allando Gordon', 'LT01', '21'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công - Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Thể Công – Viettel', 'Nguyễn Đức Chiến', 'LT01', '26'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công – Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Thể Công – Viettel', 'Phan Tuấn Tài', 'LT01', '45'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công – Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Thể Công – Viettel', 'Abdumuminov Jakhongir', 'LT01', '45'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công – Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Thể Công – Viettel', 'Bùi Tiến Dũng', 'LT01', '55'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công – Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Thể Công – Viettel', 'Bruno Cunha Cantanhede', 'LT01', '58'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công – Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Đông Á Thanh Hóa', 'Gustavo Sant Ana Santos', 'LT01', '58'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công – Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Thể Công – Viettel', 'Nguyễn Xuân Kiên', 'LT01', '90'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-27 19:15' AND club_one_name = 'Thể Công – Viettel' AND club_two_name = 'Đông Á Thanh Hóa'), 'Đông Á Thanh Hóa', 'Đoàn Ngọc Hà', 'LT02', '87'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 17:00' AND club_one_name = 'LPBank Hoàng Anh Gia Lai' AND club_two_name = 'Công An Hà Nội'), 'Công An Hà Nội', 'Phạm Văn Luân', 'LT01', '30'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 17:00' AND club_one_name = 'LPBank Hoàng Anh Gia Lai' AND club_two_name = 'Công An Hà Nội'), 'Công An Hà Nội', 'Hoàng Văn Toản', 'LT01', '61'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 17:00' AND club_one_name = 'LPBank Hoàng Anh Gia Lai' AND club_two_name = 'Công An Hà Nội'), 'LPBank Hoàng Anh Gia Lai', 'Diakite Pape Abdoulaye', 'LT01', '61'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 17:00' AND club_one_name = 'LPBank Hoàng Anh Gia Lai' AND club_two_name = 'Công An Hà Nội'), 'Công An Hà Nội', 'Nguyễn Xuân Nam', 'LT01', '87'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Khánh Hòa' AND club_two_name = 'Thép Xanh Nam Định'), 'Khánh Hòa', 'Nguyễn Đình Mạnh', 'LT01', '51'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Khánh Hòa' AND club_two_name = 'Thép Xanh Nam Định'), 'Thép Xanh Nam Định', 'Trần Văn Kiên', 'LT01', '73'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Khánh Hòa' AND club_two_name = 'Thép Xanh Nam Định'), 'Thép Xanh Nam Định', 'Hồ Khắc Ngọc', 'LT01', '78'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Hồng Lĩnh Hà Tĩnh' AND club_two_name = 'Sông Lam Nghệ An'), 'Sông Lam Nghệ An', 'Mai Sỹ Hoàng', 'LT01', '23'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Hồng Lĩnh Hà Tĩnh' AND club_two_name = 'Sông Lam Nghệ An'), 'Sông Lam Nghệ An', 'Lê Nguyên Hoàng', 'LT01', '40'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Hồng Lĩnh Hà Tĩnh' AND club_two_name = 'Sông Lam Nghệ An'), 'Sông Lam Nghệ An', 'Phan Bá Quyền', 'LT01', '46'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Hồng Lĩnh Hà Tĩnh' AND club_two_name = 'Sông Lam Nghệ An'), 'Sông Lam Nghệ An', 'Nguyễn Văn Việt', 'LT01', '81'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-28 18:00' AND club_one_name = 'Hồng Lĩnh Hà Tĩnh' AND club_two_name = 'Sông Lam Nghệ An'), 'Hồng Lĩnh Hà Tĩnh', 'Ngô Xuân Toàn', 'LT01', '90'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 17:00' AND club_one_name = 'Quảng Nam' AND club_two_name = 'TP Hồ Chí Minh'), 'Quảng Nam', 'Lê Xuân Tú', 'LT01', '18'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 17:00' AND club_one_name = 'Quảng Nam' AND club_two_name = 'TP Hồ Chí Minh'),'TP Hồ Chí Minh', 'Timite Cheick Aymar', 'LT01', '31'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 17:00' AND club_one_name = 'Quảng Nam' AND club_two_name = 'TP Hồ Chí Minh'), 'Quảng Nam', 'Mạch Ngọc Hà', 'LT01', '34'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 17:00' AND club_one_name = 'Quảng Nam' AND club_two_name = 'TP Hồ Chí Minh'), 'Quảng Nam', 'Eze Stephen', 'LT01', '66'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 17:00' AND club_one_name = 'Quảng Nam' AND club_two_name = 'TP Hồ Chí Minh'), 'TP Hồ Chí Minh', 'Da Silva Estevam Brendon Lucas', 'LT01', '41'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 17:00' AND club_one_name = 'Quảng Nam' AND club_two_name = 'TP Hồ Chí Minh'), 'TP Hồ Chí Minh', 'Nguyễn Thanh Thảo', 'LT01', '60'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 17:00' AND club_one_name = 'Quảng Nam' AND club_two_name = 'TP Hồ Chí Minh'), 'TP Hồ Chí Minh', 'Nguyễn Vũ Tín', 'LT01', '90'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 18:00' AND club_one_name = 'MerryLand Quy Nhơn Bình Định' AND club_two_name = 'Becamex Bình Dương'), 'MerryLand Quy Nhơn Bình Định', 'Ngô Hồng Phước', 'LT01', '62'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 18:00' AND club_one_name = 'MerryLand Quy Nhơn Bình Định' AND club_two_name = 'Becamex Bình Dương'), 'Becamex Bình Dương', 'Janclesio Almeida Santos', 'LT01', '64'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 18:00' AND club_one_name = 'MerryLand Quy Nhơn Bình Định' AND club_two_name = 'Becamex Bình Dương'), 'Becamex Bình Dương', 'Nguyễn Thành Lộc', 'LT01', '80'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 19:15' AND club_one_name = 'Hà Nội' AND club_two_name = 'Hải Phòng'), 'Hà Nội', 'Nguyễn Hai Long', 'LT01', '24'),
                          (SHA1(UUID()), NOW(), NOW(), NULL, (SELECT id FROM matches WHERE real_time = '2023-10-29 19:15' AND club_one_name = 'Hà Nội' AND club_two_name = 'Hải Phòng'), 'Hà Nội' ,'Trương Văn Thái Quý', 'LT01', '87');