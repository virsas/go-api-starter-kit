package health

import (
	"context"
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

	Routes(env.Router, "", env.DB, env.Logger)

	return env.Router
}

func TestHealth(t *testing.T) {
	testServer := httptest.NewServer(setupServer())
	defer testServer.Close()

	tests := []struct {
		object               string
		testType             string
		result               string
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
		{object: "Health", testType: "Check response", result: "Success", name: "Get /health", method: "GET", authorizationEnabled: false, email: "", role: []string{}, body: "", endpoint: "/health", status: 200, expected: `{"database":"OK","server":"OK"}`},
		{object: "Health", testType: "Check response", result: "Success", name: "Get /v1/status", method: "GET", authorizationEnabled: false, email: "", role: []string{}, body: "", endpoint: "/v1/status", status: 200, expected: `{"database":"OK","server":"OK"}`},
		{object: "Health", testType: "Unknown path", result: "Fail", name: "Get /aaaa", method: "GET", authorizationEnabled: false, email: "", role: []string{}, body: "", endpoint: "/aaaa", status: 404, expected: `404 page not found`},
	}

	for _, st := range tests {
		var authorization string
		if st.authorizationEnabled {
			authorization = test.SetupJWT("../../keys", st.email, st.role)
		}

		resp := test.GetResponse(testServer.URL, st.endpoint, st.method, st.body, st.authorizationEnabled, authorization)
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)

		assert.Equal(t, st.status, resp.StatusCode, st.object+" - "+st.testType+" - "+st.result+" - "+st.name)
		assert.Equal(t, st.expected, bodyStr, st.object+" - "+st.testType+" - "+st.result+" - "+st.name)
	}
}
