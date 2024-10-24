// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NamespaceInfo Namespace info
// NODESCRIPTION NamespaceInfo
//
// Deprecated: This model is deprecated.
type NamespaceInfo struct {
	// Namespace
	Namespace string `json:"namespace"`
	// Clusters in the namespace
	Clusters int32 `json:"clusters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNamespaceInfo instantiates a new NamespaceInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNamespaceInfo(namespace string, clusters int32) *NamespaceInfo {
	this := NamespaceInfo{}
	this.Namespace = namespace
	this.Clusters = clusters
	return &this
}

// NewNamespaceInfoWithDefaults instantiates a new NamespaceInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNamespaceInfoWithDefaults() *NamespaceInfo {
	this := NamespaceInfo{}
	return &this
}

// GetNamespace returns the Namespace field value.
func (o *NamespaceInfo) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *NamespaceInfo) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *NamespaceInfo) SetNamespace(v string) {
	o.Namespace = v
}

// GetClusters returns the Clusters field value.
func (o *NamespaceInfo) GetClusters() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *NamespaceInfo) GetClustersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// SetClusters sets field value.
func (o *NamespaceInfo) SetClusters(v int32) {
	o.Clusters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NamespaceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["namespace"] = o.Namespace
	toSerialize["clusters"] = o.Clusters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NamespaceInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Namespace *string `json:"namespace"`
		Clusters  *int32  `json:"clusters"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.Clusters == nil {
		return fmt.Errorf("required field clusters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"namespace", "clusters"})
	} else {
		return err
	}
	o.Namespace = *all.Namespace
	o.Clusters = *all.Clusters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
