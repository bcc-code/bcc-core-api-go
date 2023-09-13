package coreapi

type ClientCredentialsEnv struct {
	TokenUrl string
	Audience string
}

const UrlProd = "https://api.bcc.no"
const UrlSandbox = "https://sandbox-api.bcc.no"
const UrlDev = "https://dev-api.bcc.no"

var CredEnvDev = ClientCredentialsEnv{
	TokenUrl: "https://bcc-sso-dev.eu.auth0.com/oauth/token",
	Audience: "dev-api.bcc.no",
}

var CredEnvSandbox = ClientCredentialsEnv{
	TokenUrl: "https://bcc-sso-sandbox.eu.auth0.com/oauth/token",
	Audience: "sandbox-api.bcc.no",
}

var CredEnvProd = ClientCredentialsEnv{
	TokenUrl: "https://bcc-sso.eu.auth0.com/oauth/token",
	Audience: "api.bcc.no",
}
