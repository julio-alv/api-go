package threads

import "database/sql"

type Thread struct {
	ID        int    `json:"-"`
	PublicID  string `json:"public_id"`
	Message   string `json:"message"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
}

func GetMany(db *sql.DB) ([]Thread, error) {
	rows, err := db.Query(`SELECT * FROM threads`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	threads := []Thread{}

	for rows.Next() {
		var t Thread
		if err := rows.Scan(&t.ID, &t.PublicID, &t.Message, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		threads = append(threads, t)
	}
	return threads, nil
}

func GetOne(db *sql.DB, publicID string) (*Thread, error) {
	row := db.QueryRow(`SELECT * FROM threads WHERE public_id=$1`, publicID)

	var t Thread
	if err := row.Scan(&t.ID, &t.PublicID, &t.Message, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}
	return &t, nil
}
