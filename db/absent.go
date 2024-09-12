package db

import (
	"context"
	"database/sql"
)

const getAbsentUsers = `-- name: GetAbsentUsers :many
SELECT 
    u.user_id,
    u.first_name,
    u.last_name,
    u.phone_number,
    e.entry_time
FROM 
    users u
LEFT JOIN 
    entrance e ON u.user_id = e.user_id 
    AND e.entry_time BETWEEN ? AND ? 
WHERE 
    e.entry_time IS NULL
`

type GetAbsentUsersRow struct {
	UserID      int64         `json:"user_id"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	PhoneNumber string        `json:"phone_number"`
	EntryTime   sql.NullInt64 `json:"entry_time"`
}

func (q *Queries) GetAbsentUsers(ctx context.Context, startTime, endTime int64) ([]GetAbsentUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getAbsentUsers, startTime, endTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []GetAbsentUsersRow
	for rows.Next() {
		var i GetAbsentUsersRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.EntryTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
