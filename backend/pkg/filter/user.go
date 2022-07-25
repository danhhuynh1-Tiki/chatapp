package filter

import "chat/services/models"

func FilteredResponse(user *models.DBResponse) models.UserResponse {
	return models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		RequestAt: user.RequestAt,
		Status:    user.Status,
	}
}
