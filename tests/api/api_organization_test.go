package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/apecloud/kb-cloud-client-go/api/kbcloud"
	"github.com/apecloud/kb-cloud-client-go/tests"
)

func TestListOrgs(t *testing.T) {
	ctx := WithClient(WithTestAuth(NewDefaultContext(context.Background())))
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := kbcloud.NewOrganizationApi(client)

	orgs, resp, err := api.ListOrg(ctx)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	assert.NotNil(orgs)
}
