package email

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"rest/api/internals/config"
	"rest/api/internals/logger"
	"strings"
	"time"
)

type SendEmailHandler struct {
	Config *config.AppConfig
	Logger *logger.Logger
	// server *api.Server
}

func NewSendEmailHandler(config *config.AppConfig, logger *logger.Logger) *SendEmailHandler {
	return &SendEmailHandler{
		Config: config,
		Logger: logger,
	}
}

func (s *SendEmailHandler) sendEmail(to []string, message string) error {
	httpHost := "https://sandbox.api.mailtrap.io/api/send/3137985"
	// 	payload := strings.NewReader(`{
	// 	"from":{"email":"hello@example.com"},
	// 	"to":[{"email":"jeggsatony@gmail.com"}],
	// 	"subject":"You are awesome!",
	// 	"text":"Congrats for sending test email with Mailtrap!",
	// 	"category":"Integration Test"
	// }`)

	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest(http.MethodPost, httpHost, strings.NewReader(message))
	if err != nil {
		s.Logger.Error("[request]: %v", err)
		return err
	}
	request.Header.Set("Authorization", "Bearer "+s.Config.MailtrapApiKey)
	request.Header.Set("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		s.Logger.Error("[httpClient]: %v", err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.Logger.Error("[Reading Body]: %v", err)
		return err
	}

	fmt.Println("Response", string(body))

	return nil
}

func (s *SendEmailHandler) SendPasswordToken(token string, toEmail string) error {
	to := []string{"kate.doe@example.com"}

	message := `{
		"from":{"email":"hello@example.com"},
		"to":[{"email":"jeggsatony@gmail.com"}],
		"subject":"Password Reset",
		"text":"Here's your reset token: #token",
		"category":"Integration Test"
	}`

	message = strings.Replace(message, "#token", token, 1)

	err := s.sendEmail(to, message)
	if err != nil {
		s.Logger.Error("[SendPasswordToken]: %v", err)
		log.Fatal(err)
		return errors.New("failed to send [password reset] mail")
	}

	return nil
}

func (s *SendEmailHandler) SendOTPEmail(token string, toEmail string) error {
	to := []string{"kate.doe@example.com"}

	message := `{
		"from":{"email":"hello@example.com"},
		"to":[{"email":"jeggsatony@gmail.com"}],
		"subject":"Verification code",
		"text":"Hurray! Good to have you onboard! Your one time verification code is #token",
		"category":"Integration Test"
	}`

	message = strings.Replace(message, "#token", token, 1)

	err := s.sendEmail(to, message)
	if err != nil {
		s.Logger.Error("[SendOTPEmail]: %v", err)
		// log.Fatal(err)
		return errors.New("failed to send [OTP] verification mail")
	}

	return nil
}

func (s *SendEmailHandler) SendPasswordResetMail(to string) error {
	message := `{
		"from":{"email":"hello@example.com"},
		"to":[{"email":"{recepient}"}],
		"subject":"Password Reset Succesfully",
		"text":"Your password reset request has been completed successfully. You can now proceed to login.",
		"category":"Integration Test"
	}`

	message = strings.Replace(message, "{recepient}", to, 1)

	err := s.sendEmail([]string{to}, message)
	if err != nil {
		s.Logger.Error("[SendPasswordResetEmail]: %v", err)
		log.Fatal(err)
		return nil
	}

	return nil
}
