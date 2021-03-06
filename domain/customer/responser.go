package customer

type CustomerRegisterPayload struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CustomerResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token,omitempty"`
	Avatar   string `json:"avatar"`
	Status   bool   `json:"status"`
}

type CustomerLoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func MarshalResponse(customer Customer) CustomerResponse {
	r := CustomerResponse{}

	r.ID = customer.ID
	r.FullName = customer.FullName
	r.Username = customer.Username
	r.Email = customer.Email
	r.Avatar = customer.Avatar
	r.Status = customer.Status

	return r
}

func (c *CustomerResponse) SetToken(token string) {
	c.Token = token
}
