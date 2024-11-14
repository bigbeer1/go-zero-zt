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

 Date: 14/11/2024 17:00:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_auth
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth`;
CREATE TABLE `sys_auth`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '第三方用户ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  `nick_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '机构名',
  `auth_token` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '令牌',
  `state` tinyint(0) NOT NULL DEFAULT 0 COMMENT '状态 1:正常 2:停用 3:封禁',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auth
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '字典ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典类型',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典标签',
  `dict_value` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典键值',
  `sort` bigint(0) NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 271 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
INSERT INTO `sys_dict` VALUES (269, '2024-10-17 15:29:39', NULL, NULL, '超级管理员', NULL, NULL, 'com_port', 'tty/dev01', 'tty/dev01', 55, '1', 1);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '字典类型ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典名称',
  `dict_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典类型',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `sort` bigint(0) NOT NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (64, '2024-10-17 15:28:43', NULL, NULL, '超级管理员', NULL, NULL, '串口类型', 'com_port', 1, '串口', 66);

-- ----------------------------
-- Table structure for sys_interface
-- ----------------------------
DROP TABLE IF EXISTS `sys_interface`;
CREATE TABLE `sys_interface`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '接口ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接口名称',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接口地址',
  `interface_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接口类型',
  `interface_group_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '接口分组名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `sort` int(0) NOT NULL DEFAULT 0 COMMENT 'sort',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 62 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_interface
-- ----------------------------
INSERT INTO `sys_interface` VALUES (1, '2024-10-17 15:28:43', NULL, NULL, '', NULL, NULL, '删除sysUser', '/tpmt/sysUser', 'delete', '系统微服务', NULL, 1);
INSERT INTO `sys_interface` VALUES (2, '2024-10-17 15:29:44', NULL, NULL, '', NULL, NULL, '根据ID查询sysUser', '/tpmt/sysUserInfo', 'get', '系统微服务', NULL, 2);
INSERT INTO `sys_interface` VALUES (3, '2024-10-17 15:30:45', NULL, NULL, '', NULL, NULL, '添加sysUser', '/tpmt/sysUser', 'post', '系统微服务', NULL, 3);
INSERT INTO `sys_interface` VALUES (4, '2024-10-17 15:31:46', NULL, NULL, '', NULL, NULL, '分页查询sysUser', '/tpmt/sysUser', 'get', '系统微服务', NULL, 4);
INSERT INTO `sys_interface` VALUES (5, '2024-10-17 15:32:47', NULL, NULL, '', NULL, NULL, '修改sysUser', '/tpmt/sysUser', 'put', '系统微服务', NULL, 5);
INSERT INTO `sys_interface` VALUES (6, '2024-10-17 15:33:48', NULL, NULL, '', NULL, NULL, '管理员重置用户密码', '/tpmt/sysUserResetPwd', 'put', '系统微服务', NULL, 6);
INSERT INTO `sys_interface` VALUES (7, '2024-10-17 15:34:49', NULL, NULL, '', NULL, NULL, '用户修改自己密码', '/tpmt/sysUserUpMyPwd', 'put', '系统微服务', NULL, 7);
INSERT INTO `sys_interface` VALUES (8, '2024-10-17 15:35:50', NULL, NULL, '', NULL, NULL, '修改自己的用户信息', '/tpmt/sysUserUpMyInfo', 'put', '系统微服务', NULL, 8);
INSERT INTO `sys_interface` VALUES (9, '2024-10-17 15:36:51', NULL, NULL, '', NULL, NULL, '登录完成后获取用户详细信息', '/tpmt/sysUserLoginInfo', 'get', '系统微服务', NULL, 9);
INSERT INTO `sys_interface` VALUES (10, '2024-10-17 15:37:52', NULL, NULL, '', NULL, NULL, '删除sysRole', '/tpmt/sysRole', 'delete', '系统微服务', NULL, 10);
INSERT INTO `sys_interface` VALUES (11, '2024-10-17 15:38:53', NULL, NULL, '', NULL, NULL, '根据ID查询sysRole', '/tpmt/sysRoleInfo', 'get', '系统微服务', NULL, 11);
INSERT INTO `sys_interface` VALUES (12, '2024-10-17 15:39:54', NULL, NULL, '', NULL, NULL, '添加sysRole', '/tpmt/sysRole', 'post', '系统微服务', NULL, 12);
INSERT INTO `sys_interface` VALUES (13, '2024-10-17 15:40:55', NULL, NULL, '', NULL, NULL, '分页查询sysRole', '/tpmt/sysRole', 'get', '系统微服务', NULL, 13);
INSERT INTO `sys_interface` VALUES (14, '2024-10-17 15:41:56', NULL, NULL, '', NULL, NULL, '修改sysRole', '/tpmt/sysRole', 'put', '系统微服务', NULL, 14);
INSERT INTO `sys_interface` VALUES (15, '2024-10-17 15:42:57', NULL, NULL, '', NULL, NULL, '删除sysMenu', '/tpmt/sysMenu', 'delete', '系统微服务', NULL, 15);
INSERT INTO `sys_interface` VALUES (16, '2024-10-17 15:43:58', NULL, NULL, '', NULL, NULL, '根据ID查询sysMenu', '/tpmt/sysMenuInfo', 'get', '系统微服务', NULL, 16);
INSERT INTO `sys_interface` VALUES (17, '2024-10-17 15:44:59', NULL, NULL, '', NULL, NULL, '添加sysMenu', '/tpmt/sysMenu', 'post', '系统微服务', NULL, 17);
INSERT INTO `sys_interface` VALUES (18, '2024-10-17 15:46:00', NULL, NULL, '', NULL, NULL, '分页查询sysMenu', '/tpmt/sysMenu', 'get', '系统微服务', NULL, 18);
INSERT INTO `sys_interface` VALUES (19, '2024-10-17 15:47:01', NULL, NULL, '', NULL, NULL, '修改sysMenu', '/tpmt/sysMenu', 'put', '系统微服务', NULL, 19);
INSERT INTO `sys_interface` VALUES (20, '2024-10-17 15:48:02', NULL, NULL, '', NULL, NULL, '删除sysInterface', '/tpmt/sysInterface', 'delete', '系统微服务', NULL, 20);
INSERT INTO `sys_interface` VALUES (21, '2024-10-17 15:49:03', NULL, NULL, '', NULL, NULL, '根据ID查询sysInterface', '/tpmt/sysInterfaceInfo', 'get', '系统微服务', NULL, 21);
INSERT INTO `sys_interface` VALUES (22, '2024-10-17 15:50:04', NULL, NULL, '', NULL, NULL, '添加sysInterface', '/tpmt/sysInterface', 'post', '系统微服务', NULL, 22);
INSERT INTO `sys_interface` VALUES (23, '2024-10-17 15:51:05', NULL, NULL, '', NULL, NULL, '分页查询sysInterface', '/tpmt/sysInterface', 'get', '系统微服务', NULL, 23);
INSERT INTO `sys_interface` VALUES (24, '2024-10-17 15:52:06', NULL, NULL, '', NULL, NULL, '修改sysInterface', '/tpmt/sysInterface', 'put', '系统微服务', NULL, 24);
INSERT INTO `sys_interface` VALUES (25, '2024-10-17 15:53:07', NULL, NULL, '', NULL, NULL, '删除sysAuth', '/tpmt/sysAuth', 'delete', '系统微服务', NULL, 25);
INSERT INTO `sys_interface` VALUES (26, '2024-10-17 15:54:08', NULL, NULL, '', NULL, NULL, '根据ID查询sysAuth', '/tpmt/sysAuthInfo', 'get', '系统微服务', NULL, 26);
INSERT INTO `sys_interface` VALUES (27, '2024-10-17 15:55:09', NULL, NULL, '', NULL, NULL, '添加sysAuth', '/tpmt/sysAuth', 'post', '系统微服务', NULL, 27);
INSERT INTO `sys_interface` VALUES (28, '2024-10-17 15:56:10', NULL, NULL, '', NULL, NULL, '分页查询sysAuth', '/tpmt/sysAuth', 'get', '系统微服务', NULL, 28);
INSERT INTO `sys_interface` VALUES (29, '2024-10-17 15:57:11', NULL, NULL, '', NULL, NULL, '修改sysAuth', '/tpmt/sysAuth', 'put', '系统微服务', NULL, 29);
INSERT INTO `sys_interface` VALUES (30, '2024-10-17 15:58:12', NULL, NULL, '', NULL, NULL, '删除sysDict', '/tpmt/sysDict', 'delete', '系统微服务', NULL, 30);
INSERT INTO `sys_interface` VALUES (31, '2024-10-17 15:59:13', NULL, NULL, '', NULL, NULL, '根据ID查询sysDict', '/tpmt/sysDictInfo', 'get', '系统微服务', NULL, 31);
INSERT INTO `sys_interface` VALUES (32, '2024-10-17 16:00:14', NULL, NULL, '', NULL, NULL, '添加sysDict', '/tpmt/sysDict', 'post', '系统微服务', NULL, 32);
INSERT INTO `sys_interface` VALUES (33, '2024-10-17 16:01:15', NULL, NULL, '', NULL, NULL, '分页查询sysDict', '/tpmt/sysDict', 'get', '系统微服务', NULL, 33);
INSERT INTO `sys_interface` VALUES (34, '2024-10-17 16:02:16', NULL, NULL, '', NULL, NULL, '修改sysDict', '/tpmt/sysDict', 'put', '系统微服务', NULL, 34);
INSERT INTO `sys_interface` VALUES (35, '2024-10-17 16:03:17', NULL, NULL, '', NULL, NULL, '删除sysDictType', '/tpmt/sysDictType', 'delete', '系统微服务', NULL, 35);
INSERT INTO `sys_interface` VALUES (36, '2024-10-17 16:04:18', NULL, NULL, '', NULL, NULL, '根据ID查询sysDictType', '/tpmt/sysDictTypeInfo', 'get', '系统微服务', NULL, 36);
INSERT INTO `sys_interface` VALUES (37, '2024-10-17 16:05:19', NULL, NULL, '', NULL, NULL, '添加sysDictType', '/tpmt/sysDictType', 'post', '系统微服务', NULL, 37);
INSERT INTO `sys_interface` VALUES (38, '2024-10-17 16:06:20', NULL, NULL, '', NULL, NULL, '分页查询sysDictType', '/tpmt/sysDictType', 'get', '系统微服务', NULL, 38);
INSERT INTO `sys_interface` VALUES (39, '2024-10-17 16:07:21', NULL, NULL, '', NULL, NULL, '修改sysDictType', '/tpmt/sysDictType', 'put', '系统微服务', NULL, 39);
INSERT INTO `sys_interface` VALUES (40, '2024-10-17 16:08:22', NULL, NULL, '', NULL, NULL, '删除tpmtAsset', '/tpmt/tpmtAsset', 'delete', '业务微服务', NULL, 40);
INSERT INTO `sys_interface` VALUES (41, '2024-10-17 16:09:23', NULL, NULL, '', NULL, NULL, '根据ID查询tpmtAsset', '/tpmt/tpmtAssetInfo', 'get', '业务微服务', NULL, 41);
INSERT INTO `sys_interface` VALUES (42, '2024-10-17 16:10:24', NULL, NULL, '', NULL, NULL, '添加tpmtAsset', '/tpmt/tpmtAsset', 'post', '业务微服务', NULL, 42);
INSERT INTO `sys_interface` VALUES (43, '2024-10-17 16:11:25', NULL, NULL, '', NULL, NULL, '分页查询tpmtAsset', '/tpmt/tpmtAsset', 'get', '业务微服务', NULL, 43);
INSERT INTO `sys_interface` VALUES (44, '2024-10-17 16:12:26', NULL, NULL, '', NULL, NULL, '修改tpmtAsset', '/tpmt/tpmtAsset', 'put', '业务微服务', NULL, 44);
INSERT INTO `sys_interface` VALUES (45, '2024-10-17 16:13:27', NULL, NULL, '', NULL, NULL, '删除tpmtGateway', '/tpmt/tpmtGateway', 'delete', '业务微服务', NULL, 45);
INSERT INTO `sys_interface` VALUES (46, '2024-10-17 16:14:28', NULL, NULL, '', NULL, NULL, '根据ID查询tpmtGateway', '/tpmt/tpmtGatewayInfo', 'get', '业务微服务', NULL, 46);
INSERT INTO `sys_interface` VALUES (47, '2024-10-17 16:15:29', NULL, NULL, '', NULL, NULL, '添加tpmtGateway', '/tpmt/tpmtGateway', 'post', '业务微服务', NULL, 47);
INSERT INTO `sys_interface` VALUES (48, '2024-10-17 16:16:30', NULL, NULL, '', NULL, NULL, '分页查询tpmtGateway', '/tpmt/tpmtGateway', 'get', '业务微服务', NULL, 48);
INSERT INTO `sys_interface` VALUES (49, '2024-10-17 16:17:31', NULL, NULL, '', NULL, NULL, '修改tpmtGateway', '/tpmt/tpmtGateway', 'put', '业务微服务', NULL, 49);
INSERT INTO `sys_interface` VALUES (50, '2024-10-17 16:18:32', NULL, NULL, '', NULL, NULL, '删除tpmtMonitorPoint', '/tpmt/tpmtMonitorPoint', 'delete', '业务微服务', NULL, 50);
INSERT INTO `sys_interface` VALUES (51, '2024-10-17 16:19:33', NULL, NULL, '', NULL, NULL, '根据ID查询tpmtMonitorPoint', '/tpmt/tpmtMonitorPointInfo', 'get', '业务微服务', NULL, 51);
INSERT INTO `sys_interface` VALUES (52, '2024-10-17 16:20:34', NULL, NULL, '', NULL, NULL, '添加tpmtMonitorPoint', '/tpmt/tpmtMonitorPoint', 'post', '业务微服务', NULL, 52);
INSERT INTO `sys_interface` VALUES (53, '2024-10-17 16:21:35', NULL, NULL, '', NULL, NULL, '分页查询tpmtMonitorPoint', '/tpmt/tpmtMonitorPoint', 'get', '业务微服务', NULL, 53);
INSERT INTO `sys_interface` VALUES (54, '2024-10-17 16:22:36', NULL, NULL, '', NULL, NULL, '修改tpmtMonitorPoint', '/tpmt/tpmtMonitorPoint', 'put', '业务微服务', NULL, 54);
INSERT INTO `sys_interface` VALUES (55, '2024-10-17 16:23:37', NULL, NULL, '', NULL, NULL, '查询监测点实时数据', '/tpmt/tpmtMonitorPointRealTime', 'get', '业务微服务', NULL, 55);
INSERT INTO `sys_interface` VALUES (56, '2024-10-17 16:24:38', NULL, NULL, '', NULL, NULL, '查询监测点历史数据', '/tpmt/tpmtMonitorPointHistorical', 'get', '业务微服务', NULL, 56);
INSERT INTO `sys_interface` VALUES (57, '2024-10-17 16:25:39', NULL, NULL, '', NULL, NULL, '操作日志', '/appLog', 'post', '日志微服务', NULL, 57);
INSERT INTO `sys_interface` VALUES (58, '2024-10-17 16:26:40', NULL, NULL, '', NULL, NULL, '定时任务日志', '/scheduledTasksLog', 'post', '日志微服务', NULL, 58);
INSERT INTO `sys_interface` VALUES (59, '2024-10-17 16:27:41', NULL, NULL, '', NULL, NULL, '重试任务日志', '/scheduledTasksFailureRecord', 'post', '日志微服务', NULL, 59);
INSERT INTO `sys_interface` VALUES (60, '2024-10-17 16:28:42', NULL, NULL, '', NULL, NULL, '告警和提醒查询', '/alarmLog', 'get', '日志微服务', NULL, 60);
INSERT INTO `sys_interface` VALUES (61, '2024-10-17 16:29:43', NULL, NULL, '', NULL, NULL, '更新告警状态', '/alarmLogUpState', 'get', '日志微服务', NULL, 61);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_type` tinyint(0) NOT NULL COMMENT '菜单类型(层级关系)',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路径',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '本地路径',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '跳转',
  `sort` int(0) NOT NULL COMMENT 'sort',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `is_hide` tinyint(1) NOT NULL COMMENT '是否隐藏',
  `is_keep_alive` tinyint(1) NOT NULL COMMENT '是否缓存',
  `parent_id` int(0) NOT NULL DEFAULT 0 COMMENT '父ID',
  `is_home` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否首页',
  `is_main` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否主菜单',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代理商-菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (13, 1, '菜单1', '菜单111', 'qui', 'in Ut', 'Lorem', 23, 'http://dummyimage.com/100x100', 32, 69, 12, 18, 20, '超级管理员', '2024-10-09 16:02:21', NULL, NULL, NULL, NULL);
