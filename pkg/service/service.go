package service

import (
	ttlchecker "github.com/marisasha/ttl-check-app"
	"github.com/marisasha/ttl-check-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user ttlchecker.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Certificate interface {
	AddCertificate(certificate ttlchecker.Certificate) (int, error)
	GetAllCertificates(userId int) ([]ttlchecker.Certificate, error)
	GetCertificateById(certificateId int) (ttlchecker.Certificate, error)
	DeleteCertificate(certificateId int) error
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
