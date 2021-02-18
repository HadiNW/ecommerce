package campaign

type CampaignFormatter struct {
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

func FormatCampaign(campaigns []Campaign) []CampaignFormatter {
	var res []CampaignFormatter

	for _, c := range campaigns {
		var imageURL string
		if len(c.Images) > 0 {
			imageURL = c.Images[0].FileName
		}
		d := CampaignFormatter{
			ID:               c.ID,
			UserID:           c.UserID,
			Name:             c.Name,
			ShortDescription: c.ShortDescription,
			Description:      c.Description,
			Perks:            c.Perks,
			BackerCount:      c.BackerCount,
			GoalAmount:       c.GoalAmounnt,
			CurrentAmount:    c.CurrentAmount,
			Slug:             c.Slug,
			ImageURL:         imageURL,
		}
		res = append(res, d)
	}
	if len(res) == 0 {
		var d []CampaignFormatter
		return d
	}
	return res
}
