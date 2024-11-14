// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentRestore struct {
	// Environment info
	Environment Environment `json:"environment"`
	// whether to install victoria metrics
	VictoriaMetrics *bool `json:"victoriaMetrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentRestore instantiates a new EnvironmentRestore object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentRestore(environment Environment) *EnvironmentRestore {
	this := EnvironmentRestore{}
	this.Environment = environment
	return &this
}

// NewEnvironmentRestoreWithDefaults instantiates a new EnvironmentRestore object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentRestoreWithDefaults() *EnvironmentRestore {
	this := EnvironmentRestore{}
	return &this
}

// GetEnvironment returns the Environment field value.
func (o *EnvironmentRestore) GetEnvironment() Environment {
	if o == nil {
		var ret Environment
		return ret
	}
	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *EnvironmentRestore) GetEnvironmentOk() (*Environment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value.
func (o *EnvironmentRestore) SetEnvironment(v Environment) {
	o.Environment = v
}

// GetVictoriaMetrics returns the VictoriaMetrics field value if set, zero value otherwise.
func (o *EnvironmentRestore) GetVictoriaMetrics() bool {
	if o == nil || o.VictoriaMetrics == nil {
		var ret bool
		return ret
	}
	return *o.VictoriaMetrics
}

// GetVictoriaMetricsOk returns a tuple with the VictoriaMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRestore) GetVictoriaMetricsOk() (*bool, bool) {
	if o == nil || o.VictoriaMetrics == nil {
		return nil, false
	}
	return o.VictoriaMetrics, true
}

// HasVictoriaMetrics returns a boolean if a field has been set.
func (o *EnvironmentRestore) HasVictoriaMetrics() bool {
	return o != nil && o.VictoriaMetrics != nil
}

// SetVictoriaMetrics gets a reference to the given bool and assigns it to the VictoriaMetrics field.
func (o *EnvironmentRestore) SetVictoriaMetrics(v bool) {
	o.VictoriaMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentRestore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["environment"] = o.Environment
	if o.VictoriaMetrics != nil {
		toSerialize["victoriaMetrics"] = o.VictoriaMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentRestore) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Environment     *Environment `json:"environment"`
		VictoriaMetrics *bool        `json:"victoriaMetrics,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Environment == nil {
		return fmt.Errorf("required field environment missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environment", "victoriaMetrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Environment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Environment = *all.Environment
	o.VictoriaMetrics = all.VictoriaMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
