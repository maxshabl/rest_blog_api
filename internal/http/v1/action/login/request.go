package login

type Request struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}
