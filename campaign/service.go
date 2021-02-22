package campaign

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	ListCampaignByUserID(userID int) ([]Campaign, error)
	ListCampaign(userID int) ([]Campaign, error)
	GetCampaignByID(ID int) (Campaign, error)
	CreateCampaign(campaign CreateCampaignPayload, userID int) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) ListCampaign(userID int) ([]Campaign, error) {
	var campaigns = []Campaign{}
	var err error
	if userID != 0 {
		campaigns, err = s.repository.ListCampaignByUserID(userID)
	} else {
		campaigns, err = s.repository.ListCampaign()
	}
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) ListCampaignByUserID(userID int) ([]Campaign, error) {
	campaigns, err := s.repository.ListCampaignByUserID(userID)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(ID int) (Campaign, error) {
	campaign, err := s.repository.GetCampaignByID(ID)
	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, errors.New("Campaign not found")
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignPayload, userID int) (Campaign, error) {
	var data Campaign

	data.Name = input.Name
	data.Description = input.Description
	data.ShortDescription = input.ShortDescription
	data.Perks = input.Perks
	data.GoalAmount = input.GoalAmount
	data.UserID = userID
	rand.Seed(time.Now().UnixNano())
	data.Slug = slug.Make(fmt.Sprintf("%s-%d-%d", data.Name, userID, (1 + rand.Intn(999-1))))

	campaign, err := s.repository.CreateCampaign(data)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) UpdateCampaign(input UpdateCampaignPayload) (Campaign, error) {

}
