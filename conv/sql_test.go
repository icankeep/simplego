package conv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SQL2Struct_JSONTag(t *testing.T) {
	a := assert.New(t)

	sql := "CREATE TABLE `system_user` (\n  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '登录用户名',\n  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',\n  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '密码盐值',\n  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',\n  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信标识',\n  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',\n  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '注册方式',\n  `created_at` int(10) NOT NULL DEFAULT '0' COMMENT '注册时间',\n  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',\n  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',\n  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',\n  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',\n  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',\n  `role_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户角色',\n  PRIMARY KEY (`id`),\n  KEY `UsernameIndex` (`username`),\n  KEY `EmailIndex` (`email`),\n  KEY `PhoneIndex` (`phone`),\n  KEY `OpenidIndex` (`openid`)\n) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'"
	expected := "// SystemUser 用户表\ntype SystemUser struct {\n\tID        int64  `json:\"ID\" gorm:\"column:id\"`                // 主键id\n\tUsername  string `json:\"Username\" gorm:\"column:username\"`    // 登录用户名\n\tPassword  string `json:\"Password\" gorm:\"column:password\"`    // 登录密码\n\tSalt      string `json:\"Salt\" gorm:\"column:salt\"`            // 密码盐值\n\tEmail     string `json:\"Email\" gorm:\"column:email\"`          // 邮箱\n\tOpenid    string `json:\"Openid\" gorm:\"column:openid\"`        // 微信标识\n\tPhone     string `json:\"Phone\" gorm:\"column:phone\"`          // 手机号\n\tType      bool   `json:\"Type\" gorm:\"column:type\"`            // 注册方式\n\tCreatedAt int64  `json:\"CreatedAt\" gorm:\"column:created_at\"` // 注册时间\n\tStatus    bool   `json:\"Status\" gorm:\"column:status\"`        // 状态\n\tNickname  string `json:\"Nickname\" gorm:\"column:nickname\"`    // 昵称\n\tAvatarUrl string `json:\"AvatarUrl\" gorm:\"column:avatar_url\"` // 头像\n\tGender    bool   `json:\"Gender\" gorm:\"column:gender\"`        // 性别\n\tUpdatedAt int64  `json:\"UpdatedAt\" gorm:\"column:updated_at\"` // 更新时间\n\tRoleID    int64  `json:\"RoleID\" gorm:\"column:role_id\"`       // 用户角色\n}\n"
	structStr, err := ToGoStruct(sql, []string{"json", "gorm"})
	a.NoError(err)
	a.Equal(expected, structStr)
	//fmt.Println(structStr)
}

func Test_SQL2Struct_GormTag(t *testing.T) {
	a := assert.New(t)

	sql := "CREATE TABLE `system_user` (\n  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '登录用户名',\n  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',\n  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '密码盐值',\n  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',\n  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信标识',\n  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',\n  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '注册方式',\n  `created_at` int(10) NOT NULL DEFAULT '0' COMMENT '注册时间',\n  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',\n  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',\n  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',\n  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',\n  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',\n  `role_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户角色',\n  PRIMARY KEY (`id`),\n  KEY `UsernameIndex` (`username`),\n  KEY `EmailIndex` (`email`),\n  KEY `PhoneIndex` (`phone`),\n  KEY `OpenidIndex` (`openid`)\n) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'"
	expected := "// SystemUser 用户表\ntype SystemUser struct {\n\tID        int64  `gorm:\"column:id\"`         // 主键id\n\tUsername  string `gorm:\"column:username\"`   // 登录用户名\n\tPassword  string `gorm:\"column:password\"`   // 登录密码\n\tSalt      string `gorm:\"column:salt\"`       // 密码盐值\n\tEmail     string `gorm:\"column:email\"`      // 邮箱\n\tOpenid    string `gorm:\"column:openid\"`     // 微信标识\n\tPhone     string `gorm:\"column:phone\"`      // 手机号\n\tType      bool   `gorm:\"column:type\"`       // 注册方式\n\tCreatedAt int64  `gorm:\"column:created_at\"` // 注册时间\n\tStatus    bool   `gorm:\"column:status\"`     // 状态\n\tNickname  string `gorm:\"column:nickname\"`   // 昵称\n\tAvatarUrl string `gorm:\"column:avatar_url\"` // 头像\n\tGender    bool   `gorm:\"column:gender\"`     // 性别\n\tUpdatedAt int64  `gorm:\"column:updated_at\"` // 更新时间\n\tRoleID    int64  `gorm:\"column:role_id\"`    // 用户角色\n}\n"
	structStr, err := ToGoStruct(sql, []string{"gorm"})
	a.NoError(err)
	a.Equal(expected, structStr)
	//fmt.Println(structStr)
}

func Test_SQL2Struct_MultiTags(t *testing.T) {
	a := assert.New(t)

	sql := "CREATE TABLE `system_user` (\n  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '登录用户名',\n  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',\n  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '密码盐值',\n  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',\n  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信标识',\n  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',\n  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '注册方式',\n  `created_at` int(10) NOT NULL DEFAULT '0' COMMENT '注册时间',\n  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',\n  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',\n  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',\n  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',\n  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',\n  `role_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户角色',\n  PRIMARY KEY (`id`),\n  KEY `UsernameIndex` (`username`),\n  KEY `EmailIndex` (`email`),\n  KEY `PhoneIndex` (`phone`),\n  KEY `OpenidIndex` (`openid`)\n) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'"
	expected := "// SystemUser 用户表\ntype SystemUser struct {\n\tID        int64  `json:\"ID\" gorm:\"column:id\" xml:\"id\"`                        // 主键id\n\tUsername  string `json:\"Username\" gorm:\"column:username\" xml:\"username\"`      // 登录用户名\n\tPassword  string `json:\"Password\" gorm:\"column:password\" xml:\"password\"`      // 登录密码\n\tSalt      string `json:\"Salt\" gorm:\"column:salt\" xml:\"salt\"`                  // 密码盐值\n\tEmail     string `json:\"Email\" gorm:\"column:email\" xml:\"email\"`               // 邮箱\n\tOpenid    string `json:\"Openid\" gorm:\"column:openid\" xml:\"openid\"`            // 微信标识\n\tPhone     string `json:\"Phone\" gorm:\"column:phone\" xml:\"phone\"`               // 手机号\n\tType      bool   `json:\"Type\" gorm:\"column:type\" xml:\"type\"`                  // 注册方式\n\tCreatedAt int64  `json:\"CreatedAt\" gorm:\"column:created_at\" xml:\"created_at\"` // 注册时间\n\tStatus    bool   `json:\"Status\" gorm:\"column:status\" xml:\"status\"`            // 状态\n\tNickname  string `json:\"Nickname\" gorm:\"column:nickname\" xml:\"nickname\"`      // 昵称\n\tAvatarUrl string `json:\"AvatarUrl\" gorm:\"column:avatar_url\" xml:\"avatar_url\"` // 头像\n\tGender    bool   `json:\"Gender\" gorm:\"column:gender\" xml:\"gender\"`            // 性别\n\tUpdatedAt int64  `json:\"UpdatedAt\" gorm:\"column:updated_at\" xml:\"updated_at\"` // 更新时间\n\tRoleID    int64  `json:\"RoleID\" gorm:\"column:role_id\" xml:\"role_id\"`          // 用户角色\n}\n"
	structStr, err := ToGoStruct(sql, []string{"json", "gorm", "xml"})
	a.NoError(err)
	a.Equal(expected, structStr)
	//fmt.Println(structStr)
}
