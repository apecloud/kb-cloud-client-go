// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EnvironmentDelete Environment deletion option
type EnvironmentDelete struct {
	// clean up resources in the cloud (only valid if creation is done by role ARN)
	CleanCloudResources *bool `json:"cleanCloudResources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentDelete instantiates a new EnvironmentDelete object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentDelete() *EnvironmentDelete {
	this := EnvironmentDelete{}
	return &this
}

// NewEnvironmentDeleteWithDefaults instantiates a new EnvironmentDelete object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentDeleteWithDefaults() *EnvironmentDelete {
	this := EnvironmentDelete{}
	return &this
}

// GetCleanCloudResources returns the CleanCloudResources field value if set, zero value otherwise.
func (o *EnvironmentDelete) GetCleanCloudResources() bool {
	if o == nil || o.CleanCloudResources == nil {
		var ret bool
		return ret
	}
	return *o.CleanCloudResources
}

// GetCleanCloudResourcesOk returns a tuple with the CleanCloudResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetCleanCloudResourcesOk() (*bool, bool) {
	if o == nil || o.CleanCloudResources == nil {
		return nil, false
	}
	return o.CleanCloudResources, true
}

// HasCleanCloudResources returns a boolean if a field has been set.
func (o *EnvironmentDelete) HasCleanCloudResources() bool {
	return o != nil && o.CleanCloudResources != nil
}

// SetCleanCloudResources gets a reference to the given bool and assigns it to the CleanCloudResources field.
func (o *EnvironmentDelete) SetCleanCloudResources(v bool) {
	o.CleanCloudResources = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CleanCloudResources != nil {
		toSerialize["cleanCloudResources"] = o.CleanCloudResources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentDelete) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CleanCloudResources *bool `json:"cleanCloudResources,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cleanCloudResources"})
	} else {
		return err
	}
	o.CleanCloudResources = all.CleanCloudResources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
