package schemas

type LocalAuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SlackAuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GoogleAuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TeamsAuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
