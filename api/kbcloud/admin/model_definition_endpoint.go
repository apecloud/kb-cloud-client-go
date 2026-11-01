// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DefinitionEndpoint Reference to an existing data endpoint in the same organization.
type DefinitionEndpoint struct {
	EndpointId string `json:"endpointID"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDefinitionEndpoint instantiates a new DefinitionEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDefinitionEndpoint(endpointId string) *DefinitionEndpoint {
	this := DefinitionEndpoint{}
	this.EndpointId = endpointId
	return &this
}

// NewDefinitionEndpointWithDefaults instantiates a new DefinitionEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDefinitionEndpointWithDefaults() *DefinitionEndpoint {
	this := DefinitionEndpoint{}
	return &this
}

// GetEndpointId returns the EndpointId field value.
func (o *DefinitionEndpoint) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *DefinitionEndpoint) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value.
func (o *DefinitionEndpoint) SetEndpointId(v string) {
	o.EndpointId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DefinitionEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["endpointID"] = o.EndpointId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DefinitionEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndpointId *string `json:"endpointID"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EndpointId == nil {
		return fmt.Errorf("required field endpointID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"endpointID"})
	} else {
		return err
	}
	o.EndpointId = *all.EndpointId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
