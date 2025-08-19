package validation

import "svi-be/internal/model"

func ValidatePost(post *model.Posts) (bool, string) {

	if len(post.Title) < 20 {
		return false, "Title must be at least 20 characters long"
	}

	if len(post.Content) < 200 {
		return false, "Content must be at least 200 characters long"
	}

	if len(post.Category) < 3 {
		return false, "Category must be at least 3 characters long"
	}

	validStatuses := []string{"publish", "draft", "thrash"}
	if !Contains(validStatuses, post.Status) {
		return false, "Status must be one of 'Publish', 'Draft', or 'Thrash'"
	}

	return true, ""
}

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
