/*
 Navicat Premium Data Transfer

 Source Server         : mysql54
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 10.132.105.54:33069
 Source Schema         : tpmt_tw

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 12/10/2024 14:20:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tpmt_gateway
-- ----------------------------
DROP TABLE IF EXISTS `tpmt_gateway`;
CREATE TABLE `tpmt_gateway`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采集器ID/网关',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `gateway_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '网关名称',
  `gateway_model` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '网关型号',
  `manu_facturer` varchar(100) CHARACTER SET utf16 COLLATE utf16_general_ci NOT NULL COMMENT '生产厂家',
  `agreement` int(0) NOT NULL DEFAULT 1 COMMENT '协议 默认1:modbus',
  `baud_rate` int(0) NOT NULL DEFAULT 1 COMMENT '波特率',
  `parity` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '2' COMMENT '校验',
  `data_bits` int(0) NOT NULL DEFAULT 8 COMMENT '数据位',
  `stop_bits` int(0) NOT NULL DEFAULT 1 COMMENT '停止位',
  `com_port` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT 'com端口',
  `address_code` int(0) NOT NULL DEFAULT 1 COMMENT '地址码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
