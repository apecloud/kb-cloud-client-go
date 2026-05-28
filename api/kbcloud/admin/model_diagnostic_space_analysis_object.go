// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisObject One PostgreSQL relation in the Top space object list.
type DiagnosticSpaceAnalysisObject struct {
	// Rank after sorting by sizeBytes descending.
	Rank int64 `json:"rank"`
	// PostgreSQL database name.
	DatabaseName string `json:"databaseName"`
	// PostgreSQL schema name.
	SchemaName string `json:"schemaName"`
	// Relation name. This is an object identifier, not SQL text.
	ObjectName string `json:"objectName"`
	// PostgreSQL relation type bucket used by space analysis.
	ObjectType DiagnosticSpaceAnalysisObjectType `json:"objectType"`
	// Parent table name for indexes, when PostgreSQL exposes it.
	ParentObjectName *string `json:"parentObjectName,omitempty"`
	// Relation size in bytes.
	SizeBytes int64 `json:"sizeBytes"`
	// Human-readable relation size produced from sizeBytes.
	SizeHuman string `json:"sizeHuman"`
	// Relation size percentage of total database size.
	PercentOfTotal float64 `json:"percentOfTotal"`
	// Object space risk level derived only from percentOfTotal.
	RiskLevel DiagnosticSpaceAnalysisRiskLevel `json:"riskLevel"`
	// Time when the readonly collector produced this object row.
	CollectedAt time.Time `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSpaceAnalysisObject instantiates a new DiagnosticSpaceAnalysisObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSpaceAnalysisObject(rank int64, databaseName string, schemaName string, objectName string, objectType DiagnosticSpaceAnalysisObjectType, sizeBytes int64, sizeHuman string, percentOfTotal float64, riskLevel DiagnosticSpaceAnalysisRiskLevel, collectedAt time.Time) *DiagnosticSpaceAnalysisObject {
	this := DiagnosticSpaceAnalysisObject{}
	this.Rank = rank
	this.DatabaseName = databaseName
	this.SchemaName = schemaName
	this.ObjectName = objectName
	this.ObjectType = objectType
	this.SizeBytes = sizeBytes
	this.SizeHuman = sizeHuman
	this.PercentOfTotal = percentOfTotal
	this.RiskLevel = riskLevel
	this.CollectedAt = collectedAt
	return &this
}

// NewDiagnosticSpaceAnalysisObjectWithDefaults instantiates a new DiagnosticSpaceAnalysisObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSpaceAnalysisObjectWithDefaults() *DiagnosticSpaceAnalysisObject {
	this := DiagnosticSpaceAnalysisObject{}
	return &this
}

// GetRank returns the Rank field value.
func (o *DiagnosticSpaceAnalysisObject) GetRank() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetRankOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetRank(v int64) {
	o.Rank = v
}

// GetDatabaseName returns the DatabaseName field value.
func (o *DiagnosticSpaceAnalysisObject) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetSchemaName returns the SchemaName field value.
func (o *DiagnosticSpaceAnalysisObject) GetSchemaName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetSchemaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaName, true
}

// SetSchemaName sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetSchemaName(v string) {
	o.SchemaName = v
}

// GetObjectName returns the ObjectName field value.
func (o *DiagnosticSpaceAnalysisObject) GetObjectName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetObjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectName, true
}

// SetObjectName sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetObjectName(v string) {
	o.ObjectName = v
}

// GetObjectType returns the ObjectType field value.
func (o *DiagnosticSpaceAnalysisObject) GetObjectType() DiagnosticSpaceAnalysisObjectType {
	if o == nil {
		var ret DiagnosticSpaceAnalysisObjectType
		return ret
	}
	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetObjectTypeOk() (*DiagnosticSpaceAnalysisObjectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetObjectType(v DiagnosticSpaceAnalysisObjectType) {
	o.ObjectType = v
}

// GetParentObjectName returns the ParentObjectName field value if set, zero value otherwise.
func (o *DiagnosticSpaceAnalysisObject) GetParentObjectName() string {
	if o == nil || o.ParentObjectName == nil {
		var ret string
		return ret
	}
	return *o.ParentObjectName
}

// GetParentObjectNameOk returns a tuple with the ParentObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetParentObjectNameOk() (*string, bool) {
	if o == nil || o.ParentObjectName == nil {
		return nil, false
	}
	return o.ParentObjectName, true
}

// HasParentObjectName returns a boolean if a field has been set.
func (o *DiagnosticSpaceAnalysisObject) HasParentObjectName() bool {
	return o != nil && o.ParentObjectName != nil
}

// SetParentObjectName gets a reference to the given string and assigns it to the ParentObjectName field.
func (o *DiagnosticSpaceAnalysisObject) SetParentObjectName(v string) {
	o.ParentObjectName = &v
}

// GetSizeBytes returns the SizeBytes field value.
func (o *DiagnosticSpaceAnalysisObject) GetSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetSizeBytes(v int64) {
	o.SizeBytes = v
}

// GetSizeHuman returns the SizeHuman field value.
func (o *DiagnosticSpaceAnalysisObject) GetSizeHuman() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SizeHuman
}

// GetSizeHumanOk returns a tuple with the SizeHuman field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetSizeHumanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeHuman, true
}

// SetSizeHuman sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetSizeHuman(v string) {
	o.SizeHuman = v
}

// GetPercentOfTotal returns the PercentOfTotal field value.
func (o *DiagnosticSpaceAnalysisObject) GetPercentOfTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.PercentOfTotal
}

// GetPercentOfTotalOk returns a tuple with the PercentOfTotal field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetPercentOfTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentOfTotal, true
}

// SetPercentOfTotal sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetPercentOfTotal(v float64) {
	o.PercentOfTotal = v
}

// GetRiskLevel returns the RiskLevel field value.
func (o *DiagnosticSpaceAnalysisObject) GetRiskLevel() DiagnosticSpaceAnalysisRiskLevel {
	if o == nil {
		var ret DiagnosticSpaceAnalysisRiskLevel
		return ret
	}
	return o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetRiskLevelOk() (*DiagnosticSpaceAnalysisRiskLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskLevel, true
}

// SetRiskLevel sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetRiskLevel(v DiagnosticSpaceAnalysisRiskLevel) {
	o.RiskLevel = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticSpaceAnalysisObject) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisObject) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticSpaceAnalysisObject) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSpaceAnalysisObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["rank"] = o.Rank
	toSerialize["databaseName"] = o.DatabaseName
	toSerialize["schemaName"] = o.SchemaName
	toSerialize["objectName"] = o.ObjectName
	toSerialize["objectType"] = o.ObjectType
	if o.ParentObjectName != nil {
		toSerialize["parentObjectName"] = o.ParentObjectName
	}
	toSerialize["sizeBytes"] = o.SizeBytes
	toSerialize["sizeHuman"] = o.SizeHuman
	toSerialize["percentOfTotal"] = o.PercentOfTotal
	toSerialize["riskLevel"] = o.RiskLevel
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSpaceAnalysisObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rank             *int64                             `json:"rank"`
		DatabaseName     *string                            `json:"databaseName"`
		SchemaName       *string                            `json:"schemaName"`
		ObjectName       *string                            `json:"objectName"`
		ObjectType       *DiagnosticSpaceAnalysisObjectType `json:"objectType"`
		ParentObjectName *string                            `json:"parentObjectName,omitempty"`
		SizeBytes        *int64                             `json:"sizeBytes"`
		SizeHuman        *string                            `json:"sizeHuman"`
		PercentOfTotal   *float64                           `json:"percentOfTotal"`
		RiskLevel        *DiagnosticSpaceAnalysisRiskLevel  `json:"riskLevel"`
		CollectedAt      *time.Time                         `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Rank == nil {
		return fmt.Errorf("required field rank missing")
	}
	if all.DatabaseName == nil {
		return fmt.Errorf("required field databaseName missing")
	}
	if all.SchemaName == nil {
		return fmt.Errorf("required field schemaName missing")
	}
	if all.ObjectName == nil {
		return fmt.Errorf("required field objectName missing")
	}
	if all.ObjectType == nil {
		return fmt.Errorf("required field objectType missing")
	}
	if all.SizeBytes == nil {
		return fmt.Errorf("required field sizeBytes missing")
	}
	if all.SizeHuman == nil {
		return fmt.Errorf("required field sizeHuman missing")
	}
	if all.PercentOfTotal == nil {
		return fmt.Errorf("required field percentOfTotal missing")
	}
	if all.RiskLevel == nil {
		return fmt.Errorf("required field riskLevel missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"rank", "databaseName", "schemaName", "objectName", "objectType", "parentObjectName", "sizeBytes", "sizeHuman", "percentOfTotal", "riskLevel", "collectedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Rank = *all.Rank
	o.DatabaseName = *all.DatabaseName
	o.SchemaName = *all.SchemaName
	o.ObjectName = *all.ObjectName
	if !all.ObjectType.IsValid() {
		hasInvalidField = true
	} else {
		o.ObjectType = *all.ObjectType
	}
	o.ParentObjectName = all.ParentObjectName
	o.SizeBytes = *all.SizeBytes
	o.SizeHuman = *all.SizeHuman
	o.PercentOfTotal = *all.PercentOfTotal
	if !all.RiskLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.RiskLevel = *all.RiskLevel
	}
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
