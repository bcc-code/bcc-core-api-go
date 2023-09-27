package bcccoreapi

import (
	"context"
	"net/http"
)

// Resets test data, only available if the SDK is run against the emulator
func (c *Client) RefreshTestData() error {
	return c.Request(context.Background(), http.MethodPost, c.URL("refresh-data"), nil, nil)
}
