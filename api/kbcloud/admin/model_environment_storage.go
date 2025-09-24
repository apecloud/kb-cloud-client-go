// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentStorage Storage config
type EnvironmentStorage struct {
	// the storage name
	Name *string `json:"name,omitempty"`
	// the storage type
	Type *EnvironmentStorageType `json:"type,omitempty"`
	// the existed cluster name for creating storage
	ReusedClusterName *string `json:"reusedClusterName,omitempty"`
	// the existed cluster namespace for creating storage
	ReusedClusterNamespace *string `json:"reusedClusterNamespace,omitempty"`
	// storageCreate is the schema for the storage create request
	Config  *StorageCreate `json:"config,omitempty"`
	Cluster *StaticCluster `json:"cluster,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStorage instantiates a new EnvironmentStorage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStorage() *EnvironmentStorage {
	this := EnvironmentStorage{}
	return &this
}

// NewEnvironmentStorageWithDefaults instantiates a new EnvironmentStorage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStorageWithDefaults() *EnvironmentStorage {
	this := EnvironmentStorage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentStorage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStorage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentStorage) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentStorage) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentStorage) GetType() EnvironmentStorageType {
	if o == nil || o.Type == nil {
		var ret EnvironmentStorageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStorage) GetTypeOk() (*EnvironmentStorageType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentStorage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given EnvironmentStorageType and assigns it to the Type field.
func (o *EnvironmentStorage) SetType(v EnvironmentStorageType) {
	o.Type = &v
}

// GetReusedClusterName returns the ReusedClusterName field value if set, zero value otherwise.
func (o *EnvironmentStorage) GetReusedClusterName() string {
	if o == nil || o.ReusedClusterName == nil {
		var ret string
		return ret
	}
	return *o.ReusedClusterName
}

// GetReusedClusterNameOk returns a tuple with the ReusedClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStorage) GetReusedClusterNameOk() (*string, bool) {
	if o == nil || o.ReusedClusterName == nil {
		return nil, false
	}
	return o.ReusedClusterName, true
}

// HasReusedClusterName returns a boolean if a field has been set.
func (o *EnvironmentStorage) HasReusedClusterName() bool {
	return o != nil && o.ReusedClusterName != nil
}

// SetReusedClusterName gets a reference to the given string and assigns it to the ReusedClusterName field.
func (o *EnvironmentStorage) SetReusedClusterName(v string) {
	o.ReusedClusterName = &v
}

// GetReusedClusterNamespace returns the ReusedClusterNamespace field value if set, zero value otherwise.
func (o *EnvironmentStorage) GetReusedClusterNamespace() string {
	if o == nil || o.ReusedClusterNamespace == nil {
		var ret string
		return ret
	}
	return *o.ReusedClusterNamespace
}

// GetReusedClusterNamespaceOk returns a tuple with the ReusedClusterNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStorage) GetReusedClusterNamespaceOk() (*string, bool) {
	if o == nil || o.ReusedClusterNamespace == nil {
		return nil, false
	}
	return o.ReusedClusterNamespace, true
}

// HasReusedClusterNamespace returns a boolean if a field has been set.
func (o *EnvironmentStorage) HasReusedClusterNamespace() bool {
	return o != nil && o.ReusedClusterNamespace != nil
}

// SetReusedClusterNamespace gets a reference to the given string and assigns it to the ReusedClusterNamespace field.
func (o *EnvironmentStorage) SetReusedClusterNamespace(v string) {
	o.ReusedClusterNamespace = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *EnvironmentStorage) GetConfig() StorageCreate {
	if o == nil || o.Config == nil {
		var ret StorageCreate
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStorage) GetConfigOk() (*StorageCreate, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *EnvironmentStorage) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given StorageCreate and assigns it to the Config field.
func (o *EnvironmentStorage) SetConfig(v StorageCreate) {
	o.Config = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *EnvironmentStorage) GetCluster() StaticCluster {
	if o == nil || o.Cluster == nil {
		var ret StaticCluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStorage) GetClusterOk() (*StaticCluster, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *EnvironmentStorage) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given StaticCluster and assigns it to the Cluster field.
func (o *EnvironmentStorage) SetCluster(v StaticCluster) {
	o.Cluster = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ReusedClusterName != nil {
		toSerialize["reusedClusterName"] = o.ReusedClusterName
	}
	if o.ReusedClusterNamespace != nil {
		toSerialize["reusedClusterNamespace"] = o.ReusedClusterNamespace
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                   *string                 `json:"name,omitempty"`
		Type                   *EnvironmentStorageType `json:"type,omitempty"`
		ReusedClusterName      *string                 `json:"reusedClusterName,omitempty"`
		ReusedClusterNamespace *string                 `json:"reusedClusterNamespace,omitempty"`
		Config                 *StorageCreate          `json:"config,omitempty"`
		Cluster                *StaticCluster          `json:"cluster,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "reusedClusterName", "reusedClusterNamespace", "config", "cluster"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.ReusedClusterName = all.ReusedClusterName
	o.ReusedClusterNamespace = all.ReusedClusterNamespace
	if all.Config != nil && all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = all.Config
	if all.Cluster != nil && all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = all.Cluster

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
