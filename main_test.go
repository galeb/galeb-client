package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestRenderWithReturn(c *C) {
	b := []byte(`{"_embedded":{"pool":[{"id":123,"name":"pool-test-1","_status":"OK"}]}}`)
	d := jsonData{Embedded{make([]Pool, 0)}}
	p, _ := render(b, d)
	c.Assert(len(p), Equals, 1)
}

func (s *S) TestRenderWithEmptyReturn(c *C) {
	b := []byte(`{"_embedded":{"pool":[]}}`)
	d := jsonData{Embedded{make([]Pool, 0)}}
	p, _ := render(b, d)
	c.Assert(len(p), Equals, 0)
}

func (s *S) TestRenderWithParseError(c *C) {
	b := []byte(`{"_embedded":{pool:[]}}`)
	d := jsonData{Embedded{make([]Pool, 0)}}
	_, err := render(b, d)
	c.Assert(err, ErrorMatches, "error while parsing body")
}

func (s *S) TestGetPool(c *C) {
	result := []byte(`{"_embedded":{"pool":[{"id":123,"name":"pool-test-1","_status":"OK"}]}}`)
	expected := Pool{Id:123, Name:"pool-test-1", Status:"OK"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(result)
	}))
	defer ts.Close()

	body, _ := getPool(ts.URL, "123456789")

	c.Assert(len(body), Equals, 1)
	c.Assert(body[0], Equals, expected)
}
