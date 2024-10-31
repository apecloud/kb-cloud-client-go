package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/apecloud/kb-cloud-client-go/api/kbcloud"
	"github.com/apecloud/kb-cloud-client-go/tests"
)

func TestCreateCluster(t *testing.T) {
	ctx := WithClient(WithTestAuth(NewDefaultContext(context.Background())))
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := kbcloud.NewClusterApi(client)

	orgName := "cloud-test"
	cluster := kbcloud.Cluster{
		Name:              "test-cluster",
		Engine:            "mysql",
		Mode:              common.PtrString("standalone"),
		Version:           common.PtrString("8.4.2"),
		Namespace:         common.PtrString("kubeblocks-cloud-ns"),
		SingleZone:        common.PtrBool(true),
		TerminationPolicy: common.Ptr(kbcloud.ClusterTerminationPolicyDelete),
		ProxyEnabled:      common.PtrBool(true),
		EnvironmentName:   "test",
		Components: []kbcloud.ComponentsItem{
			{
				Component: common.PtrString("mysql"),
				Replicas:  common.PtrInt32(1),
				Storage:   common.PtrString("20Gi"),
				ClassCode: common.PtrString("mysql.standalone.mysql.1c1g.g"),
			},
		},
		Backup: &kbcloud.ClusterBackup{
			RetentionPeriod: common.PtrString("7d"),
			CronExpression:  common.PtrString("0 2 * * *"),
			AutoBackup:      common.PtrBool(true),
			PitrEnabled:     common.PtrBool(true),
		},
	}

	createdCluster, resp, err := api.CreateCluster(ctx, orgName, cluster)

	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	assert.NotNil(createdCluster)
	assert.Equal(cluster.Name, createdCluster.Name)
}

func TestDeleteCluster(t *testing.T) {
	ctx := WithClient(WithTestAuth(NewDefaultContext(context.Background())))
	assert := tests.Assert(ctx, t)
	client := Client(ctx)
	api := kbcloud.NewClusterApi(client)

	orgName := "cloud-test"
	clusterName := "test-cluster"

	_, resp, err := api.DeleteCluster(ctx, orgName, clusterName)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
