package ctype

import "encoding/json"

//创建角色权限类型

type Role int

//创建角色常量

const (
	PermissionAdmin       Role = 1 //管理员
	PermissionUser        Role = 2 //普通用户
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //禁用用户
)

//创建Json解析

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	var roleStr string
	switch r {
	case PermissionAdmin:
		roleStr = "管理员"
	case PermissionUser:
		roleStr = "普通用户"
	case PermissionVisitor:
		roleStr = "游客"
	case PermissionDisableUser:
		roleStr = "禁用用户"
	default:
		roleStr = "其他"
	}
	return roleStr
}
