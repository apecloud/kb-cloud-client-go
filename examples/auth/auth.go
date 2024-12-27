package auth

import (
	"context"
	"os"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NewAuthContext creates a new context with authentication
func NewAuthContext() context.Context {
	ctx := context.WithValue(
		context.Background(),
		common.ContextDigestAuth,
		common.DigestAuth{
			UserName: os.Getenv("KB_CLOUD_API_KEY_NAME"),
			Password: os.Getenv("KB_CLOUD_API_KEY_SECRET"),
		},
	)

	return context.WithValue(
		ctx,
		common.ContextServerVariables,
		map[string]string{"site": os.Getenv("KB_CLOUD_SITE")},
	)
}
