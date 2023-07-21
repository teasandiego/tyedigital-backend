package util

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func StringToArray(str string) []string {
	var str_array []string
	for _, v := range strings.Split(str, ",") {
		// fmt.Println(v)
		str_array = append(str_array, v)
	}
	return str_array
}

func AutoUPdateEverySecond(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
func GenerateSixDigitesRandomNumber() int {
	random_number := rand.Intn(7890-1234) + 1234
	return random_number
}
