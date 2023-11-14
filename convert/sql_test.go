package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SQL2Struct_JSONTag(t *testing.T) {
	a := assert.New(t)

	sql := "CREATE TABLE `system_user` (\n  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '登录用户名',\n  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',\n  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '密码盐值',\n  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',\n  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信标识',\n  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',\n  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '注册方式',\n  `created_at` int(10) NOT NULL DEFAULT '0' COMMENT '注册时间',\n  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',\n  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',\n  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',\n  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',\n  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',\n  `role_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户角色',\n  PRIMARY KEY (`id`),\n  KEY `UsernameIndex` (`username`),\n  KEY `EmailIndex` (`email`),\n  KEY `PhoneIndex` (`phone`),\n  KEY `OpenidIndex` (`openid`)\n) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'"
	expected := "// SystemUser 用户表\ntype SystemUser struct {\n\tId        int64  `json:\"Id\"`        // 主键id\n\tUsername  string `json:\"Username\"`  // 登录用户名\n\tPassword  string `json:\"Password\"`  // 登录密码\n\tSalt      string `json:\"Salt\"`      // 密码盐值\n\tEmail     string `json:\"Email\"`     // 邮箱\n\tOpenid    string `json:\"Openid\"`    // 微信标识\n\tPhone     string `json:\"Phone\"`     // 手机号\n\tType      int8   `json:\"Type\"`      // 注册方式\n\tCreatedAt int64  `json:\"CreatedAt\"` // 注册时间\n\tStatus    int8   `json:\"Status\"`    // 状态\n\tNickname  string `json:\"Nickname\"`  // 昵称\n\tAvatarUrl string `json:\"AvatarUrl\"` // 头像\n\tGender    int8   `json:\"Gender\"`    // 性别\n\tUpdatedAt int64  `json:\"UpdatedAt\"` // 更新时间\n\tRoleId    int64  `json:\"RoleId\"`    // 用户角色\n}\n"
	structStr, err := ToGoStruct(sql, []string{"json"})
	a.NoError(err)
	a.Equal(expected, structStr)
	//fmt.Println(structStr)
}

func Test_SQL2Struct_GormTag(t *testing.T) {
	a := assert.New(t)

	sql := "CREATE TABLE `system_user` (\n  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '登录用户名',\n  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',\n  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '密码盐值',\n  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',\n  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信标识',\n  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',\n  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '注册方式',\n  `created_at` int(10) NOT NULL DEFAULT '0' COMMENT '注册时间',\n  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',\n  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',\n  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',\n  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',\n  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',\n  `role_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户角色',\n  PRIMARY KEY (`id`),\n  KEY `UsernameIndex` (`username`),\n  KEY `EmailIndex` (`email`),\n  KEY `PhoneIndex` (`phone`),\n  KEY `OpenidIndex` (`openid`)\n) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'"
	expected := "// SystemUser 用户表\ntype SystemUser struct {\n\tId        int64  `gorm:\"id\"`         // 主键id\n\tUsername  string `gorm:\"username\"`   // 登录用户名\n\tPassword  string `gorm:\"password\"`   // 登录密码\n\tSalt      string `gorm:\"salt\"`       // 密码盐值\n\tEmail     string `gorm:\"email\"`      // 邮箱\n\tOpenid    string `gorm:\"openid\"`     // 微信标识\n\tPhone     string `gorm:\"phone\"`      // 手机号\n\tType      int8   `gorm:\"type\"`       // 注册方式\n\tCreatedAt int64  `gorm:\"created_at\"` // 注册时间\n\tStatus    int8   `gorm:\"status\"`     // 状态\n\tNickname  string `gorm:\"nickname\"`   // 昵称\n\tAvatarUrl string `gorm:\"avatar_url\"` // 头像\n\tGender    int8   `gorm:\"gender\"`     // 性别\n\tUpdatedAt int64  `gorm:\"updated_at\"` // 更新时间\n\tRoleId    int64  `gorm:\"role_id\"`    // 用户角色\n}\n"
	structStr, err := ToGoStruct(sql, []string{"gorm"})
	a.NoError(err)
	a.Equal(expected, structStr)
	//fmt.Println(structStr)
}

func Test_SQL2Struct_MultiTags(t *testing.T) {
	a := assert.New(t)

	sql := "CREATE TABLE `system_user` (\n  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '登录用户名',\n  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',\n  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '密码盐值',\n  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',\n  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信标识',\n  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',\n  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '注册方式',\n  `created_at` int(10) NOT NULL DEFAULT '0' COMMENT '注册时间',\n  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',\n  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',\n  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',\n  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',\n  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',\n  `role_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户角色',\n  PRIMARY KEY (`id`),\n  KEY `UsernameIndex` (`username`),\n  KEY `EmailIndex` (`email`),\n  KEY `PhoneIndex` (`phone`),\n  KEY `OpenidIndex` (`openid`)\n) ENGINE=InnoDB AUTO_INCREMENT=652 DEFAULT CHARSET=utf8mb4 COMMENT='用户表'"
	expected := "// SystemUser 用户表\ntype SystemUser struct {\n\tId        int64  `json:\"Id\" gorm:\"id\" xml:\"id\"`                        // 主键id\n\tUsername  string `json:\"Username\" gorm:\"username\" xml:\"username\"`      // 登录用户名\n\tPassword  string `json:\"Password\" gorm:\"password\" xml:\"password\"`      // 登录密码\n\tSalt      string `json:\"Salt\" gorm:\"salt\" xml:\"salt\"`                  // 密码盐值\n\tEmail     string `json:\"Email\" gorm:\"email\" xml:\"email\"`               // 邮箱\n\tOpenid    string `json:\"Openid\" gorm:\"openid\" xml:\"openid\"`            // 微信标识\n\tPhone     string `json:\"Phone\" gorm:\"phone\" xml:\"phone\"`               // 手机号\n\tType      int8   `json:\"Type\" gorm:\"type\" xml:\"type\"`                  // 注册方式\n\tCreatedAt int64  `json:\"CreatedAt\" gorm:\"created_at\" xml:\"created_at\"` // 注册时间\n\tStatus    int8   `json:\"Status\" gorm:\"status\" xml:\"status\"`            // 状态\n\tNickname  string `json:\"Nickname\" gorm:\"nickname\" xml:\"nickname\"`      // 昵称\n\tAvatarUrl string `json:\"AvatarUrl\" gorm:\"avatar_url\" xml:\"avatar_url\"` // 头像\n\tGender    int8   `json:\"Gender\" gorm:\"gender\" xml:\"gender\"`            // 性别\n\tUpdatedAt int64  `json:\"UpdatedAt\" gorm:\"updated_at\" xml:\"updated_at\"` // 更新时间\n\tRoleId    int64  `json:\"RoleId\" gorm:\"role_id\" xml:\"role_id\"`          // 用户角色\n}\n"
	structStr, err := ToGoStruct(sql, []string{"json", "gorm", "xml"})
	a.NoError(err)
	a.Equal(expected, structStr)
	//fmt.Println(structStr)
}
