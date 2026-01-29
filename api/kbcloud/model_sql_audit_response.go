// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SqlAuditResponse Response after updating SQL audit log status
type SqlAuditResponse struct {
	// Whether SQL audit log is enabled after the operation
	Enabled bool `json:"enabled"`
	// Task ID if the operation is asynchronous (e.g. reconfigure)
	TaskId *string `json:"taskId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlAuditResponse instantiates a new SqlAuditResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAuditResponse(enabled bool) *SqlAuditResponse {
	this := SqlAuditResponse{}
	this.Enabled = enabled
	return &this
}

// NewSqlAuditResponseWithDefaults instantiates a new SqlAuditResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAuditResponseWithDefaults() *SqlAuditResponse {
	this := SqlAuditResponse{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *SqlAuditResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SqlAuditResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *SqlAuditResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *SqlAuditResponse) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditResponse) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *SqlAuditResponse) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *SqlAuditResponse) SetTaskId(v string) {
	o.TaskId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAuditResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAuditResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool   `json:"enabled"`
		TaskId  *string `json:"taskId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "taskId"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.TaskId = all.TaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
