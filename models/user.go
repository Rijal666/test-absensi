package models

type User struct {
	ID          int                   `json:"id" gorm:"primary_key:auto_increment"`
	Fullname    string                `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Email       string                `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password    string                `json:"password" form:"password" gorm:"type: varchar(255)"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	ClockIn string `json:"clock_in"`
	ClockOut string `json:"-"`
}

type UsersResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Image string `json:"image"`
	ClockIn string `json:"clock_in"`
}

func (UsersResponse) TableName() string {
	return "users"
}
