package bean

import "time"

const (
	TOKEN        = "TOKEN_"
	BRUTE        = "BRUTE_"
	RefreshToken = "REFRESH_TOKEN_"

	TokenPast           = time.Hour * 12 * 12
	RefreshTokenKeyPast = time.Hour * 12 * 15
	ViewUser            = "VIEW_"
	CommentCount        = "Comment_Count:"
	ViewCommentCount    = "View_Comment_Count:"
)
