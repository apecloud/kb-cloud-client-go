// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchStorageAnalysis struct {
	Indices     []ElasticsearchIndexStorage   `json:"indices"`
	Allocations []ElasticsearchNodeAllocation `json:"allocations"`
	Watermarks  ElasticsearchDiskWatermarks   `json:"watermarks"`
	Sources     []PerformanceTrendSource      `json:"sources"`
	Warnings    []string                      `json:"warnings"`
	CollectedAt string                        `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchStorageAnalysis instantiates a new ElasticsearchStorageAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchStorageAnalysis(indices []ElasticsearchIndexStorage, allocations []ElasticsearchNodeAllocation, watermarks ElasticsearchDiskWatermarks, sources []PerformanceTrendSource, warnings []string, collectedAt string) *ElasticsearchStorageAnalysis {
	this := ElasticsearchStorageAnalysis{}
	this.Indices = indices
	this.Allocations = allocations
	this.Watermarks = watermarks
	this.Sources = sources
	this.Warnings = warnings
	this.CollectedAt = collectedAt
	return &this
}

// NewElasticsearchStorageAnalysisWithDefaults instantiates a new ElasticsearchStorageAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchStorageAnalysisWithDefaults() *ElasticsearchStorageAnalysis {
	this := ElasticsearchStorageAnalysis{}
	return &this
}

// GetIndices returns the Indices field value.
func (o *ElasticsearchStorageAnalysis) GetIndices() []ElasticsearchIndexStorage {
	if o == nil {
		var ret []ElasticsearchIndexStorage
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchStorageAnalysis) GetIndicesOk() (*[]ElasticsearchIndexStorage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indices, true
}

// SetIndices sets field value.
func (o *ElasticsearchStorageAnalysis) SetIndices(v []ElasticsearchIndexStorage) {
	o.Indices = v
}

// GetAllocations returns the Allocations field value.
func (o *ElasticsearchStorageAnalysis) GetAllocations() []ElasticsearchNodeAllocation {
	if o == nil {
		var ret []ElasticsearchNodeAllocation
		return ret
	}
	return o.Allocations
}

// GetAllocationsOk returns a tuple with the Allocations field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchStorageAnalysis) GetAllocationsOk() (*[]ElasticsearchNodeAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allocations, true
}

// SetAllocations sets field value.
func (o *ElasticsearchStorageAnalysis) SetAllocations(v []ElasticsearchNodeAllocation) {
	o.Allocations = v
}

// GetWatermarks returns the Watermarks field value.
func (o *ElasticsearchStorageAnalysis) GetWatermarks() ElasticsearchDiskWatermarks {
	if o == nil {
		var ret ElasticsearchDiskWatermarks
		return ret
	}
	return o.Watermarks
}

// GetWatermarksOk returns a tuple with the Watermarks field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchStorageAnalysis) GetWatermarksOk() (*ElasticsearchDiskWatermarks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Watermarks, true
}

// SetWatermarks sets field value.
func (o *ElasticsearchStorageAnalysis) SetWatermarks(v ElasticsearchDiskWatermarks) {
	o.Watermarks = v
}

// GetSources returns the Sources field value.
func (o *ElasticsearchStorageAnalysis) GetSources() []PerformanceTrendSource {
	if o == nil {
		var ret []PerformanceTrendSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchStorageAnalysis) GetSourcesOk() (*[]PerformanceTrendSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ElasticsearchStorageAnalysis) SetSources(v []PerformanceTrendSource) {
	o.Sources = v
}

// GetWarnings returns the Warnings field value.
func (o *ElasticsearchStorageAnalysis) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchStorageAnalysis) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *ElasticsearchStorageAnalysis) SetWarnings(v []string) {
	o.Warnings = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ElasticsearchStorageAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchStorageAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ElasticsearchStorageAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchStorageAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["indices"] = o.Indices
	toSerialize["allocations"] = o.Allocations
	toSerialize["watermarks"] = o.Watermarks
	toSerialize["sources"] = o.Sources
	toSerialize["warnings"] = o.Warnings
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchStorageAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Indices     *[]ElasticsearchIndexStorage   `json:"indices"`
		Allocations *[]ElasticsearchNodeAllocation `json:"allocations"`
		Watermarks  *ElasticsearchDiskWatermarks   `json:"watermarks"`
		Sources     *[]PerformanceTrendSource      `json:"sources"`
		Warnings    *[]string                      `json:"warnings"`
		CollectedAt *string                        `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Indices == nil {
		return fmt.Errorf("required field indices missing")
	}
	if all.Allocations == nil {
		return fmt.Errorf("required field allocations missing")
	}
	if all.Watermarks == nil {
		return fmt.Errorf("required field watermarks missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"indices", "allocations", "watermarks", "sources", "warnings", "collectedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Indices = *all.Indices
	o.Allocations = *all.Allocations
	if all.Watermarks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Watermarks = *all.Watermarks
	o.Sources = *all.Sources
	o.Warnings = *all.Warnings
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
