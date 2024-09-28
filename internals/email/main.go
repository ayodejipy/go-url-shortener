package email

type Email interface {
	SendPasswordToken(token string) error
	SendPasswordResetMail(to string) error
}