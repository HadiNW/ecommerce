package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	ListCampaign() ([]Campaign, error)
	ListCampaignByUserID(userID int) ([]Campaign, error)
	GetCampaignByID(campaignID int) (Campaign, error)
	CreateCampaign(campaign Campaign) (Campaign, error)
	UpdateCampaign(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ListCampaign() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("Images", "is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) ListCampaignByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("Images", "is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) GetCampaignByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("id = ?", ID).Preload("User").Preload("Images").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) CreateCampaign(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) UpdateCampaign(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
