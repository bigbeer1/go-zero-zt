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

 Date: 14/11/2024 13:34:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tpmt_scheduled_tasks
-- ----------------------------
DROP TABLE IF EXISTS `tpmt_scheduled_tasks`;
CREATE TABLE `tpmt_scheduled_tasks`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '定时任务ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `scheduler_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `scheduler_category` tinyint(0) NOT NULL COMMENT '类别 1:已接入任务, 2:自定义任务',
  `scheduler_task_number` tinyint(0) NOT NULL COMMENT '已接入任务号',
  `scheduler_type` tinyint(0) NOT NULL COMMENT '类型 1:Http任务,2:Webservices任务',
  `interval_time` int(0) NOT NULL COMMENT '间隔时间按秒',
  `error_order` int(0) NOT NULL DEFAULT 1 COMMENT '失败重新发送次数1-10次 不可超过10次',
  `fail_interval_time` int(0) NOT NULL COMMENT '失败间隔时间按秒',
  `state` tinyint(0) NOT NULL DEFAULT 1 COMMENT '状态 1:启动  2:暂停',
  `scheduler_data` varbinary(30000) NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
