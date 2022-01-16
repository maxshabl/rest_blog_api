package request

type Registration struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}
