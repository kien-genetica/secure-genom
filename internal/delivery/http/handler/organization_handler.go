package handler

import (
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
	panic("not implemented")
}
