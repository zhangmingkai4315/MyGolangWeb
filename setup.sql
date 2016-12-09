SHOW DATABASES;
DROP DATABASE IF EXISTS `golangchina`;
CREATE DATABASE IF NOT EXISTS `golangchina`  CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `golangchina`;
-- SHOW TABLES;

-- Enable client program to communicate with the server using utf8 character set
SET NAMES 'utf8';

DROP TABLE IF EXISTS `users`;
create table IF NOT EXISTS `users`(
    `id`            INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_name`     varchar(48) not null,
    `user_guid`     varchar(256) not null default '',
    `user_email`    varchar(128) not null,
    `user_password` varchar(128) not null,
    `user_salt`     varchar(128) not null default '',
    `user_joined_timestamp`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(`user_email`),
    UNIQUE(`user_name`),
    PRIMARY KEY (`id`)
)DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

EXPLAIN `users`;

DROP TABLE IF EXISTS `usersinfo`;
create table IF NOT EXISTS `usersinfo`(
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT(10) UNSIGNED,
    `user_phone`       varchar(100),
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `big_photo`   varchar(255) default "/static/images/big-header-default.jpg",
    `small_photo` varchar(255) default "/static/images/small-header-default.png",
    `status`    ENUM('Active', 'Deleted','Nonactive','Banned') DEFAULT 'Active',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
)DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

EXPLAIN `usersinfo`;



DROP TABLE IF EXISTS `sessions`;
create table IF NOT EXISTS `sessions` (
    `id`          INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `session_id`  varchar(255) not null default '',
    `user_id`     INT(10) UNSIGNED,
    `session_start` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `session_update` TIMESTAMP Not Null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `session_active` ENUM('Active','Nonactive') DEFAULT 'Nonactive',
    PRIMARY KEY (`id`),
    UNIQUE KEY `session_id` (`session_id`),
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




