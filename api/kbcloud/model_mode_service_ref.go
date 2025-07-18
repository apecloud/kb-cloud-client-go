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
	// The name will be referenced in clusterCreate request. The name should also be defined in `.components`
	// so that frontend can use it to get proper localized title.
	//
	Name string `json:"name"`
	// The engine to be used in serviceRef. This field is used to filter clusters.
	EngineName string `json:"engineName"`
	// The path to be used in values. Separated with commas. ClusterCreate API will use these path to override values in the cluster chart.
	HelmValuePath ModeServiceRefHelmValuePath `json:"helmValuePath"`
	// ServiceSelectors will map cluster's mode to a serviceSelector. The serviceSelector
	// will be used to provide the corresponding helm values.
	// If no serviceSelector is matched, the corresponding helm value will not be set.
	//
	ServiceSelectors []ServiceSelector `json:"serviceSelectors,omitempty"`
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

// GetServiceSelectors returns the ServiceSelectors field value if set, zero value otherwise.
func (o *ModeServiceRef) GetServiceSelectors() []ServiceSelector {
	if o == nil || o.ServiceSelectors == nil {
		var ret []ServiceSelector
		return ret
	}
	return o.ServiceSelectors
}

// GetServiceSelectorsOk returns a tuple with the ServiceSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetServiceSelectorsOk() (*[]ServiceSelector, bool) {
	if o == nil || o.ServiceSelectors == nil {
		return nil, false
	}
	return &o.ServiceSelectors, true
}

// HasServiceSelectors returns a boolean if a field has been set.
func (o *ModeServiceRef) HasServiceSelectors() bool {
	return o != nil && o.ServiceSelectors != nil
}

// SetServiceSelectors gets a reference to the given []ServiceSelector and assigns it to the ServiceSelectors field.
func (o *ModeServiceRef) SetServiceSelectors(v []ServiceSelector) {
	o.ServiceSelectors = v
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
	if o.ServiceSelectors != nil {
		toSerialize["serviceSelectors"] = o.ServiceSelectors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeServiceRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string                      `json:"name"`
		EngineName       *string                      `json:"engineName"`
		HelmValuePath    *ModeServiceRefHelmValuePath `json:"helmValuePath"`
		ServiceSelectors []ServiceSelector            `json:"serviceSelectors,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "engineName", "helmValuePath", "serviceSelectors"})
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
	o.ServiceSelectors = all.ServiceSelectors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
