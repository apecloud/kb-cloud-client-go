// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type InspectionTaskByEngine struct {
	Engine  *string                        `json:"engine,omitempty"`
	Summary *InspectionTaskByEngineSummary `json:"summary,omitempty"`
	Tasks   []InspectionTask               `json:"tasks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionTaskByEngine instantiates a new InspectionTaskByEngine object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionTaskByEngine() *InspectionTaskByEngine {
	this := InspectionTaskByEngine{}
	return &this
}

// NewInspectionTaskByEngineWithDefaults instantiates a new InspectionTaskByEngine object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionTaskByEngineWithDefaults() *InspectionTaskByEngine {
	this := InspectionTaskByEngine{}
	return &this
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *InspectionTaskByEngine) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskByEngine) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *InspectionTaskByEngine) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *InspectionTaskByEngine) SetEngine(v string) {
	o.Engine = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *InspectionTaskByEngine) GetSummary() InspectionTaskByEngineSummary {
	if o == nil || o.Summary == nil {
		var ret InspectionTaskByEngineSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskByEngine) GetSummaryOk() (*InspectionTaskByEngineSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *InspectionTaskByEngine) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given InspectionTaskByEngineSummary and assigns it to the Summary field.
func (o *InspectionTaskByEngine) SetSummary(v InspectionTaskByEngineSummary) {
	o.Summary = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *InspectionTaskByEngine) GetTasks() []InspectionTask {
	if o == nil || o.Tasks == nil {
		var ret []InspectionTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskByEngine) GetTasksOk() (*[]InspectionTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *InspectionTaskByEngine) HasTasks() bool {
	return o != nil && o.Tasks != nil
}

// SetTasks gets a reference to the given []InspectionTask and assigns it to the Tasks field.
func (o *InspectionTaskByEngine) SetTasks(v []InspectionTask) {
	o.Tasks = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InspectionTaskByEngine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InspectionTaskByEngine) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine  *string                        `json:"engine,omitempty"`
		Summary *InspectionTaskByEngineSummary `json:"summary,omitempty"`
		Tasks   []InspectionTask               `json:"tasks,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "summary", "tasks"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Engine = all.Engine
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary
	o.Tasks = all.Tasks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
