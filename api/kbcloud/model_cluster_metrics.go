// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterMetrics Cluster metrics

type ClusterMetrics struct {
	// NODESCRIPTION Value
	Value []interface{} `json:"value,omitempty"`
	// NODESCRIPTION Values
	Values [][]interface{} `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterMetrics instantiates a new ClusterMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterMetrics() *ClusterMetrics {
	this := ClusterMetrics{}
	return &this
}

// NewClusterMetricsWithDefaults instantiates a new ClusterMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterMetricsWithDefaults() *ClusterMetrics {
	this := ClusterMetrics{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterMetrics) GetValue() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterMetrics) GetValueOk() (*[]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ClusterMetrics) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given []interface{} and assigns it to the Value field.
func (o *ClusterMetrics) SetValue(v []interface{}) {
	o.Value = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ClusterMetrics) GetValues() [][]interface{} {
	if o == nil || o.Values == nil {
		var ret [][]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMetrics) GetValuesOk() (*[][]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ClusterMetrics) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given [][]interface{} and assigns it to the Values field.
func (o *ClusterMetrics) SetValues(v [][]interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value  []interface{}   `json:"value,omitempty"`
		Values [][]interface{} `json:"values,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"value", "values"})
	} else {
		return err
	}
	o.Value = all.Value
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
