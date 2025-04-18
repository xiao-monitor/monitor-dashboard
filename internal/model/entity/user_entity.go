package entity

import "time"

// UserEntity 用户实体
type UserEntity struct {
	UUID      string    `json:"uuid" gorm:"primaryKey;type:uuid" comment:"用户主键"`
	Username  string    `json:"username" gorm:"type:varchar(32);not null;comment:用户名" validate:"required,min=3,max=32"`
	Password  string    `json:"password" gorm:"type:varchar(60);not null;comment:用户密码" validate:"required,min=6,max=60"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;comment:邮箱" validate:"required,email"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()" comment:"更新时间"`
}
