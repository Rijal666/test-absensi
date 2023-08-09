package userdto

type CreateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

type UpdateUserRequest struct {
	UserID int `json:"user_id" form:"user_id"`
	ClockOut string `json:"clock_out" form:"clock_out"`
}