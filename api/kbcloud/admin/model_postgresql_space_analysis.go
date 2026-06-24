// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSpaceAnalysis struct {
	Summary         PostgresqlSpaceSummary    `json:"summary"`
	StorageOverview PostgresqlStorageOverview `json:"storageOverview"`
	// Sizes of all connectable non-template databases in the PostgreSQL instance.
	Databases []PostgresqlDatabaseSpace `json:"databases"`
	// Database used for tables, indexes, and toastRelations in this response.
	SelectedDatabase string `json:"selectedDatabase"`
	// Top tables from selectedDatabase.
	Tables []PostgresqlTableSpace `json:"tables"`
	// Top indexes from selectedDatabase.
	Indexes []PostgresqlIndexSpace `json:"indexes"`
	// Toast relations derived from selectedDatabase table space data.
	ToastRelations []PostgresqlToastRelationSpace `json:"toastRelations"`
	Sources        []PostgresqlSpaceSource        `json:"sources"`
	// Backend collection timestamp in UTC.
	CollectedAt string `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSpaceAnalysis instantiates a new PostgresqlSpaceAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSpaceAnalysis(summary PostgresqlSpaceSummary, storageOverview PostgresqlStorageOverview, databases []PostgresqlDatabaseSpace, selectedDatabase string, tables []PostgresqlTableSpace, indexes []PostgresqlIndexSpace, toastRelations []PostgresqlToastRelationSpace, sources []PostgresqlSpaceSource, collectedAt string) *PostgresqlSpaceAnalysis {
	this := PostgresqlSpaceAnalysis{}
	this.Summary = summary
	this.StorageOverview = storageOverview
	this.Databases = databases
	this.SelectedDatabase = selectedDatabase
	this.Tables = tables
	this.Indexes = indexes
	this.ToastRelations = toastRelations
	this.Sources = sources
	this.CollectedAt = collectedAt
	return &this
}

// NewPostgresqlSpaceAnalysisWithDefaults instantiates a new PostgresqlSpaceAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSpaceAnalysisWithDefaults() *PostgresqlSpaceAnalysis {
	this := PostgresqlSpaceAnalysis{}
	return &this
}

// GetSummary returns the Summary field value.
func (o *PostgresqlSpaceAnalysis) GetSummary() PostgresqlSpaceSummary {
	if o == nil {
		var ret PostgresqlSpaceSummary
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetSummaryOk() (*PostgresqlSpaceSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *PostgresqlSpaceAnalysis) SetSummary(v PostgresqlSpaceSummary) {
	o.Summary = v
}

// GetStorageOverview returns the StorageOverview field value.
func (o *PostgresqlSpaceAnalysis) GetStorageOverview() PostgresqlStorageOverview {
	if o == nil {
		var ret PostgresqlStorageOverview
		return ret
	}
	return o.StorageOverview
}

// GetStorageOverviewOk returns a tuple with the StorageOverview field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetStorageOverviewOk() (*PostgresqlStorageOverview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageOverview, true
}

// SetStorageOverview sets field value.
func (o *PostgresqlSpaceAnalysis) SetStorageOverview(v PostgresqlStorageOverview) {
	o.StorageOverview = v
}

// GetDatabases returns the Databases field value.
func (o *PostgresqlSpaceAnalysis) GetDatabases() []PostgresqlDatabaseSpace {
	if o == nil {
		var ret []PostgresqlDatabaseSpace
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetDatabasesOk() (*[]PostgresqlDatabaseSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Databases, true
}

// SetDatabases sets field value.
func (o *PostgresqlSpaceAnalysis) SetDatabases(v []PostgresqlDatabaseSpace) {
	o.Databases = v
}

// GetSelectedDatabase returns the SelectedDatabase field value.
func (o *PostgresqlSpaceAnalysis) GetSelectedDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SelectedDatabase
}

// GetSelectedDatabaseOk returns a tuple with the SelectedDatabase field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetSelectedDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedDatabase, true
}

// SetSelectedDatabase sets field value.
func (o *PostgresqlSpaceAnalysis) SetSelectedDatabase(v string) {
	o.SelectedDatabase = v
}

// GetTables returns the Tables field value.
func (o *PostgresqlSpaceAnalysis) GetTables() []PostgresqlTableSpace {
	if o == nil {
		var ret []PostgresqlTableSpace
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetTablesOk() (*[]PostgresqlTableSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tables, true
}

// SetTables sets field value.
func (o *PostgresqlSpaceAnalysis) SetTables(v []PostgresqlTableSpace) {
	o.Tables = v
}

// GetIndexes returns the Indexes field value.
func (o *PostgresqlSpaceAnalysis) GetIndexes() []PostgresqlIndexSpace {
	if o == nil {
		var ret []PostgresqlIndexSpace
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetIndexesOk() (*[]PostgresqlIndexSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// SetIndexes sets field value.
func (o *PostgresqlSpaceAnalysis) SetIndexes(v []PostgresqlIndexSpace) {
	o.Indexes = v
}

// GetToastRelations returns the ToastRelations field value.
func (o *PostgresqlSpaceAnalysis) GetToastRelations() []PostgresqlToastRelationSpace {
	if o == nil {
		var ret []PostgresqlToastRelationSpace
		return ret
	}
	return o.ToastRelations
}

// GetToastRelationsOk returns a tuple with the ToastRelations field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetToastRelationsOk() (*[]PostgresqlToastRelationSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToastRelations, true
}

// SetToastRelations sets field value.
func (o *PostgresqlSpaceAnalysis) SetToastRelations(v []PostgresqlToastRelationSpace) {
	o.ToastRelations = v
}

// GetSources returns the Sources field value.
func (o *PostgresqlSpaceAnalysis) GetSources() []PostgresqlSpaceSource {
	if o == nil {
		var ret []PostgresqlSpaceSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetSourcesOk() (*[]PostgresqlSpaceSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *PostgresqlSpaceAnalysis) SetSources(v []PostgresqlSpaceSource) {
	o.Sources = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *PostgresqlSpaceAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *PostgresqlSpaceAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSpaceAnalysis) MarshalJSON() ([]byte, error) {
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
	toSerialize["toastRelations"] = o.ToastRelations
	toSerialize["sources"] = o.Sources
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSpaceAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Summary          *PostgresqlSpaceSummary         `json:"summary"`
		StorageOverview  *PostgresqlStorageOverview      `json:"storageOverview"`
		Databases        *[]PostgresqlDatabaseSpace      `json:"databases"`
		SelectedDatabase *string                         `json:"selectedDatabase"`
		Tables           *[]PostgresqlTableSpace         `json:"tables"`
		Indexes          *[]PostgresqlIndexSpace         `json:"indexes"`
		ToastRelations   *[]PostgresqlToastRelationSpace `json:"toastRelations"`
		Sources          *[]PostgresqlSpaceSource        `json:"sources"`
		CollectedAt      *string                         `json:"collectedAt"`
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
	if all.ToastRelations == nil {
		return fmt.Errorf("required field toastRelations missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"summary", "storageOverview", "databases", "selectedDatabase", "tables", "indexes", "toastRelations", "sources", "collectedAt"})
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
	o.ToastRelations = *all.ToastRelations
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
