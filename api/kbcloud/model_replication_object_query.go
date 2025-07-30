// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ReplicationObjectQuery struct {
	Endpoint  *DataChannelEndpointCreate `json:"endpoint,omitempty"`
	NodeType  common.NullableString      `json:"nodeType,omitempty"`
	NodeValue common.NullableString      `json:"nodeValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplicationObjectQuery instantiates a new ReplicationObjectQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplicationObjectQuery() *ReplicationObjectQuery {
	this := ReplicationObjectQuery{}
	return &this
}

// NewReplicationObjectQueryWithDefaults instantiates a new ReplicationObjectQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplicationObjectQueryWithDefaults() *ReplicationObjectQuery {
	this := ReplicationObjectQuery{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ReplicationObjectQuery) GetEndpoint() DataChannelEndpointCreate {
	if o == nil || o.Endpoint == nil {
		var ret DataChannelEndpointCreate
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationObjectQuery) GetEndpointOk() (*DataChannelEndpointCreate, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ReplicationObjectQuery) HasEndpoint() bool {
	return o != nil && o.Endpoint != nil
}

// SetEndpoint gets a reference to the given DataChannelEndpointCreate and assigns it to the Endpoint field.
func (o *ReplicationObjectQuery) SetEndpoint(v DataChannelEndpointCreate) {
	o.Endpoint = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationObjectQuery) GetNodeType() string {
	if o == nil || o.NodeType.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeType.Get()
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplicationObjectQuery) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeType.Get(), o.NodeType.IsSet()
}

// HasNodeType returns a boolean if a field has been set.
func (o *ReplicationObjectQuery) HasNodeType() bool {
	return o != nil && o.NodeType.IsSet()
}

// SetNodeType gets a reference to the given common.NullableString and assigns it to the NodeType field.
func (o *ReplicationObjectQuery) SetNodeType(v string) {
	o.NodeType.Set(&v)
}

// SetNodeTypeNil sets the value for NodeType to be an explicit nil.
func (o *ReplicationObjectQuery) SetNodeTypeNil() {
	o.NodeType.Set(nil)
}

// UnsetNodeType ensures that no value is present for NodeType, not even an explicit nil.
func (o *ReplicationObjectQuery) UnsetNodeType() {
	o.NodeType.Unset()
}

// GetNodeValue returns the NodeValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationObjectQuery) GetNodeValue() string {
	if o == nil || o.NodeValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeValue.Get()
}

// GetNodeValueOk returns a tuple with the NodeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReplicationObjectQuery) GetNodeValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeValue.Get(), o.NodeValue.IsSet()
}

// HasNodeValue returns a boolean if a field has been set.
func (o *ReplicationObjectQuery) HasNodeValue() bool {
	return o != nil && o.NodeValue.IsSet()
}

// SetNodeValue gets a reference to the given common.NullableString and assigns it to the NodeValue field.
func (o *ReplicationObjectQuery) SetNodeValue(v string) {
	o.NodeValue.Set(&v)
}

// SetNodeValueNil sets the value for NodeValue to be an explicit nil.
func (o *ReplicationObjectQuery) SetNodeValueNil() {
	o.NodeValue.Set(nil)
}

// UnsetNodeValue ensures that no value is present for NodeValue, not even an explicit nil.
func (o *ReplicationObjectQuery) UnsetNodeValue() {
	o.NodeValue.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplicationObjectQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.NodeType.IsSet() {
		toSerialize["nodeType"] = o.NodeType.Get()
	}
	if o.NodeValue.IsSet() {
		toSerialize["nodeValue"] = o.NodeValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplicationObjectQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Endpoint  *DataChannelEndpointCreate `json:"endpoint,omitempty"`
		NodeType  common.NullableString      `json:"nodeType,omitempty"`
		NodeValue common.NullableString      `json:"nodeValue,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"endpoint", "nodeType", "nodeValue"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Endpoint != nil && all.Endpoint.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Endpoint = all.Endpoint
	o.NodeType = all.NodeType
	o.NodeValue = all.NodeValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
