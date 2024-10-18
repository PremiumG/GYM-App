package models

import (
	"GYM-App/db"
	"errors"
	"fmt"
	"strconv"
)

type ChallangeName struct {
	Name string
}

func GetRandomChallenge() string {
	var challengeName string
	query := "SELECT name FROM challange ORDER BY RANDOM() LIMIT 1"
	row := db.DB.QueryRow(query)
	err := row.Scan(&challengeName)
	if err != nil {
		return "No Challenge"
	}
	return challengeName

}

func SetRandomChallenge(name string) error {
	if name == "" {
		return errors.New("empty string")
	}

	query := "INSERT INTO challange(name) VALUES (?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name)

	if err != nil {
		return err
	}
	return nil

}

func GetAllChallenges() ([]string, error) {
	var chalangesTable []string
	var name string
	var id string
	query := "SELECT id, name FROM challange"
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("no Challange")
	}

	for row.Next() {
		if err := row.Scan(&id, &name); err != nil {
			return nil, errors.New("no Challange")
		}
		result := id + " " + name
		chalangesTable = append(chalangesTable, result)
	}

	return chalangesTable, nil

}

func RemoveChallengesByID(id string) error {
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(num)

	query := "DELETE FROM challange WHERE id = (?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil

}
