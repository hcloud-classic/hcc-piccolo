package usertool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"hcc/piccolo/action/graphql/queryparserExt"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"strings"
	"time"
)

type claims struct {
	ID       string
	Password string
	jwt.StandardClaims
}

var jwtKey = []byte("hccJWTKey")
var errLoginMismatch = errors.New("can not login with provided token")

// GenerateToken : Generate token by ID and password.
func GenerateToken(id string, password string) (string, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(time.Minute * time.Duration(config.User.TokenExpirationTimeMinutes))
	// Create the JWT claims, which includes the user id and password with expiry time
	claims := &claims{
		ID:       id,
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
	}

	return tokenString, nil
}

func getGroupOfUser(id string) *model.Group {
	var group model.Group

	sql := "select piccolo.group.id, piccolo.group.name from piccolo.group where id in (select piccolo.user.group_id as id from piccolo.user where id = ?)"
	row := mysql.Db.QueryRow(sql, id)
	err := mysql.QueryRowScan(row, &group.ID, &group.Name)
	if err != nil {
		logger.Logger.Println("getGroupOfUser(): " + err.Error())
		return nil
	}

	return &group
}

// ValidateToken : Validate given token string
func ValidateToken(args map[string]interface{}, checkForAdmin bool) (err error, isAdmin bool, isMaster bool, id string, groupID int64) {
	var _groupID int

	tokenString, tokenStringOk := args["token"].(string)
	_groupID, _groupIDOk := args["group_id"].(int)

	if !tokenStringOk {
		return errors.New("need a token argument"), false, false, "", 0
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
		return errors.New("invalid token"), false, false, "", 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["iss"].(string) != "piccolo" &&
			claims["sub"].(string) != "Auth" {
			return errors.New("invalid token"), false, false, "", 0
		}

		if time.Now().Unix() >= int64(claims["exp"].(float64)) {
			return errors.New("token is expired"), false, false, "", 0
		}

		id := claims["ID"].(string)
		queryArgs := make(map[string]interface{})
		queryArgs["id"] = id
		user, err := queryparserExt.User(queryArgs)

		var userIsAdminOrMaster = user.(model.User).Authentication == "admin" ||
			user.(model.User).Authentication == "master"
		if checkForAdmin && !userIsAdminOrMaster {
			return errors.New("hey there, you are not the admin or a master"), false, false, "", 0
		}

		var dbPassword string
		sql := "select password from user where id = ?"
		row := mysql.Db.QueryRow(sql, claims["ID"].(string))
		err = mysql.QueryRowScan(row, &dbPassword)
		if err != nil {
			logger.Logger.Println(err)
			return errLoginMismatch, false, false, "", 0
		}

		// Given password is hashed password with bcrypt
		err = bcrypt.CompareHashAndPassword([]byte(claims["Password"].(string)), []byte(dbPassword))
		if err != nil {
			return errLoginMismatch, false, false, "", 0
		}

		if !_groupIDOk {
			group := getGroupOfUser(claims["ID"].(string))
			if group != nil {
				_groupID = int(group.ID)
			}
		}

		if strings.ToLower(id) == "master" {
			return nil, false, true, "", 0
		}

		return nil, userIsAdminOrMaster, false, id, int64(_groupID)
	}

	return errors.New("invalid token"), false, false, "", 0
}

// GetUserID : Get the user ID from the token
func GetUserID(tokenString string) (ID string, err error) {

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
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["iss"].(string) != "piccolo" &&
			claims["sub"].(string) != "Auth" {
			return "", errors.New("invalid token")
		}

		if time.Now().Unix() >= int64(claims["exp"].(float64)) {
			return "", errors.New("token is expired")
		}

		var dbPassword string
		sql := "select password from user where id = ?"
		row := mysql.Db.QueryRow(sql, claims["ID"].(string))
		err := mysql.QueryRowScan(row, &dbPassword)
		if err != nil {
			return "", errors.New("invalid token")
		}

		// Given password is hashed password with bcrypt
		err = bcrypt.CompareHashAndPassword([]byte(claims["Password"].(string)), []byte(dbPassword))
		if err != nil {
			return "", errors.New("invalid token")
		}

		return claims["ID"].(string), nil
	}

	return "", errors.New("invalid token")
}
