package model

import (
	"errors"
	"fmt"
	bcrypto "gin/model/crypto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   string
	Password string
}

func init() {
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(User{})
}

func (u *User) LoggedIn() bool {
	return u.ID != 0
}

func Signup(userId, password string) (*User, error) {
	user := User{}
	Db.Where("user_id = ?", userId).First(&user)
	if user.ID != 0 {
		err := errors.New("同一名のUserIdが既に登録されています。")
		fmt.Println(err)
		return nil, err
	}

	encryptPw, err := bcrypto.PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return nil, err
	}
	user = User{UserId: userId, Password: encryptPw}
	Db.Create(&user)
	return &user, nil
}

func Login(userId, password string) (*User, error) {
	user := User{}
	Db.Where("user_id = ?", userId).First(&user)
	if user.ID == 0 {
		err := errors.New("UserIdが一致するユーザーが存在しません。")
		fmt.Println(err)
		return nil, err
	}

	err := bcrypto.CompareHashAndPassword(user.Password, password)
	if err != nil {
		fmt.Println("パスワードが一致しませんでした。：", err)
		return nil, err
	}

	return &user, nil
}

func GetOneUser(userId string) User {
	user := User{}
	Db.Where("user_id = ?", userId).First(&user)
	if user.ID == 0 {
		err := errors.New("UserIdが一致するユーザーが存在しません。")
		fmt.Println(err)
		return user
	}
	return user
}
