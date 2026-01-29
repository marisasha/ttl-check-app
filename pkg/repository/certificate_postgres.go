package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	ttlchecker "github.com/marisasha/ttl-check-app"
)

type CertificatePostgres struct {
	db *sqlx.DB
}

func NewCertificatePostgres(db *sqlx.DB) *CertificatePostgres {
	return &CertificatePostgres{db: db}
}

func (r *CertificatePostgres) AddCertificate(certificate *ttlchecker.Certificate) error {

	query := fmt.Sprintf("INSERT INTO %s (user_id,url,valid_from,valid_to) VALUES ($1, $2, $3, $4)", certificatesTable)
	_, err := r.db.Exec(query, certificate.UserId, certificate.Url, certificate.ValidFrom, certificate.ValidTo)
	if err != nil {
		return err
	}
	return nil
}

func (r *CertificatePostgres) GetAllCertificates(userId int) (*[]ttlchecker.Certificate, error) {

	var certificates []ttlchecker.Certificate

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", certificatesTable)
	err := r.db.Select(&certificates, query, userId)

	return &certificates, err

}

func (r *CertificatePostgres) GetCertificateById(certificateId int) (*ttlchecker.Certificate, error) {

	var certificates ttlchecker.Certificate

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", certificatesTable)
	err := r.db.Get(&certificates, query, certificateId)

	return &certificates, err

}

func (r *CertificatePostgres) DeleteCertificate(certificateId int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", certificatesTable)
	_, err := r.db.Exec(query, certificateId)
	if err != nil {
		return err
	}
	return nil

}
