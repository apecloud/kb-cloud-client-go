// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterModeTransition Parameters for switching the engine mode of a cluster.
type ClusterModeTransition struct {
	// Target engine mode.
	Mode string `json:"mode"`
	// OpsHScale is the payload to horizontally scale a KubeBlocks cluster. It requires specifying either the number of replicas or the number of shards.
	HScale OpsHScale `json:"hScale"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterModeTransition instantiates a new ClusterModeTransition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterModeTransition(mode string, hScale OpsHScale) *ClusterModeTransition {
	this := ClusterModeTransition{}
	this.Mode = mode
	this.HScale = hScale
	return &this
}

// NewClusterModeTransitionWithDefaults instantiates a new ClusterModeTransition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterModeTransitionWithDefaults() *ClusterModeTransition {
	this := ClusterModeTransition{}
	return &this
}

// GetMode returns the Mode field value.
func (o *ClusterModeTransition) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ClusterModeTransition) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ClusterModeTransition) SetMode(v string) {
	o.Mode = v
}

// GetHScale returns the HScale field value.
func (o *ClusterModeTransition) GetHScale() OpsHScale {
	if o == nil {
		var ret OpsHScale
		return ret
	}
	return o.HScale
}

// GetHScaleOk returns a tuple with the HScale field value
// and a boolean to check if the value has been set.
func (o *ClusterModeTransition) GetHScaleOk() (*OpsHScale, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HScale, true
}

// SetHScale sets field value.
func (o *ClusterModeTransition) SetHScale(v OpsHScale) {
	o.HScale = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterModeTransition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	toSerialize["hScale"] = o.HScale

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterModeTransition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode   *string    `json:"mode"`
		HScale *OpsHScale `json:"hScale"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.HScale == nil {
		return fmt.Errorf("required field hScale missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mode", "hScale"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Mode = *all.Mode
	if all.HScale.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.HScale = *all.HScale

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
