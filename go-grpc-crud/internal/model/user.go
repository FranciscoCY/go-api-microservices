package model

// User representa un usuario en la tabla `users`
type User struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Email string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
}

// TableName le dice a GORM que use la tabla "users"
func (User) TableName() string {
	return "users"
}
