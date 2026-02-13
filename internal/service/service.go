package service

import (
	"github.com/marisasha/ttl-check-app/internal/models"
	"github.com/marisasha/ttl-check-app/internal/repository"
)

type Authorization interface {
	CreateUser(user *models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Certificate interface {
	GetAllCertificates(userId int) (*[]models.CertificateResponse, error)
	GetCertificateById(certificateId int) (*models.CertificateResponse, error)
	AddCertificate(certificate *models.Certificate) error
	DeleteCertificate(certificateId int) error
	CheckCertificate(inputUrl string) (*models.CertificateInfo, error)
}

type Service struct {
	Authorization
	Certificate
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Certificate:   NewCertificateService(repos.Certificate),
	}
}
