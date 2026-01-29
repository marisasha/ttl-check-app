package ttlchecker

import (
	"time"
)

type Certificate struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id" db:"user_id"`
	Url       string    `json:"url" db:"url"`
	ValidFrom time.Time `json:"valid_from" db:"valid_from"`
	ValidTo   time.Time `json:"valid_to" db:"valid_to"`
}

type CertificateInfo struct {
	URL       string    `json:"url"`
	ValidFrom time.Time `json:"valid_from"`
	ValidTo   time.Time `json:"valid_to"`
	DaysLeft  int       `json:"days_left"`
}

type CertificateResponse struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id" `
	Url       string    `json:"url" `
	ValidFrom time.Time `json:"valid_from" `
	ValidTo   time.Time `json:"valid_to" `
	DaysLeft  int       `json:"days_left" `
}
