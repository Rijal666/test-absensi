package authdto

type LoginResponse struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Image string `json:"image"`
	Token   string `json:"token"`
	ClockIn string`json:"clock_in"`
}