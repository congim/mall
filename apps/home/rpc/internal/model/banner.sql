CREATE DATABASE `test-mall`;
USE `test-mall`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
-- `test-mall`.banner definition

CREATE TABLE `banner` (
      `id` int NOT NULL AUTO_INCREMENT COMMENT '活动id',
      `name` varchar(255) NOT NULL COMMENT '活动名',
      `image` varchar(255) NOT NULL COMMENT '活动图片',
      `url` varchar(255) NOT NULL COMMENT '活动url',
      `state` varchar(100) NOT NULL COMMENT '活动状态: 1 使用中， 2 未使用',
      PRIMARY KEY (`id`),
      KEY `banner_state_IDX` (`state`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;