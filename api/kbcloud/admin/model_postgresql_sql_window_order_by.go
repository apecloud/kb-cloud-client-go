// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowOrderBy string

// List of PostgresqlSQLWindowOrderBy.
const (
	PostgresqlSQLWindowOrderByTotalTime            PostgresqlSQLWindowOrderBy = "totalTime"
	PostgresqlSQLWindowOrderByMeanTime             PostgresqlSQLWindowOrderBy = "meanTime"
	PostgresqlSQLWindowOrderByCalls                PostgresqlSQLWindowOrderBy = "calls"
	PostgresqlSQLWindowOrderByRows                 PostgresqlSQLWindowOrderBy = "rows"
	PostgresqlSQLWindowOrderBySharedBlksHit        PostgresqlSQLWindowOrderBy = "sharedBlksHit"
	PostgresqlSQLWindowOrderBySharedBlksRead       PostgresqlSQLWindowOrderBy = "sharedBlksRead"
	PostgresqlSQLWindowOrderBySharedBlksDirtied    PostgresqlSQLWindowOrderBy = "sharedBlksDirtied"
	PostgresqlSQLWindowOrderBySharedBlksWritten    PostgresqlSQLWindowOrderBy = "sharedBlksWritten"
	PostgresqlSQLWindowOrderByLocalBlksHit         PostgresqlSQLWindowOrderBy = "localBlksHit"
	PostgresqlSQLWindowOrderByLocalBlksRead        PostgresqlSQLWindowOrderBy = "localBlksRead"
	PostgresqlSQLWindowOrderByLocalBlksDirtied     PostgresqlSQLWindowOrderBy = "localBlksDirtied"
	PostgresqlSQLWindowOrderByLocalBlksWritten     PostgresqlSQLWindowOrderBy = "localBlksWritten"
	PostgresqlSQLWindowOrderByTempBlksRead         PostgresqlSQLWindowOrderBy = "tempBlksRead"
	PostgresqlSQLWindowOrderByTempBlksWritten      PostgresqlSQLWindowOrderBy = "tempBlksWritten"
	PostgresqlSQLWindowOrderByPlans                PostgresqlSQLWindowOrderBy = "plans"
	PostgresqlSQLWindowOrderByPlanningTimeMs       PostgresqlSQLWindowOrderBy = "planningTimeMs"
	PostgresqlSQLWindowOrderByWalRecords           PostgresqlSQLWindowOrderBy = "walRecords"
	PostgresqlSQLWindowOrderByWalFpi               PostgresqlSQLWindowOrderBy = "walFpi"
	PostgresqlSQLWindowOrderByWalBytes             PostgresqlSQLWindowOrderBy = "walBytes"
	PostgresqlSQLWindowOrderBySharedBlkReadTimeMs  PostgresqlSQLWindowOrderBy = "sharedBlkReadTimeMs"
	PostgresqlSQLWindowOrderBySharedBlkWriteTimeMs PostgresqlSQLWindowOrderBy = "sharedBlkWriteTimeMs"
	PostgresqlSQLWindowOrderByLocalBlkReadTimeMs   PostgresqlSQLWindowOrderBy = "localBlkReadTimeMs"
	PostgresqlSQLWindowOrderByLocalBlkWriteTimeMs  PostgresqlSQLWindowOrderBy = "localBlkWriteTimeMs"
	PostgresqlSQLWindowOrderByTempBlkReadTimeMs    PostgresqlSQLWindowOrderBy = "tempBlkReadTimeMs"
	PostgresqlSQLWindowOrderByTempBlkWriteTimeMs   PostgresqlSQLWindowOrderBy = "tempBlkWriteTimeMs"
	PostgresqlSQLWindowOrderByBlkReadTimeMs        PostgresqlSQLWindowOrderBy = "blkReadTimeMs"
	PostgresqlSQLWindowOrderByBlkWriteTimeMs       PostgresqlSQLWindowOrderBy = "blkWriteTimeMs"
)

var allowedPostgresqlSQLWindowOrderByEnumValues = []PostgresqlSQLWindowOrderBy{
	PostgresqlSQLWindowOrderByTotalTime,
	PostgresqlSQLWindowOrderByMeanTime,
	PostgresqlSQLWindowOrderByCalls,
	PostgresqlSQLWindowOrderByRows,
	PostgresqlSQLWindowOrderBySharedBlksHit,
	PostgresqlSQLWindowOrderBySharedBlksRead,
	PostgresqlSQLWindowOrderBySharedBlksDirtied,
	PostgresqlSQLWindowOrderBySharedBlksWritten,
	PostgresqlSQLWindowOrderByLocalBlksHit,
	PostgresqlSQLWindowOrderByLocalBlksRead,
	PostgresqlSQLWindowOrderByLocalBlksDirtied,
	PostgresqlSQLWindowOrderByLocalBlksWritten,
	PostgresqlSQLWindowOrderByTempBlksRead,
	PostgresqlSQLWindowOrderByTempBlksWritten,
	PostgresqlSQLWindowOrderByPlans,
	PostgresqlSQLWindowOrderByPlanningTimeMs,
	PostgresqlSQLWindowOrderByWalRecords,
	PostgresqlSQLWindowOrderByWalFpi,
	PostgresqlSQLWindowOrderByWalBytes,
	PostgresqlSQLWindowOrderBySharedBlkReadTimeMs,
	PostgresqlSQLWindowOrderBySharedBlkWriteTimeMs,
	PostgresqlSQLWindowOrderByLocalBlkReadTimeMs,
	PostgresqlSQLWindowOrderByLocalBlkWriteTimeMs,
	PostgresqlSQLWindowOrderByTempBlkReadTimeMs,
	PostgresqlSQLWindowOrderByTempBlkWriteTimeMs,
	PostgresqlSQLWindowOrderByBlkReadTimeMs,
	PostgresqlSQLWindowOrderByBlkWriteTimeMs,
}

// GetAllowedValues returns the list of possible values.
func (v *PostgresqlSQLWindowOrderBy) GetAllowedValues() []PostgresqlSQLWindowOrderBy {
	return allowedPostgresqlSQLWindowOrderByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostgresqlSQLWindowOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostgresqlSQLWindowOrderBy(value)
	return nil
}

// NewPostgresqlSQLWindowOrderByFromValue returns a pointer to a valid PostgresqlSQLWindowOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostgresqlSQLWindowOrderByFromValue(v string) (*PostgresqlSQLWindowOrderBy, error) {
	ev := PostgresqlSQLWindowOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostgresqlSQLWindowOrderBy: valid values are %v", v, allowedPostgresqlSQLWindowOrderByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostgresqlSQLWindowOrderBy) IsValid() bool {
	for _, existing := range allowedPostgresqlSQLWindowOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostgresqlSQLWindowOrderBy value.
func (v PostgresqlSQLWindowOrderBy) Ptr() *PostgresqlSQLWindowOrderBy {
	return &v
}
