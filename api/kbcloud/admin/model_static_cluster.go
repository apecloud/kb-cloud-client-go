// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type StaticCluster struct {
	// the number of replicas
	Replicas *int32 `json:"replicas,omitempty"`
	// Storage size, the unit is Gi.
	Storage   *float64 `json:"storage,omitempty"`
	ClassCode *string  `json:"classCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStaticCluster instantiates a new StaticCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStaticCluster() *StaticCluster {
	this := StaticCluster{}
	return &this
}

// NewStaticClusterWithDefaults instantiates a new StaticCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStaticClusterWithDefaults() *StaticCluster {
	this := StaticCluster{}
	return &this
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *StaticCluster) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCluster) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *StaticCluster) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *StaticCluster) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *StaticCluster) GetStorage() float64 {
	if o == nil || o.Storage == nil {
		var ret float64
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCluster) GetStorageOk() (*float64, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *StaticCluster) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given float64 and assigns it to the Storage field.
func (o *StaticCluster) SetStorage(v float64) {
	o.Storage = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *StaticCluster) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCluster) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *StaticCluster) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *StaticCluster) SetClassCode(v string) {
	o.ClassCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StaticCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StaticCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Replicas  *int32   `json:"replicas,omitempty"`
		Storage   *float64 `json:"storage,omitempty"`
		ClassCode *string  `json:"classCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"replicas", "storage", "classCode"})
	} else {
		return err
	}
	o.Replicas = all.Replicas
	o.Storage = all.Storage
	o.ClassCode = all.ClassCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
