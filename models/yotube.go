package models

import (
	"GYM-App/db"
	"errors"
)

type YT struct {
	Title string
	Link  string
}

func GetRandomYTLink() YT {
	var ytName YT
	query := "SELECT title, link FROM ytLinks ORDER BY RANDOM() LIMIT 1"
	row := db.DB.QueryRow(query)
	err := row.Scan(&ytName.Title, &ytName.Link)
	if err != nil {
		return ytName
	}
	return ytName

}

func SetYTLink(title string, link string) error {
	if title == "" {
		return errors.New("empty string")
	}

	query := "INSERT INTO ytLinks(title,link) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, link)

	if err != nil {
		return err
	}
	return nil

}

func GetAllYTLink() ([]string, error) {
	var ytLinksTable []string
	var ytID string
	var ytTitle string
	query := "SELECT id, title FROM ytLinks"
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("no ytLinks")
	}

	for row.Next() {
		if err := row.Scan(&ytID, &ytTitle); err != nil {
			return nil, errors.New("no ytLinks")
		}
		result := ytID + " " + ytTitle
		ytLinksTable = append(ytLinksTable, result)
	}

	return ytLinksTable, nil

}

func RemoveYTLinkByID(id string) error {
	// num, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(num)

	query := "DELETE FROM ytLinks WHERE id = (?)"

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
