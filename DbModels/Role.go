package DbModels

// ControlRole 角色表
type ControlRole struct {
	ID         int64  `gorm:"type:bigint(13);primaryKey"`
	Title      string `gorm:"type:varchar(64);comment:'角色名'"`
	TitleEn    string `gorm:"type:varchar(64);comment:'角色英文名'"`
	IsDel      int64  `gorm:"comment:'是否删除，1是2否'"`
	CreateTime int64  `gorm:"comment:'创建时间'"`
	Sort       int64  `gorm:"comment:'排序'"`
	Contents   string `gorm:"comment:'简介'"`
	// 角色类型 1是超管 2是其他
	RoleType int64  `gorm:"comment:'角色类型 1是超管 2是其他'"`
	ParentID string `gorm:"type:varchar(64);comment:'父ID'"`
	IsEnable int64  `gorm:"comment:'是否启用，1是2否'"`
}
