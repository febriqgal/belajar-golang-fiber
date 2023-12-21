package entity

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique;column:id;type:bigint;"`
	Username string `json:"username" gorm:"not null;unique;column:username;type:varchar(100);"`
	Password string `json:"password" gorm:"not null;column:password;type:varchar(255);"`
	Email    string `json:"email" gorm:"not null;unique;column:email;type:varchar(100);" validate:"email" required:"true"`
}
