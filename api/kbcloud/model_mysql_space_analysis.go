// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlSpaceAnalysis struct {
	Summary         MysqlSpaceSummary    `json:"summary"`
	StorageOverview MysqlStorageOverview `json:"storageOverview"`
	// Sizes of all non-system MySQL databases.
	Databases []MysqlDatabaseSpace `json:"databases"`
	// Database used for table and index details.
	SelectedDatabase string `json:"selectedDatabase"`
	// Top tables from selectedDatabase.
	Tables []MysqlTableSpace `json:"tables"`
	// Top indexes from selectedDatabase. Individual index size can be unavailable when mysql.innodb_index_stats is inaccessible.
	Indexes []MysqlIndexSpace  `json:"indexes"`
	Sources []MysqlSpaceSource `json:"sources"`
	// Backend collection timestamp in UTC.
	CollectedAt string `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlSpaceAnalysis instantiates a new MysqlSpaceAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlSpaceAnalysis(summary MysqlSpaceSummary, storageOverview MysqlStorageOverview, databases []MysqlDatabaseSpace, selectedDatabase string, tables []MysqlTableSpace, indexes []MysqlIndexSpace, sources []MysqlSpaceSource, collectedAt string) *MysqlSpaceAnalysis {
	this := MysqlSpaceAnalysis{}
	this.Summary = summary
	this.StorageOverview = storageOverview
	this.Databases = databases
	this.SelectedDatabase = selectedDatabase
	this.Tables = tables
	this.Indexes = indexes
	this.Sources = sources
	this.CollectedAt = collectedAt
	return &this
}

// NewMysqlSpaceAnalysisWithDefaults instantiates a new MysqlSpaceAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlSpaceAnalysisWithDefaults() *MysqlSpaceAnalysis {
	this := MysqlSpaceAnalysis{}
	return &this
}

// GetSummary returns the Summary field value.
func (o *MysqlSpaceAnalysis) GetSummary() MysqlSpaceSummary {
	if o == nil {
		var ret MysqlSpaceSummary
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetSummaryOk() (*MysqlSpaceSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *MysqlSpaceAnalysis) SetSummary(v MysqlSpaceSummary) {
	o.Summary = v
}

// GetStorageOverview returns the StorageOverview field value.
func (o *MysqlSpaceAnalysis) GetStorageOverview() MysqlStorageOverview {
	if o == nil {
		var ret MysqlStorageOverview
		return ret
	}
	return o.StorageOverview
}

// GetStorageOverviewOk returns a tuple with the StorageOverview field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetStorageOverviewOk() (*MysqlStorageOverview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageOverview, true
}

// SetStorageOverview sets field value.
func (o *MysqlSpaceAnalysis) SetStorageOverview(v MysqlStorageOverview) {
	o.StorageOverview = v
}

// GetDatabases returns the Databases field value.
func (o *MysqlSpaceAnalysis) GetDatabases() []MysqlDatabaseSpace {
	if o == nil {
		var ret []MysqlDatabaseSpace
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetDatabasesOk() (*[]MysqlDatabaseSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Databases, true
}

// SetDatabases sets field value.
func (o *MysqlSpaceAnalysis) SetDatabases(v []MysqlDatabaseSpace) {
	o.Databases = v
}

// GetSelectedDatabase returns the SelectedDatabase field value.
func (o *MysqlSpaceAnalysis) GetSelectedDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SelectedDatabase
}

// GetSelectedDatabaseOk returns a tuple with the SelectedDatabase field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetSelectedDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedDatabase, true
}

// SetSelectedDatabase sets field value.
func (o *MysqlSpaceAnalysis) SetSelectedDatabase(v string) {
	o.SelectedDatabase = v
}

// GetTables returns the Tables field value.
func (o *MysqlSpaceAnalysis) GetTables() []MysqlTableSpace {
	if o == nil {
		var ret []MysqlTableSpace
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetTablesOk() (*[]MysqlTableSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tables, true
}

// SetTables sets field value.
func (o *MysqlSpaceAnalysis) SetTables(v []MysqlTableSpace) {
	o.Tables = v
}

// GetIndexes returns the Indexes field value.
func (o *MysqlSpaceAnalysis) GetIndexes() []MysqlIndexSpace {
	if o == nil {
		var ret []MysqlIndexSpace
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetIndexesOk() (*[]MysqlIndexSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// SetIndexes sets field value.
func (o *MysqlSpaceAnalysis) SetIndexes(v []MysqlIndexSpace) {
	o.Indexes = v
}

// GetSources returns the Sources field value.
func (o *MysqlSpaceAnalysis) GetSources() []MysqlSpaceSource {
	if o == nil {
		var ret []MysqlSpaceSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetSourcesOk() (*[]MysqlSpaceSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *MysqlSpaceAnalysis) SetSources(v []MysqlSpaceSource) {
	o.Sources = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *MysqlSpaceAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *MysqlSpaceAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *MysqlSpaceAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlSpaceAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["summary"] = o.Summary
	toSerialize["storageOverview"] = o.StorageOverview
	toSerialize["databases"] = o.Databases
	toSerialize["selectedDatabase"] = o.SelectedDatabase
	toSerialize["tables"] = o.Tables
	toSerialize["indexes"] = o.Indexes
	toSerialize["sources"] = o.Sources
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlSpaceAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Summary          *MysqlSpaceSummary    `json:"summary"`
		StorageOverview  *MysqlStorageOverview `json:"storageOverview"`
		Databases        *[]MysqlDatabaseSpace `json:"databases"`
		SelectedDatabase *string               `json:"selectedDatabase"`
		Tables           *[]MysqlTableSpace    `json:"tables"`
		Indexes          *[]MysqlIndexSpace    `json:"indexes"`
		Sources          *[]MysqlSpaceSource   `json:"sources"`
		CollectedAt      *string               `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.StorageOverview == nil {
		return fmt.Errorf("required field storageOverview missing")
	}
	if all.Databases == nil {
		return fmt.Errorf("required field databases missing")
	}
	if all.SelectedDatabase == nil {
		return fmt.Errorf("required field selectedDatabase missing")
	}
	if all.Tables == nil {
		return fmt.Errorf("required field tables missing")
	}
	if all.Indexes == nil {
		return fmt.Errorf("required field indexes missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"summary", "storageOverview", "databases", "selectedDatabase", "tables", "indexes", "sources", "collectedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = *all.Summary
	if all.StorageOverview.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StorageOverview = *all.StorageOverview
	o.Databases = *all.Databases
	o.SelectedDatabase = *all.SelectedDatabase
	o.Tables = *all.Tables
	o.Indexes = *all.Indexes
	o.Sources = *all.Sources
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
