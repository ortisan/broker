package vo

import (
	"math/rand"
	errapp "ortisan-broker/go-commons/error"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const secretLength = 36
const minCharIdx = 0
const maxCharIdx = 94

type Secret interface {
	Value() string
}

type secret struct {
	value string
}

func (p *secret) Value() string {
	return p.value
}

func NewSecret() Secret {
	secretVal := randomPassword(secretLength)
	return &secret{
		value: string(secretVal),
	}
}

func NewSecretFromValue(value string) (Name, error) {
	if value == "" {
		return nil, errapp.NewBadArgumentError("secret is invalid")
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(value), 14)
	return &secret{
		value: string(bytes),
	}, nil
}

func NewSecretFromValueCrypted(value string) (Name, error) {
	return &secret{
		value: value,
	}, nil
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomPassword(length int) string {
	seed := time.Now().Unix()
	rand.Seed(seed)
	startChar := "!"
	var i int = 1
	var sb strings.Builder
	for {
		myRand := random(minCharIdx, maxCharIdx)
		newChar := string(startChar[0] + byte(myRand))
		sb.WriteString(newChar)
		if i == int(length) {
			break
		}
		i++
	}
	return sb.String()
}
