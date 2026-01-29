package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ttlchecker "github.com/marisasha/ttl-check-app"
)

type AddCertificateRequest struct {
	Url string `json:"url" binding:"required,url"`
}

// @Summary Добавление отслеживаемого сертификата
// @Tags certificate
// @Description Добавление нового сертификата
// @ID add-certificate
// @Accept json
// @Produce json
// @Param input body AddCertificateRequest true "Данные пользователя"
// @Security ApiKeyAuth
// @Router /api/certificates/add [post]
func (h *Handler) addCertificate(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input AddCertificateRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newCertificate := &ttlchecker.Certificate{
		UserId: userId,
		Url:    input.Url,
	}

	err = h.services.Certificate.AddCertificate(newCertificate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"message": "URL successfully added !",
	})

}

type getAllCertificatesResponse struct {
	Data []ttlchecker.CertificateResponse `json:"data"`
}

// @Summary Просмотр сертификатов
// @Tags certificate
// @Description Просмотр всех сертификатов
// @ID get-all-certificates
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/certificates [get]
func (h *Handler) getAllCertificates(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	certificates, err := h.services.Certificate.GetAllCertificates(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, getAllCertificatesResponse{
		Data: *certificates,
	})

}

type getCertificateResponse struct {
	Data ttlchecker.CertificateResponse `json:"data"`
}

// @Summary Просмотр сертификата
// @Tags certificate
// @Description Просмотр сертификата по id
// @ID get-certificate
// @Accept json
// @Produce json
// @Param id path int true "ID сертификата"
// @Security ApiKeyAuth
// @Router /api/certificates/{id} [get]
func (h *Handler) getCertificate(c *gin.Context) {

	certificateId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	certificate, err := h.services.Certificate.GetCertificateById(certificateId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	c.JSON(http.StatusAccepted, getCertificateResponse{
		Data: *certificate,
	})

}

// @Summary Удаление сертификата
// @Tags certificate
// @Description Просмотр сертификата по id
// @ID get-certificate
// @Accept json
// @Produce json
// @Param id path int true "ID сертификата"
// @Security ApiKeyAuth
// @Router /api/certificates/{id} [delete]
func (h *Handler) deleteCertificate(c *gin.Context) {
	certificateId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	err = h.services.Certificate.DeleteCertificate(certificateId)

	c.JSON(http.StatusAccepted, map[string]string{
		"message": "certificate successfully delete !",
	})
}

type checkCertificateInfoRequest struct {
	Url string `json:"url" binding:"required,url"`
}
type checkCetrificateInfoResponse struct {
	Data ttlchecker.CertificateInfo `json:"data"`
}

// @Summary Проверка сертификата
// @Tags certificate
// @Description Проверка сертификата
// @ID check-certificate
// @Accept json
// @Produce json
// @Param url body checkCertificateInfoRequest true
// @Security ApiKeyAuth
// @Router /api/certificates/check [post]
func (h *Handler) checkCertificate(c *gin.Context) {
	var input checkCertificateInfoRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	certificateInfo, err := h.services.Certificate.CheckCertificate(input.Url)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, checkCetrificateInfoResponse{
		Data: *certificateInfo,
	})
}
