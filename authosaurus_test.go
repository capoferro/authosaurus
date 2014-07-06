package main_test

import (
	"testing"

	"strings"
	"io/ioutil"
	"net/http"
	
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type AuthosaurusSuite struct{}

var _ = Suite(&AuthosaurusSuite{})

func errCheck(c *C, err error) {
	if err != nil {
		c.Fatal("Error: " + err.Error())
	}
}

// Endpoint Tests
const TARGET = "http://localhost:8080"

type assertResponseFunc func(*http.Response, string)

func assertOnGet(c *C, path string, f assertResponseFunc) {
	response, err := http.Get(TARGET + path)
	errCheck(c, err)
	assertOnResponse(c, response, f)
}

func assertOnPost(c *C, path string, postData string, f assertResponseFunc) {
	response, err := http.Post(TARGET + path, "application/json", strings.NewReader(postData))
	errCheck(c, err)
	assertOnResponse(c, response, f)
}

func assertOnResponse(c *C, response *http.Response, f assertResponseFunc) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	errCheck(c, err)

	f(response, string(body))
}

func matchable(str string) string {
	// Matches will not match across newlines so replace \n with space
	return strings.Replace(str, "\n", " ", -1)
}

func (s *AuthosaurusSuite) TestApiDocs_ServesSwaggerMetadata(c *C) {
	assertOnGet(c, "/api-docs", func(response *http.Response, body string) {
		c.Assert(response.StatusCode, Equals, 200)
		c.Assert(matchable(body),
			Matches,
			".*swaggerVersion.*")
	})
}

func (s *AuthosaurusSuite) TestUsers_Create(c *C) {
	assertOnPost(c, "/users", `{"name": "Capo"}`, func(response *http.Response, body string) {
		c.Assert(response.StatusCode, Equals, 201)
		c.Assert(body, Equals, "{\n" +
     "  \"Id\": \"1\",\n" +
     "  \"Name\": \"Capo\"\n" +
     " }")
	})
}

func (s *AuthosaurusSuite) TestUsers_CreateError(c *C) {
	assertOnPost(c, "/users", `bumblebee`, func(response *http.Response, body string) {
		c.Assert(response.StatusCode, Equals, 400)
		c.Assert(body, Equals, "Error parsing user JSON: invalid character 'b' looking for beginning of value")
	})
}
