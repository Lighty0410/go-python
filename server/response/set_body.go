package response

import "github.com/valyala/fasthttp"

// SetBody sets body for a correct request.
func SetBody(ctx *fasthttp.RequestCtx, body []byte, statusCode int) {
	ctx.Response.Header.Set("Content-type", "application/json")
	ctx.SetBody(body)
	ctx.SetStatusCode(statusCode)
}
