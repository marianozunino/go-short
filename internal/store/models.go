// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package store

import (
	"database/sql"
)

type Url struct {
	ID        int64
	Url       string
	Code      string
	Md5       string
	Hits      sql.NullInt64
	LastUsed  sql.NullTime
	CreatedAt sql.NullTime
}