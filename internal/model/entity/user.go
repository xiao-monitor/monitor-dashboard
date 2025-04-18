package entity

import "time"

type User struct {
	UUID      string    `json:"uuid" gorm:"primaryKey" comment:"用户主键"`
	Username  string    `json:"username" gorm:"type:varchar(32);not null;comment:用户名" validate:"required,min=3,max=255"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null;comment:用户密码" validate:"required,min=6,max=255"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;comment:邮箱" validate:"required,email"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" comment:"更新时间"`
}
