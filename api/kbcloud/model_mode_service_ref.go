// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeServiceRef Defines a ServiceRef for a cluster, enabling access to both external services and
// Services provided by other Clusters. The defined serviceRef must be provided when creating cluster.
type ModeServiceRef struct {
	// The name will be referenced in clusterCreate request.
	Name string `json:"name"`
	// The engine to be used in serviceRef. Frontend can use this field to filter clusters.
	EngineName string `json:"engineName"`
	// separate values with commas.
	HelmValuePath ModeServiceRefHelmValuePath `json:"helmValuePath"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeServiceRef instantiates a new ModeServiceRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeServiceRef(name string, engineName string, helmValuePath ModeServiceRefHelmValuePath) *ModeServiceRef {
	this := ModeServiceRef{}
	this.Name = name
	this.EngineName = engineName
	this.HelmValuePath = helmValuePath
	return &this
}

// NewModeServiceRefWithDefaults instantiates a new ModeServiceRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeServiceRefWithDefaults() *ModeServiceRef {
	this := ModeServiceRef{}
	return &this
}

// GetName returns the Name field value.
func (o *ModeServiceRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ModeServiceRef) SetName(v string) {
	o.Name = v
}

// GetEngineName returns the EngineName field value.
func (o *ModeServiceRef) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *ModeServiceRef) SetEngineName(v string) {
	o.EngineName = v
}

// GetHelmValuePath returns the HelmValuePath field value.
func (o *ModeServiceRef) GetHelmValuePath() ModeServiceRefHelmValuePath {
	if o == nil {
		var ret ModeServiceRefHelmValuePath
		return ret
	}
	return o.HelmValuePath
}

// GetHelmValuePathOk returns a tuple with the HelmValuePath field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetHelmValuePathOk() (*ModeServiceRefHelmValuePath, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HelmValuePath, true
}

// SetHelmValuePath sets field value.
func (o *ModeServiceRef) SetHelmValuePath(v ModeServiceRefHelmValuePath) {
	o.HelmValuePath = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeServiceRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["engineName"] = o.EngineName
	toSerialize["helmValuePath"] = o.HelmValuePath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeServiceRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name          *string                      `json:"name"`
		EngineName    *string                      `json:"engineName"`
		HelmValuePath *ModeServiceRefHelmValuePath `json:"helmValuePath"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.HelmValuePath == nil {
		return fmt.Errorf("required field helmValuePath missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "engineName", "helmValuePath"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.EngineName = *all.EngineName
	if all.HelmValuePath.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.HelmValuePath = *all.HelmValuePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
