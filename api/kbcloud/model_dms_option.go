// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION DmsOption
type DmsOption struct {
	// NODESCRIPTION Enabled
	Enabled bool `json:"enabled"`
	// NODESCRIPTION Protocol
	Protocol string `json:"protocol"`
	// NODESCRIPTION Feature
	Feature map[string]interface{} `json:"feature,omitempty"`
	// NODESCRIPTION TableMetadata
	TableMetadata []interface{} `json:"tableMetadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsOption instantiates a new DmsOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsOption(enabled bool, protocol string) *DmsOption {
	this := DmsOption{}
	this.Enabled = enabled
	this.Protocol = protocol
	return &this
}

// NewDmsOptionWithDefaults instantiates a new DmsOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsOptionWithDefaults() *DmsOption {
	this := DmsOption{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *DmsOption) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DmsOption) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *DmsOption) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProtocol returns the Protocol field value.
func (o *DmsOption) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *DmsOption) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value.
func (o *DmsOption) SetProtocol(v string) {
	o.Protocol = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *DmsOption) GetFeature() map[string]interface{} {
	if o == nil || o.Feature == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsOption) GetFeatureOk() (*map[string]interface{}, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return &o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *DmsOption) HasFeature() bool {
	return o != nil && o.Feature != nil
}

// SetFeature gets a reference to the given map[string]interface{} and assigns it to the Feature field.
func (o *DmsOption) SetFeature(v map[string]interface{}) {
	o.Feature = v
}

// GetTableMetadata returns the TableMetadata field value if set, zero value otherwise.
func (o *DmsOption) GetTableMetadata() []interface{} {
	if o == nil || o.TableMetadata == nil {
		var ret []interface{}
		return ret
	}
	return o.TableMetadata
}

// GetTableMetadataOk returns a tuple with the TableMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsOption) GetTableMetadataOk() (*[]interface{}, bool) {
	if o == nil || o.TableMetadata == nil {
		return nil, false
	}
	return &o.TableMetadata, true
}

// HasTableMetadata returns a boolean if a field has been set.
func (o *DmsOption) HasTableMetadata() bool {
	return o != nil && o.TableMetadata != nil
}

// SetTableMetadata gets a reference to the given []interface{} and assigns it to the TableMetadata field.
func (o *DmsOption) SetTableMetadata(v []interface{}) {
	o.TableMetadata = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["protocol"] = o.Protocol
	if o.Feature != nil {
		toSerialize["feature"] = o.Feature
	}
	if o.TableMetadata != nil {
		toSerialize["tableMetadata"] = o.TableMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       *bool                  `json:"enabled"`
		Protocol      *string                `json:"protocol"`
		Feature       map[string]interface{} `json:"feature,omitempty"`
		TableMetadata []interface{}          `json:"tableMetadata,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Protocol == nil {
		return fmt.Errorf("required field protocol missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "protocol", "feature", "tableMetadata"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Protocol = *all.Protocol
	o.Feature = all.Feature
	o.TableMetadata = all.TableMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
