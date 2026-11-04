// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBDrainRequest struct {
	Source         string   `json:"source"`
	RemainingNodes []string `json:"remainingNodes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBDrainRequest instantiates a new DolphinDBDrainRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBDrainRequest(source string, remainingNodes []string) *DolphinDBDrainRequest {
	this := DolphinDBDrainRequest{}
	this.Source = source
	this.RemainingNodes = remainingNodes
	return &this
}

// NewDolphinDBDrainRequestWithDefaults instantiates a new DolphinDBDrainRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBDrainRequestWithDefaults() *DolphinDBDrainRequest {
	this := DolphinDBDrainRequest{}
	return &this
}

// GetSource returns the Source field value.
func (o *DolphinDBDrainRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DolphinDBDrainRequest) SetSource(v string) {
	o.Source = v
}

// GetRemainingNodes returns the RemainingNodes field value.
func (o *DolphinDBDrainRequest) GetRemainingNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RemainingNodes
}

// GetRemainingNodesOk returns a tuple with the RemainingNodes field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainRequest) GetRemainingNodesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingNodes, true
}

// SetRemainingNodes sets field value.
func (o *DolphinDBDrainRequest) SetRemainingNodes(v []string) {
	o.RemainingNodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBDrainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	toSerialize["remainingNodes"] = o.RemainingNodes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBDrainRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source         *string   `json:"source"`
		RemainingNodes *[]string `json:"remainingNodes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.RemainingNodes == nil {
		return fmt.Errorf("required field remainingNodes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "remainingNodes"})
	} else {
		return err
	}
	o.Source = *all.Source
	o.RemainingNodes = *all.RemainingNodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
