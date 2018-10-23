-- 创建数据库
create database `test` default character set utf8 collate utf8_general_ci;

use test;

-- 建表
DROP TABLE IF EXISTS `test`;

CREATE TABLE `member` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- 插入数据
INSERT INTO `member` (`id`, `name`)
VALUES
    (1,'user');