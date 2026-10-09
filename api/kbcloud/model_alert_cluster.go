// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// AlertCluster Cluster-scoped alert settings. PATCH updates only supplied fields; settings apply to both tenant and administrator notifications.
type AlertCluster struct {
	// Disable all cluster alerts. Existing rule and silence settings are preserved.
	Disabled *bool `json:"disabled,omitempty"`
	// Alert names disabled only for this cluster. An empty array clears all rule overrides; omission preserves them.
	DisabledRules []string `json:"disabledRules,omitempty"`
	// Pause notifications while continuing to record alerts. Set enabled to false to cancel. Time windows include their start and exclude their end.
	Silence *AlertClusterSilence `json:"silence,omitempty"`
	// Whether the configured silence window is active now, regardless of the alert master switch.
	SilenceActive *bool `json:"silenceActive,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertCluster instantiates a new AlertCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertCluster() *AlertCluster {
	this := AlertCluster{}
	return &this
}

// NewAlertClusterWithDefaults instantiates a new AlertCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertClusterWithDefaults() *AlertCluster {
	this := AlertCluster{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AlertCluster) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCluster) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AlertCluster) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AlertCluster) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDisabledRules returns the DisabledRules field value if set, zero value otherwise.
func (o *AlertCluster) GetDisabledRules() []string {
	if o == nil || o.DisabledRules == nil {
		var ret []string
		return ret
	}
	return o.DisabledRules
}

// GetDisabledRulesOk returns a tuple with the DisabledRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCluster) GetDisabledRulesOk() (*[]string, bool) {
	if o == nil || o.DisabledRules == nil {
		return nil, false
	}
	return &o.DisabledRules, true
}

// HasDisabledRules returns a boolean if a field has been set.
func (o *AlertCluster) HasDisabledRules() bool {
	return o != nil && o.DisabledRules != nil
}

// SetDisabledRules gets a reference to the given []string and assigns it to the DisabledRules field.
func (o *AlertCluster) SetDisabledRules(v []string) {
	o.DisabledRules = v
}

// GetSilence returns the Silence field value if set, zero value otherwise.
func (o *AlertCluster) GetSilence() AlertClusterSilence {
	if o == nil || o.Silence == nil {
		var ret AlertClusterSilence
		return ret
	}
	return *o.Silence
}

// GetSilenceOk returns a tuple with the Silence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCluster) GetSilenceOk() (*AlertClusterSilence, bool) {
	if o == nil || o.Silence == nil {
		return nil, false
	}
	return o.Silence, true
}

// HasSilence returns a boolean if a field has been set.
func (o *AlertCluster) HasSilence() bool {
	return o != nil && o.Silence != nil
}

// SetSilence gets a reference to the given AlertClusterSilence and assigns it to the Silence field.
func (o *AlertCluster) SetSilence(v AlertClusterSilence) {
	o.Silence = &v
}

// GetSilenceActive returns the SilenceActive field value if set, zero value otherwise.
func (o *AlertCluster) GetSilenceActive() bool {
	if o == nil || o.SilenceActive == nil {
		var ret bool
		return ret
	}
	return *o.SilenceActive
}

// GetSilenceActiveOk returns a tuple with the SilenceActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCluster) GetSilenceActiveOk() (*bool, bool) {
	if o == nil || o.SilenceActive == nil {
		return nil, false
	}
	return o.SilenceActive, true
}

// HasSilenceActive returns a boolean if a field has been set.
func (o *AlertCluster) HasSilenceActive() bool {
	return o != nil && o.SilenceActive != nil
}

// SetSilenceActive gets a reference to the given bool and assigns it to the SilenceActive field.
func (o *AlertCluster) SetSilenceActive(v bool) {
	o.SilenceActive = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.DisabledRules != nil {
		toSerialize["disabledRules"] = o.DisabledRules
	}
	if o.Silence != nil {
		toSerialize["silence"] = o.Silence
	}
	if o.SilenceActive != nil {
		toSerialize["silenceActive"] = o.SilenceActive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Disabled      *bool                `json:"disabled,omitempty"`
		DisabledRules []string             `json:"disabledRules,omitempty"`
		Silence       *AlertClusterSilence `json:"silence,omitempty"`
		SilenceActive *bool                `json:"silenceActive,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"disabled", "disabledRules", "silence", "silenceActive"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Disabled = all.Disabled
	o.DisabledRules = all.DisabledRules
	if all.Silence != nil && all.Silence.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Silence = all.Silence
	o.SilenceActive = all.SilenceActive

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
