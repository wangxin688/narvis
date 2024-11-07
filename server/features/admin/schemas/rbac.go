package schemas

type Oauth2PasswordRequest struct {
	Username       string  `form:"username" binding:"required,email"`
	Password       string  `form:"password" binding:"required"`
	EnterpriseCode *string `form:"enterpriseCode" binding:"omitempty"`
	DomainName     *string `form:"domainName" binding:"omitempty"`
}

type Oauth2RefreshRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}
