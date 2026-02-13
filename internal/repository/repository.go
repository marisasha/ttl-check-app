package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/marisasha/ttl-check-app/internal/models"
)

type Authorization interface {
	CreateUser(user *models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Certificate interface {
	GetAllCertificates(userId int) (*[]models.Certificate, error)
	GetCertificateById(certificateId int) (*models.Certificate, error)
	AddCertificate(certificate *models.Certificate) error
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