INSERT INTO `sys_menu` VALUES (14, 1, '菜单2', '菜单2', 'qui', 'in Ut', 'Lorem', 23, 'http://dummyimage.com/100x100', 32, 69, 12, 18, 20, '超级管理员', '2024-10-09 16:02:30', NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `role_type` tinyint(0) NOT NULL COMMENT '角色类型 1:管理员角色  2:普通角色  3:第三方角色',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (13, '测试角色', 'in labore sunt Lorem', 1, '超级管理员', '2024-11-06 16:41:16', NULL, NULL, NULL, NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 61 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代理商表-角色与菜单中间表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_interface
-- ----------------------------
INSERT INTO `sys_role_interface` VALUES (1, 13, 1, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (2, 13, 2, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (3, 13, 3, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (4, 13, 4, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (5, 13, 5, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (6, 13, 6, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (7, 13, 7, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (8, 13, 8, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (9, 13, 9, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (10, 13, 10, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (11, 13, 11, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (12, 13, 12, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (13, 13, 13, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (14, 13, 14, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (15, 13, 15, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (16, 13, 16, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (17, 13, 17, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (18, 13, 18, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (19, 13, 19, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (20, 13, 20, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (21, 13, 21, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (22, 13, 22, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (23, 13, 23, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (24, 13, 24, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (25, 13, 25, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (26, 13, 26, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (27, 13, 27, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (28, 13, 28, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (29, 13, 29, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (30, 13, 30, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (31, 13, 31, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (32, 13, 32, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (33, 13, 33, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (34, 13, 34, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (35, 13, 35, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (36, 13, 36, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (37, 13, 37, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (38, 13, 38, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (39, 13, 39, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (40, 13, 40, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (41, 13, 41, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (42, 13, 42, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (43, 13, 43, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (44, 13, 44, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (45, 13, 45, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (46, 13, 46, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (47, 13, 47, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (48, 13, 48, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (49, 13, 49, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (50, 13, 50, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (51, 13, 51, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (52, 13, 52, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (53, 13, 53, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (54, 13, 54, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (55, 13, 55, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (56, 13, 56, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (57, 13, 57, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (58, 13, 58, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (59, 13, 59, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (60, 13, 60, '超级管理员', '2024-11-06 17:09:17');
INSERT INTO `sys_role_interface` VALUES (61, 13, 61, '超级管理员', '2024-11-06 17:09:17');

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` int(0) NOT NULL COMMENT '角色ID',
  `menu_id` int(0) NOT NULL COMMENT '菜单ID',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1009 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代理商表-角色与菜单中间表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (1008, 13, 13, '超级管理员', '2024-11-06 16:41:16');
INSERT INTO `sys_role_menu` VALUES (1009, 13, 14, '超级管理员', '2024-11-06 16:41:16');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ID',
  `account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `nick_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `state` tinyint(0) NOT NULL DEFAULT 0 COMMENT '状态 1:正常 2:停用 3:封禁',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `deleted_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代理商-用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('776b9904-95db-44a1-a56f-9616a2fd4b6c', 'admin', '超级管理员', '7c46e0ec4bfdd5800247d45d8369a638757665a30ed8fc9e193ad4976889403e', 1, '超级管理员', '2024-09-26 16:36:28', '袁伟1', '2024-10-12 17:05:13', NULL, NULL);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ID',
  `role_id` int(0) NOT NULL COMMENT '角色ID',
  `user_type` tinyint(0) NULL DEFAULT NULL COMMENT '用户类型 1:普通用户 2:第三方用户',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 60 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代理商表-角色与菜单中间表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, '776b9904-95db-44a1-a56f-9616a2fd4b6c', 13, 1, '超级管理员', '2024-11-06 16:44:17');

-- ----------------------------
-- Table structure for tpmt_asset
-- ----------------------------
DROP TABLE IF EXISTS `tpmt_asset`;
CREATE TABLE `tpmt_asset`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资产ID',
  `asset_type` int(0) NOT NULL DEFAULT 1 COMMENT '资产类型',
  `asset_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资产编号',
  `asset_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资产名称',
  `asset_model` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '资产型号',
  `manu_facturer` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '生产厂家',
  `voltage` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '电压',
  `capacity` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '容量',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tpmt_asset
-- ----------------------------
INSERT INTO `tpmt_asset` VALUES ('2b87ae13-b426-4600-9e49-1732b509ee61', 1, '111', '变压器', '11', '正泰工厂', '220kv', '10L', '2024-10-17 15:33:47', '超级管理员', NULL, NULL);

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
  `address_code` int(0) NOT NULL DEFAULT 1 COMMENT '地址码   com里第几个设备的连接',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tpmt_gateway
-- ----------------------------
INSERT INTO `tpmt_gateway` VALUES ('387e6f1c-af82-47fd-bde9-2b773fb2b142', '2024-10-17 15:38:27', NULL, '超级管理员', NULL, '网关1', '1', '11', 1, 9600, '2', 4, 2, 'tty/dev01  ', 256);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1097 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tpmt_monitor_point
-- ----------------------------
INSERT INTO `tpmt_monitor_point` VALUES (2, '2024-10-23 14:35:01', NULL, '超级管理员', NULL, '0965', '断路器状态', '0x0028', 2, 1, 2, 1, 29, 18, '0', 1, 1.0000000000, 0, NULL, 0, 0, 0, 0, 0, 1, '387e6f1c-af82-47fd-bde9-2b773fb2b142', '2b87ae13-b426-4600-9e49-1732b509ee61', 1);
INSERT INTO `tpmt_monitor_point` VALUES (3, '2024-10-23 16:30:50', NULL, '超级管理员', NULL, '0930', 'C相柜前', '0x0693', 2, 1, 1, 2, 25, 24, '0', 1, 10.0000000000, 2, '℃ ', 40, 100, -20, 70, -10, 0, '387e6f1c-af82-47fd-bde9-2b773fb2b142', '2b87ae13-b426-4600-9e49-1732b509ee61', 2);

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
  `scheduler_task_number` tinyint(0) NOT NULL COMMENT '已接入任务号 1: 北向传输mqtt',
  `scheduler_type` tinyint(0) NOT NULL COMMENT '类型 1:Http任务,2:Webservices任务',
  `interval_time` int(0) NOT NULL COMMENT '间隔时间按秒',
  `error_order` int(0) NOT NULL DEFAULT 1 COMMENT '失败重新发送次数1-10次 不可超过10次',
  `fail_interval_time` int(0) NOT NULL COMMENT '失败间隔时间按秒',
  `state` tinyint(0) NOT NULL DEFAULT 1 COMMENT '状态 1:启动  2:暂停',
  `scheduler_data` varbinary(30000) NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tpmt_scheduled_tasks
-- ----------------------------
INSERT INTO `tpmt_scheduled_tasks` VALUES ('c830bc6f-2231-4a84-a69a-d391b0186116', '2024-11-14 16:43:14', NULL, '超级管理员', NULL, '测试定时任务', 1, 1, 1, 10, 5, 20, 1, 0x7B2267775F736E223A223838383838222C226D73675F696E666F223A223131313131222C226D7174745F686F7374223A223132372E302E302E313A31383833222C226D7174745F75736572223A2261646D696E222C226D7174745F70617373223A223171617A32777378222C2273656E645F746F706963223A226465765F73656E642F66333961373265632D663435642D343565622D626631322D373930616563646665633831227D);

-- ----------------------------
-- Table structure for tpmt_scheduled_tasks_failure_record
-- ----------------------------
DROP TABLE IF EXISTS `tpmt_scheduled_tasks_failure_record`;
CREATE TABLE `tpmt_scheduled_tasks_failure_record`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '失败记录ID',
  `scheduled_tasks_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务ID',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `created_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `updated_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  `scheduler_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `scheduler_category` tinyint(0) NOT NULL COMMENT '类别 1:已接入任务, 2:自定义任务',
  `scheduler_task_number` tinyint(0) NOT NULL COMMENT '已接入任务号',
  `scheduler_type` tinyint(0) NOT NULL COMMENT '类型 1:Http任务,2:Webservices任务',
  `error_order` int(0) NOT NULL DEFAULT 1 COMMENT '失败重新发送次数1-10次 不可超过10次',
  `fail_interval_time` int(0) NOT NULL COMMENT '失败间隔时间按秒',
  `fail_order` int(0) NOT NULL DEFAULT 1 COMMENT '失败次数',
  `scheduler_data` varbinary(30000) NOT NULL COMMENT '内容',
  `request_data` varbinary(30000) NOT NULL COMMENT '请求内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tpmt_scheduled_tasks_failure_record
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
