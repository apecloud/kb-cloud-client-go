// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsVolumeExpand OpsVolumeExpand is the payload to expand volume for a KubeBlocks cluster
type OpsVolumeExpand struct {
	// Requires storageExpansion in the current mode component instanceTemplate.ops. Every configured template must explicitly specify the same volume names and target capacities as component-level volumes. Names identify templates, not Pods. Targets are forwarded together in one OpsRequest without filling in, merging, or removing template targets. KubeBlocks may materialize template volumeClaimTemplates for explicitly targeted inherited volumes.
	InstanceTemplates []InstanceTemplateVolumeExpand `json:"instanceTemplates,omitempty"`
	Component         string                         `json:"component"`
	Volumes           []OpsVolumeExpandVolumesItems0 `json:"volumes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsVolumeExpand instantiates a new OpsVolumeExpand object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsVolumeExpand(component string, volumes []OpsVolumeExpandVolumesItems0) *OpsVolumeExpand {
	this := OpsVolumeExpand{}
	this.Component = component
	this.Volumes = volumes
	return &this
}

// NewOpsVolumeExpandWithDefaults instantiates a new OpsVolumeExpand object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsVolumeExpandWithDefaults() *OpsVolumeExpand {
	this := OpsVolumeExpand{}
	return &this
}

// GetInstanceTemplates returns the InstanceTemplates field value if set, zero value otherwise.
func (o *OpsVolumeExpand) GetInstanceTemplates() []InstanceTemplateVolumeExpand {
	if o == nil || o.InstanceTemplates == nil {
		var ret []InstanceTemplateVolumeExpand
		return ret
	}
	return o.InstanceTemplates
}

// GetInstanceTemplatesOk returns a tuple with the InstanceTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsVolumeExpand) GetInstanceTemplatesOk() (*[]InstanceTemplateVolumeExpand, bool) {
	if o == nil || o.InstanceTemplates == nil {
		return nil, false
	}
	return &o.InstanceTemplates, true
}

// HasInstanceTemplates returns a boolean if a field has been set.
func (o *OpsVolumeExpand) HasInstanceTemplates() bool {
	return o != nil && o.InstanceTemplates != nil
}

// SetInstanceTemplates gets a reference to the given []InstanceTemplateVolumeExpand and assigns it to the InstanceTemplates field.
func (o *OpsVolumeExpand) SetInstanceTemplates(v []InstanceTemplateVolumeExpand) {
	o.InstanceTemplates = v
}

// GetComponent returns the Component field value.
func (o *OpsVolumeExpand) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OpsVolumeExpand) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OpsVolumeExpand) SetComponent(v string) {
	o.Component = v
}

// GetVolumes returns the Volumes field value.
func (o *OpsVolumeExpand) GetVolumes() []OpsVolumeExpandVolumesItems0 {
	if o == nil {
		var ret []OpsVolumeExpandVolumesItems0
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *OpsVolumeExpand) GetVolumesOk() (*[]OpsVolumeExpandVolumesItems0, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// SetVolumes sets field value.
func (o *OpsVolumeExpand) SetVolumes(v []OpsVolumeExpandVolumesItems0) {
	o.Volumes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsVolumeExpand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.InstanceTemplates != nil {
		toSerialize["instanceTemplates"] = o.InstanceTemplates
	}
	toSerialize["component"] = o.Component
	toSerialize["volumes"] = o.Volumes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsVolumeExpand) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceTemplates []InstanceTemplateVolumeExpand  `json:"instanceTemplates,omitempty"`
		Component         *string                         `json:"component"`
		Volumes           *[]OpsVolumeExpandVolumesItems0 `json:"volumes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Volumes == nil {
		return fmt.Errorf("required field volumes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceTemplates", "component", "volumes"})
	} else {
		return err
	}
	o.InstanceTemplates = all.InstanceTemplates
	o.Component = *all.Component
	o.Volumes = *all.Volumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
