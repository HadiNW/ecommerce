package campaign

import "errors"

type Service interface {
	ListCampaignByUserID(userID int) ([]Campaign, error)
	ListCampaign(userID int) ([]Campaign, error)
	GetCampaignByID(ID int) (Campaign, error)
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
