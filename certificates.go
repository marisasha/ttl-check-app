package ttlchecker

import (
	"time"
)

type Certificate struct {
	Id        int       `json:"-"`
	UserId    int       `json:"user_id" db:"user_id"`
	Url       string    `json:"url" db:"url"`
	ValidFrom time.Time `json:"valid_from" db:"valid_from"`
	ValidTo   time.Time `json:"valid_to" db:"valid_to"`
}
