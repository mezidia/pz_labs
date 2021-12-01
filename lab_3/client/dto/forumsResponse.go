package dto

type ForumsResponse struct {
	ForumName string `json:"ForumName"`
	TopicKeyword string `json:"InterestName"`
	Users []string `json:"UserName"`
}
