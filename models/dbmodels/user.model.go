package dbmodels

type User struct {
	ID           int64     `json:"id" gorm:"column:id"`
	UserName     string    `json:"userName" gorm:"column:username"`
	Password     string    `json:"password" gorm:"column:password" `
}

// TableName ...
func (t *User) TableName() string {
	return "public.user"
}
