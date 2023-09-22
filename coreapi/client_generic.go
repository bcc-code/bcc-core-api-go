package coreapi

import (
	"context"
	"net/http"

	"github.com/bcc-code/bcc-core-api-go/models"
	"github.com/go-openapi/strfmt"
)

type genericClient[R any, W any] struct {
	client *Client
	path   string
}

func (c *genericClient[R, W]) Get(ctx context.Context, uid strfmt.UUID, opts ...RequestOption) (models.Wrapped[R], error) {
	var res models.Wrapped[R]
	err := c.client.Request(ctx, http.MethodGet, c.client.URL(c.path, string(uid)), nil, &res, opts...)
	return res, err
}

func (c *genericClient[R, W]) Find(ctx context.Context, opts ...RequestOption) (models.WrappedWithMeta[[]R], error) {
	var res models.WrappedWithMeta[[]R]
	err := c.client.Request(ctx, http.MethodGet, c.client.URL(c.path), nil, &res, opts...)
	return res, err
}

func (c *genericClient[R, W]) Create(ctx context.Context, aff W) (models.Wrapped[R], error) {
	var res models.Wrapped[R]
	err := c.client.Request(ctx, http.MethodPost, c.client.URL(c.path), aff, &res)
	return res, err
}

func (c *genericClient[R, W]) Update(ctx context.Context, uid strfmt.UUID, aff W) (models.Wrapped[R], error) {
	var res models.Wrapped[R]
	err := c.client.Request(ctx, http.MethodPut, c.client.URL(c.path, uid.String()), aff, &res)
	return res, err
}

func (c *genericClient[R, W]) Remove(ctx context.Context, uid strfmt.UUID) (models.Wrapped[R], error) {
	var res models.Wrapped[R]
	err := c.client.Request(ctx, http.MethodDelete, c.client.URL(c.path, uid.String()), nil, &res)
	return res, err
}
