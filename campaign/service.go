package campaign

type Service interface {
	ListCampaignByUserID(userID int) ([]Campaign, error)
	ListCampaign(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) ListCampaign(userID int) ([]Campaign, error) {
	var campaigns []Campaign
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
