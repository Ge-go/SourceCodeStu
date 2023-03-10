package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	marshal, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(marshal)
	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SysErrJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadReqJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

// NewContext 后续使用池化技术优化Context
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{w, r}
}
