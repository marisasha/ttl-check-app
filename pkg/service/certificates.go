package service

import (
	"crypto/tls"
	"errors"
	"math"
	"net"
	"net/url"
	"time"

	ttlchecker "github.com/marisasha/ttl-check-app"
	"github.com/marisasha/ttl-check-app/pkg/repository"
)

type CertificateService struct {
	repos repository.Certificate
}

func NewCertificateService(repos repository.Certificate) *CertificateService {
	return &CertificateService{repos: repos}
}

func (s *CertificateService) AddCertificate(certificate *ttlchecker.Certificate) error {

	validInfo, err := checkTTL(certificate.Url)
	if err != nil {
		return err
	}
	certificate.ValidFrom = validInfo.ValidFrom
	certificate.ValidTo = validInfo.ValidTo

	return s.repos.AddCertificate(certificate)
}

func (s *CertificateService) GetAllCertificates(userId int) (*[]ttlchecker.CertificateResponse, error) {
	certificatesWithoutDaysleft, err := s.repos.GetAllCertificates(userId)
	if err != nil {
		return nil, err
	}

	certificates := make([]ttlchecker.CertificateResponse, len(*certificatesWithoutDaysleft))

	for i, c := range *certificatesWithoutDaysleft {
		certificates[i] = ttlchecker.CertificateResponse{
			Id:        c.Id,
			UserId:    c.UserId,
			Url:       c.Url,
			ValidFrom: c.ValidFrom,
			ValidTo:   c.ValidTo,
			DaysLeft:  getDaysLeft(c.ValidTo),
		}
	}

	return &certificates, nil
}

func (s *CertificateService) GetCertificateById(certificateId int) (*ttlchecker.CertificateResponse, error) {
	certificatesWithoutDaysleft, err := s.repos.GetCertificateById(certificateId)
	if err != nil {
		return nil, err
	}

	certificate := &ttlchecker.CertificateResponse{
		Id:        certificatesWithoutDaysleft.Id,
		UserId:    certificatesWithoutDaysleft.UserId,
		Url:       certificatesWithoutDaysleft.Url,
		ValidFrom: certificatesWithoutDaysleft.ValidFrom,
		DaysLeft:  getDaysLeft(certificatesWithoutDaysleft.ValidTo),
	}

	return certificate, nil
}

func (s *CertificateService) DeleteCertificate(certificateId int) error {
	return s.repos.DeleteCertificate(certificateId)
}

func (s *CertificateService) CheckCertificate(inputURL string) (*ttlchecker.CertificateInfo, error) {
	return checkTTL(inputURL)
}

func checkTTL(inputURL string) (*ttlchecker.CertificateInfo, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return nil, errors.New("invalid url")
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}
	if parsedURL.Scheme != "https" {
		return nil, errors.New("only https scheme is supported")
	}

	host := parsedURL.Hostname()
	if host == "" {
		return nil, errors.New("invalid host")
	}

	port := parsedURL.Port()
	if port == "" {
		port = "443"
	}

	address := net.JoinHostPort(host, port)

	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", address, &tls.Config{
		ServerName: host,
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(10 * time.Second))

	state := conn.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		return nil, errors.New("certificates not founded")
	}

	cert := state.PeerCertificates[0]

	daysLeft := getDaysLeft(cert.NotAfter)

	return &ttlchecker.CertificateInfo{
		URL:       inputURL,
		ValidFrom: cert.NotBefore,
		ValidTo:   cert.NotAfter,
		DaysLeft:  daysLeft,
	}, nil
}

func getDaysLeft(ValidTo time.Time) int {
	remaining := time.Until(ValidTo)
	daysLeft := int(math.Ceil(remaining.Hours() / 24))
	if daysLeft < 0 {
		daysLeft = 0
	}
	return daysLeft
}
