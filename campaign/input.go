package campaign

type CreateCampaignPayload struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
}

type UpdateCampaignPayload struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Perks            string `json:"perks"`
	GoalAmount       int    `json:"goal_amount"`
}
