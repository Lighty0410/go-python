package controller

import (
	"context"
	"test/controller/mocks"
	api "test/server/grpc/api"
	"test/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	jsonMap := make(map[string]map[string]int)
	jsonMap["hello"] = map[string]int{
		"hello": 1,
	}
	client := mocks.MockClient{}
	ctrl := Controller{
		JSONMap: jsonMap,
		Client:  &client,
	}
	require.Equal(t, New(jsonMap, &client), &ctrl)
}

func TestGetTopics(t *testing.T) {
	tt := []struct {
		name           string
		word           string
		expectedTopics *api.Topics
	}{
		{
			name: "1st element",
			word: "1",
			expectedTopics: &api.Topics{
				Topics: []*api.Topic{
					{
						Key:   "0",
						Value: 1,
					},
				},
			},
		},
		{
			name: "test element",
			word: "test",
			expectedTopics: &api.Topics{
				Topics: []*api.Topic{
					{
						Key:   "0",
						Value: 1,
					},
					{
						Key:   "12",
						Value: 2,
					},
				},
			},
		},
		{
			name: "69 element",
			word: "69",
			expectedTopics: &api.Topics{
				Topics: []*api.Topic{
					{
						Key:   "0",
						Value: 2,
					},
					{
						Key:   "4",
						Value: 3,
					},
				},
			},
		},
	}
	jsonMap, err := utils.CreateJSONMap("../data.json")
	if err != nil {
		t.Fatalf("cannot fill json map: %v\n", err)
	}
	ctrl := Controller{
		JSONMap: jsonMap,
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			topics := ctrl.GetTopics(tc.word)
			for _, topic := range topics.Topics {
				for _, tp := range tc.expectedTopics.Topics {
					if topic.Key == tp.Key {
						require.Equal(t, topic.Value, tp.Value)
					}
				}
			}
		})
	}
}

func TestFilterText(t *testing.T) {
	tt := []struct {
		name          string
		textCase      string
		expectedMap   map[string]int64
		expectedError string
	}{
		{
			name:          "bad request",
			textCase:      "bad text",
			expectedMap:   nil,
			expectedError: "unacceptable text",
		},
		{
			name:     "ok",
			textCase: "ok",
			expectedMap: map[string]int64{
				"ok": 1,
			},
			expectedError: "",
		},
	}
	client := mocks.MockClient{}
	ctrl := Controller{
		Client: &client,
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ctrl.FilterText(context.Background(), tc.textCase)
			if err != nil {
				require.EqualError(t, err, tc.expectedError)
			} else {
				require.Equal(t, result, tc.expectedMap)
			}
		})
	}
}
