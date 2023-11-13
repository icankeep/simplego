package fmtx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFmtx(t *testing.T) {
	a := assert.New(t)

	str := "// SystemUser 用户表\ntype SystemUser struct {\n\tId int64 `json:\"id\"` // 主键id\n\tUsername string `json:\"username\"` // 登录用户名\n\tPassword string `json:\"password\"` // 登录密码\n\tSalt string `json:\"salt\"` // 密码盐值\n\tEmail string `json:\"email\"` // 邮箱\n\tOpenid string `json:\"openid\"` // 微信标识\n\tPhone string `json:\"phone\"` // 手机号\n\tType int8 `json:\"type\"` // 注册方式\n\tCreatedAt int64 `json:\"createdAt\"` // 注册时间\n\tStatus int8 `json:\"status\"` // 状态\n\tNickname string `json:\"nickname\"` // 昵称\n\tAvatarUrl string `json:\"avatarUrl\"` // 头像\n\tGender int8 `json:\"gender\"` // 性别\n\tUpdatedAt int64 `json:\"updatedAt\"` // 更新时间\n\tRoleId int64 `json:\"roleId\"` // 用户角色\n}"
	expected := "// SystemUser 用户表\ntype SystemUser struct {\n\tId        int64  `json:\"id\"`        // 主键id\n\tUsername  string `json:\"username\"`  // 登录用户名\n\tPassword  string `json:\"password\"`  // 登录密码\n\tSalt      string `json:\"salt\"`      // 密码盐值\n\tEmail     string `json:\"email\"`     // 邮箱\n\tOpenid    string `json:\"openid\"`    // 微信标识\n\tPhone     string `json:\"phone\"`     // 手机号\n\tType      int8   `json:\"type\"`      // 注册方式\n\tCreatedAt int64  `json:\"createdAt\"` // 注册时间\n\tStatus    int8   `json:\"status\"`    // 状态\n\tNickname  string `json:\"nickname\"`  // 昵称\n\tAvatarUrl string `json:\"avatarUrl\"` // 头像\n\tGender    int8   `json:\"gender\"`    // 性别\n\tUpdatedAt int64  `json:\"updatedAt\"` // 更新时间\n\tRoleId    int64  `json:\"roleId\"`    // 用户角色\n}"

	fmtStr, err := FormatGoCode(str)
	a.NoError(err)

	//file, err := os.Create("/tmp/tmp2.txt")
	//a.NoError(err)
	//file.Write([]byte(fmtStr))

	a.Equal(expected, fmtStr)
	fmt.Println(fmtStr)
}
