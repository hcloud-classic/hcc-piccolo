package userTool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"time"
)

type claims struct {
	Id       string
	Password string
	jwt.StandardClaims
}

var jwtKey = []byte("hccJWTKey")
var loginMismatchError = errors.New("can not login with provided token")

func GenerateToken(id string, password string) (string, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(time.Minute * time.Duration(config.User.TokenExpirationTimeMinutes))
	// Create the JWT claims, which includes the user id and password with expiry time
	claims := &claims{
		Id:       id,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			Issuer:  "piccolo",
			Subject: "Auth",
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		// w.WriteHeader(http.StatusInternalServerError)
		return "", errors.New("token signing error")
	} else {
		return tokenString, nil
	}
}

// ValidateToken : Validate given token string
func ValidateToken(args map[string]interface{}) error {
	tokenString, tokenStringOk := args["token"].(string)
	if !tokenStringOk {
		return errors.New("need a token argument")
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtKey, nil
	})
	if err != nil {
		return errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["iss"].(string) != "piccolo" &&
			claims["sub"].(string) != "Auth" {
			return errors.New("invalid token")
		}

		if time.Now().Unix() >= int64(claims["exp"].(float64)) {
			return errors.New("token is expired")
		}

		var dbPassword string

		sql := "select password from user where id = ?"
		err := mysql.Db.QueryRow(sql, claims["ID"].(string)).Scan(&dbPassword)
		if err != nil {
			logger.Logger.Println(err)
			return loginMismatchError
		}

		// Given password is hashed password with bcrypt
		err = bcrypt.CompareHashAndPassword([]byte(claims["Password"].(string)), []byte(dbPassword))
		if err != nil {
			return loginMismatchError
		}

		logger.Logger.Println("Token validated for user [" + claims["ID"].(string) + "]")

		return nil
	}

	return errors.New("invalid token")
}
