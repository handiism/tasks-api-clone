CREATE DATABASE tasks_clone;

USE tasks_clone;

CREATE TABLE `user`
(
    `id`         bigint(20) unsigned                    NOT NULL AUTO_INCREMENT,
    `name`       varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `email`      varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password`   varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` datetime(3)                            NOT NULL DEFAULT current_timestamp(3),
    `updated_at` datetime(3)                                     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `list`
(
    `id`         bigint(20) unsigned                    NOT NULL AUTO_INCREMENT,
    `user_id`    bigint(20) unsigned                    NOT NULL,
    `title`      varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` datetime(3)                            NOT NULL DEFAULT current_timestamp(3),
    `updated_at` datetime(3)                                     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_user_lists` (`user_id`),
    CONSTRAINT `fk_user_lists` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `task`
(
    `id`         bigint(20) unsigned                    NOT NULL AUTO_INCREMENT,
    `list_id`    bigint(20) unsigned                    NOT NULL,
    `name`       varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `detail`     varchar(100) COLLATE utf8mb4_unicode_ci         DEFAULT NULL,
    `due_date`   datetime(3)                                     DEFAULT NULL,
    `is_done`    tinyint(1)                             NOT NULL DEFAULT 0,
    `created_at` datetime(3)                            NOT NULL DEFAULT current_timestamp(3),
    `updated_at` datetime(3)                                     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_list_tasks` (`list_id`),
    CONSTRAINT `fk_list_tasks` FOREIGN KEY (`list_id`) REFERENCES `list` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `subtask`
(
    `id`         bigint(20) unsigned                    NOT NULL AUTO_INCREMENT,
    `task_id`    bigint(20) unsigned                    NOT NULL,
    `name`       varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `is_done`    tinyint(1)                             NOT NULL DEFAULT 0,
    `created_at` datetime(3)                            NOT NULL DEFAULT current_timestamp(3),
    `updated_at` datetime(3)                                     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_task_subtasks` (`task_id`),
    CONSTRAINT `fk_task_subtasks` FOREIGN KEY (`task_id`) REFERENCES `task` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;