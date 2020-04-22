package http

import (
	"context"
	"encoding/json"
	"net/http"
	"test/model"
	"test/server/response"

	"github.com/valyala/fasthttp"
)

func (s *Server) testText(ctx *fasthttp.RequestCtx) {
	var textCase model.TextCase
	err := json.Unmarshal(ctx.PostBody(), &textCase)
	if err != nil {
		response.Incorrect(ctx, http.StatusBadRequest, err)
		return
	}
	topics, err := s.Ctrl.FilterText(context.Background(), textCase.Text)
	if err != nil {
		response.Incorrect(ctx, http.StatusInternalServerError, err)
		return
	}
	resp := model.TextResp{
		Text:   textCase.Text,
		Topics: topics,
	}
	body, err := json.Marshal(&resp)
	if err != nil {
		response.Incorrect(ctx, http.StatusInternalServerError, err)
		return
	}
	response.SetBody(ctx, body, http.StatusOK)
}
