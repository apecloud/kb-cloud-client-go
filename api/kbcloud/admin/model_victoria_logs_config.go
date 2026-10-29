// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type VictoriaLogsConfig struct {
	// Optional component overrides. Unknown, duplicate or other-engine components are rejected.
	Components []StaticClusterComponent `json:"components,omitempty"`
	// Environment installation log retention, explicitly overriding the addon default. Defaults to 100y when omitted. Accepts a number with s (seconds), h (hours), d (days), w (weeks), M (months), or y (years); a bare number means months. Must be between 1d and 100y. Disk usage retention limits still apply.
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
	// Maximum disk usage percentage before VictoriaLogs deletes older log partitions. Defaults to 80 when omitted. Applies independently of retentionPeriod.
	RetentionMaxDiskUsagePercent common.NullableInt32 `json:"retentionMaxDiskUsagePercent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVictoriaLogsConfig instantiates a new VictoriaLogsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVictoriaLogsConfig() *VictoriaLogsConfig {
	this := VictoriaLogsConfig{}
	var retentionPeriod string = "100y"
	this.RetentionPeriod = &retentionPeriod
	var retentionMaxDiskUsagePercent int32 = 80
	this.RetentionMaxDiskUsagePercent = *common.NewNullableInt32(&retentionMaxDiskUsagePercent)
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *VictoriaLogsConfig) GetComponents() []StaticClusterComponent {
	if o == nil || o.Components == nil {
		var ret []StaticClusterComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictoriaLogsConfig) GetComponentsOk() (*[]StaticClusterComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *VictoriaLogsConfig) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []StaticClusterComponent and assigns it to the Components field.
func (o *VictoriaLogsConfig) SetComponents(v []StaticClusterComponent) {
	o.Components = v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *VictoriaLogsConfig) GetRetentionPeriod() string {
	if o == nil || o.RetentionPeriod == nil {
		var ret string
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictoriaLogsConfig) GetRetentionPeriodOk() (*string, bool) {
	if o == nil || o.RetentionPeriod == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *VictoriaLogsConfig) HasRetentionPeriod() bool {
	return o != nil && o.RetentionPeriod != nil
}

// SetRetentionPeriod gets a reference to the given string and assigns it to the RetentionPeriod field.
func (o *VictoriaLogsConfig) SetRetentionPeriod(v string) {
	o.RetentionPeriod = &v
}

// GetRetentionMaxDiskUsagePercent returns the RetentionMaxDiskUsagePercent field value if set, zero value otherwise.
func (o *VictoriaLogsConfig) GetRetentionMaxDiskUsagePercent() int32 {
	if o == nil || o.RetentionMaxDiskUsagePercent.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RetentionMaxDiskUsagePercent.Get()
}

// GetRetentionMaxDiskUsagePercentOk returns a tuple with the RetentionMaxDiskUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictoriaLogsConfig) GetRetentionMaxDiskUsagePercentOk() (*int32, bool) {
	if o == nil || o.RetentionMaxDiskUsagePercent.Get() == nil {
		return nil, false
	}
	return o.RetentionMaxDiskUsagePercent.Get(), o.RetentionMaxDiskUsagePercent.IsSet()
}

// HasRetentionMaxDiskUsagePercent returns a boolean if a field has been set.
func (o *VictoriaLogsConfig) HasRetentionMaxDiskUsagePercent() bool {
	return o != nil && o.RetentionMaxDiskUsagePercent.IsSet()
}

// SetRetentionMaxDiskUsagePercent gets a reference to the given common.NullableInt32 and assigns it to the RetentionMaxDiskUsagePercent field.
func (o *VictoriaLogsConfig) SetRetentionMaxDiskUsagePercent(v int32) {
	o.RetentionMaxDiskUsagePercent.Set(&v)
}

// SetRetentionMaxDiskUsagePercentNil sets the value for RetentionMaxDiskUsagePercent to be an explicit nil.
func (o *VictoriaLogsConfig) SetRetentionMaxDiskUsagePercentNil() {
	o.RetentionMaxDiskUsagePercent.Set(nil)
}

// UnsetRetentionMaxDiskUsagePercent ensures that no value is present for RetentionMaxDiskUsagePercent, not even an explicit nil.
func (o *VictoriaLogsConfig) UnsetRetentionMaxDiskUsagePercent() {
	o.RetentionMaxDiskUsagePercent.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o VictoriaLogsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.RetentionPeriod != nil {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if o.RetentionMaxDiskUsagePercent.IsSet() {
		toSerialize["retentionMaxDiskUsagePercent"] = o.RetentionMaxDiskUsagePercent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VictoriaLogsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components                   []StaticClusterComponent `json:"components,omitempty"`
		RetentionPeriod              *string                  `json:"retentionPeriod,omitempty"`
		RetentionMaxDiskUsagePercent common.NullableInt32     `json:"retentionMaxDiskUsagePercent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"components", "retentionPeriod", "retentionMaxDiskUsagePercent"})
	} else {
		return err
	}
	o.Components = all.Components
	o.RetentionPeriod = all.RetentionPeriod
	o.RetentionMaxDiskUsagePercent = all.RetentionMaxDiskUsagePercent
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
