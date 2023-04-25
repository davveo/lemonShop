package entity

import "github.com/davveo/lemonShop/app/consts"

type User struct {
	Uid      int64         `json:"uid"`
	Uuid     string        `json:"uuid"`
	UserName string        `json:"username"`
	Roles    []consts.Role `json:"roles"`
}

func (u *User) Add(roles ...consts.Role) {
	for _, role := range roles {
		u.Roles = append(u.Roles, role)
	}
}
