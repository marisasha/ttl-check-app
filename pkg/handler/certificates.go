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

type getAllCertificatesResponse struct {
	Data []ttlchecker.CertificateResponse `json:"data"`
}

type getCertificateResponse struct {
	Data ttlchecker.CertificateResponse `json:"data"`
}

type checkCertificateInfoRequest struct {
	Url string `json:"url" binding:"required,url"`
}

type checkCertificateInfoResponse struct {
	Data ttlchecker.CertificateInfo `json:"data"`
}

//
// ADD CERTIFICATE
//

// @Summary Добавить сертификат
// @Tags certificates
// @Description Добавляет новый сайт для отслеживания SSL сертификата
// @ID add-certificate
// @Accept json
// @Produce json
// @Param request body AddCertificateRequest true "URL сайта"
// @Success 201 {object} map[string]string
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
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

	if err := h.services.Certificate.AddCertificate(newCertificate); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"message": "URL successfully added!",
	})
}

//
// GET ALL CERTIFICATES
//

// @Summary Получить все сертификаты
// @Tags certificates
// @Description Возвращает список сертификатов пользователя
// @ID get-all-certificates
// @Accept json
// @Produce json
// @Success 200 {object} getAllCertificatesResponse
// @Failure 500 {object} errorResponse
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

	c.JSON(http.StatusOK, getAllCertificatesResponse{
		Data: *certificates,
	})
}

//
// GET ONE CERTIFICATE
//

// @Summary Получить сертификат
// @Tags certificates
// @Description Возвращает сертификат по ID
// @ID get-certificate
// @Accept json
// @Produce json
// @Param id path int true "ID сертификата"
// @Success 200 {object} getCertificateResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Security ApiKeyAuth
// @Router /api/certificates/{id} [get]
func (h *Handler) getCertificate(c *gin.Context) {
	certificateId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	certificate, err := h.services.Certificate.GetCertificateById(certificateId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCertificateResponse{
		Data: *certificate,
	})
}

//
// DELETE CERTIFICATE
//

// @Summary Удалить сертификат
// @Tags certificates
// @Description Удаляет сертификат по ID
// @ID delete-certificate
// @Accept json
// @Produce json
// @Param id path int true "ID сертификата"
// @Success 200 {object} map[string]string
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Security ApiKeyAuth
// @Router /api/certificates/{id} [delete]
func (h *Handler) deleteCertificate(c *gin.Context) {
	certificateId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Certificate.DeleteCertificate(certificateId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "certificate successfully deleted!",
	})
}

//
// CHECK CERTIFICATE TTL
//

// @Summary Проверить SSL сертификат
// @Tags certificates
// @Description Проверяет срок действия SSL сертификата сайта
// @ID check-certificate
// @Accept json
// @Produce json
// @Param request body checkCertificateInfoRequest true "URL сайта"
// @Success 200 {object} checkCertificateInfoResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
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

	c.JSON(http.StatusOK, checkCertificateInfoResponse{
		Data: *certificateInfo,
	})
}
