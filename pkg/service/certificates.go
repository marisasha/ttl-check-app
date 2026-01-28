package service

import (
	ttlchecker "github.com/marisasha/ttl-check-app"
	"github.com/marisasha/ttl-check-app/pkg/repository"
)

type CertificateService struct {
	repos repository.Certificate
}

func NewCertificateService(repos repository.Certificate) *CertificateService {
	return &CertificateService{repos: repos}
}

func (s *CertificateService) AddCertificate(certificate ttlchecker.Certificate) (int, error) {
	return s.repos.AddCertificate(certificate)
}

func (s *CertificateService) GetAllCertificates(userId int) ([]ttlchecker.Certificate, error) {
	return s.repos.GetAllCertificates(userId)
}

func (s *CertificateService) GetCertificateById(certificateId int) (ttlchecker.Certificate, error) {
	return s.repos.GetCertificateById(certificateId)
}

func (s *CertificateService) DeleteCertificate(certificateId int) error {
	return s.repos.DeleteCertificate(certificateId)
}
