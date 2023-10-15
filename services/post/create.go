package post

import "account-transaction-api/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
