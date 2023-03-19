package models

//Search criteria model
type SearchCriteria struct {
	Title        *string `json:"title"`
	Tags         *string `json:"tags"`
	Description  *string `json:"description"`
	Error        *string `json:"error"`
	Status       *bool   `json:"status"` //para asver si esta resuelto o no
	UserNickname *string `json:"nickname"`
	From         int64   `json:"from"`
	Size         int64   `json:"size"`
}

func (s SearchCriteria) IsZero() bool {
	if s.Title == nil &&
		s.Status == nil &&
		s.UserNickname == nil &&
		s.Tags == nil &&
		s.Description == nil &&
		s.Error == nil {
		return true
	}
	return false
}
