// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SqlAuditStatus SQL audit log status
type SqlAuditStatus struct {
	// Whether SQL audit log is enabled
	Enabled bool `json:"enabled"`
	// Additional configuration options for SQL audit log (engine-specific)
	Options map[string]interface{} `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlAuditStatus instantiates a new SqlAuditStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAuditStatus(enabled bool) *SqlAuditStatus {
	this := SqlAuditStatus{}
	this.Enabled = enabled
	return &this
}

// NewSqlAuditStatusWithDefaults instantiates a new SqlAuditStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAuditStatusWithDefaults() *SqlAuditStatus {
	this := SqlAuditStatus{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *SqlAuditStatus) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SqlAuditStatus) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *SqlAuditStatus) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SqlAuditStatus) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditStatus) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SqlAuditStatus) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *SqlAuditStatus) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAuditStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAuditStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                  `json:"enabled"`
		Options map[string]interface{} `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "options"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
