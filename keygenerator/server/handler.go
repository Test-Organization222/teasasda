package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type KeyHandler struct {
}

func (c *KeyHandler) GetKey(ctx context.Context, in *empty.Empty) (*Key, error) {
	return nil, nil
}
