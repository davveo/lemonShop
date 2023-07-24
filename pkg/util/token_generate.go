package util

import (
	"strconv"

	"github.com/davveo/lemonShop/app/consts"
)

// GenerateBuyerAccessToken 生成token KEY
func GenerateBuyerAccessToken(uuid string, memberId int) string {
	return consts.AccessToken + strconv.Itoa(memberId) + "_" + uuid
}

// GenerateBuyerRefreshToken 生成模糊token 提供删除使用
func GenerateBuyerRefreshToken(uuid string, memberId int) string {
	return consts.RefreshToken + strconv.Itoa(memberId) + "_" + uuid
}

func GenerateVagueBuyerAccessToken(memberId int) string {
	return consts.AccessToken + strconv.Itoa(memberId) + "_"
}

func GenerateVagueBuyerRefreshToken(memberId int) string {
	return consts.RefreshToken + strconv.Itoa(memberId) + "_"
}

func GenerateAdminAccessToken(uuid string, memberId int) string {
	return consts.AccessToken + "ADMIN_" + strconv.Itoa(memberId) + "_" + uuid
}

func GenerateAdminRefreshToken(uuid string, memberId int) string {
	return consts.RefreshToken + "ADMIN_" + strconv.Itoa(memberId) + "_" + uuid
}
