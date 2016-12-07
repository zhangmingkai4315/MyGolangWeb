SHOW DATABASES;
DROP DATABASE IF EXISTS `golangchina`;
CREATE DATABASE IF NOT EXISTS `golangchina`  CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `golangchina`;
-- SHOW TABLES;

-- Enable client program to communicate with the server using utf8 character set
SET NAMES 'utf8';

DROP TABLE IF EXISTS `users`;
create table IF NOT EXISTS `users`(
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        varchar(255),
    `email`       varchar(255) not null,
    `password`    varchar(255) not null,
    `created_at`  DATE not null,
    UNIQUE(`email`),
    PRIMARY KEY (`id`)
)DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

EXPLAIN `users`;

DROP TABLE IF EXISTS `usersinfo`;
create table IF NOT EXISTS `usersinfo`(
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT(10) UNSIGNED,
    `phone`       varchar(100),
    `created_at`  DATE not null,
    `big_photo`   varchar(255) default "/static/images/big-header-default.jpg",
    `small_photo` varchar(255) default "/static/images/small-header-default.png",
    `status`    ENUM('active', 'deleted','nonactive','banned') DEFAULT 'active',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
)DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

EXPLAIN `usersinfo`;



DROP TABLE IF EXISTS `sessions`;
create table IF NOT EXISTS `sessions` (
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `email`       varchar(255),
    `uuid`        varchar(255) not null,
    `user_id`     INT(10) UNSIGNED,
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
)DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EXPLAIN `sessions`;

DROP TABLE IF EXISTS `threads`;
create table threads (
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title`       varchar(200) not null,
    `content`     text,
    `type`        varchar(100),
    `uuid`        varchar(255) not null,
    `user_id`     INT(10) UNSIGNED,
    `status`    ENUM('passed', 'pending','deleted') DEFAULT 'passed',
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `modification_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
)DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EXPLAIN `threads`;



DROP TABLE IF EXISTS `comments`;
create table comments (
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `body`        text,
    `thread_id`   INT(10) UNSIGNED,
    `user_id`     INT(10) UNSIGNED,
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`),
    FOREIGN KEY (`thread_id`) REFERENCES threads(`id`)
)DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EXPLAIN `comments`;




