package graphql

import (
	"context"

	"github.com/machinebox/graphql"
)

// IClient is the interface for the client
type IClient interface {
	Run(req *graphql.Request, variables map[string]interface{}, response interface{}) error
}

// Client is the client
type Client struct {
	c   *graphql.Client
	key string
}

// Run query data from the graphql client
func (c *Client) Run(req *graphql.Request, variables map[string]interface{}, response interface{}) error {
	req.Header.Set("Api-Key", c.key)

	//set variables
	for k, v := range variables {
		req.Var(k, v)
	}

	ctx := context.Background()
	if err := c.c.Run(ctx, req, response); err != nil {
		return err
	}

	return nil
}

// Query converts a string into a graphql.Request
func Query(q string) *graphql.Request {
	return graphql.NewRequest(q)
}

// NewClient creates a new client
func NewClient(url, key string) *Client {
	client := &Client{
		c:   graphql.NewClient(url),
		key: key,
	}

	return client
}
