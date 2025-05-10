package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

var JWTKEY []byte

// Loading environment variable JWTKEY which is the secret to used for encoding the JWT
// GO init function will load when the program is ran and run this automatically
func init() {
	var err error = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var jwtSecret string = os.Getenv("JWT")
	if jwtSecret == "" {
		log.Fatalf("JWT environment variable not set")
	}
	JWTKEY = []byte(jwtSecret)
	fmt.Println("JWT Key loaded successfully.")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateJWT(username string) string {
	var (
		t *jwt.Token
		s string
	)

	var now time.Time = time.Now()
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "seismicstudios",
		"username": username,
		// setting a time of when this was issued to when this is going to expire
		// Does not work at the moment due to how it is implemented
		// I had many issues trying to set JWT as a cookie but it would not set
		// even though request was seen
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 12).Unix(),
	})

	s, _ = t.SignedString(JWTKEY)
	return s

}

func VerifyToken(token string) (*jwt.Token, bool) {
	parsedToken, _ := jwt.Parse(token, parseToken)
	if !parsedToken.Valid {
		return nil, false
	} else {
		return parsedToken, true
	}
}

func parseToken(jwt *jwt.Token) (interface{}, error) {
	return JWTKEY, nil
}

func TestPasswords() {
	password := "test"
	hash, _ := HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
