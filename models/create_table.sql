DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           bigint(20)                             NOT NULL AUTO_INCREMENT,
    `user_id`      bigint(20)                             NOT NULL,
    `username`     varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`     varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email`        varchar(64) COLLATE utf8mb4_general_ci,
    `gender`       tinyint(4)                             NOT NULL DEFAULT '0',
    `created_at` datetime                               NULL     DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime                               NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime                               NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;


