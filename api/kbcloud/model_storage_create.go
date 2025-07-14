// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// StorageCreate storageCreate is the schema for the storage create request
type StorageCreate struct {
	// Name of the storage
	Name *string `json:"name,omitempty"`
	// Name of the storage provider
	StorageProvider *string `json:"storageProvider,omitempty"`
	// the parameters to create the storage
	Params map[string]string `json:"params,omitempty"`
	// the id of cluster that storage used
	ClusterId *string `json:"clusterID,omitempty"`
	// the tags for the storage
	Tags    map[string]string `json:"tags,omitempty"`
	Engines []string          `json:"engines,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageCreate instantiates a new StorageCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageCreate() *StorageCreate {
	this := StorageCreate{}
	return &this
}

// NewStorageCreateWithDefaults instantiates a new StorageCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageCreateWithDefaults() *StorageCreate {
	this := StorageCreate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageCreate) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageCreate) SetName(v string) {
	o.Name = &v
}

// GetStorageProvider returns the StorageProvider field value if set, zero value otherwise.
func (o *StorageCreate) GetStorageProvider() string {
	if o == nil || o.StorageProvider == nil {
		var ret string
		return ret
	}
	return *o.StorageProvider
}

// GetStorageProviderOk returns a tuple with the StorageProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCreate) GetStorageProviderOk() (*string, bool) {
	if o == nil || o.StorageProvider == nil {
		return nil, false
	}
	return o.StorageProvider, true
}

// HasStorageProvider returns a boolean if a field has been set.
func (o *StorageCreate) HasStorageProvider() bool {
	return o != nil && o.StorageProvider != nil
}

// SetStorageProvider gets a reference to the given string and assigns it to the StorageProvider field.
func (o *StorageCreate) SetStorageProvider(v string) {
	o.StorageProvider = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *StorageCreate) GetParams() map[string]string {
	if o == nil || o.Params == nil {
		var ret map[string]string
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCreate) GetParamsOk() (*map[string]string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *StorageCreate) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *StorageCreate) SetParams(v map[string]string) {
	o.Params = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *StorageCreate) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCreate) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *StorageCreate) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *StorageCreate) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *StorageCreate) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCreate) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *StorageCreate) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *StorageCreate) SetTags(v map[string]string) {
	o.Tags = v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *StorageCreate) GetEngines() []string {
	if o == nil || o.Engines == nil {
		var ret []string
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageCreate) GetEnginesOk() (*[]string, bool) {
	if o == nil || o.Engines == nil {
		return nil, false
	}
	return &o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *StorageCreate) HasEngines() bool {
	return o != nil && o.Engines != nil
}

// SetEngines gets a reference to the given []string and assigns it to the Engines field.
func (o *StorageCreate) SetEngines(v []string) {
	o.Engines = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StorageProvider != nil {
		toSerialize["storageProvider"] = o.StorageProvider
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.ClusterId != nil {
		toSerialize["clusterID"] = o.ClusterId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Engines != nil {
		toSerialize["engines"] = o.Engines
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string           `json:"name,omitempty"`
		StorageProvider *string           `json:"storageProvider,omitempty"`
		Params          map[string]string `json:"params,omitempty"`
		ClusterId       *string           `json:"clusterID,omitempty"`
		Tags            map[string]string `json:"tags,omitempty"`
		Engines         []string          `json:"engines,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "storageProvider", "params", "clusterID", "tags", "engines"})
	} else {
		return err
	}
	o.Name = all.Name
	o.StorageProvider = all.StorageProvider
	o.Params = all.Params
	o.ClusterId = all.ClusterId
	o.Tags = all.Tags
	o.Engines = all.Engines

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
