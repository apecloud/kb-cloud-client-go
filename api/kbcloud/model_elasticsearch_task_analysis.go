// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchTaskAnalysis struct {
	// List of currently running tasks, sorted by running time in descending order
	Tasks       []ElasticsearchTask       `json:"tasks"`
	HotThreads  []ElasticsearchHotThreads `json:"hotThreads"`
	Sources     []PerformanceTrendSource  `json:"sources"`
	Warnings    []string                  `json:"warnings"`
	CollectedAt string                    `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchTaskAnalysis instantiates a new ElasticsearchTaskAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchTaskAnalysis(tasks []ElasticsearchTask, hotThreads []ElasticsearchHotThreads, sources []PerformanceTrendSource, warnings []string, collectedAt string) *ElasticsearchTaskAnalysis {
	this := ElasticsearchTaskAnalysis{}
	this.Tasks = tasks
	this.HotThreads = hotThreads
	this.Sources = sources
	this.Warnings = warnings
	this.CollectedAt = collectedAt
	return &this
}

// NewElasticsearchTaskAnalysisWithDefaults instantiates a new ElasticsearchTaskAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchTaskAnalysisWithDefaults() *ElasticsearchTaskAnalysis {
	this := ElasticsearchTaskAnalysis{}
	return &this
}

// GetTasks returns the Tasks field value.
func (o *ElasticsearchTaskAnalysis) GetTasks() []ElasticsearchTask {
	if o == nil {
		var ret []ElasticsearchTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTaskAnalysis) GetTasksOk() (*[]ElasticsearchTask, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// SetTasks sets field value.
func (o *ElasticsearchTaskAnalysis) SetTasks(v []ElasticsearchTask) {
	o.Tasks = v
}

// GetHotThreads returns the HotThreads field value.
func (o *ElasticsearchTaskAnalysis) GetHotThreads() []ElasticsearchHotThreads {
	if o == nil {
		var ret []ElasticsearchHotThreads
		return ret
	}
	return o.HotThreads
}

// GetHotThreadsOk returns a tuple with the HotThreads field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTaskAnalysis) GetHotThreadsOk() (*[]ElasticsearchHotThreads, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HotThreads, true
}

// SetHotThreads sets field value.
func (o *ElasticsearchTaskAnalysis) SetHotThreads(v []ElasticsearchHotThreads) {
	o.HotThreads = v
}

// GetSources returns the Sources field value.
func (o *ElasticsearchTaskAnalysis) GetSources() []PerformanceTrendSource {
	if o == nil {
		var ret []PerformanceTrendSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTaskAnalysis) GetSourcesOk() (*[]PerformanceTrendSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ElasticsearchTaskAnalysis) SetSources(v []PerformanceTrendSource) {
	o.Sources = v
}

// GetWarnings returns the Warnings field value.
func (o *ElasticsearchTaskAnalysis) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTaskAnalysis) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *ElasticsearchTaskAnalysis) SetWarnings(v []string) {
	o.Warnings = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ElasticsearchTaskAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchTaskAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ElasticsearchTaskAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchTaskAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["tasks"] = o.Tasks
	toSerialize["hotThreads"] = o.HotThreads
	toSerialize["sources"] = o.Sources
	toSerialize["warnings"] = o.Warnings
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchTaskAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tasks       *[]ElasticsearchTask       `json:"tasks"`
		HotThreads  *[]ElasticsearchHotThreads `json:"hotThreads"`
		Sources     *[]PerformanceTrendSource  `json:"sources"`
		Warnings    *[]string                  `json:"warnings"`
		CollectedAt *string                    `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Tasks == nil {
		return fmt.Errorf("required field tasks missing")
	}
	if all.HotThreads == nil {
		return fmt.Errorf("required field hotThreads missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"tasks", "hotThreads", "sources", "warnings", "collectedAt"})
	} else {
		return err
	}
	o.Tasks = *all.Tasks
	o.HotThreads = *all.HotThreads
	o.Sources = *all.Sources
	o.Warnings = *all.Warnings
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
