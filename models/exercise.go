package models

import (
	"GYM-App/db"
	"errors"
)

type Exercise struct {
	Name     string
	Weight   int64
	B_weight int64
	I_weight int64
	A_weight int64
	Img      string
}

func GetAllLatestExercises() ([]Exercise, error) {
	var exercises []Exercise
	var temp Exercise
	var id int64
	var name string
	var weight int64
	var B_weight int64
	var I_weight int64
	var A_weight int64
	var Img string

	query := "SELECT * FROM exercises"
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("no exercises in row")
	}
	//prizadeto v glavo to
	for row.Next() {
		if err := row.Scan(&id, &name, &weight, &B_weight, &I_weight, &A_weight, &Img); err != nil {
			return nil, errors.New("no exercises")
		}
		temp = Exercise{
			Name:     name,
			Weight:   weight,
			B_weight: B_weight,
			I_weight: I_weight,
			A_weight: A_weight,
			Img:      Img,
		}

		exercises = append(exercises, temp)

	}

	return exercises, nil

}

func RemoveExerciseByName(name string) error {

	query := "DELETE FROM exercises WHERE name = (?)"

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

func AddExercise(exerciseData Exercise, img_path string) error {

	query := "INSERT INTO exercises(name, current_weight, b_weight, i_weight, a_weight, img_path) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exerciseData.Name, exerciseData.Weight, exerciseData.B_weight, exerciseData.I_weight, exerciseData.A_weight, img_path)

	if err != nil {
		return err
	}
	return nil

}

func UpdateExerciseByName(name string, newKG int64) error {

	query := "UPDATE exercises SET current_weight = (?) WHERE name = (?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newKG, name)
	if err != nil {
		return err
	}
	return nil

}

func GetAllWeightsEX() (int64, error) {
	var weight int64
	var all int64
	query := "SELECT current_weight FROM exercises"
	row, err := db.DB.Query(query)
	if err != nil {
		return 0, errors.New("no exercises in row")
	}
	//prizadeto v glavo to

	i := 0
	all = 0
	for row.Next() {
		if err := row.Scan(&weight); err != nil {
			return 0, errors.New("no exercises")
		}

		i++
		all += weight

	}

	return all, nil

}
