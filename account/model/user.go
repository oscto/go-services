package model

type Users struct {
	BaseModel
	Username string `json:"username" gorm:"type:varchar(12);uniqueIndex;not null"`
	Email    string `json:"email" gorm:"type:varchar(64);uniqueIndex;not null"`
	Nickname string `json:"nickname" gorm:"type:varchar(24)"`
	Salt     string `json:"salt" gorm:"type:varchar(8)"`
	Password string `json:"password" gorm:"type:varchar(32)"`
	Status   uint8  `json:"status" gorm:"default:1"` // 状态(1:可用,2:不可用)
}

func (Users) TableName() string {

	return "users"
}
