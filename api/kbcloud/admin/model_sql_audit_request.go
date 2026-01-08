// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SqlAuditRequest Request to update SQL audit log status
type SqlAuditRequest struct {
	// Component type to update SQL audit log for
	Component string `json:"component"`
	// Whether to enable SQL audit log
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlAuditRequest instantiates a new SqlAuditRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAuditRequest(component string, enabled bool) *SqlAuditRequest {
	this := SqlAuditRequest{}
	this.Component = component
	this.Enabled = enabled
	return &this
}

// NewSqlAuditRequestWithDefaults instantiates a new SqlAuditRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAuditRequestWithDefaults() *SqlAuditRequest {
	this := SqlAuditRequest{}
	return &this
}

// GetComponent returns the Component field value.
func (o *SqlAuditRequest) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *SqlAuditRequest) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *SqlAuditRequest) SetComponent(v string) {
	o.Component = v
}

// GetEnabled returns the Enabled field value.
func (o *SqlAuditRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SqlAuditRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *SqlAuditRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAuditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAuditRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component"`
		Enabled   *bool   `json:"enabled"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "enabled"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
