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
target := "http://localhost:8080"

func (s *AuthosaurusSuite) TestApiDocs_ServesSwaggerMetadata(c *C) {
	response, err := http.Get(target + "/api-docs")
	defer response.Body.Close()
	errCheck(c, err)
	body, err := ioutil.ReadAll(response.Body)
	errCheck(c, err)
	c.Assert(
		// Matches will not match across newlines so split on newline and
		// join with space.
		strings.Replace(string(body), "\n", " ", -1),
		Matches,
		".*swaggerVersion.*")
}

