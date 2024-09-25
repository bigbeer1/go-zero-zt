/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : tpmt

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 25/09/2024 14:55:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_role_interface
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_interface`;
CREATE TABLE `sys_role_interface`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` int(0) NOT NULL COMMENT '角色ID',
  `interface_id` int(0) NOT NULL COMMENT '接口ID',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代理商表-角色与菜单中间表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
