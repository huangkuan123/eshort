
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for eshorts
-- ----------------------------
DROP TABLE IF EXISTS `eshorts`;
CREATE TABLE `eshorts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `short_key` char(10) NOT NULL DEFAULT '',
  `full_data` varchar(255) NOT NULL DEFAULT '' COMMENT '长链接',
  `ext` varchar(50) NOT NULL,
  `exp` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1启用，5过期，10拉黑',
  `is_deleted` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `short_key` (`short_key`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of eshorts
-- ----------------------------
