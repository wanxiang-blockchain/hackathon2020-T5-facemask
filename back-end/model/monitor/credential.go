package monitor

import (
	"context"
)

type customeCredential struct{}

func (c customeCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appkey": "this is my key",
	}, nil
}

func (c customeCredential) RequireTransportSecurity() bool {
	return true
}
