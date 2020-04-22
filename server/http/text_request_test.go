package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"test/model"
	"test/server/http/mock"

	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

func TestTestRequest(t *testing.T) {
	badText, err := json.Marshal(model.TextCase{Text: "bad text"})
	require.NoError(t, err)
	okText, err := json.Marshal(model.TextCase{Text: "1"})
	require.NoError(t, err)

	tt := []struct {
		name           string
		body           []byte
		expectedResp   []byte
		expectedStatus int
	}{
		{
			name:           "bad body",
			body:           []byte("badbody"),
			expectedStatus: http.StatusBadRequest,
			expectedResp:   []byte(`{"reason":"invalid character 'b' looking for beginning of value"}`),
		},
		{
			name:           "bad request",
			body:           badText,
			expectedResp:   []byte(`{"reason":"bad text"}`),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "ok",
			body:           okText,
			expectedResp:   []byte(`{"text":"1","topics":{"ok":1}}`),
			expectedStatus: http.StatusOK,
		},
	}
	ctrl := mock.Service{}
	server := NewHTTPServer(&ctrl)
	fastServer := fasthttp.Server{Handler: server.Run}
	go func() {
		err := fastServer.ListenAndServe(":8081")
		if err != nil {
			log.Println(err)
		}
	}()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client := http.Client{}
			body := bytes.NewReader(tc.body)
			resp, err := client.Post("http://localhost:8081/test-text", "application/json", body)
			require.NoError(t, err)
			defer resp.Body.Close()
			JSON, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, JSON, tc.expectedResp)
			require.Equal(t, resp.StatusCode, tc.expectedStatus)
		})
	}
}
