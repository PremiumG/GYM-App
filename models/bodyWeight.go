package models

import (
	"GYM-App/db"
	"errors"
	"fmt"
	"time"
)

func InsertWeight(kg string) error {
	if kg == "" {
		return errors.New("empty string")
	}
	eventDate := time.Now()
	query := "INSERT INTO bodyweights(date, weight_kg) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventDate, kg)
	fmt.Println(err)

	if err != nil {
		return err
	}
	return nil

}

type WeightEntry struct {
	ID     string
	Date   string
	Weight string
}

func GetAllWeightDB() ([]WeightEntry, error) {
	var allWeights []WeightEntry
	var oneWeight WeightEntry
	query := "SELECT id, date, weight_kg FROM bodyweights"
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("no ytLinks")
	}

	for row.Next() {
		if err := row.Scan(&oneWeight.ID, &oneWeight.Date, &oneWeight.Weight); err != nil {
			return nil, errors.New("no ytLinks")
		}
		allWeights = append(allWeights, oneWeight)
	}

	return allWeights, nil

}

func RemoveWeightID(id string) error {
	// num, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(num)

	query := "DELETE FROM bodyweights WHERE id = (?)"

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

//CODE BELLOW WOULD BE USED FOR DB IF WE WOULD HAVE MUTIPLE USERS.
// type BodyWeight struct {
// 	Id     int64
// 	UserId int64
// 	Date   time.Time
// 	Weight int64
// }

// func (bw BodyWeight) InsertWeight() error {
// 	query := "INSERT INTO bodyweights(user_id, date, weight_kg) VALUES (?, ?, ?)"
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(bw.UserId, bw.Date, bw.Weight)
// 	return err
// }

// func GetWeight(userId int64) ([]BodyWeight, error) {
// 	query := "SELECT * FROM bodyweights WHERE user_id = ?"
// 	var weights []BodyWeight

// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	rows, err := db.DB.Query(query, userId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer stmt.Close()

// 	for rows.Next() {
// 		var oneWeight BodyWeight
// 		err := rows.Scan(&oneWeight.Id, &oneWeight.UserId, &oneWeight.Date, &oneWeight.Weight)
// 		if err != nil {
// 			return nil, err
// 		}
// 		weights = append(weights, oneWeight)
// 	}
// 	return weights, nil
// }
