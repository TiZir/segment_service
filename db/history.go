package db

import (
	"time"
)

type History struct {
	IdUser      int    `json:"id_user"`
	NameSegment string `json:"name_segment"`
	TSadd       []byte `json:"time_add"`
	TSdel       []byte `json:"time_del"`
}

func InsertHistory(compliance Compliance) error {
	database, err := GetDB()
	if err != nil {
		return err
	}
	_, err = database.Exec(
		"INSERT INTO history (id_user, name_segment, ts_add) "+
			"VALUES (?, ?, ?);",
		compliance.IdUser,
		compliance.NameSegment,
		time.Now(),
	)
	return err
}

// func DeleteHistory(compliance Compliance) error {
// 	database, err := GetDB()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = database.Exec("TRUNCATE TABLE history;")
// 	return err
// }

func UpdateHistory(compliance Compliance, flag bool) error {
	// true - update add
	// false - update del
	database, err := GetDB()
	if err != nil {
		return err
	}
	if flag {
		_, err = database.Exec("UPDATE history "+
			"SET `ts_add` = ?, `ts_del` = NULL "+
			"WHERE id_user = ? AND name_segment = ?;",
			time.Now(),
			compliance.IdUser,
			compliance.NameSegment,
		)
	} else {
		_, err = database.Exec("UPDATE history "+
			"SET `ts_del` = ? "+
			"WHERE id_user = ? AND name_segment = ?;",
			time.Now(),
			compliance.IdUser,
			compliance.NameSegment,
		)
	}
	return err
}

func SelectHistory() ([]History, error) {
	history := []History{}
	database, err := GetDB()
	if err != nil {
		return history, err
	}

	rows, err := database.Query("SELECT id_user, name_segment, ts_add, ts_del FROM history;")
	if err != nil {
		return history, err
	}
	defer rows.Close()

	for rows.Next() {
		var h History
		err := rows.Scan(&h.IdUser, &h.NameSegment, &h.TSadd, &h.TSdel)
		if err != nil {
			return history, err
		}
		history = append(history, h)
	}
	return history, nil
}

func ExistDataInHistory(compliance Compliance) (bool, error) {
	database, err := GetDB()
	if err != nil {
		return false, err
	}
	rows, err := database.Query(
		"SELECT * "+
			"FROM history "+
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
