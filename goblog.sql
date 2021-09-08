/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : goblog

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 08/09/2021 12:12:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blogs
-- ----------------------------
DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `context` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `update_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pre_id` int NULL DEFAULT NULL,
  `next_id` int NULL DEFAULT NULL,
  `like_count` int NULL DEFAULT NULL,
  `view_count` int NULL DEFAULT NULL,
  `pircture` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `comment_count` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blogs
-- ----------------------------
INSERT INTO `blogs` VALUES (5, 'admin', '测试1', '## 测试\n![](/static/upload/image/5.jpg)', 'admin', '原创', '2021-08-24 09:39:35', '2021-08-24 09:39:35', 0, 0, 103, 13, '../../static/images/backimg1.jpg', '测试1', 8);
INSERT INTO `blogs` VALUES (8, 'admin', '测试2', '# 测试2', 'admin', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 7, 13, '', '测试2', 0);
INSERT INTO `blogs` VALUES (9, 'admin', '测试3', '# 测试3', 'admin', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 52, 0, '../../static/images/backimg1.jpg', '测试3', 0);
INSERT INTO `blogs` VALUES (10, 'admin', '测试3', '# 测试2', 'admin', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 3, 13, '', '测试2', 0);
INSERT INTO `blogs` VALUES (11, 'admin', '测试4', '# 测试2', 'admin', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 1003, 1, '../../static/images/backimg1.jpg', '测试2', 0);
INSERT INTO `blogs` VALUES (12, 'admin', '测试5', '# 测试2', 'user', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 3, 3, '', '测试2', 1);
INSERT INTO `blogs` VALUES (13, 'admin', '测试6', '# 测试2', 'admin', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 11, 5, '../../static/upload/background/3b.jpg', '测试2', 0);
INSERT INTO `blogs` VALUES (14, 'admin', '测试7', '# 测试2', 'admin', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 9, 13, '', '测试2', 0);
INSERT INTO `blogs` VALUES (15, 'admin', '测试8', '# 测试2', 'user', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 3, 7, '', '测试2', 1);
INSERT INTO `blogs` VALUES (16, 'admin', '测试9', '# 测试2', 'user', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 2, 0, '', '测试2', 0);
INSERT INTO `blogs` VALUES (17, 'admin', '测试10', '# 测试2', 'admin', '原创', '2021-08-24 09:41:29', '2021-08-24 09:41:29', 0, 0, 3, 5, '', '测试2', 1);
INSERT INTO `blogs` VALUES (18, 'admin', '详情测试', '# 详情测试\n![](/static/upload/image/5.jpg)', 'admin', '原创', '2021-08-25 10:40:51', '2021-08-25 10:40:51', 0, 0, 2, 10, '', '详情测试', 0);
INSERT INTO `blogs` VALUES (19, 'admin', '第一篇', '#yes\n![](/static/upload/image/315afa6f8f1247dfe019e9cae0884b2c_gamersky_08origin_15_2019391555F8C.jpg)', 'jack', '原创', '2021-08-25 16:51:42', '2021-08-25 16:51:42', 0, 0, 2, 18, '../../static/upload/background/28.jpg', '测试下', 0);

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `blog_id` int NULL DEFAULT NULL,
  `parent_comment_id` int NULL DEFAULT NULL,
  `admin` tinyint(1) NULL DEFAULT NULL,
  `create_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, 'user1', '测试', -1, -1, 0, '2021-09-01 16:36:34');
INSERT INTO `comments` VALUES (3, 'admin', '测试3', -1, -1, 1, '2021-09-01 17:02:18');
INSERT INTO `comments` VALUES (4, 'admin', '测试4', -1, -1, 1, '2021-09-01 17:02:25');
INSERT INTO `comments` VALUES (5, 'admin', '啦啦啦', -1, -1, 1, '2021-09-01 17:02:29');
INSERT INTO `comments` VALUES (6, 'admin', 'gogogo', -1, -1, 1, '2021-09-01 17:02:34');
INSERT INTO `comments` VALUES (9, 'admin', '测试234', -1, -1, 1, '2021-09-01 17:20:37');
INSERT INTO `comments` VALUES (11, 'admin', '测试回复', -1, 9, 1, '2021-09-02 11:27:52');
INSERT INTO `comments` VALUES (12, 'admin', '测试回复2', -1, 11, 1, '2021-09-02 11:30:47');
INSERT INTO `comments` VALUES (13, 'admin', '测试回复3', -1, 12, 1, '2021-09-02 11:32:23');
INSERT INTO `comments` VALUES (14, 'user1', 'yes', -1, 5, 0, '2021-09-02 14:04:36');
INSERT INTO `comments` VALUES (17, 'user1', '回复6', -1, 6, 0, '2021-09-02 14:38:42');
INSERT INTO `comments` VALUES (18, 'user1', '测试回复回复', -1, 11, 0, '2021-09-02 14:40:39');
INSERT INTO `comments` VALUES (19, 'user1', '测试234回复', -1, 9, 0, '2021-09-02 14:41:07');
INSERT INTO `comments` VALUES (20, 'user1', '回复gogoggo', -1, 6, 0, '2021-09-02 14:41:42');
INSERT INTO `comments` VALUES (21, 'admin', '人生就是不公平的,你慢慢习惯吧你', -1, -1, 1, '2021-09-02 15:34:45');
INSERT INTO `comments` VALUES (22, 'admin', 'asdsafsdfds', 5, -1, 1, '2021-09-02 16:34:38');
INSERT INTO `comments` VALUES (23, 'admin', 'googog', 5, -1, 1, '2021-09-02 16:35:56');
INSERT INTO `comments` VALUES (24, 'admin', '回复gogogo', 5, 23, 1, '2021-09-02 16:36:15');
INSERT INTO `comments` VALUES (25, 'admin', 'yes', 5, -1, 1, '2021-09-07 11:24:46');
INSERT INTO `comments` VALUES (26, 'admin', '不错', 5, -1, 1, '2021-09-07 11:25:07');
INSERT INTO `comments` VALUES (27, 'admin', 'yess', 5, 24, 1, '2021-09-07 11:25:57');
INSERT INTO `comments` VALUES (28, 'admin', 'gu', 17, -1, 1, '2021-09-07 11:27:13');
INSERT INTO `comments` VALUES (29, 'admin', 'sd', 5, -1, 1, '2021-09-07 11:27:42');
INSERT INTO `comments` VALUES (30, 'admin', 's', 5, -1, 1, '2021-09-07 11:28:57');
INSERT INTO `comments` VALUES (31, 'admin', 'test', 12, -1, 1, '2021-09-07 14:38:21');
INSERT INTO `comments` VALUES (32, 'user1', 's', 15, -1, 0, '2021-09-07 15:31:32');

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `blog_id` int NULL DEFAULT NULL,
  `user_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of likes
