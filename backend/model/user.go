package model

type Credentials struct {
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Credentials Credentials `gorm:"embedded"`
}

type Token struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	TokenString string `json:"token"`
}
