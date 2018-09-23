package subscriber

import (
	"context"
	"encoding/base64"

	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-log"
	"golang.org/x/crypto/bcrypt"

	account "github.com/ahmadnurus/go-micro-bookstore/account/srv/proto/account"
)

// AccountCreated subscriber struct
type AccountCreated struct {
	Redis *redis.Pool
}

// SendEmail verification
func SendEmail(email string, url string) {}

// Process AccountCreated subscriber
func (sub *AccountCreated) Process(ctx context.Context, account *account.Credential) error {
	token, err := bcrypt.GenerateFromPassword([]byte(account.Email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	conn := sub.Redis.Get()
	defer conn.Close()
	_, err = conn.Do("SETEX", account.Email, 600, token)
	if err != nil {
		return err
	}

	log.Log("Sending email to: ", account.Email)
	log.Log("Account ID: ", account.Id)
	log.Log("Token: ", base64.StdEncoding.EncodeToString(token))
	return nil
}
