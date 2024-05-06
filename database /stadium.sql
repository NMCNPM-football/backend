USE SE104;

INSERT INTO `stadia` (id, created_at, updated_at, deleted_at, club_id, stadium_name, stadium_address, season, capacity)
VALUES
    (
        SHA1(UUID()),
        NOW(),
        NOW(),
        NULL,
        (SELECT `id` FROM `clubs` WHERE `name_club` = 'MerryLand Quy Nhơn Bình Định' AND `sea_son` = season LIMIT 1),
        'Old Trafford',
        'Sir Matt Busby Way, Old Trafford, Manchester M16 0RA, United Kingdom',
        '2023-2024',
        '74879'
    );