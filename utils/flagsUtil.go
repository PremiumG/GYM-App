package utils

import (
	"GYM-App/db"
	"flag"
	"fmt"
)

type flagsVal struct {
	UserName  string
	Password  string
	resetFlag bool
}

func RunFlags() {
	var flagData flagsVal
	initFlags(&flagData)
	resetFlag(&flagData)
	newUserFlag(&flagData)
}

func initFlags(flagData *flagsVal) {
	flag.BoolVar(&flagData.resetFlag, "reset", false, "Type <true> to reset DB (Default is false)")
	flag.StringVar(&flagData.UserName, "username", "", "Username for new user")
	flag.StringVar(&flagData.Password, "pass", "", "pass for new user")
	flag.Parse()
}

func resetFlag(flagData *flagsVal) {

	if flagData.resetFlag {
		//call db reset
		panic("lol argument true dela")
	} else {
		return
	}
}

func newUserFlag(flagData *flagsVal) error {
	fmt.Print(flagData.UserName)
	if flagData.UserName == "" || flagData.Password == "" {
		fmt.Println("All fields (go run . -username=xxxxx -pass=xxxxx) must be provided")
		return nil
	}

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("User NOT created")

		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(flagData.UserName, flagData.Password)
	if err != nil {
		fmt.Println("User NOT created")
		return err
	}

	// err := user.Save()
	// if err != nil {
	// 	return err
	// }
	fmt.Println("User CREATED")
	return nil
}
