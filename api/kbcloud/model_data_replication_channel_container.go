// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DataReplicationChannelContainer struct {
	Name            *string `json:"name,omitempty"`
	IsMainContainer *bool   `json:"isMainContainer,omitempty"`
	IsInitContainer *bool   `json:"isInitContainer,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataReplicationChannelContainer instantiates a new DataReplicationChannelContainer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataReplicationChannelContainer() *DataReplicationChannelContainer {
	this := DataReplicationChannelContainer{}
	var isMainContainer bool = false
	this.IsMainContainer = &isMainContainer
	var isInitContainer bool = false
	this.IsInitContainer = &isInitContainer
	return &this
}

// NewDataReplicationChannelContainerWithDefaults instantiates a new DataReplicationChannelContainer object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataReplicationChannelContainerWithDefaults() *DataReplicationChannelContainer {
	this := DataReplicationChannelContainer{}
	var isMainContainer bool = false
	this.IsMainContainer = &isMainContainer
	var isInitContainer bool = false
	this.IsInitContainer = &isInitContainer
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataReplicationChannelContainer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationChannelContainer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataReplicationChannelContainer) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataReplicationChannelContainer) SetName(v string) {
	o.Name = &v
}

// GetIsMainContainer returns the IsMainContainer field value if set, zero value otherwise.
func (o *DataReplicationChannelContainer) GetIsMainContainer() bool {
	if o == nil || o.IsMainContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsMainContainer
}

// GetIsMainContainerOk returns a tuple with the IsMainContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationChannelContainer) GetIsMainContainerOk() (*bool, bool) {
	if o == nil || o.IsMainContainer == nil {
		return nil, false
	}
	return o.IsMainContainer, true
}

// HasIsMainContainer returns a boolean if a field has been set.
func (o *DataReplicationChannelContainer) HasIsMainContainer() bool {
	return o != nil && o.IsMainContainer != nil
}

// SetIsMainContainer gets a reference to the given bool and assigns it to the IsMainContainer field.
func (o *DataReplicationChannelContainer) SetIsMainContainer(v bool) {
	o.IsMainContainer = &v
}

// GetIsInitContainer returns the IsInitContainer field value if set, zero value otherwise.
func (o *DataReplicationChannelContainer) GetIsInitContainer() bool {
	if o == nil || o.IsInitContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsInitContainer
}

// GetIsInitContainerOk returns a tuple with the IsInitContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReplicationChannelContainer) GetIsInitContainerOk() (*bool, bool) {
	if o == nil || o.IsInitContainer == nil {
		return nil, false
	}
	return o.IsInitContainer, true
}

// HasIsInitContainer returns a boolean if a field has been set.
func (o *DataReplicationChannelContainer) HasIsInitContainer() bool {
	return o != nil && o.IsInitContainer != nil
}

// SetIsInitContainer gets a reference to the given bool and assigns it to the IsInitContainer field.
func (o *DataReplicationChannelContainer) SetIsInitContainer(v bool) {
	o.IsInitContainer = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataReplicationChannelContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IsMainContainer != nil {
		toSerialize["isMainContainer"] = o.IsMainContainer
	}
	if o.IsInitContainer != nil {
		toSerialize["isInitContainer"] = o.IsInitContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataReplicationChannelContainer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string `json:"name,omitempty"`
		IsMainContainer *bool   `json:"isMainContainer,omitempty"`
		IsInitContainer *bool   `json:"isInitContainer,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "isMainContainer", "isInitContainer"})
	} else {
		return err
	}
	o.Name = all.Name
	o.IsMainContainer = all.IsMainContainer
	o.IsInitContainer = all.IsInitContainer

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
