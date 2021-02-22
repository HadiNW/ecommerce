package campaign

type CampaignResponse struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Perks            string `json:"perks"`
	BackerCount      int    `json:"backer_count"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	ImageURL         string `json:"image_url"`
}

type CampaignDetailResponse struct {
	ID               int                     `json:"id"`
	User             *CampaignUserResponse   `json:"user"`
	Name             string                  `json:"name"`
	ImageURL         string                  `json:"image_url"`
	ShortDescription string                  `json:"short_description"`
	Description      string                  `json:"description"`
	Perks            string                  `json:"perks"`
	BackerCount      int                     `json:"backer_count"`
	GoalAmount       int                     `json:"goal_amount"`
	CurrentAmount    int                     `json:"current_amount"`
	Slug             string                  `json:"slug"`
	Images           []CampaignImageResponse `json:"images"`
}

type CampaignImageResponse struct {
	ImageURL  string `json:"image_url"`
	IsPrimary int    `json:"is_primary"`
}

type CampaignUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatCampaign(campaign Campaign) CampaignResponse {
	res := CampaignResponse{}
	res.ID = campaign.ID
	res.Name = campaign.Name
	res.ShortDescription = campaign.ShortDescription
	res.Description = campaign.Description
	res.Perks = campaign.Perks
	res.BackerCount = campaign.BackerCount
	res.GoalAmount = campaign.GoalAmount
	res.CurrentAmount = campaign.CurrentAmount
	res.Slug = campaign.Slug
	res.UserID = campaign.UserID

	for _, img := range campaign.Images {
		if img.IsPrimary == 1 {
			res.ImageURL = img.FileName
		}
	}
	return res
}

func FormatCampaigns(campaigns []Campaign) []CampaignResponse {
	res := []CampaignResponse{}
	for _, c := range campaigns {
		campaign := FormatCampaign(c)
		res = append(res, campaign)
	}
	return res
}

func FormatCampaignDetail(c Campaign) CampaignDetailResponse {
	d := CampaignDetailResponse{}
	d.ID = c.ID
	d.Name = c.Name
	d.ShortDescription = c.ShortDescription
	d.Description = c.Description
	d.Perks = c.Perks
	d.BackerCount = c.BackerCount
	d.GoalAmount = c.GoalAmount
	d.CurrentAmount = c.CurrentAmount
	d.Slug = c.Slug

	if c.User.ID != 0 {
		d.User = &CampaignUserResponse{
			ID:   c.User.ID,
			Name: c.User.FullName,
		}
	}
	d.Images = []CampaignImageResponse{}

	for _, img := range c.Images {
		d.Images = append(d.Images, CampaignImageResponse{
			ImageURL:  img.FileName,
			IsPrimary: img.IsPrimary,
		})
		if img.IsPrimary == 1 {
			d.ImageURL = img.FileName
		}
	}

	return d
}
