// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClickhouseSpaceAnalysis struct {
	// Collection time in RFC3339 UTC.
	CollectedAt string `json:"collectedAt"`
	// success when all sources were collected, partial when at least one best-effort source failed.
	Status string `json:"status"`
	// Database filter applied to the collection, empty when all databases are included.
	Database *string                 `json:"database,omitempty"`
	Summary  *ClickhouseSpaceSummary `json:"summary,omitempty"`
	// Top tables ordered by physical bytes_on_disk. Totals are physical usage across all hosts, not deduplicated logical usage.
	Tables []ClickhouseTableSpace `json:"tables"`
	// Top partitions ordered by physical bytes_on_disk.
	Partitions []ClickhousePartitionSpace `json:"partitions"`
	// In-progress background merges from system.merges ordered by elapsed time.
	Merges []ClickhouseMerge `json:"merges"`
	// Unfinished mutations from system.mutations ordered by creation time.
	Mutations []ClickhouseMutation `json:"mutations"`
	// Per-source collection status.
	Sources []ClickhouseSource `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSpaceAnalysis instantiates a new ClickhouseSpaceAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSpaceAnalysis(collectedAt string, status string, tables []ClickhouseTableSpace, partitions []ClickhousePartitionSpace, merges []ClickhouseMerge, mutations []ClickhouseMutation, sources []ClickhouseSource) *ClickhouseSpaceAnalysis {
	this := ClickhouseSpaceAnalysis{}
	this.CollectedAt = collectedAt
	this.Status = status
	this.Tables = tables
	this.Partitions = partitions
	this.Merges = merges
	this.Mutations = mutations
	this.Sources = sources
	return &this
}

// NewClickhouseSpaceAnalysisWithDefaults instantiates a new ClickhouseSpaceAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSpaceAnalysisWithDefaults() *ClickhouseSpaceAnalysis {
	this := ClickhouseSpaceAnalysis{}
	return &this
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ClickhouseSpaceAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ClickhouseSpaceAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetStatus returns the Status field value.
func (o *ClickhouseSpaceAnalysis) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ClickhouseSpaceAnalysis) SetStatus(v string) {
	o.Status = v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClickhouseSpaceAnalysis) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClickhouseSpaceAnalysis) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClickhouseSpaceAnalysis) SetDatabase(v string) {
	o.Database = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ClickhouseSpaceAnalysis) GetSummary() ClickhouseSpaceSummary {
	if o == nil || o.Summary == nil {
		var ret ClickhouseSpaceSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetSummaryOk() (*ClickhouseSpaceSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ClickhouseSpaceAnalysis) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given ClickhouseSpaceSummary and assigns it to the Summary field.
func (o *ClickhouseSpaceAnalysis) SetSummary(v ClickhouseSpaceSummary) {
	o.Summary = &v
}

// GetTables returns the Tables field value.
func (o *ClickhouseSpaceAnalysis) GetTables() []ClickhouseTableSpace {
	if o == nil {
		var ret []ClickhouseTableSpace
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetTablesOk() (*[]ClickhouseTableSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tables, true
}

// SetTables sets field value.
func (o *ClickhouseSpaceAnalysis) SetTables(v []ClickhouseTableSpace) {
	o.Tables = v
}

// GetPartitions returns the Partitions field value.
func (o *ClickhouseSpaceAnalysis) GetPartitions() []ClickhousePartitionSpace {
	if o == nil {
		var ret []ClickhousePartitionSpace
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetPartitionsOk() (*[]ClickhousePartitionSpace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// SetPartitions sets field value.
func (o *ClickhouseSpaceAnalysis) SetPartitions(v []ClickhousePartitionSpace) {
	o.Partitions = v
}

// GetMerges returns the Merges field value.
func (o *ClickhouseSpaceAnalysis) GetMerges() []ClickhouseMerge {
	if o == nil {
		var ret []ClickhouseMerge
		return ret
	}
	return o.Merges
}

// GetMergesOk returns a tuple with the Merges field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetMergesOk() (*[]ClickhouseMerge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merges, true
}

// SetMerges sets field value.
func (o *ClickhouseSpaceAnalysis) SetMerges(v []ClickhouseMerge) {
	o.Merges = v
}

// GetMutations returns the Mutations field value.
func (o *ClickhouseSpaceAnalysis) GetMutations() []ClickhouseMutation {
	if o == nil {
		var ret []ClickhouseMutation
		return ret
	}
	return o.Mutations
}

// GetMutationsOk returns a tuple with the Mutations field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetMutationsOk() (*[]ClickhouseMutation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mutations, true
}

// SetMutations sets field value.
func (o *ClickhouseSpaceAnalysis) SetMutations(v []ClickhouseMutation) {
	o.Mutations = v
}

// GetSources returns the Sources field value.
func (o *ClickhouseSpaceAnalysis) GetSources() []ClickhouseSource {
	if o == nil {
		var ret []ClickhouseSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSpaceAnalysis) GetSourcesOk() (*[]ClickhouseSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ClickhouseSpaceAnalysis) SetSources(v []ClickhouseSource) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSpaceAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["collectedAt"] = o.CollectedAt
	toSerialize["status"] = o.Status
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	toSerialize["tables"] = o.Tables
	toSerialize["partitions"] = o.Partitions
	toSerialize["merges"] = o.Merges
	toSerialize["mutations"] = o.Mutations
	toSerialize["sources"] = o.Sources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSpaceAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CollectedAt *string                     `json:"collectedAt"`
		Status      *string                     `json:"status"`
		Database    *string                     `json:"database,omitempty"`
		Summary     *ClickhouseSpaceSummary     `json:"summary,omitempty"`
		Tables      *[]ClickhouseTableSpace     `json:"tables"`
		Partitions  *[]ClickhousePartitionSpace `json:"partitions"`
		Merges      *[]ClickhouseMerge          `json:"merges"`
		Mutations   *[]ClickhouseMutation       `json:"mutations"`
		Sources     *[]ClickhouseSource         `json:"sources"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Tables == nil {
		return fmt.Errorf("required field tables missing")
	}
	if all.Partitions == nil {
		return fmt.Errorf("required field partitions missing")
	}
	if all.Merges == nil {
		return fmt.Errorf("required field merges missing")
	}
	if all.Mutations == nil {
		return fmt.Errorf("required field mutations missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"collectedAt", "status", "database", "summary", "tables", "partitions", "merges", "mutations", "sources"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CollectedAt = *all.CollectedAt
	o.Status = *all.Status
	o.Database = all.Database
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary
	o.Tables = *all.Tables
	o.Partitions = *all.Partitions
	o.Merges = *all.Merges
	o.Mutations = *all.Mutations
	o.Sources = *all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
