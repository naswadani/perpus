package handler

import (
	"net/http"
	auth "perpustakaan/Auth"
	siswa "perpustakaan/User/Siswa"
	helper "perpustakaan/helper"

	"github.com/gin-gonic/gin"
)

type siswaHandler struct {
	siswaService siswa.Service
	authService  auth.Service
}

func NewSiswaHandler(SiswaService siswa.Service, authService auth.Service) *siswaHandler {
	return &siswaHandler{SiswaService, authService}
}

func (h *siswaHandler) RegisterSiswa(c *gin.Context) {
	var input siswa.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorFormater(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newSiswa, err := h.siswaService.RegisterSiswa(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newSiswa.IdUser)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, token)
}
