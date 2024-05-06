USE SE104;

INSERT INTO `users` (`id`, `name`, `email`, `password` ,`is_verified_email`) VALUES
(SHA1(UUID()), 'admin', 'admin@gmail.com','admin','1');

INSERT INTO matches
(`id`, `sea_son`, `club_one_name`, `id_club_one`, `club_two_name`, `id_club_two`, `intend_time`, `real_time`, `match_round`, `match_turn`, `stadium`)
VALUES
    (
        SHA1(UUID()),
        '2023-2024',
        'Becamex Bình Dương',
        (SELECT id FROM clubs WHERE name_club = 'Becamex Bình Dương' AND sea_son = match_calendars.sea_son),
        'Công An Hà Nội',
        (SELECT id FROM clubs WHERE name_club = 'Công An Hà Nội' AND sea_son = match_calendars.sea_son),
        '2023-10-20 18:00',
        '2023-10-20 18:00',
        '1',
        'Đi',
        (SELECT stadium_name FROM stadia WHERE club_id = match_calendars.id_club_one)
    );