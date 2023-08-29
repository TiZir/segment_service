package db

type Compliance struct {
	IdUser      int    `json:"id_user"`
	NameSegment string `json:"name_segment"`
}

func InsertCompliance(compliances []Compliance) error {
	database, err := GetDB()
	if err != nil {
		return err
	}
	for key, _ := range compliances {
		_, err = database.Exec("INSERT INTO compliance (id_user, name_segment) VALUES (?, ?)",
			compliances[key].IdUser, compliances[key].NameSegment)
		if err != nil {
			return err
		}
	}
	return err
}

func DeleteCompliance(compliances []Compliance) error {
	database, err := GetDB()
	if err != nil {
		return err
	}
	for key, _ := range compliances {
		_, err = database.Exec("DELETE FROM compliance WHERE id = ? AND name = ",
			compliances[key].IdUser, compliances[key].NameSegment)
		if err != nil {
			return err
		}
	}
	return err
}

func SelectComplianceById(id int) ([]Compliance, error) {
	compliances := []Compliance{}
	database, err := GetDB()
	if err != nil {
		return compliances, err
	}
	rows, err := database.Query("SELECT * FROM compliance WHERE id = ?", id)
	if err != nil {
		return compliances, err
	}
	for rows.Next() {
		var c Compliance
		err = rows.Scan(&c.IdUser, &c.NameSegment)
		if err != nil {
			return compliances, err
		}
		compliances = append(compliances, c)
	}
	return compliances, nil
}

func SelectCompliance() ([]Compliance, error) {
	compliances := []Compliance{}
	database, err := GetDB()
	if err != nil {
		return compliances, err
	}
	rows, err := database.Query("SELECT * FROM compliance")
	if err != nil {
		return compliances, err
	}
	for rows.Next() {
		var c Compliance
		err = rows.Scan(&c.IdUser, &c.NameSegment)
		if err != nil {
			return compliances, err
		}
		compliances = append(compliances, c)
	}
	return compliances, nil
}
