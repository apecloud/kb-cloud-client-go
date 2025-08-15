// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Storage Storage is a specification that provides guidance accessing remote storage.
type Storage struct {
	// Unique identifier for the storage
	Id *string `json:"id,omitempty"`
	// Name of the storage
	Name *string `json:"name,omitempty"`
	// Name of the storage provider
	StorageProvider *string `json:"storageProvider,omitempty"`
	// the parameters for the storage
	Params map[string]string `json:"params,omitempty"`
	// the environment name that storage belongs to
	EnvName *string `json:"envName,omitempty"`
	// User who created the storage
	CreatedBy *time.Time `json:"createdBy,omitempty"`
	// the id of cluster that storage used
	ClusterId common.NullableString `json:"clusterID,omitempty"`
	// the address of cluster
	ClusterMetadataAddress common.NullableString `json:"clusterMetadataAddress,omitempty"`
	// User who updated the storage
	UpdatedBy *time.Time `json:"updatedBy,omitempty"`
	// the tags for the storage
	Tags    map[string]string `json:"tags,omitempty"`
	Engines []string          `json:"engines,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorage instantiates a new Storage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorage() *Storage {
	this := Storage{}
	return &this
}

// NewStorageWithDefaults instantiates a new Storage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageWithDefaults() *Storage {
	this := Storage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Storage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Storage) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Storage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Storage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Storage) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Storage) SetName(v string) {
	o.Name = &v
}

// GetStorageProvider returns the StorageProvider field value if set, zero value otherwise.
func (o *Storage) GetStorageProvider() string {
	if o == nil || o.StorageProvider == nil {
		var ret string
		return ret
	}
	return *o.StorageProvider
}

// GetStorageProviderOk returns a tuple with the StorageProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetStorageProviderOk() (*string, bool) {
	if o == nil || o.StorageProvider == nil {
		return nil, false
	}
	return o.StorageProvider, true
}

// HasStorageProvider returns a boolean if a field has been set.
func (o *Storage) HasStorageProvider() bool {
	return o != nil && o.StorageProvider != nil
}

// SetStorageProvider gets a reference to the given string and assigns it to the StorageProvider field.
func (o *Storage) SetStorageProvider(v string) {
	o.StorageProvider = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *Storage) GetParams() map[string]string {
	if o == nil || o.Params == nil {
		var ret map[string]string
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetParamsOk() (*map[string]string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *Storage) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *Storage) SetParams(v map[string]string) {
	o.Params = v
}

// GetEnvName returns the EnvName field value if set, zero value otherwise.
func (o *Storage) GetEnvName() string {
	if o == nil || o.EnvName == nil {
		var ret string
		return ret
	}
	return *o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetEnvNameOk() (*string, bool) {
	if o == nil || o.EnvName == nil {
		return nil, false
	}
	return o.EnvName, true
}

// HasEnvName returns a boolean if a field has been set.
func (o *Storage) HasEnvName() bool {
	return o != nil && o.EnvName != nil
}

// SetEnvName gets a reference to the given string and assigns it to the EnvName field.
func (o *Storage) SetEnvName(v string) {
	o.EnvName = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Storage) GetCreatedBy() time.Time {
	if o == nil || o.CreatedBy == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetCreatedByOk() (*time.Time, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Storage) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given time.Time and assigns it to the CreatedBy field.
func (o *Storage) SetCreatedBy(v time.Time) {
	o.CreatedBy = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Storage) GetClusterId() string {
	if o == nil || o.ClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Storage) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *Storage) HasClusterId() bool {
	return o != nil && o.ClusterId.IsSet()
}

// SetClusterId gets a reference to the given common.NullableString and assigns it to the ClusterId field.
func (o *Storage) SetClusterId(v string) {
	o.ClusterId.Set(&v)
}

// SetClusterIdNil sets the value for ClusterId to be an explicit nil.
func (o *Storage) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil.
func (o *Storage) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetClusterMetadataAddress returns the ClusterMetadataAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Storage) GetClusterMetadataAddress() string {
	if o == nil || o.ClusterMetadataAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterMetadataAddress.Get()
}

// GetClusterMetadataAddressOk returns a tuple with the ClusterMetadataAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Storage) GetClusterMetadataAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterMetadataAddress.Get(), o.ClusterMetadataAddress.IsSet()
}

// HasClusterMetadataAddress returns a boolean if a field has been set.
func (o *Storage) HasClusterMetadataAddress() bool {
	return o != nil && o.ClusterMetadataAddress.IsSet()
}

// SetClusterMetadataAddress gets a reference to the given common.NullableString and assigns it to the ClusterMetadataAddress field.
func (o *Storage) SetClusterMetadataAddress(v string) {
	o.ClusterMetadataAddress.Set(&v)
}

// SetClusterMetadataAddressNil sets the value for ClusterMetadataAddress to be an explicit nil.
func (o *Storage) SetClusterMetadataAddressNil() {
	o.ClusterMetadataAddress.Set(nil)
}

// UnsetClusterMetadataAddress ensures that no value is present for ClusterMetadataAddress, not even an explicit nil.
func (o *Storage) UnsetClusterMetadataAddress() {
	o.ClusterMetadataAddress.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Storage) GetUpdatedBy() time.Time {
	if o == nil || o.UpdatedBy == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetUpdatedByOk() (*time.Time, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Storage) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given time.Time and assigns it to the UpdatedBy field.
func (o *Storage) SetUpdatedBy(v time.Time) {
	o.UpdatedBy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Storage) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Storage) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Storage) SetTags(v map[string]string) {
	o.Tags = v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *Storage) GetEngines() []string {
	if o == nil || o.Engines == nil {
		var ret []string
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetEnginesOk() (*[]string, bool) {
	if o == nil || o.Engines == nil {
		return nil, false
	}
	return &o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *Storage) HasEngines() bool {
	return o != nil && o.Engines != nil
}

// SetEngines gets a reference to the given []string and assigns it to the Engines field.
func (o *Storage) SetEngines(v []string) {
	o.Engines = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Storage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.EnvName != nil {
		toSerialize["envName"] = o.EnvName
	}
	if o.CreatedBy != nil {
		if o.CreatedBy.Nanosecond() == 0 {
			toSerialize["createdBy"] = o.CreatedBy.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdBy"] = o.CreatedBy.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ClusterId.IsSet() {
		toSerialize["clusterID"] = o.ClusterId.Get()
	}
	if o.ClusterMetadataAddress.IsSet() {
		toSerialize["clusterMetadataAddress"] = o.ClusterMetadataAddress.Get()
	}
	if o.UpdatedBy != nil {
		if o.UpdatedBy.Nanosecond() == 0 {
			toSerialize["updatedBy"] = o.UpdatedBy.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedBy"] = o.UpdatedBy.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *Storage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                     *string               `json:"id,omitempty"`
		Name                   *string               `json:"name,omitempty"`
		StorageProvider        *string               `json:"storageProvider,omitempty"`
		Params                 map[string]string     `json:"params,omitempty"`
		EnvName                *string               `json:"envName,omitempty"`
		CreatedBy              *time.Time            `json:"createdBy,omitempty"`
		ClusterId              common.NullableString `json:"clusterID,omitempty"`
		ClusterMetadataAddress common.NullableString `json:"clusterMetadataAddress,omitempty"`
		UpdatedBy              *time.Time            `json:"updatedBy,omitempty"`
		Tags                   map[string]string     `json:"tags,omitempty"`
		Engines                []string              `json:"engines,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "storageProvider", "params", "envName", "createdBy", "clusterID", "clusterMetadataAddress", "updatedBy", "tags", "engines"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.StorageProvider = all.StorageProvider
	o.Params = all.Params
	o.EnvName = all.EnvName
	o.CreatedBy = all.CreatedBy
	o.ClusterId = all.ClusterId
	o.ClusterMetadataAddress = all.ClusterMetadataAddress
	o.UpdatedBy = all.UpdatedBy
	o.Tags = all.Tags
	o.Engines = all.Engines

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
