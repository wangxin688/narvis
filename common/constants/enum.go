package constants

type CorsMode string

const (
	CorsAllowAllMode  CorsMode = "allow-all"
	CorsWhiteListMode CorsMode = "whitelist"
	CorsStrictMode    CorsMode = "strict-whitelist"
)

