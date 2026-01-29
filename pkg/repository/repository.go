package repository

import (
	"github.com/jmoiron/sqlx"
	ttlchecker "github.com/marisasha/ttl-check-app"
)

type Authorization interface {
	CreateUser(user *ttlchecker.User) (int, error)
	GetUser(username, password string) (ttlchecker.User, error)
}

type Certificate interface {
	GetAllCertificates(userId int) (*[]ttlchecker.Certificate, error)
	GetCertificateById(certificateId int) (*ttlchecker.Certificate, error)
	AddCertificate(certificate *ttlchecker.Certificate) error
	DeleteCertificate(certificateId int) error
}

type Repository struct {
	Authorization
	Certificate
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Certificate:   NewCertificatePostgres(db),
	}
}
