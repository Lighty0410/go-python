package response

import (
	"encoding/json"
	"fmt"
	"test/model"

	"github.com/valyala/fasthttp"
)

// Incorrect is used for incorrect response.
func Incorrect(ctx *fasthttp.RequestCtx, statusCode int, err error) {
	body, e := json.Marshal(model.IncorrectResp{Reason: err.Error()})
	if e != nil {
		ctx.SetBody([]byte(fmt.Sprintf("cannot marshal body: %v", err)))
	}
	ctx.Response.Header.Set("Content-type", "application/json")
	ctx.SetBody(body)
	ctx.SetStatusCode(statusCode)
}
