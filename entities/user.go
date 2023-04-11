package entities

import (
	"context"
	"fmt"
	"html"
	"log"
	"shabo_edge/utils"
	"strings"

	"github.com/edgedb/edgedb-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         edgedb.OptionalUUID
	EthAddress string `edgedb:"ethAddress"`
	UserName   string `edgedb:"userName"`
	Email      string `edgedb:"email"`
	FoAddress  string `edgedb:"foAddress"`
	XcpAddress string `edgedb:"xcpAddress"`
	HasType    Type   `edgedb:"hasType"`
	Password   string `edgedb:"password"`
}

func GetAllUsers() ([]User, error) {
	var users []User
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
		return users, err
	}
	defer client.Close()

	query := "select User{email, ethAddress, userName, foAddress}"
	err = client.Query(ctx, query, &users)
	if err != nil {
		log.Fatalln(err)
		return users, err
	}
	return users, err
}

func GetUser(property string, value string, user *User) error {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer client.Close()

	// var user User
	query := fmt.Sprintf("select distinct User{email, ethAddress, userName, foAddress} filter .'%s'='%s'", property, value)
	err = client.Query(ctx, query, &user)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	var er error
	return er
}

func CreateUsers(users *[]User) error {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var er error

	var query string
	for _, user := range *users {
		query += fmt.Sprintf(
			"insert User{email='%s', userName='%s', ethAddress='%s', foAddress='%s', xcpAddress='%s'}unless conflict on .email", user.Email, user.UserName, user.EthAddress, user.FoAddress, user.XcpAddress)
	}

	var result []User
	if len(*users) > 1 {
		err = client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
			e := tx.Execute(ctx, query)
			return e
		})
	} else {
		err = client.Query(ctx, query, &result)
	}

	if err != nil {
		fmt.Println(err)
	}
	return er
}

func (u *User) SaveUser() (*User, error) {
	user := User{}
	err := GetUser("email", *&u.Email, &user)
	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

func (u *User) BeforeSave() error {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPw)
	// remove spaces
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error
	u := User{}
	err = GetUser("email", username, &u)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.EthAddress)
	if err != nil {
		return "", err
	}

	return token, nil
}
