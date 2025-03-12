// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DatabaseOption struct {
	Enabled             bool    `json:"enabled"`
	DatabaseNamePattern *string `json:"databaseNamePattern,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseOption instantiates a new DatabaseOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseOption(enabled bool) *DatabaseOption {
	this := DatabaseOption{}
	this.Enabled = enabled
	return &this
}

// NewDatabaseOptionWithDefaults instantiates a new DatabaseOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseOptionWithDefaults() *DatabaseOption {
	this := DatabaseOption{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *DatabaseOption) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *DatabaseOption) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDatabaseNamePattern returns the DatabaseNamePattern field value if set, zero value otherwise.
func (o *DatabaseOption) GetDatabaseNamePattern() string {
	if o == nil || o.DatabaseNamePattern == nil {
		var ret string
		return ret
	}
	return *o.DatabaseNamePattern
}

// GetDatabaseNamePatternOk returns a tuple with the DatabaseNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOption) GetDatabaseNamePatternOk() (*string, bool) {
	if o == nil || o.DatabaseNamePattern == nil {
		return nil, false
	}
	return o.DatabaseNamePattern, true
}

// HasDatabaseNamePattern returns a boolean if a field has been set.
func (o *DatabaseOption) HasDatabaseNamePattern() bool {
	return o != nil && o.DatabaseNamePattern != nil
}

// SetDatabaseNamePattern gets a reference to the given string and assigns it to the DatabaseNamePattern field.
func (o *DatabaseOption) SetDatabaseNamePattern(v string) {
	o.DatabaseNamePattern = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.DatabaseNamePattern != nil {
		toSerialize["databaseNamePattern"] = o.DatabaseNamePattern
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabaseOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled             *bool   `json:"enabled"`
		DatabaseNamePattern *string `json:"databaseNamePattern,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "databaseNamePattern"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.DatabaseNamePattern = all.DatabaseNamePattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
