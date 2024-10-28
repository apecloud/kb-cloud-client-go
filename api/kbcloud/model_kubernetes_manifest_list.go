// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// KubernetesManifestList KubernetesManifestList is the list of Kubernetes objects in the list

type KubernetesManifestList struct {
	// Items is the list of Kubernetes objects in the list
	KubernetesManifests []map[string]interface{} `json:"kubernetesManifests,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKubernetesManifestList instantiates a new KubernetesManifestList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKubernetesManifestList() *KubernetesManifestList {
	this := KubernetesManifestList{}
	return &this
}

// NewKubernetesManifestListWithDefaults instantiates a new KubernetesManifestList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKubernetesManifestListWithDefaults() *KubernetesManifestList {
	this := KubernetesManifestList{}
	return &this
}

// GetKubernetesManifests returns the KubernetesManifests field value if set, zero value otherwise.
func (o *KubernetesManifestList) GetKubernetesManifests() []map[string]interface{} {
	if o == nil || o.KubernetesManifests == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.KubernetesManifests
}

// GetKubernetesManifestsOk returns a tuple with the KubernetesManifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesManifestList) GetKubernetesManifestsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.KubernetesManifests == nil {
		return nil, false
	}
	return &o.KubernetesManifests, true
}

// HasKubernetesManifests returns a boolean if a field has been set.
func (o *KubernetesManifestList) HasKubernetesManifests() bool {
	return o != nil && o.KubernetesManifests != nil
}

// SetKubernetesManifests gets a reference to the given []map[string]interface{} and assigns it to the KubernetesManifests field.
func (o *KubernetesManifestList) SetKubernetesManifests(v []map[string]interface{}) {
	o.KubernetesManifests = v
}

// MarshalJSON serializes the struct using spec logic.
func (o KubernetesManifestList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.KubernetesManifests != nil {
		toSerialize["kubernetesManifests"] = o.KubernetesManifests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KubernetesManifestList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KubernetesManifests []map[string]interface{} `json:"kubernetesManifests,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kubernetesManifests"})
	} else {
		return err
	}
	o.KubernetesManifests = all.KubernetesManifests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
