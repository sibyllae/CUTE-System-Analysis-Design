-- 用户表（User）
INSERT INTO User (username, email, password)
VALUES
    ('user1', 'user1@example.com', 'password1'),
    ('user2', 'user2@example.com', 'password2'),
    ('user3', 'user3@example.com', 'password3'),
    ('user4', 'user4@example.com', 'password4'),
    ('user5', 'user5@example.com', 'password5');

-- 游戏表（Game）
INSERT INTO Game (title, description, release_date, game_type_id)
VALUES
    ('Game Title 1', 'Description for Game 1', '2023-01-01', 1),
    ('Game Title 2', 'Description for Game 2', '2023-02-01', 2),
    ('Game Title 3', 'Description for Game 3', '2023-03-01', 1),
    ('Game Title 4', 'Description for Game 4', '2023-04-01', 2),
    ('Game Title 5', 'Description for Game 5', '2023-05-01', 1);

-- 消费订单表（PurchaseOrder）
INSERT INTO PurchaseOrder (user_id, game_id, amount)
VALUES
    (1, 1, 19.99),
    (2, 2, 29.99),
    (3, 3, 14.99),
    (4, 4, 24.99),
    (5, 5, 9.99);

-- 登录记录表（LoginRecord）
INSERT INTO LoginRecord (user_id, login_time)
VALUES
    (1, '2023-01-01 10:00:00'),
    (2, '2023-01-02 11:00:00'),
    (3, '2023-01-03 12:00:00'),
    (4, '2023-01-04 13:00:00'),
    (5, '2023-01-05 14:00:00');

-- 游戏类型表（GameType）
INSERT INTO GameType (type_name)
VALUES
    ('Action'),
    ('Adventure'),
    ('Puzzle'),
    ('Strategy'),
    ('Simulation');
