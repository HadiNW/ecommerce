package handler

import (
	"ecommerce/campaign"
	"ecommerce/helper"
	"ecommerce/user"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) ListMyCampaign(c *gin.Context) {
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
	c.JSON(http.StatusOK, helper.APIResponseOK("success", campaign.FormatCampaign(campaigns)))
}

func (h *campaignHandler) ListCampaign(c *gin.Context) {
	userIDStr := c.Query("user_id")
	var userID int
	if userIDStr != "" {
		ID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("error", err))
			return
		}
		userID = ID
	}

	campaigns, err := h.campaignService.ListCampaign(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("List campaign failed", err))
		return
	}
	c.JSON(http.StatusOK, helper.APIResponseOK("success", campaign.FormatCampaign(campaigns)))
}

func (h *campaignHandler) GetCampaignByID(c *gin.Context) {
	strID := c.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("error", err))
		return
	}

	data, err := h.campaignService.GetCampaignByID(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponseBadRequest("error", err))
		return
	}
	c.JSON(http.StatusOK, helper.APIResponseOK("success", campaign.FormatCampaignDetail(data)))
}
