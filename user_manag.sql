/*
 Navicat Premium Data Transfer

 Source Server         : Bootcamp_MySql
 Source Server Type    : MySQL
 Source Server Version : 100411
 Source Host           : localhost:3306
 Source Schema         : user_manag

 Target Server Type    : MySQL
 Target Server Version : 100411
 File Encoding         : 65001

 Date: 06/05/2021 22:48:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'testB', 'Bismillah', 'testB');
INSERT INTO `users` VALUES (3, 'ridwanfebrianas', 'Bismillah', 'ridwanfebrianas');
INSERT INTO `users` VALUES (4, 'johnmayer', 'Bismillah', 'ridwan');
INSERT INTO `users` VALUES (5, 'john', '1c1b72aee1bf62759f34d254e0fb8ce2dafb480d42bc5ab1f5af0bdb6e53d108', 'johnmayer');
INSERT INTO `users` VALUES (6, 'sd', 'f718af9c9f6673b5cbea714527ec29a737dee4ffd97634f94a3f1828564aa3c3', 'as');
INSERT INTO `users` VALUES (7, 'aaa', '1982ed786929c7ddfa43cd88909eddfcf1fb8746b1b3aab346af3b7e1ecc274d', 'as');
INSERT INTO `users` VALUES (8, 'aaaa', '1982ed786929c7ddfa43cd88909eddfcf1fb8746b1b3aab346af3b7e1ecc274d', 'asa');

SET FOREIGN_KEY_CHECKS = 1;
