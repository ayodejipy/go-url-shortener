package email

type Email interface {
	SendPasswordToken(token string, toEmail string) error
	SendOTPEmail(token string, toEmail string) error
	SendPasswordResetMail(to string) error
}