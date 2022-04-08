package test

import "os"

func TestSetupDBEnv(port string, db string, user string, pass string, host string) {
	os.Setenv("DB_PORT", port)
	os.Setenv("DB_USER", user)
	os.Setenv("DB_NAME", db)
	os.Setenv("DB_PASS", pass)
	os.Setenv("DB_HOST", host)
}
