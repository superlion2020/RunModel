package DbModels

// ControlMenu 功能模块表
type Menu struct {
	ID         int64  `gorm:"type:bigint(13);primaryKey"`
	Title      string `gorm:"type:varchar(64);comment:'标题'"`
	IsDel      int64  `gorm:"comment:'是否删除，1是0否'"`
	CreateTime int64  `gorm:"comment:'创建时间'"`
	Sort       int64  `gorm:"comment:'排序'"`
	Contents   string `gorm:"comment:'简介'"`
	//是否是显示的菜单 1是0否
	IsShowMenu int64  `gorm:"comment:'是否是菜单按钮，1是2否'"`
	ParentID   int64  `gorm:"type:bigint(13);comment:'父ID'"`
	MenuUrl    string `gorm:"comment:'路由URL'"`
	MenuIco    string `gorm:"comment:'菜单按钮url'"`
	//是否是基本模块 1是0否 基本模块不需要权限
	IsBase   int64  `gorm:"comment:'是否是基本模块，1是0否'"`
	MenuSign string `gorm:"uniqueIndex;comment:'模块特征码'"`
}
