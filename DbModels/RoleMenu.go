package DbModels

// RoleMenu 管理员与角色关系表
type RoleMenu struct {
	ID        int64 `gorm:"type:bigint(13);primaryKey"`
	RoleID    int64 `gorm:"type:bigint(13);comment:'角色ID'"`
	ManagerID int64 `gorm:"type:bigint(13);comment:'管理员ID'"`
}
