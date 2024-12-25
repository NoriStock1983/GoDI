package main

import "fmt"

// interface
type UserInterface interface {
	getUserCd() string
}

// Interfaceを実装する構造体
type Usercd struct{}

// Interfaceのメソッドを実装（Usercd構造体）
func (u Usercd) getUserCd() string {
	return "a0006802"
}

// 構造体
type User struct {
	usercd    string
	user_l_nm string
	user_f_nm string
}

func main() {
	var user UserInterface

	user = &Usercd{}
	fmt.Println(user.getUserCd())

}
