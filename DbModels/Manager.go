package DbModels

// ControlManager 管理员表
type Manager struct {
	ID             int64  `gorm:"type:bigint(13);primaryKey"`
	LoginName      string `gorm:"type:varchar(64);comment:'登录用户名'"`
	LoginPwd       string `gorm:"type:varchar(64);comment:'登录密码'"`
	HeadImageUrl   string `gorm:"type:varchar(64);comment:'头像URL'"`
	CreateTime     int64  `gorm:"type:bigint(13);comment:'创建时间'"`
	LastLoginTime  int64  `gorm:"type:bigint(13);comment:'最后登录时间'"`
	LastUpdateUser string `gorm:"type:varchar(32);comment:'最后操作人'"`
	LastUpdateTime int64  `gorm:"type:bigint(13);comment:'最后操作时间'"`
	NickName       string `gorm:"type:varchar(64);comment:'昵称'"`
	IsLock         uint8  `gorm:"type:tinyint(4);comment:'是否锁定，1是2否'"`
	State          uint8  `gorm:"type:tinyint(4);comment:'状态 1:启用 2:禁用'"`
	IsGoogle       uint8  `gorm:"type:tinyint(4);comment:'是否开启谷歌验证，1是2否'"`
	GoogleKey      string `gorm:"type:varchar(32);comment:'谷歌密钥'"`
	IsDel          uint8  `gorm:"type:tinyint(4);comment:'是否已删除，1是2否'"`
	RegisterIP     string `gorm:"type:varchar(32);comment:'注册IP'"`
	LastLoginIP    string `gorm:"type:varchar(32);comment:'最后登录IP'"`
}
