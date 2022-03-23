package redis

import (
	"context"

	"github.com/go-redis/redis"
)

// Any key -> number
func (c *Client) GetValue(key string) (string, error) {

	value, err := c.client.Get(context.Background(), key).Result()
	if err == redis.Nil || value == "" {
		value = ""
		err = nil
	}

	return value, err
}

func (c *Client) SetValue(key string, value string) error {

	err := c.client.Set(context.Background(), key, value, 0).Err()

	return err
}

func (c *Client) GetAllKeys(pattern string) ([]string, error) {

	keysStringCmd := c.client.Keys(context.Background(), pattern)
	keys, err := keysStringCmd.Result()

	return keys, err
}
