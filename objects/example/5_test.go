package example

import (
	"context"
	"go-api-starter-kit/middlewares"
	"go-api-starter-kit/test"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"regexp"
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

const (
	updatedString = ",\"createdat\":\\s?\"(20.*?)\""
	createdString = ",\"updatedat\":\\s?\"(20.*?)\""
)

var (
	updatedRegex = regexp.MustCompile(updatedString)
	createdRegex = regexp.MustCompile(createdString)
)

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
		// Unauthorized test should fail
		{name: "Examples - Get - unauthorized - admin role - list", method: "GET", authorizationEnabled: false, email: "", role: []string{""}, body: "", endpoint: "/v1/examples/", status: 500, expected: `{"message":"apiError"}`},
		{name: "Examples - Get - unauthorized - admin role - show", method: "GET", authorizationEnabled: false, email: "", role: []string{""}, body: "", endpoint: "/v1/examples/1", status: 500, expected: `{"message":"apiError"}`},
		{name: "Examples - Post - unauthorized - admin role - list", method: "POST", authorizationEnabled: false, email: "", role: []string{""}, body: `{"name":"111"}`, endpoint: "/v1/examples/", status: 500, expected: `{"message":"apiError"}`},
		{name: "Examples - Patch - unauthorized - admin role - update", method: "PATCH", authorizationEnabled: false, email: "", role: []string{""}, body: `{"name":"222"}`, endpoint: "/v1/examples/3", status: 500, expected: `{"message":"apiError"}`},
		{name: "Examples - Delete - unauthorized - admin role - delete", method: "DELETE", authorizationEnabled: false, email: "", role: []string{""}, body: "", endpoint: "/v1/examples/3", status: 500, expected: `{"message":"apiError"}`},
		// Authorized user with proper role
		{name: "Examples - Get - authorized - admin role - list", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		{name: "Examples - Get - authorized - admin role - show", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/1", status: 200, expected: `{"id":1,"name":"test"}`},
		{name: "Examples - Post - authorized - admin role - list", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"111"}`, endpoint: "/v1/examples/", status: 200, expected: `{"message":"OK"}`},
		{name: "Examples - Get - authorized - admin role - list again 1", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":3,"name":"111"},{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		{name: "Examples - Patch - authorized - admin role - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"222"}`, endpoint: "/v1/examples/3", status: 200, expected: `{"message":"OK"}`},
		{name: "Examples - Get - authorized - admin role - list again 2", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":3,"name":"222"},{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		{name: "Examples - Delete - authorized - admin role - delete", method: "DELETE", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/3", status: 200, expected: `{"message":"OK"}`},
		{name: "Examples - Get - authorized - admin role - list again 3", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		// Validation test with incorrect values
		{name: "Examples - Post - authorized - admin role - list", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"111 111"}`, endpoint: "/v1/examples/", status: 400, expected: `{"message":"validationError"}`},
		{name: "Examples - Patch - authorized - admin role - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"222 222"}`, endpoint: "/v1/examples/2", status: 400, expected: `{"message":"validationError"}`},
		// Incorrect role
		{name: "Examples - Get - unauthorized - user role - list", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: "", endpoint: "/v1/examples/", status: 403, expected: `{"message":"notAllowed"}`},
		{name: "Examples - Get - unauthorized - user role - show", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: "", endpoint: "/v1/examples/1", status: 403, expected: `{"message":"notAllowed"}`},
		{name: "Examples - Post - unauthorized - user role - list", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: `{"name":"111"}`, endpoint: "/v1/examples/", status: 403, expected: `{"message":"notAllowed"}`},
		{name: "Examples - Patch - unauthorized - user role - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: `{"name":"222"}`, endpoint: "/v1/examples/3", status: 403, expected: `{"message":"notAllowed"}`},
		{name: "Examples - Delete - unauthorized - user role - delete", method: "DELETE", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: "", endpoint: "/v1/examples/3", status: 403, expected: `{"message":"notAllowed"}`},
		// Unknown email address in JWT
		{name: "Examples - Get - wrong email - admin role - list", method: "GET", authorizationEnabled: true, email: "wrong@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 403, expected: `{"message":"noUserFound"}`},
		// Not existing items
		{name: "Examples - Get - authorized - admin role - show unknown", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/99", status: 404, expected: `{"message":"notFound"}`},
		{name: "Examples - Patch - authorized - admin role - update unknown", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"222"}`, endpoint: "/v1/examples/99", status: 404, expected: `{"message":"notFound"}`},
		{name: "Examples - Delete - authorized - admin role - delete unknown", method: "DELETE", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/99", status: 404, expected: `{"message":"notFound"}`},
	}

	for _, st := range tests {
		var authorization string
		if st.authorizationEnabled {
			authorization = test.SetupJWT("../../keys", st.email, st.role)
		}

		resp := test.GetResponse(testServer.URL, st.endpoint, st.method, st.body, st.authorizationEnabled, authorization)
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		bodyStr = updatedRegex.ReplaceAllString(bodyStr, "")
		bodyStr = createdRegex.ReplaceAllString(bodyStr, "")

		assert.Equal(t, st.status, resp.StatusCode, st.name)
		assert.Equal(t, st.expected, bodyStr, st.name)
	}
}
