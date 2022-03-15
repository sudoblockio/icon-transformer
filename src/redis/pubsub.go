package redis

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func (c *Client) Publish(channelName string, data []byte) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Publish
		err := c.client.Publish(ctx, channelName, string(data)).Err()
		if err != nil {
			// Failure
			zap.S().Warn("Redis Publish: Cannot publish message...retrying in 3 second")
			time.Sleep(3 * time.Second)

			continue
		}

		// Success
		break
	}
}
