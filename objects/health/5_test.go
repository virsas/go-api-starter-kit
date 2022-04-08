package health

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
		pgc, err := test.NewTestDB(ctx)
		if err != nil {
			panic(err)
		}
		defer pgc.Terminate(ctx)
	} else {
		test.TestSetupDBEnv("3306", "test", "root", "test", "mysql")
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
		{name: "Checking /health", method: "GET", authorizationEnabled: true, email: "stefan@kephala.com", role: []string{"admin"}, body: "", endpoint: "/health", status: 200, expected: `{"database":"OK","server":"OK"}`},
		{name: "Checking /status", method: "GET", authorizationEnabled: true, email: "stefan@kephala.com", role: []string{"admin"}, body: "", endpoint: "/status", status: 200, expected: `{"database":"OK","server":"OK"}`},
		{name: "Checking /aaaa", method: "GET", authorizationEnabled: true, email: "stefan@kephala.com", role: []string{"admin"}, body: "", endpoint: "/aaaa", status: 404, expected: `404 page not found`},
	}

	for _, st := range tests {
		var authorization string
		if st.authorizationEnabled {
			authorization = test.SetupJWT("../../keys", st.email, st.role)
		}

		resp := test.GetResponse(testServer.URL, st.endpoint, st.method, st.body, st.authorizationEnabled, authorization)
		body, _ := ioutil.ReadAll(resp.Body)

		assert.Equal(t, st.status, resp.StatusCode, "Failed to read endpoint "+st.endpoint)
		assert.Equal(t, st.expected, string(body), "Not expected body on endpoint "+st.endpoint)
	}
}
