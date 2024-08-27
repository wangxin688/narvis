package schemas


type PlatformQuery struct {
	Platform *string `form:"platform" binding:"omitempty"`
	Keyword  *string `form:"keyword" binding:"omitempty"`
}
