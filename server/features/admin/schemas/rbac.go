package schemas

type Oauth2PasswordRequest struct {
	Username       string  `form:"username" binding:"required,email"`
	Password       string  `form:"password" binding:"required"`
	EnterpriseCode *string `form:"enterprise_code" binding:"omitempty"`
	DomainName     *string `form:"domain_name" binding:"omitempty"`
}
