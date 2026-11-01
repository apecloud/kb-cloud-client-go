// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type StaticCluster struct {
	// Optional component overrides. Unknown, duplicate or other-engine components are rejected.
	Components []StaticClusterComponent `json:"components,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStaticCluster instantiates a new StaticCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStaticCluster() *StaticCluster {
	this := StaticCluster{}
	return &this
}

// NewStaticClusterWithDefaults instantiates a new StaticCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStaticClusterWithDefaults() *StaticCluster {
	this := StaticCluster{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *StaticCluster) GetComponents() []StaticClusterComponent {
	if o == nil || o.Components == nil {
		var ret []StaticClusterComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCluster) GetComponentsOk() (*[]StaticClusterComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *StaticCluster) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []StaticClusterComponent and assigns it to the Components field.
func (o *StaticCluster) SetComponents(v []StaticClusterComponent) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StaticCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StaticCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components []StaticClusterComponent `json:"components,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"components"})
	} else {
		return err
	}
	o.Components = all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
