package api

import (
	"context"
	"net/http"

	"customer_management_service/tracing"

	"github.com/beego/beego/v2/client/httplib"
)

type Client struct {
	Request *Request
	Type_   string
}

// SendRequest performs the request without a tracing context
// (e.g. from background jobs outside the request path).
func (c *Client) SendRequest() (*http.Response, error) {
	return c.SendRequestWithContext(context.Background())
}

// SendRequestWithContext performs the request using ctx and forwards
// the tracing identifiers already attached to it (X-Session-Id,
// X-Trace-Id, X-Request-Id) to the downstream service. The SAME
// identifiers are forwarded; no new request IDs are generated here.
func (c *Client) SendRequestWithContext(ctx context.Context) (*http.Response, error) {
	// TODO: BaseURLとPathのvalidation
	beegoRequest := httplib.NewBeegoRequestWithCtx(
		ctx,
		c.Request.BaseURL+c.Request.Path,
		c.Request.Method.String())
	for key, value := range c.Request.HeaderField {
		beegoRequest.Header(key, value)
	}
	tracing.ApplyOutboundHeaders(ctx, beegoRequest.GetRequest().Header)
	if c.Type_ == "params" {
		for key, value := range c.Request.Params {
			beegoRequest.Param(key, value)
		}
	} else if c.Type_ == "body" {
		beegoRequest.JSONBody(c.Request.Params)
	}
	res, err := beegoRequest.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}
