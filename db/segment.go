package db

import "database/sql"

type Segment struct {
	Name string `json:"name"`
}

func InsertSegment(segment Segment) error {
	database, err := GetDB()
	if err != nil {
		return err
	}
	_, err = database.Exec("INSERT INTO segment (name) VALUES (?);", segment.Name)
	return err
}

func DeleteSegment(segment Segment) error {
	database, err := GetDB()
	if err != nil {
		return err
	}
	_, err = database.Exec("DELETE FROM segment WHERE name = ?;", segment.Name)
	return err
}

func SelectSegmentByName(name string) (Segment, error) {
	var segment Segment
	database, err := GetDB()
	if err != nil {
		return segment, err
	}
	row := database.QueryRow("SELECT * FROM segment WHERE name = ?;", name)
	err = row.Scan(&segment.Name)
	if err != nil {
		return segment, err
	}
	return segment, nil
}

func SelectSegment(name []string) ([]Segment, error) {
	segments := []Segment{}
	database, err := GetDB()
	if err != nil {
		return segments, err
	}
	for key, _ := range name {
		var s Segment
		row := database.QueryRow("SELECT * FROM segment WHERE name = ?;", name[key])
		err = row.Scan(&s.Name)
		if err != nil {
			return segments, err
		}
		segments = append(segments, s)
	}
	return segments, nil
}

func SelectSegmentTest() ([]Segment, error) {
	segments := []Segment{}
	database, err := GetDB()
	if err != nil {
		return segments, err
	}

	rows, err := database.Query("SELECT name FROM segment;")
	if err != nil {
		return segments, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Segment
		err := rows.Scan(&s.Name)
		defer rows.Close()
		if err != nil {
			if err == sql.ErrNoRows {
				return segments, err
			}
			return segments, err
		}
		segments = append(segments, s)
	}

	return segments, nil
}
