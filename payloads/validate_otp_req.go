package payloads

type ValidateOtpRequest struct {
	ReferenceId string `json:"reference_id"`
	OTP         string `json:"otp"`
}