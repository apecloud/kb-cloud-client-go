// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchIndexStorage struct {
	Name              string               `json:"name"`
	Uuid              *string              `json:"uuid,omitempty"`
	Health            string               `json:"health"`
	Status            string               `json:"status"`
	PrimaryShards     common.NullableInt64 `json:"primaryShards"`
	ReplicaShards     common.NullableInt64 `json:"replicaShards"`
	Documents         common.NullableInt64 `json:"documents"`
	DeletedDocuments  common.NullableInt64 `json:"deletedDocuments"`
	PrimaryStoreBytes common.NullableInt64 `json:"primaryStoreBytes"`
	TotalStoreBytes   common.NullableInt64 `json:"totalStoreBytes"`
	Segments          common.NullableInt64 `json:"segments"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchIndexStorage instantiates a new ElasticsearchIndexStorage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchIndexStorage(name string, health string, status string, primaryShards common.NullableInt64, replicaShards common.NullableInt64, documents common.NullableInt64, deletedDocuments common.NullableInt64, primaryStoreBytes common.NullableInt64, totalStoreBytes common.NullableInt64, segments common.NullableInt64) *ElasticsearchIndexStorage {
	this := ElasticsearchIndexStorage{}
	this.Name = name
	this.Health = health
	this.Status = status
	this.PrimaryShards = primaryShards
	this.ReplicaShards = replicaShards
	this.Documents = documents
	this.DeletedDocuments = deletedDocuments
	this.PrimaryStoreBytes = primaryStoreBytes
	this.TotalStoreBytes = totalStoreBytes
	this.Segments = segments
	return &this
}

// NewElasticsearchIndexStorageWithDefaults instantiates a new ElasticsearchIndexStorage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchIndexStorageWithDefaults() *ElasticsearchIndexStorage {
	this := ElasticsearchIndexStorage{}
	return &this
}

// GetName returns the Name field value.
func (o *ElasticsearchIndexStorage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchIndexStorage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ElasticsearchIndexStorage) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ElasticsearchIndexStorage) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchIndexStorage) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ElasticsearchIndexStorage) HasUuid() bool {
	return o != nil && o.Uuid != nil
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ElasticsearchIndexStorage) SetUuid(v string) {
	o.Uuid = &v
}

// GetHealth returns the Health field value.
func (o *ElasticsearchIndexStorage) GetHealth() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchIndexStorage) GetHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value.
func (o *ElasticsearchIndexStorage) SetHealth(v string) {
	o.Health = v
}

// GetStatus returns the Status field value.
func (o *ElasticsearchIndexStorage) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchIndexStorage) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ElasticsearchIndexStorage) SetStatus(v string) {
	o.Status = v
}

// GetPrimaryShards returns the PrimaryShards field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchIndexStorage) GetPrimaryShards() int64 {
	if o == nil || o.PrimaryShards.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrimaryShards.Get()
}

// GetPrimaryShardsOk returns a tuple with the PrimaryShards field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchIndexStorage) GetPrimaryShardsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryShards.Get(), o.PrimaryShards.IsSet()
}

// SetPrimaryShards sets field value.
func (o *ElasticsearchIndexStorage) SetPrimaryShards(v int64) {
	o.PrimaryShards.Set(&v)
}

// GetReplicaShards returns the ReplicaShards field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchIndexStorage) GetReplicaShards() int64 {
	if o == nil || o.ReplicaShards.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ReplicaShards.Get()
}

// GetReplicaShardsOk returns a tuple with the ReplicaShards field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchIndexStorage) GetReplicaShardsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicaShards.Get(), o.ReplicaShards.IsSet()
}

// SetReplicaShards sets field value.
func (o *ElasticsearchIndexStorage) SetReplicaShards(v int64) {
	o.ReplicaShards.Set(&v)
}

// GetDocuments returns the Documents field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchIndexStorage) GetDocuments() int64 {
	if o == nil || o.Documents.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Documents.Get()
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchIndexStorage) GetDocumentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents.Get(), o.Documents.IsSet()
}

// SetDocuments sets field value.
func (o *ElasticsearchIndexStorage) SetDocuments(v int64) {
	o.Documents.Set(&v)
}

// GetDeletedDocuments returns the DeletedDocuments field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchIndexStorage) GetDeletedDocuments() int64 {
	if o == nil || o.DeletedDocuments.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DeletedDocuments.Get()
}

// GetDeletedDocumentsOk returns a tuple with the DeletedDocuments field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchIndexStorage) GetDeletedDocumentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedDocuments.Get(), o.DeletedDocuments.IsSet()
}

// SetDeletedDocuments sets field value.
func (o *ElasticsearchIndexStorage) SetDeletedDocuments(v int64) {
	o.DeletedDocuments.Set(&v)
}

// GetPrimaryStoreBytes returns the PrimaryStoreBytes field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchIndexStorage) GetPrimaryStoreBytes() int64 {
	if o == nil || o.PrimaryStoreBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrimaryStoreBytes.Get()
}

// GetPrimaryStoreBytesOk returns a tuple with the PrimaryStoreBytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchIndexStorage) GetPrimaryStoreBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryStoreBytes.Get(), o.PrimaryStoreBytes.IsSet()
}

// SetPrimaryStoreBytes sets field value.
func (o *ElasticsearchIndexStorage) SetPrimaryStoreBytes(v int64) {
	o.PrimaryStoreBytes.Set(&v)
}

// GetTotalStoreBytes returns the TotalStoreBytes field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchIndexStorage) GetTotalStoreBytes() int64 {
	if o == nil || o.TotalStoreBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalStoreBytes.Get()
}

// GetTotalStoreBytesOk returns a tuple with the TotalStoreBytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchIndexStorage) GetTotalStoreBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalStoreBytes.Get(), o.TotalStoreBytes.IsSet()
}

// SetTotalStoreBytes sets field value.
func (o *ElasticsearchIndexStorage) SetTotalStoreBytes(v int64) {
	o.TotalStoreBytes.Set(&v)
}

// GetSegments returns the Segments field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchIndexStorage) GetSegments() int64 {
	if o == nil || o.Segments.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Segments.Get()
}

// GetSegmentsOk returns a tuple with the Segments field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchIndexStorage) GetSegmentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segments.Get(), o.Segments.IsSet()
}

// SetSegments sets field value.
func (o *ElasticsearchIndexStorage) SetSegments(v int64) {
	o.Segments.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchIndexStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["health"] = o.Health
	toSerialize["status"] = o.Status
	toSerialize["primaryShards"] = o.PrimaryShards.Get()
	toSerialize["replicaShards"] = o.ReplicaShards.Get()
	toSerialize["documents"] = o.Documents.Get()
	toSerialize["deletedDocuments"] = o.DeletedDocuments.Get()
	toSerialize["primaryStoreBytes"] = o.PrimaryStoreBytes.Get()
	toSerialize["totalStoreBytes"] = o.TotalStoreBytes.Get()
	toSerialize["segments"] = o.Segments.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchIndexStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string              `json:"name"`
		Uuid              *string              `json:"uuid,omitempty"`
		Health            *string              `json:"health"`
		Status            *string              `json:"status"`
		PrimaryShards     common.NullableInt64 `json:"primaryShards"`
		ReplicaShards     common.NullableInt64 `json:"replicaShards"`
		Documents         common.NullableInt64 `json:"documents"`
		DeletedDocuments  common.NullableInt64 `json:"deletedDocuments"`
		PrimaryStoreBytes common.NullableInt64 `json:"primaryStoreBytes"`
		TotalStoreBytes   common.NullableInt64 `json:"totalStoreBytes"`
		Segments          common.NullableInt64 `json:"segments"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Health == nil {
		return fmt.Errorf("required field health missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if !all.PrimaryShards.IsSet() {
		return fmt.Errorf("required field primaryShards missing")
	}
	if !all.ReplicaShards.IsSet() {
		return fmt.Errorf("required field replicaShards missing")
	}
	if !all.Documents.IsSet() {
		return fmt.Errorf("required field documents missing")
	}
	if !all.DeletedDocuments.IsSet() {
		return fmt.Errorf("required field deletedDocuments missing")
	}
	if !all.PrimaryStoreBytes.IsSet() {
		return fmt.Errorf("required field primaryStoreBytes missing")
	}
	if !all.TotalStoreBytes.IsSet() {
		return fmt.Errorf("required field totalStoreBytes missing")
	}
	if !all.Segments.IsSet() {
		return fmt.Errorf("required field segments missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "uuid", "health", "status", "primaryShards", "replicaShards", "documents", "deletedDocuments", "primaryStoreBytes", "totalStoreBytes", "segments"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Uuid = all.Uuid
	o.Health = *all.Health
	o.Status = *all.Status
	o.PrimaryShards = all.PrimaryShards
	o.ReplicaShards = all.ReplicaShards
	o.Documents = all.Documents
	o.DeletedDocuments = all.DeletedDocuments
	o.PrimaryStoreBytes = all.PrimaryStoreBytes
	o.TotalStoreBytes = all.TotalStoreBytes
	o.Segments = all.Segments

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
