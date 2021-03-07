package customer

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID        int       `db:"id"`
	FullName  string    `db:"full_name"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Avatar    string    `db:"avatar"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CustomerScan struct {
	ID        sql.NullInt64  `db:"id"`
	FullName  sql.NullString `db:"full_name"`
	Username  sql.NullString `db:"username"`
	Email     sql.NullString `db:"email"`
	Password  sql.NullString `db:"password"`
	Avatar    sql.NullString `db:"avatar"`
	Status    sql.NullInt64  `db:"status"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

func (c *Customer) FromScan(scan CustomerScan) {
	c.ID = int(scan.ID.Int64)
	c.FullName = scan.FullName.String
	c.Username = scan.Username.String
	c.Email = scan.Email.String
	c.Password = scan.Password.String
	c.Avatar = scan.Avatar.String

	c.CreatedAt = scan.CreatedAt.Time
	c.UpdatedAt = scan.UpdatedAt.Time
	if scan.Status.Valid {
		if scan.Status.Int64 == 1 {
			c.Status = true
		} else {
			c.Status = false
		}
	}
}
