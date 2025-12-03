// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DisasterRecoverySettings struct {
	SupportedSourceModes []string            `json:"supportedSourceModes,omitempty"`
	PromotionTargetModes map[string][]string `json:"promotionTargetModes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoverySettings instantiates a new DisasterRecoverySettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoverySettings() *DisasterRecoverySettings {
	this := DisasterRecoverySettings{}
	return &this
}

// NewDisasterRecoverySettingsWithDefaults instantiates a new DisasterRecoverySettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoverySettingsWithDefaults() *DisasterRecoverySettings {
	this := DisasterRecoverySettings{}
	return &this
}

// GetSupportedSourceModes returns the SupportedSourceModes field value if set, zero value otherwise.
func (o *DisasterRecoverySettings) GetSupportedSourceModes() []string {
	if o == nil || o.SupportedSourceModes == nil {
		var ret []string
		return ret
	}
	return o.SupportedSourceModes
}

// GetSupportedSourceModesOk returns a tuple with the SupportedSourceModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoverySettings) GetSupportedSourceModesOk() (*[]string, bool) {
	if o == nil || o.SupportedSourceModes == nil {
		return nil, false
	}
	return &o.SupportedSourceModes, true
}

// HasSupportedSourceModes returns a boolean if a field has been set.
func (o *DisasterRecoverySettings) HasSupportedSourceModes() bool {
	return o != nil && o.SupportedSourceModes != nil
}

// SetSupportedSourceModes gets a reference to the given []string and assigns it to the SupportedSourceModes field.
func (o *DisasterRecoverySettings) SetSupportedSourceModes(v []string) {
	o.SupportedSourceModes = v
}

// GetPromotionTargetModes returns the PromotionTargetModes field value if set, zero value otherwise.
func (o *DisasterRecoverySettings) GetPromotionTargetModes() map[string][]string {
	if o == nil || o.PromotionTargetModes == nil {
		var ret map[string][]string
		return ret
	}
	return o.PromotionTargetModes
}

// GetPromotionTargetModesOk returns a tuple with the PromotionTargetModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoverySettings) GetPromotionTargetModesOk() (*map[string][]string, bool) {
	if o == nil || o.PromotionTargetModes == nil {
		return nil, false
	}
	return &o.PromotionTargetModes, true
}

// HasPromotionTargetModes returns a boolean if a field has been set.
func (o *DisasterRecoverySettings) HasPromotionTargetModes() bool {
	return o != nil && o.PromotionTargetModes != nil
}

// SetPromotionTargetModes gets a reference to the given map[string][]string and assigns it to the PromotionTargetModes field.
func (o *DisasterRecoverySettings) SetPromotionTargetModes(v map[string][]string) {
	o.PromotionTargetModes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoverySettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SupportedSourceModes != nil {
		toSerialize["supportedSourceModes"] = o.SupportedSourceModes
	}
	if o.PromotionTargetModes != nil {
		toSerialize["promotionTargetModes"] = o.PromotionTargetModes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoverySettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SupportedSourceModes []string            `json:"supportedSourceModes,omitempty"`
		PromotionTargetModes map[string][]string `json:"promotionTargetModes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supportedSourceModes", "promotionTargetModes"})
	} else {
		return err
	}
	o.SupportedSourceModes = all.SupportedSourceModes
	o.PromotionTargetModes = all.PromotionTargetModes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
