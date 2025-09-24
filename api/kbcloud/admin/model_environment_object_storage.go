// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentObjectStorage struct {
	Clusters []RawCluster `json:"clusters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentObjectStorage instantiates a new EnvironmentObjectStorage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentObjectStorage() *EnvironmentObjectStorage {
	this := EnvironmentObjectStorage{}
	return &this
}

// NewEnvironmentObjectStorageWithDefaults instantiates a new EnvironmentObjectStorage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentObjectStorageWithDefaults() *EnvironmentObjectStorage {
	this := EnvironmentObjectStorage{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *EnvironmentObjectStorage) GetClusters() []RawCluster {
	if o == nil || o.Clusters == nil {
		var ret []RawCluster
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentObjectStorage) GetClustersOk() (*[]RawCluster, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *EnvironmentObjectStorage) HasClusters() bool {
	return o != nil && o.Clusters != nil
}

// SetClusters gets a reference to the given []RawCluster and assigns it to the Clusters field.
func (o *EnvironmentObjectStorage) SetClusters(v []RawCluster) {
	o.Clusters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentObjectStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentObjectStorage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Clusters []RawCluster `json:"clusters,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusters"})
	} else {
		return err
	}
	o.Clusters = all.Clusters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
