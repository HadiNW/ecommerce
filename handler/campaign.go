package handler

import (
	"ecommerce/campaign"
	"ecommerce/helper"
	"ecommerce/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) ListCampaignByUserID(c *gin.Context) {
	decoded, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("User not found", errors.New("User not found")))
		return
	}

	user := decoded.(user.User)
	campaigns, err := h.campaignService.ListCampaignByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("List campaign failed", err))
		return
	}
	c.JSON(http.StatusOK, helper.APIResponseOK("success", campaigns))
}
