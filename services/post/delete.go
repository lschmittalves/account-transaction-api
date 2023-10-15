package post

import "account-transaction-api/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
