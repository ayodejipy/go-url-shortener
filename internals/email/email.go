package email

type Email interface {
	SendPasswordToken(token string) error
}