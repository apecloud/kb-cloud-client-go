// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// OpsPromote OpsPromote is the payload to promote a KubeBlocks cluster
// NODESCRIPTION OpsPromote
//
// Deprecated: This model is deprecated.
type OpsPromote struct {
	// NODESCRIPTION ComponentName
	ComponentName *string `json:"componentName,omitempty"`
	// NODESCRIPTION InstanceName
	InstanceName *string `json:"instanceName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsPromote instantiates a new OpsPromote object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsPromote() *OpsPromote {
	this := OpsPromote{}
	return &this
}

// NewOpsPromoteWithDefaults instantiates a new OpsPromote object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsPromoteWithDefaults() *OpsPromote {
	this := OpsPromote{}
	return &this
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *OpsPromote) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsPromote) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *OpsPromote) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *OpsPromote) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *OpsPromote) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsPromote) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *OpsPromote) HasInstanceName() bool {
	return o != nil && o.InstanceName != nil
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *OpsPromote) SetInstanceName(v string) {
	o.InstanceName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsPromote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsPromote) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentName *string `json:"componentName,omitempty"`
		InstanceName  *string `json:"instanceName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentName", "instanceName"})
	} else {
		return err
	}
	o.ComponentName = all.ComponentName
	o.InstanceName = all.InstanceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
