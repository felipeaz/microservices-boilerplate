package env

type DBEnv struct {
	Server       ServerProperties
	Credentials  StandardAuth
	DatabaseName string
}
