package example

import (
	"context"
	"go-api-starter-kit/helpers"
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
		{object: "Example", testType: "Unauthorized user", result: "Fail", name: "no user - list", method: "GET", authorizationEnabled: false, email: "", role: []string{""}, body: "", endpoint: "/v1/examples/", status: 500, expected: `{"message":"apiError"}`},
		{object: "Example", testType: "Unauthorized user", result: "Fail", name: "no user - show", method: "GET", authorizationEnabled: false, email: "", role: []string{""}, body: "", endpoint: "/v1/examples/1", status: 500, expected: `{"message":"apiError"}`},
		{object: "Example", testType: "Unauthorized user", result: "Fail", name: "no user - list", method: "POST", authorizationEnabled: false, email: "", role: []string{""}, body: `{"name":"111"}`, endpoint: "/v1/examples/", status: 500, expected: `{"message":"apiError"}`},
		{object: "Example", testType: "Unauthorized user", result: "Fail", name: "no user - update", method: "PATCH", authorizationEnabled: false, email: "", role: []string{""}, body: `{"name":"222"}`, endpoint: "/v1/examples/3", status: 500, expected: `{"message":"apiError"}`},
		{object: "Example", testType: "Unauthorized user", result: "Fail", name: "no user - delete", method: "DELETE", authorizationEnabled: false, email: "", role: []string{""}, body: "", endpoint: "/v1/examples/3", status: 500, expected: `{"message":"apiError"}`},
		{object: "Example", testType: "Unauthorized user", result: "Fail", name: "spoofed user - list", method: "GET", authorizationEnabled: true, email: "wrong@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 403, expected: `{"message":"noUserFound"}`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - list 1", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - show", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/1", status: 200, expected: `{"id":1,"name":"test"}`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - list", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"111"}`, endpoint: "/v1/examples/", status: 200, expected: `{"message":"OK"}`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - list 2", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":3,"name":"111"},{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"222"}`, endpoint: "/v1/examples/3", status: 200, expected: `{"message":"OK"}`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - list 3", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":3,"name":"222"},{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - delete", method: "DELETE", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/3", status: 200, expected: `{"message":"OK"}`},
		{object: "Example", testType: "Authorized user", result: "Success", name: "admin role - list 4", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/", status: 200, expected: `[{"id":2,"name":"test2"},{"id":1,"name":"test"}]`},
		{object: "Example", testType: "Authorized user", result: "Fail", name: "unknown role - list", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: "", endpoint: "/v1/examples/", status: 403, expected: `{"message":"notAllowed"}`},
		{object: "Example", testType: "Authorized user", result: "Fail", name: "unknown role - show", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: "", endpoint: "/v1/examples/1", status: 403, expected: `{"message":"notAllowed"}`},
		{object: "Example", testType: "Authorized user", result: "Fail", name: "unknown role - list", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: `{"name":"111"}`, endpoint: "/v1/examples/", status: 403, expected: `{"message":"notAllowed"}`},
		{object: "Example", testType: "Authorized user", result: "Fail", name: "unknown role - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: `{"name":"222"}`, endpoint: "/v1/examples/3", status: 403, expected: `{"message":"notAllowed"}`},
		{object: "Example", testType: "Authorized user", result: "Fail", name: "unknown role - delete", method: "DELETE", authorizationEnabled: true, email: "info@examples.org", role: []string{"user"}, body: "", endpoint: "/v1/examples/3", status: 403, expected: `{"message":"notAllowed"}`},
		{object: "Example", testType: "Validation", result: "Fail", name: "wrong legth - create", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"1"}`, endpoint: "/v1/examples/", status: 400, expected: `{"message":"validationError"}`},
		{object: "Example", testType: "Validation", result: "Fail", name: "wrong symbol - create", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"111*111"}`, endpoint: "/v1/examples/", status: 400, expected: `{"message":"validationError"}`},
		{object: "Example", testType: "Validation", result: "Fail", name: "wrong fields - create", method: "POST", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{}`, endpoint: "/v1/examples/", status: 400, expected: `{"message":"validationError"}`},
		{object: "Example", testType: "Validation", result: "Fail", name: "wrong length - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"2"}`, endpoint: "/v1/examples/2", status: 400, expected: `{"message":"validationError"}`},
		{object: "Example", testType: "Validation", result: "Fail", name: "wrong symbol - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"222*222"}`, endpoint: "/v1/examples/2", status: 400, expected: `{"message":"validationError"}`},
		{object: "Example", testType: "Validation", result: "Fail", name: "wrong fields - update", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{}`, endpoint: "/v1/examples/2", status: 400, expected: `{"message":"validationError"}`},
		{object: "Example", testType: "Unknown objects", result: "Fail", name: "show unknown", method: "GET", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/99", status: 404, expected: `{"message":"notFound"}`},
		{object: "Example", testType: "Unknown objects", result: "Fail", name: "update unknown", method: "PATCH", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: `{"name":"222"}`, endpoint: "/v1/examples/99", status: 404, expected: `{"message":"notFound"}`},
		{object: "Example", testType: "Unknown objects", result: "Fail", name: "delete unknown", method: "DELETE", authorizationEnabled: true, email: "info@examples.org", role: []string{"admin"}, body: "", endpoint: "/v1/examples/99", status: 404, expected: `{"message":"notFound"}`},
	}

	for _, st := range tests {
		var authorization string
		if st.authorizationEnabled {
			authorization = test.SetupJWT("../../keys", st.email, st.role)
		}

		resp := test.GetResponse(testServer.URL, st.endpoint, st.method, st.body, st.authorizationEnabled, authorization)
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		bodyStr = helpers.UpdatedRegex.ReplaceAllString(bodyStr, "")
		bodyStr = helpers.CreatedRegex.ReplaceAllString(bodyStr, "")

		assert.Equal(t, st.status, resp.StatusCode, st.object+" - "+st.testType+" - "+st.result+" - "+st.name)
		assert.Equal(t, st.expected, bodyStr, st.object+" - "+st.testType+" - "+st.result+" - "+st.name)
	}
}
