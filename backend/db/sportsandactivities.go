package db

import (
	"bufio"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func InitSportsActivities() {

	// Create table if not exists
	// Open file
	file, err := os.Open("list of activities.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Prepare insert statement
	insertSQL := `
	INSERT INTO sports_activities (activity_name)
	VALUES ($1)
	ON CONFLICT (activity_name) DO NOTHING;
	`
	stmt, err := DB.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Read + insert
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Remove leading "number. "
		if _, activity, ok := strings.Cut(line, ". "); ok {
			activity = strings.TrimSpace(activity)

			if _, err := stmt.Exec(activity); err != nil {
				continue
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetSportsActivities() ([]string, error) {
	query := `SELECT activity_name FROM sports_activities ORDER BY activity_name`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []string
	for rows.Next() {
		var activity string
		if err := rows.Scan(&activity); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return activities, nil
}
