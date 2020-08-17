package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aiman-zaki/go_justcall/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-pg/pg/v9"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
	RoleID   int64
	Role     *Role `pg:"fk:role_id"`

	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	AcessToken   string    `json:"access_token" pg:"-"`
	RefreshToken string    `json:"refresh_token" pg:"-"`
}

type UserWrapper struct {
	Single User
	Array  []User
}

func (tw *UserWrapper) Login() (error, int16) {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	plainPassword := tw.Single.Password
	count, err := db.Model(&User{}).
		Where("email = ?", tw.Single.Email).
		Count()
	if err != nil {
		return err, 0
	}
	if count > 0 {
		err := db.Model(&tw.Single).Where("email = ?", tw.Single.Email).Select()
		if err != nil {
			return err, 0
		}
		valid := tw.ComparePasswords(tw.Single.Password, []byte(plainPassword))
		if !valid {
			return errors.New("Invalid Credential"), 401
		}
		err = db.Model(&tw.Single).
			Where("email = ?", tw.Single.Email).
			Select()
		tw.GenerateToken()
		if err != nil {
			return err, 0
		}
	} else {
		return errors.New("No Data"), 0
	}
	return nil, 0
}

func (tw *UserWrapper) Register() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	db.AddQueryHook(services.DbLogger{})

	count, err := db.Model(&User{}).
		Where("email = ?", tw.Single.Email).
		Where("username = ?", tw.Single.Username).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Account Existed")
	}
	hashed := tw.HashAndSalt([]byte(tw.Single.Password))
	tw.Single.Password = hashed
	tw.Single.DateCreated = time.Now()
	tw.Single.DateUpdated = time.Now()
	tw.Single.RoleID = 2
	err = db.Insert(&tw.Single)
	if err != nil {
		return err
	}
	return nil
}
func (tw *UserWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&tw.Single).Select()
	if err != nil {
		return err
	}
	return nil
}

func (tw *UserWrapper) ReadProfile() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&tw.Single).WherePK().Select()
	if err != nil {
		return err
	}
	return nil
}

func (tw *UserWrapper) Update() error {

	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})

	defer db.Close()
	if tw.Single.Password != "" {
		hashed := tw.HashAndSalt([]byte(tw.Single.Password))
		tw.Single.Password = hashed
		_, err := db.Model(&tw.Single).WherePK().Set("name = ?", tw.Single.Name).Set("username = ?", tw.Single.Username).Set("email = ?", tw.Single.Email).Set("phone = ?", tw.Single.Phone).Set("password = ?", tw.Single.Password).Update()
		if err != nil {
			return err
		}
	} else {
		_, err := db.Model(&tw.Single).WherePK().Set("name = ?", tw.Single.Name).Set("username = ?", tw.Single.Username).Set("email = ?", tw.Single.Email).Set("phone = ?", tw.Single.Phone).Update()
		if err != nil {
			return err
		}

	}

	return nil
}
func (tw *UserWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}

func TokenSetting() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte("NcRfUjXn2r5u8x/A?D(G+KbPdSgVkYp3"), nil)
}

// GenerateToken : Generate JWT Token
func (aw *UserWrapper) GenerateToken() {
	claims := jwt.MapClaims{"user": aw.Single.ID}
	tokenAuth := TokenSetting()

	duration, err := time.ParseDuration("120h")
	if err != nil {
		return
	}
	jwtauth.SetExpiryIn(claims, duration)
	_, tokenString, _ := tokenAuth.Encode(claims)
	duration, err = time.ParseDuration("24h")
	if err != nil {
		return
	}
	jwtauth.SetExpiryIn(claims, duration)
	_, refreshToken, _ := tokenAuth.Encode(claims)

	aw.Single.AcessToken = tokenString
	aw.Single.RefreshToken = refreshToken
}

// HashAndSalt : generate hashed password
func (auth UserWrapper) HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func (auth UserWrapper) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (auth *UserWrapper) RefreshToken() error {
	tokenAuth := TokenSetting()
	t, err := tokenAuth.Decode(auth.Single.RefreshToken)
	if err != nil {
		return err
	}
	email := t.Claims.(jwt.MapClaims)["user"]
	fmt.Println()
	auth.Single.Email = email.(string)
	auth.GenerateToken()
	return nil
}
