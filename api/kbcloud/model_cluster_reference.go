// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterReference Cluster reference information returned by cluster detail APIs.
type ClusterReference struct {
	// Name of the cluster that references this cluster.
	Cluster *string `json:"cluster,omitempty"`
	// Cluster definition of the referencing cluster.
	ClusterDefinition *string `json:"clusterDefinition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterReference instantiates a new ClusterReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterReference() *ClusterReference {
	this := ClusterReference{}
	return &this
}

// NewClusterReferenceWithDefaults instantiates a new ClusterReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterReferenceWithDefaults() *ClusterReference {
	this := ClusterReference{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ClusterReference) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterReference) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ClusterReference) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *ClusterReference) SetCluster(v string) {
	o.Cluster = &v
}

// GetClusterDefinition returns the ClusterDefinition field value if set, zero value otherwise.
func (o *ClusterReference) GetClusterDefinition() string {
	if o == nil || o.ClusterDefinition == nil {
		var ret string
		return ret
	}
	return *o.ClusterDefinition
}

// GetClusterDefinitionOk returns a tuple with the ClusterDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterReference) GetClusterDefinitionOk() (*string, bool) {
	if o == nil || o.ClusterDefinition == nil {
		return nil, false
	}
	return o.ClusterDefinition, true
}

// HasClusterDefinition returns a boolean if a field has been set.
func (o *ClusterReference) HasClusterDefinition() bool {
	return o != nil && o.ClusterDefinition != nil
}

// SetClusterDefinition gets a reference to the given string and assigns it to the ClusterDefinition field.
func (o *ClusterReference) SetClusterDefinition(v string) {
	o.ClusterDefinition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.ClusterDefinition != nil {
		toSerialize["clusterDefinition"] = o.ClusterDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster           *string `json:"cluster,omitempty"`
		ClusterDefinition *string `json:"clusterDefinition,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster", "clusterDefinition"})
	} else {
		return err
	}
	o.Cluster = all.Cluster
	o.ClusterDefinition = all.ClusterDefinition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
