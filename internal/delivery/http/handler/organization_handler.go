package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kien6034/secure-genom/internal/usecase/organization"
)

type OrganizationHandler struct {
	orgService organization.Service
}

func NewOrganizationHandler(orgService organization.Service) *OrganizationHandler {
	return &OrganizationHandler{orgService: orgService}
}

func (h *OrganizationHandler) Register(c *gin.Context) {
	var req organization.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orgService.RegisterOrganization(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization registered successfully"})
}
