package payloads

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}