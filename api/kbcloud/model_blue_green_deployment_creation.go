// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentCreation blueGreenDeploymentCreation defines the request to create a blue-green deployment.
type BlueGreenDeploymentCreation struct {
	// The name of the blue-green deployment.
	DeploymentName *string `json:"deploymentName,omitempty"`
	// KubeBlocks cluster information
	GreenCluster ClusterCreate `json:"greenCluster"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentCreation instantiates a new BlueGreenDeploymentCreation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentCreation(greenCluster ClusterCreate) *BlueGreenDeploymentCreation {
	this := BlueGreenDeploymentCreation{}
	this.GreenCluster = greenCluster
	return &this
}

// NewBlueGreenDeploymentCreationWithDefaults instantiates a new BlueGreenDeploymentCreation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentCreationWithDefaults() *BlueGreenDeploymentCreation {
	this := BlueGreenDeploymentCreation{}
	return &this
}

// GetDeploymentName returns the DeploymentName field value if set, zero value otherwise.
func (o *BlueGreenDeploymentCreation) GetDeploymentName() string {
	if o == nil || o.DeploymentName == nil {
		var ret string
		return ret
	}
	return *o.DeploymentName
}

// GetDeploymentNameOk returns a tuple with the DeploymentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentCreation) GetDeploymentNameOk() (*string, bool) {
	if o == nil || o.DeploymentName == nil {
		return nil, false
	}
	return o.DeploymentName, true
}

// HasDeploymentName returns a boolean if a field has been set.
func (o *BlueGreenDeploymentCreation) HasDeploymentName() bool {
	return o != nil && o.DeploymentName != nil
}

// SetDeploymentName gets a reference to the given string and assigns it to the DeploymentName field.
func (o *BlueGreenDeploymentCreation) SetDeploymentName(v string) {
	o.DeploymentName = &v
}

// GetGreenCluster returns the GreenCluster field value.
func (o *BlueGreenDeploymentCreation) GetGreenCluster() ClusterCreate {
	if o == nil {
		var ret ClusterCreate
		return ret
	}
	return o.GreenCluster
}

// GetGreenClusterOk returns a tuple with the GreenCluster field value
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentCreation) GetGreenClusterOk() (*ClusterCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GreenCluster, true
}

// SetGreenCluster sets field value.
func (o *BlueGreenDeploymentCreation) SetGreenCluster(v ClusterCreate) {
	o.GreenCluster = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DeploymentName != nil {
		toSerialize["deploymentName"] = o.DeploymentName
	}
	toSerialize["greenCluster"] = o.GreenCluster

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentCreation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeploymentName *string        `json:"deploymentName,omitempty"`
		GreenCluster   *ClusterCreate `json:"greenCluster"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.GreenCluster == nil {
		return fmt.Errorf("required field greenCluster missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"deploymentName", "greenCluster"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DeploymentName = all.DeploymentName
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
