// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionSupportedEngines Specifies the engine type for the inspection task. Supports all database engines.
type InspectionSupportedEngines string

// List of InspectionSupportedEngines.
const (
	InspectionSupportedEnginesNode               InspectionSupportedEngines = "node"
	InspectionSupportedEnginesCommon             InspectionSupportedEngines = "common"
	InspectionSupportedEnginesCamelliaRedisProxy InspectionSupportedEngines = "camellia-redis-proxy"
	InspectionSupportedEnginesClickhouse         InspectionSupportedEngines = "clickhouse"
	InspectionSupportedEnginesCubetran           InspectionSupportedEngines = "cubetran"
	InspectionSupportedEnginesDamengdb           InspectionSupportedEngines = "damengdb"
	InspectionSupportedEnginesDoris              InspectionSupportedEngines = "doris"
	InspectionSupportedEnginesElasticsearch      InspectionSupportedEngines = "elasticsearch"
	InspectionSupportedEnginesEtcd               InspectionSupportedEngines = "etcd"
	InspectionSupportedEnginesFoundationdb       InspectionSupportedEngines = "foundationdb"
	InspectionSupportedEnginesGaussdb            InspectionSupportedEngines = "gaussdb"
	InspectionSupportedEnginesGoldendb           InspectionSupportedEngines = "goldendb"
	InspectionSupportedEnginesGreatdb            InspectionSupportedEngines = "greatdb"
	InspectionSupportedEnginesHadoop             InspectionSupportedEngines = "hadoop"
	InspectionSupportedEnginesHive               InspectionSupportedEngines = "hive"
	InspectionSupportedEnginesInfluxdb           InspectionSupportedEngines = "influxdb"
	InspectionSupportedEnginesKafka              InspectionSupportedEngines = "kafka"
	InspectionSupportedEnginesKingbase           InspectionSupportedEngines = "kingbase"
	InspectionSupportedEnginesLoki               InspectionSupportedEngines = "loki"
	InspectionSupportedEnginesMilvus             InspectionSupportedEngines = "milvus"
	InspectionSupportedEnginesMinio              InspectionSupportedEngines = "minio"
	InspectionSupportedEnginesMogdb              InspectionSupportedEngines = "mogdb"
	InspectionSupportedEnginesMongodb            InspectionSupportedEngines = "mongodb"
	InspectionSupportedEnginesMssql              InspectionSupportedEngines = "mssql"
	InspectionSupportedEnginesMysql              InspectionSupportedEngines = "mysql"
	InspectionSupportedEnginesNacos              InspectionSupportedEngines = "nacos"
	InspectionSupportedEnginesNebula             InspectionSupportedEngines = "nebula"
	InspectionSupportedEnginesOceanbase          InspectionSupportedEngines = "oceanbase"
	InspectionSupportedEnginesOceanbaseProxy     InspectionSupportedEngines = "oceanbase-proxy"
	InspectionSupportedEnginesOracle             InspectionSupportedEngines = "oracle"
	InspectionSupportedEnginesPostgresql         InspectionSupportedEngines = "postgresql"
	InspectionSupportedEnginesPulsar             InspectionSupportedEngines = "pulsar"
	InspectionSupportedEnginesQdrant             InspectionSupportedEngines = "qdrant"
	InspectionSupportedEnginesRabbitmq           InspectionSupportedEngines = "rabbitmq"
	InspectionSupportedEnginesRedis              InspectionSupportedEngines = "redis"
	InspectionSupportedEnginesRocketmq           InspectionSupportedEngines = "rocketmq"
	InspectionSupportedEnginesSelectdb           InspectionSupportedEngines = "selectdb"
	InspectionSupportedEnginesStarrocks          InspectionSupportedEngines = "starrocks"
	InspectionSupportedEnginesTdengine           InspectionSupportedEngines = "tdengine"
	InspectionSupportedEnginesTdsql              InspectionSupportedEngines = "tdsql"
	InspectionSupportedEnginesTidb               InspectionSupportedEngines = "tidb"
	InspectionSupportedEnginesVastbase           InspectionSupportedEngines = "vastbase"
	InspectionSupportedEnginesVault              InspectionSupportedEngines = "vault"
	InspectionSupportedEnginesVictoriaLogs       InspectionSupportedEngines = "victoria-logs"
	InspectionSupportedEnginesVictoriametrics    InspectionSupportedEngines = "victoriametrics"
	InspectionSupportedEnginesZookeeper          InspectionSupportedEngines = "zookeeper"
)

var allowedInspectionSupportedEnginesEnumValues = []InspectionSupportedEngines{
	InspectionSupportedEnginesNode,
	InspectionSupportedEnginesCommon,
	InspectionSupportedEnginesCamelliaRedisProxy,
	InspectionSupportedEnginesClickhouse,
	InspectionSupportedEnginesCubetran,
	InspectionSupportedEnginesDamengdb,
	InspectionSupportedEnginesDoris,
	InspectionSupportedEnginesElasticsearch,
	InspectionSupportedEnginesEtcd,
	InspectionSupportedEnginesFoundationdb,
	InspectionSupportedEnginesGaussdb,
	InspectionSupportedEnginesGoldendb,
	InspectionSupportedEnginesGreatdb,
	InspectionSupportedEnginesHadoop,
	InspectionSupportedEnginesHive,
	InspectionSupportedEnginesInfluxdb,
	InspectionSupportedEnginesKafka,
	InspectionSupportedEnginesKingbase,
	InspectionSupportedEnginesLoki,
	InspectionSupportedEnginesMilvus,
	InspectionSupportedEnginesMinio,
	InspectionSupportedEnginesMogdb,
	InspectionSupportedEnginesMongodb,
	InspectionSupportedEnginesMssql,
	InspectionSupportedEnginesMysql,
	InspectionSupportedEnginesNacos,
	InspectionSupportedEnginesNebula,
	InspectionSupportedEnginesOceanbase,
	InspectionSupportedEnginesOceanbaseProxy,
	InspectionSupportedEnginesOracle,
	InspectionSupportedEnginesPostgresql,
	InspectionSupportedEnginesPulsar,
	InspectionSupportedEnginesQdrant,
	InspectionSupportedEnginesRabbitmq,
	InspectionSupportedEnginesRedis,
	InspectionSupportedEnginesRocketmq,
	InspectionSupportedEnginesSelectdb,
	InspectionSupportedEnginesStarrocks,
	InspectionSupportedEnginesTdengine,
	InspectionSupportedEnginesTdsql,
	InspectionSupportedEnginesTidb,
	InspectionSupportedEnginesVastbase,
	InspectionSupportedEnginesVault,
	InspectionSupportedEnginesVictoriaLogs,
	InspectionSupportedEnginesVictoriametrics,
	InspectionSupportedEnginesZookeeper,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionSupportedEngines) GetAllowedValues() []InspectionSupportedEngines {
	return allowedInspectionSupportedEnginesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionSupportedEngines) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionSupportedEngines(value)
	return nil
}

// NewInspectionSupportedEnginesFromValue returns a pointer to a valid InspectionSupportedEngines
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionSupportedEnginesFromValue(v string) (*InspectionSupportedEngines, error) {
	ev := InspectionSupportedEngines(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionSupportedEngines: valid values are %v", v, allowedInspectionSupportedEnginesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionSupportedEngines) IsValid() bool {
	for _, existing := range allowedInspectionSupportedEnginesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionSupportedEngines value.
func (v InspectionSupportedEngines) Ptr() *InspectionSupportedEngines {
	return &v
}
