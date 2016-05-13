package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestRenderWithReturn(c *C) {
	fakeBody := []byte(`{"_embedded":{"pool":[{"id":123,"name":"pool-test-1","_status":"OK"}]}}`)
	fakeData := jsonData{Embedded{make([]Pool, 0)}}
	pools,_ := render(fakeBody, fakeData)
	c.Assert(len(pools), Equals, 1)
}

func (s *S) TestRenderWithEmptyReturn(c *C) {
	fakeBody := []byte(`{"_embedded":{"pool":[]}}`)
	fakeData := jsonData{Embedded{make([]Pool, 0)}}
	pools,_ := render(fakeBody, fakeData)
	c.Assert(len(pools), Equals, 0)
}

func (s *S) TestRenderWithParseError(c *C) {
	fakeBody := []byte(`{"_embedded":{pool:[]}}`)
	fakeData := jsonData{Embedded{make([]Pool, 0)}}
	_,err := render(fakeBody, fakeData)
	c.Assert(err, ErrorMatches, "Error while parsing body!")
}
