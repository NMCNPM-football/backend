# CREATE DATABASE SE104;
USE SE104;

-- Create information of clubs name Hoang Anh Gia Lai
INSERT INTO `players` (`id`, `club_id`,`club_name`,`sea_son`,`name`,`birth_day`, `position`, `height`, `weight`, `nationality`,`kit`)  VALUES
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Dương Văn Lợi', '2000-12-02', 'GK', '177', '73', 'VIE', '1'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Lê Văn Sơn', '1996-12-20', 'DF', '169', '64', 'VIE', '2'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Trần Hữu Đông Triều', '1995-08-20', 'DF', '170', '70', 'VIE', '5'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Lương Xuân Trường', '1995-04-28', 'MF', '178', '72', 'VIE', '6'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Phong Hồng Duy', '1996-06-13', 'DF', '169', '66', 'VIE', 7),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Trần Minh Vương', '1995-03-28', 'MF', '167', '67', 'VIE', '8'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Văn Toàn', '1996-04-12', 'ST', '170', '58', 'VIE', '9'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Công Phượng', '1995-01-21', 'ST', '168', '65', 'VIE', '10'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Tuấn Anh', '1995-05-16', 'MF', '175', '66', 'VIE', '11'),
    (SHA1(UUID()), (SELECT clubs.`id`FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Tiêu Ê Xal', '2000-08-14', 'DF', '180', '74', 'VIE', '12'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Hữu Tuấn', '1992-05-06', 'DF', '178', '70', 'VIE', '15'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Vũ Văn Thanh', '1996-04-14', 'DF', '172', '73', 'VIE', '17'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Trần Bảo Toàn', '2000-07-14', 'ST', '169', '63', 'VIE', '20'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Lê Huy Kiệt', '2003-10-20', 'MF', '168', '65', 'VIE', '21'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Nhĩ Khang', '2001-02-16', 'ST', '169', '63', 'VIE', '22'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Thanh Nhân', '2000-10-25', 'DF', '168', '65', 'VIE', '23'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Đức Việt', '2004-01-01', 'MF', '174', '70', 'VIE', '24'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Trần Trung Kiên', '2003-02-09', 'GK', '191', '82', 'VIE', '25'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Huỳnh Tuấn Linh', '1991-04-12', 'GK', '180', '79', 'VIE', '26'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Văn Triệu', '2003-01-17', 'MF', '188', '73', 'VIE', '27'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Nguyễn Văn Việt', '1989-01-17', 'DF', '177', '75', 'VIE', '28'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Washington Brandão Dos Santos', '1990-08-18', 'ST', '183', '83', 'BRA', '30'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Lê Hữu Phước', '2001-05-07', 'MF', '174', '62', 'VIE', '34'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'An Sae Hee', '1991-02-08', 'MF', '186', '80', 'KR', '40'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Cao Hoàng Tú', '2001-06-09', 'ST', '172', '60', 'VIE', '47'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Võ Đình Lâm', '2000-01-10', 'MF', '169', '68', 'VIE', '60'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'A Hoàng', '1995-07-31', 'MF', '174', '67', 'VIE', '82'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Bruno Henrique', '1992-10-25', 'ST', '185', '81', 'BRA', '92'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'LPBHA' AND `sea_son` = '2023-2024'), 'LPBank Hoàng Anh Gia Lai', '2023-2024', 'Barbosa Teixeira Mauricio', '1994-12-08', 'DF', '198', '94', 'BRA', '94');

-- Create information of clubs name Becamex Bình Dương
INSERT INTO `players` (`id`, `club_id`,`club_name`,`sea_son`,`name`,`birth_day`, `position`, `height`, `weight`, `nationality`,`kit`)
VALUES
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Sơn Hải', '1994-05-06', 'GK', '178', '86', 'VIE', '1'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Hùng Thiện Đức', '1999-12-08', 'DF', '169', '66', 'VIE', '2'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Thanh Thảo', '1995-05-14', 'DF', '178', '73', 'VIE', '3'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Ndiaye Guy Olivier', '1994-05-05', 'DF', '186', '88', 'VIE', '4'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Uông Ngọc Tiến', '2000-03-23', 'MF', '166', '63', 'VIE', '5'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Trọng Huy', '1997-06-25', 'MF', '180', '74', 'VIE', '6'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Thanh Long', '1993-01-10', 'DF', '180', '79', 'VIE', '7'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Anh Tài', '1996-02-05', 'DF', '168', '70', 'VIE', '8'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Đoàn Tuấn Cảnh', '1998-07-27', 'DF', '176', '69', 'VIE', '9'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Teofilo Soares Eydison', '1988-05-30', 'ST', '178', '82', 'BRA', '10'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Bùi Vĩ Hào', '2003-02-24', 'ST', '175', '68', 'VIE', '11'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Trần Duy Khánh', '1997-07-20', 'MF', '172', '77', 'VIE', '12'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Huỳnh Văn Bin', '2002-01-12', 'GK', '178', '69,' ,'VIE', '13'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Trần Hoàng Phương', '1994-04-11', 'MF', '168', '63', 'VIE', '14'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Trương Dũ Đạt', '1997-07-25', 'DF', '180', '75', 'VIE', '15'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Tống Anh Tỷ', '1997-01-24', 'MF', '173', '73', 'VIE', '17'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Hồ Sỹ Giáp', '1994-04-18', 'ST', '176', '67', 'VIE', '18'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Lưu Tự Nhân', '2001-12-18', 'MF', '180', '73', 'VIE', '19'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Đoàn Văn Quý', '1998-01-02', 'DF', '171', '64', 'VIE', '20'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Tiến Linh', '1997-10-20', 'ST', '181', '75', 'VIE', '22'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Hà Trung Hậu', '2002-10-03', 'DF', '177', '80', 'VIE', '23'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Trần Hoàng Bảo', '2001-05-15', 'MF', '176', '58', 'VIE', '24'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Trần Đức Cường', '1985-05-20', 'GK', '183', '75', 'VIE', '25'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Lê Vũ Quốc Nhật', '1996-10-16', 'MF', '169', '62', 'VIE', '26'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Trung Tín', '1991-09-14', 'DF', '178', '75', 'VIE', '27'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Tô Văn Vũ', '1993-10-20', 'MF', '175', '60', 'VIE', '28'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Võ Hoàng Minh Khoa', '2001-03-12', 'MF', '175', '70', 'VIE', '29'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Lại Tuấn Vũ', '1993-02-03', 'GK', '183', '76', 'VIE', '30'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Thành Kiên', '2003-01-16', 'DF', '182', '74', 'VIE', '31'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Lê Văn Đại', '1996-08-02', 'DF', '182', '78', 'VIE', '33'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Nguyễn Đoàn Trung Nhân', '1998-03-05', 'MF', '170', '63', 'VIE', '34'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Adão Cruz Welington', NULL, 'ST', '180', '74', 'BRA', '37'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Phan Minh Thành', '1998-11-22', 'GK', '192', '80', 'VIE', '46'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'BDFC' AND `sea_son` = '2023-2024'), 'Becamex Bình Dương', '2023-2024', 'Huỳnh Kesley Alves', '1981-12-23', 'ST', '175', '70', 'VIE', '77');

-- Create information of clubs name Đông Á Thanh Hóa
INSERT INTO `players` (`id`, `club_id`,`club_name`,`sea_son`,`name`,`birth_day`, `position`, `height`, `weight`, `nationality`,`kit`)
VALUES
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Hoàng Đình Tùng', '1988-08-24', 'MF', '168', '62', 'VIE', '2'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Vũ Xuân Cường', '1992-08-06', 'DF', '170', '65', 'VIE', '3'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Đàm Tiến Dũng', '1996-01-10', 'DF', '180', '79', 'VIE', '4'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-20241', 'Nguyễn Minh Tùng', '1992-08-09', 'MF', '185', '75', 'VIE', '5'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Nguyễn Hữu Dũng', '1995-08-28', 'MF', '171', '64', 'VIE', '7'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Jose Paulo De Oliveira Pinto', '1994-03-26', 'ST', '189', '82', 'BRA', '8'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Lê Xuân Hùng', '1991-11-14', 'ST', '173', '69', 'VIE', '9'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Lê Phạm Thành Long', '1996-06-05', 'MF', '165', '64', 'VIE', '11'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Nguyễn Thái Sơn', '2003-07-13', 'MF', '171', '61', 'VIE', '12'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Trịnh Văn Lợi', '1995-05-26', 'DF', '180', '71', 'VIE', '15'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Đinh Tiến Thành', '1991-01-24', 'MF', '183', '71', 'VIE', '16'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Nguyễn Văn Tiếp', '2003-11-21', 'MF', '174', '64', 'VIE', '18'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Lê Quốc Phương', '1991-05-19', 'MF', '165', '61', 'VIE', '19'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Nguyễn Trọng Hùng', '1997-10-03', 'MF', '172', '68', 'VIE', '20'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Nguyễn Đình Huyên', '2001-08-12', 'DF', '178', '70', 'VIE', '21'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Trịnh Xuân Hoàng', '2000-11-06', 'GK', '184', '78', 'VIE', '23'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Nguyễn Thanh Diệp', '1991-09-06', 'GK', '180', '75', 'VIE', '25'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Trần Bữu Ngọc', '1991-02-26', 'GK', '190', '94', 'VIE', '26'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'A Mít', '1997-07-24', 'MF', '168', '67', 'VIE', '27'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Đoàn Ngọc Hà', '2004-02-22', 'DF', '182', '72', 'VIE', '29'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Lê Văn Cường', '2003-06-17', 'ST', '178', '69', 'VIE', '31'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Lê Ngọc Nam', '1993-02-26', 'MF', '167', '66', 'VIE', '32'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Doãn Ngọc Tân', '1994-08-15', 'DF', '169', '57', 'VIE', '34'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Lục Xuân Hưng', '1995-04-15', 'DF', '181', '75', 'VIE', '89'),
    (SHA1(UUID()), (SELECT clubs.`id` FROM `clubs` WHERE `shorthand` = 'THFC' AND `sea_son` = '2023-2024'), 'Đông Á Thanh Hóa', '2023-2024', 'Gustavo Sant Ana Santos', '1995-02-23', 'CB', '195', '80', 'BRA', '95');