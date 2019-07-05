/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : zadmin

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-07-05 17:16:08
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES ('1', '123', '123');

-- ----------------------------
-- Table structure for configuration
-- ----------------------------
DROP TABLE IF EXISTS `configuration`;
CREATE TABLE `configuration` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) NOT NULL COMMENT '配置键',
  `value` varchar(255) NOT NULL COMMENT '配置值',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `mark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`,`key`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of configuration
-- ----------------------------
INSERT INTO `configuration` VALUES ('4', '567', '5467', '4567', '4567');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `mobile` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', '1', '1', '1');
INSERT INTO `users` VALUES ('2', '2', '2', '2');
INSERT INTO `users` VALUES ('3', '3', '3', '3');
