package controller

import (
	"context"
	api "test/server/grpc/api"
)

// Controller represents controller.
type Controller struct {
	JSONMap map[string]map[string]int
	Client  api.FilterTextClient
}

// New creates a new instance of the controller.
func New(jsonMap map[string]map[string]int, client api.FilterTextClient) *Controller {
	return &Controller{
		JSONMap: jsonMap,
		Client:  client,
	}
}

// GetTopics gets topic for a given word.
func (c *Controller) GetTopics(word string) *api.Topics {
	words := c.JSONMap[word]
	topics := make([]*api.Topic, 0)
	for k, v := range words {
		topics = append(topics, &api.Topic{
			Key:   k,
			Value: int32(v),
		})
	}
	return &api.Topics{Topics: topics}
}

// FilterText filters text.
func (c *Controller) FilterText(ctx context.Context, textCase string) (map[string]int64, error) {
	result, err := c.Client.FilterText(ctx, &api.Text{Text: textCase})
	if err != nil {
		return nil, err
	}
	return result.Result, nil
}
