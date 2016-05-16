package main

import (
	"testing"

	"github.com/Jeffail/gabs"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestRenderWithReturn(c *C) {
	b := []byte(`{"_embedded":{"pool":[{"id":123,"name":"pool-test-1","_status":"OK"}]}}`)
	p, _ := parseJson(b, "pool")
	c.Assert(p, HasLen, 1)
}

func (s *S) TestRenderWithEmptyReturn(c *C) {
	b := []byte(`{"_embedded":{"pool":[]}}`)
	p, _ := parseJson(b, "pool")
	c.Assert(p, HasLen, 0)
}

func (s *S) TestRenderWithParseError(c *C) {
	b := []byte(`{"_embedded":{pool:[]}}`)
	_, err := parseJson(b, "pool")
	c.Assert(err, ErrorMatches, "error while parsing body")
}

func (s *S) TestRenderWithGettingError(c *C) {
	b := []byte(`{"_embedded":{"pool":[{"id":123,"name":"pool-test-1","_status":"OK"}]}}`)
	_, err := parseJson(b, "test")
	c.Assert(err, ErrorMatches, "error while getting entity")
}

func (s *S) TestGetPool(c *C) {
	result := []byte(`{"_embedded":{"pool":[{"id":123,"name":"pool-test-1","_status":"OK"}]}}`)
	jsonObj, _ := gabs.ParseJSON(result)
	expected, _ := jsonObj.S("_embedded", "pool").Children()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(result)
	}))
	defer ts.Close()

	body, _ := getEntity(ts.URL, "123456789", "pool")

	c.Assert(body, HasLen, 1)
	c.Assert(body, DeepEquals, expected)
}
