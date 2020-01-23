package models

import (
	"backend/utils"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token is a JWT token that will returned to the frontend
type Token struct {
	UserID   uint
	UserName string
	Exp      int64
	jwt.StandardClaims
}

// Account represents a user to be saved in the DB
type Account struct {
	gorm.Model
	Email    string
	UserName string
	Password string
	Token    string
}

// Validate account that have yet to be created. This is called when a new account object has to be created
func (acc *Account) Validate() (map[string]interface{}, bool) {
	resp := make(map[string]interface{})
	if !strings.Contains(acc.Email, "@") {
		return utils.Message(false, "Email Address Required"), false
	}
	if len(acc.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		resp["error"] = "Connection Error. Please retry"
		return resp, false
	}
	//Email must be unique
	if temp.Email != "" {
		resp["error"] = "Email already in use"
		return resp, false
	}
	//Username must be unique
	if temp.UserName != "" {
		resp["error"] = "username already is use"
		return resp, false
	}
	return utils.Message(false, "Requirement passed"), true
}

//ValidateLogin is used to validate accounts that are attempting to login
func (acc *Account) ValidateLogin() (map[string]interface{}, bool) {
	resp := make(map[string]interface{})
	// Get the account object from the DB
	// Allocate a temp account object
	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		resp["error"] = "Connection Error. Please retry"
		return resp, false
	}
	// let's validate the old password
	// temp is the record that exists in the database. The password is hashed using bcrypt earlier during creation.
	err = bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(acc.Password))
	if err != nil {
		resp["error"] = "Invalid Password"
		return resp, false
	}
	// Lets change the password
	return resp, true
}

//Create account
func (acc *Account) Create() (map[string]interface{}, bool) {
	response, ok := acc.Validate()
	if !ok {
		return response, ok
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	acc.Password = string(hashedPassword)
	//stores the account into the database
	GetDB().Create(acc)

	if acc.ID <= 0 {
		response["error"] = "Failed to create new account dew to database error of ID less than 0"
		return response, false
	}
	//create a new JWT token
	tk := &Token{UserID: acc.ID, UserName: acc.Email, Exp: getExpiryDate()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString
	acc.Password = ""

	response = utils.Message(true, "Account has been created")
	log.Print("Succesfully created new account for user: " + acc.Email)
	response["account"] = acc
	return response, true
}

// ChangePassword will change the password of the account
func (acc *Account) ChangePassword(newPassword string) (map[string]interface{}, bool) {
	response, ok := acc.ValidateLogin()
	if !ok {
		return response, ok
	}
	// If ok
	// Get new hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	// Update the password field only
	GetDB().Model(acc).Update("Password", string(hashedPassword))
	//create JWT TOKEN
	tk := &Token{UserID: acc.ID, UserName: acc.Email, Exp: getExpiryDate()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString
	fmt.Println("User: " + acc.Email + " has succesfully changed password")
	resp := utils.Message(true, "Succesfully changed password")
	resp["account"] = acc
	return resp, ok
}

// Login : Will attemp a login
func (acc *Account) Login() (map[string]interface{}, bool) {
	response, ok := acc.ValidateLogin()
	if !ok {
		return response, ok
	}
	tk := &Token{UserID: acc.ID, UserName: acc.Email, Exp: getExpiryDate()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Password = ""
	acc.Token = tokenString
	fmt.Println("User: " + acc.Email + " has succesfully logged in.")
	resp := utils.Message(true, "Succesfully logged in")
	resp["account"] = acc
	return resp, ok
}

func getExpiryDate() int64 {
	start := time.Now()
	end := start.Add(time.Second * 15)
	return end.Unix()
}
