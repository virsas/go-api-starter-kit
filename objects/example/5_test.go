package example

import (
	"context"
	"go-api-starter-kit/middlewares"
	"go-api-starter-kit/test"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	_, ciPresent := os.LookupEnv("CI")
	if !ciPresent {
		ctx := context.Background()
		// Postgres setup
		dbC, err := test.StartPostgresTestDB(ctx)
		// Mysql setup
		//dbC, err := test.StartMysqlTestDB(ctx)
		if err != nil {
			panic(err)
		}
		defer dbC.Terminate(ctx)
	} else {
		// Postgres setup
		test.TestSetupDBEnv("5432", "test", "test", "test", "postgres")
		// Mysql setup
		//test.TestSetupDBEnv("5432", "test", "root", "test", "mysql")
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

func setupServer() *gin.Engine {
	env, err := test.NewTestAPI()
	if err != nil {
		panic(err)
	}
	env.Router.Use(middlewares.Auth("../../keys", "test", env.Logger))
	env.Router.Use(middlewares.User(env.DB, env.Logger))
	Routes(env.Router, "", env.DB, env.Logger)
	return env.Router
}

func TestHealth(t *testing.T) {
	testServer := httptest.NewServer(setupServer())
	defer testServer.Close()

	tests := []struct {
		name                 string
		method               string
		authorizationEnabled bool
		email                string
		role                 []string
		body                 string
		endpoint             string
		status               int
		expected             string
	}{
		{name: "Examples - Get - authorized - admin role - list", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":1,"name":"test","updatedat":"2022-04-08T12:00:00Z","createdat":"2022-04-08T12:00:00Z"}]`},
		{name: "Examples - Get - authorized - admin role - list", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/1", status: 200, expected: `{"id":1,"name":"test","updatedat":"2022-04-08T12:00:00Z","createdat":"2022-04-08T12:00:00Z"}`},
	}

	for _, st := range tests {
		var authorization string
		if st.authorizationEnabled {
			authorization = test.SetupJWT("../../keys", st.email, st.role)
		}

		resp := test.GetResponse(testServer.URL, st.endpoint, st.method, st.body, st.authorizationEnabled, authorization)
		body, _ := ioutil.ReadAll(resp.Body)

		assert.Equal(t, st.status, resp.StatusCode, st.name)
		assert.Equal(t, st.expected, string(body), st.name)
	}
}
