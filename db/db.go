package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic("No DB connection")
	}
	// defer DB.Close()
	// lol := DB.Ping()
	// fmt.Println(lol.Error())
	DB.SetMaxOpenConns(0)
	DB.SetMaxIdleConns(5)

	creteTables()
	firstInsert()
}

func creteTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createUsersWeightTable := `
		CREATE TABLE IF NOT EXISTS bodyweights (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date DATE NOT NULL,
			weight_kg INTEGER
		)
		`
	_, err = DB.Exec(createUsersWeightTable)
	if err != nil {
		panic("Could not create users table")
	}

	createChallangeWeightTable := `
		CREATE TABLE IF NOT EXISTS challange (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name text NOT NULL
		)
		`
	_, err = DB.Exec(createChallangeWeightTable)
	if err != nil {
		panic("Could not create challanges table")
	}

	createYTLinkTable := `
		CREATE TABLE IF NOT EXISTS ytLinks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title text NOT NULL,
			link text NOT NULL
		)
		`
	_, err = DB.Exec(createYTLinkTable)
	if err != nil {
		panic("Could not create ytLinks table")
	}

	createExercise := `
		CREATE TABLE IF NOT EXISTS exercises (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			current_weight INTEGER NOT NULL,
			b_weight INTEGER NOT NULL,
			i_weight INTEGER NOT NULL,
			a_weight INTEGER NOT NULL,
			img_path TEXT NOT NULL
		)
		`
	_, err = DB.Exec(createExercise)
	if err != nil {
		panic("Could not create ytLinks table")
	}

	// x := `
	// DROP TABLE IF EXISTS exercises
	// `
	// _, err = DB.Exec(x)
	// if err != nil {
	// 	panic("Could not")
	// }

	// createEventsTable := `
	// CREATE TABLE IF NOT EXISTS events (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name TEXT NOT NULL,
	// 	description TEXT NOT NULL,
	// 	location TEXT NOT NULL,
	// 	dateTime DATETIME NOT NULL,
	// 	user_id INTEGER,
	// 	FOREIGN KEY(user_id) REFERENCES users(id)
	// )
	// `

	// _, err = DB.Exec(createEventsTable)
	// if err != nil {
	// 	panic("Could not create events table")
	// }

	// createRegistrationsTable := `
	// CREATE TABLE IF NOT EXISTS registrations (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	event_id INTEGER NOT NULL,
	// 	user_id INTEGER,
	// 	FOREIGN KEY(event_id) REFERENCES events(id)
	// 	FOREIGN KEY(user_id) REFERENCES users(id)

	// )
	// `
	// _, err = DB.Exec(createRegistrationsTable)
	// if err != nil {
	// 	panic("Could not create events table")
	// }
}

func firstInsert() error {
	yt := map[string]string{
		"Minimalist Workout Plan": "https://www.youtube.com/watch?v=eMjyvIQbn9M&t=436s",
		"QUAD Exercises":          "https://www.youtube.com/watch?v=kIXcoivzGf8",
		"TRICEPS Exercises ":      "https://www.youtube.com/watch?v=OpRMRhr0Ycc",
	}

	for key, value := range yt {
		query := "INSERT INTO ytLinks(title, link) VALUES (?, ?)"

		stmt, err := DB.Prepare(query)

		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(key, value)

		if err != nil {
			return err
		}
	}

	challange := []string{"100 Pushups", "50 Pullups", "200 Squats", "1-Minute Plank Hold", "5K Run", "100 Burpees", "100 Jumping Jacks", "5-Minute Wall Sit", "1-Minute Mountain Climbers", "50 Dips", "100 Lunges", "500m Row", "1-Minute Jump Rope", "50 Box Jumps", "10km Bike Ride", "50 Sit-ups", "1-Minute Handstand Hold", "100 Russian Twists", "10 Sprints", "1-Minute Side Plank (each side)", "100 Calf Raises", "50 V-Ups", "30 Tuck Jumps", "50 Bulgarian Split Squats", "100 Deadlifts (bodyweight)", "1-Minute Hollow Hold", "1-Minute Superman Hold", "100 Shoulder Taps", "50 Step-ups", "200m Farmer's Carry", "50 Chest-to-Bar Pullups", "1-Minute L-Sit Hold", "100 Kettlebell Swings", "50 Thrusters", "1-Minute Box Shuffle", "200m Sprint Row", "30 Second Push-Up Hold", "10 Rope Climbs", "100m Walking Lunges", "50 Man Makers", "100 Bench Presses", "10km Row", "1-Minute Flutter Kicks", "30 Dead Hang", "100 Bicycle Crunches", "5-Minute Jump Rope", "500m Swim", "100m Bear Crawl", "50 Turkish Get-Ups", "100 Weighted Squats"}

	for _, value := range challange {
		query := "INSERT INTO challange(name) VALUES (?)"

		stmt, err := DB.Prepare(query)

		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(value)

		if err != nil {
			return err
		}
	}

	return nil

}