-- ----------------------------
INSERT INTO `likes` VALUES (1, 19, '7,8');
INSERT INTO `likes` VALUES (2, 8, '7,8');
INSERT INTO `likes` VALUES (3, 13, '7,8');
INSERT INTO `likes` VALUES (4, 14, '7,8');
INSERT INTO `likes` VALUES (5, 10, '8');
INSERT INTO `likes` VALUES (6, 15, '8');
INSERT INTO `likes` VALUES (7, 5, '7');
INSERT INTO `likes` VALUES (8, 17, '7');
INSERT INTO `likes` VALUES (9, 12, '7');

-- ----------------------------
-- Table structure for links
-- ----------------------------
DROP TABLE IF EXISTS `links`;
CREATE TABLE `links`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `update_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of links
-- ----------------------------
INSERT INTO `links` VALUES (3, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (4, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (5, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (6, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (7, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (9, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (10, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (11, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (12, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (13, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (15, '测试', '内网', '第一张图', '2021-08-30 21:39:07', '2021-08-30 21:39:07');
INSERT INTO `links` VALUES (17, '测试', '2', '2', '2021-08-30 21:42:52', '2021-08-30 21:42:52');

-- ----------------------------
-- Table structure for pictures
-- ----------------------------
DROP TABLE IF EXISTS `pictures`;
CREATE TABLE `pictures`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `time_place` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pictures
-- ----------------------------
INSERT INTO `pictures` VALUES (1, '测试1', '2021年8月31号  武汉', 'static/images/me.jpg', '测试1');
INSERT INTO `pictures` VALUES (2, '测试2', '2021年8月31号  武汉', 'static/images/me.jpg', '测试2');
INSERT INTO `pictures` VALUES (3, '测试3', '2021年8月31号  武汉', 'static/images/me.jpg', '测试3');
INSERT INTO `pictures` VALUES (4, '测试4', '2021年8月31号  武汉', 'static/images/me.jpg', '测试4');
INSERT INTO `pictures` VALUES (5, '测试5', '2021年8月31号  武汉', 'static/images/me.jpg', '测试5');

-- ----------------------------
-- Table structure for types
-- ----------------------------
DROP TABLE IF EXISTS `types`;
CREATE TABLE `types`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of types
-- ----------------------------
INSERT INTO `types` VALUES (1, 'admin');
INSERT INTO `types` VALUES (3, 'user');
INSERT INTO `types` VALUES (9, 'test');
INSERT INTO `types` VALUES (10, 'jack');
INSERT INTO `types` VALUES (11, '测试');
INSERT INTO `types` VALUES (12, 'admin2');
INSERT INTO `types` VALUES (13, 'user3');
INSERT INTO `types` VALUES (14, 'type2');
INSERT INTO `types` VALUES (15, 'user5');
INSERT INTO `types` VALUES (16, 'yun2');
INSERT INTO `types` VALUES (17, 'num2');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `recent_login` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_date` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `role` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (7, 'admin', 'hp9VjJcWPvvji6hlMqJi4Q==', '2021-09-08 12:01:34', '2021-07-16 14:06:53', 'admin');
INSERT INTO `users` VALUES (8, 'user1', 'hp9VjJcWPvvji6hlMqJi4Q==', '2021-09-07 16:31:09', '2021-08-31 15:48:26', 'user');

SET FOREIGN_KEY_CHECKS = 1;
