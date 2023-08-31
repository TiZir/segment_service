package db

import (
	"database/sql"
)

type Compliance struct {
	IdUser      int    `json:"id_user"`
	NameSegment string `json:"name_segment"`
}

func InsertCompliance(data []string, id int) error {
	database, err := GetDB()
	if err != nil {
		return err
	}
	for key, _ := range data {
		var compliance Compliance
		var flag bool
		compliance.IdUser = id
		compliance.NameSegment = data[key]
		if flag, err = ExistDataInCompliance(compliance); !flag && err == nil {
			_, err = database.Exec("INSERT INTO compliance (id_user, name_segment) VALUES (?, ?);",
				compliance.IdUser, compliance.NameSegment)
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}

		//history
		if flag, err = ExistDataInHistory(compliance); !flag && err == nil {
			err := InsertHistory(compliance)
			if err != nil {
				return err
			}
		} else {
			if err != nil {
				return err
			}
			err := UpdateHistory(compliance, true)
			if err != nil {
				return err
			}
		}

	}
	return err
}

func DeleteCompliance(data []string, id int) error {
	database, err := GetDB()
	if err != nil {
		return err
	}
	for key, _ := range data {
		var compliance Compliance
		var flag bool
		compliance.IdUser = id
		compliance.NameSegment = data[key]
		if flag, err = ExistDataInCompliance(compliance); flag && err == nil {
			_, err = database.Exec("DELETE FROM compliance WHERE id_user = ? AND name_segment = ?;",
				compliance.IdUser, compliance.NameSegment)
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}

		//history
		if flag, err = ExistDataInHistory(compliance); flag && err == nil {
			err := UpdateHistory(compliance, false)
			if err != nil {
				return err
			}
		}
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
	rows, err := database.Query("SELECT * FROM compliance WHERE id_user = ? ORDER BY name_segment;", id)
	if err != nil {
		return compliances, err
	}
	defer rows.Close()

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

	rows, err := database.Query("SELECT * FROM compliance ORDER BY id_user;")
	if err != nil {
		return compliances, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Compliance
		err := rows.Scan(&c.IdUser, &c.NameSegment)
		if err != nil {
			if err == sql.ErrNoRows {
				return compliances, err
			}
			return compliances, err
		}
		compliances = append(compliances, c)
	}

	return compliances, nil
}

func ExistDataInCompliance(compliance Compliance) (bool, error) {
	database, err := GetDB()
	if err != nil {
		return false, err
	}
	rows, err := database.Query(
		"SELECT * "+
			"FROM compliance "+
			"WHERE id_user = ? AND name_segment = ? "+
			"ORDER BY id_user;",
		compliance.IdUser,
		compliance.NameSegment,
	)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
