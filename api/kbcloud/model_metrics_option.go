// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MetricsOption struct {
	ReplicationLag *MetricsOptionReplicationLag `json:"replicationLag,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricsOption instantiates a new MetricsOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricsOption() *MetricsOption {
	this := MetricsOption{}
	return &this
}

// NewMetricsOptionWithDefaults instantiates a new MetricsOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricsOptionWithDefaults() *MetricsOption {
	this := MetricsOption{}
	return &this
}

// GetReplicationLag returns the ReplicationLag field value if set, zero value otherwise.
func (o *MetricsOption) GetReplicationLag() MetricsOptionReplicationLag {
	if o == nil || o.ReplicationLag == nil {
		var ret MetricsOptionReplicationLag
		return ret
	}
	return *o.ReplicationLag
}

// GetReplicationLagOk returns a tuple with the ReplicationLag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOption) GetReplicationLagOk() (*MetricsOptionReplicationLag, bool) {
	if o == nil || o.ReplicationLag == nil {
		return nil, false
	}
	return o.ReplicationLag, true
}

// HasReplicationLag returns a boolean if a field has been set.
func (o *MetricsOption) HasReplicationLag() bool {
	return o != nil && o.ReplicationLag != nil
}

// SetReplicationLag gets a reference to the given MetricsOptionReplicationLag and assigns it to the ReplicationLag field.
func (o *MetricsOption) SetReplicationLag(v MetricsOptionReplicationLag) {
	o.ReplicationLag = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricsOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ReplicationLag != nil {
		toSerialize["replicationLag"] = o.ReplicationLag
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricsOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReplicationLag *MetricsOptionReplicationLag `json:"replicationLag,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"replicationLag"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ReplicationLag != nil && all.ReplicationLag.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReplicationLag = all.ReplicationLag

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
