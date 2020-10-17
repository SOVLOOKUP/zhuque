package role

type Role struct {
	ID      int    `json:"id"`
	RoleNum int    `json:"roleNum"` //超管1，管理员2，开发3
	Name    string `json:"name"`    //说明
}

// 设置Permission的表名为`permission`
func (Role) TableName() string {
	return "role"
}
