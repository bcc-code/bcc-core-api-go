package bcccoreapi

type Environment int

const (
	EnvironmentProd Environment = iota
	EnvironmentSandbox
	EnvironmentDev
)

type EnvironmentConfig struct {
	BaseUrl  string
	TokenUrl string
	Audience string
}

var envMap = map[Environment]EnvironmentConfig{
	EnvironmentProd:    envProd,
	EnvironmentSandbox: envSandbox,
	EnvironmentDev:     envDev,
}

var envProd = EnvironmentConfig{
	BaseUrl:  "https://api.bcc.no",
	TokenUrl: "https://bcc-sso.eu.auth0.com/oauth/token",
	Audience: "api.bcc.no",
}

var envSandbox = EnvironmentConfig{
	BaseUrl:  "https://sandbox-api.bcc.no",
	TokenUrl: "https://bcc-sso-sandbox.eu.auth0.com/oauth/token",
	Audience: "sandbox-api.bcc.no",
}
var envDev = EnvironmentConfig{
	BaseUrl:  "https://dev-api.bcc.no",
	TokenUrl: "https://bcc-sso-dev.eu.auth0.com/oauth/token",
	Audience: "dev-api.bcc.no",
}
