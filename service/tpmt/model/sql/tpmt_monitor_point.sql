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

 Date: 12/10/2024 14:21:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tpmt_monitor_point
-- ----------------------------
DROP TABLE IF EXISTS `tpmt_monitor_point`;
CREATE TABLE `tpmt_monitor_point`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '监测点ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `serial_number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '编号',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '监测点名称',
  `register_address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '寄存器地址',
  `point_collector_instruction` int(0) NOT NULL DEFAULT 2 COMMENT '采集器指令  1: 01  2: 02  3:03  4:04',
  `point_analysis_rule` int(0) NOT NULL DEFAULT 1 COMMENT '采集器解析规则 1: 16位无符号/2:单精度浮点数',
  `point_type` int(0) NOT NULL DEFAULT 1 COMMENT '类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他',
  `point_category` int(0) NOT NULL DEFAULT 1 COMMENT '类别：1:遥信/2:遥测/3:遥脉',
  `point_group` int(0) NOT NULL DEFAULT 1 COMMENT '分组',
  `circuit_type` int(0) NOT NULL DEFAULT 1 COMMENT '回路类型',
  `yx_decode` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '遥信解译',
  `data_bits` int(0) NOT NULL DEFAULT 2 COMMENT '数据位',
  `coefficient` double(20, 10) NOT NULL DEFAULT 1.0000000000 COMMENT '系数',
  `retain_decimals` int(0) NOT NULL DEFAULT 2 COMMENT '保留小数位',
  `unit` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单位',
  `alarm_duration` bigint(0) NOT NULL DEFAULT 20 COMMENT '持续时间',
  `alarm_up_value` double NOT NULL DEFAULT 0 COMMENT '告警上限',
  `alarm_down_value` double NOT NULL DEFAULT 0 COMMENT '告警下限',
  `warning_up_value` double NOT NULL DEFAULT 0 COMMENT '预警上限',
  `warning_down_value` double NOT NULL DEFAULT 0 COMMENT '预警下限',
  `is_displacement_warning` tinyint(0) NOT NULL DEFAULT 0 COMMENT '变位预警 0 不启用 1:启用',
  `tpmt_gateway_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '网关ID',
  `asset_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资产ID',
  `sort` int(0) NOT NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
