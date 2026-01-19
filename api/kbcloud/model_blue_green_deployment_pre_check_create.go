// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentPreCheckCreate blueGreenDeploymentPreCheckCreate defines the request to create a precheck task for a blue-green deployment.
type BlueGreenDeploymentPreCheckCreate struct {
	// KubeBlocks cluster information
	GreenCluster ClusterCreate `json:"greenCluster"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentPreCheckCreate instantiates a new BlueGreenDeploymentPreCheckCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentPreCheckCreate(greenCluster ClusterCreate) *BlueGreenDeploymentPreCheckCreate {
	this := BlueGreenDeploymentPreCheckCreate{}
	this.GreenCluster = greenCluster
	return &this
}

// NewBlueGreenDeploymentPreCheckCreateWithDefaults instantiates a new BlueGreenDeploymentPreCheckCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentPreCheckCreateWithDefaults() *BlueGreenDeploymentPreCheckCreate {
	this := BlueGreenDeploymentPreCheckCreate{}
	return &this
}

// GetGreenCluster returns the GreenCluster field value.
func (o *BlueGreenDeploymentPreCheckCreate) GetGreenCluster() ClusterCreate {
	if o == nil {
		var ret ClusterCreate
		return ret
	}
	return o.GreenCluster
}

// GetGreenClusterOk returns a tuple with the GreenCluster field value
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentPreCheckCreate) GetGreenClusterOk() (*ClusterCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GreenCluster, true
}

// SetGreenCluster sets field value.
func (o *BlueGreenDeploymentPreCheckCreate) SetGreenCluster(v ClusterCreate) {
	o.GreenCluster = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentPreCheckCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["greenCluster"] = o.GreenCluster

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentPreCheckCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GreenCluster *ClusterCreate `json:"greenCluster"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.GreenCluster == nil {
		return fmt.Errorf("required field greenCluster missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"greenCluster"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GreenCluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GreenCluster = *all.GreenCluster

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
