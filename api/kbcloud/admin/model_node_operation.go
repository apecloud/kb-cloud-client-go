// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type NodeOperation struct {
	// Node name (e.g. 'MyName',  or 'my.name',  or '123-abc')
	Name *string `json:"name,omitempty"`
	// operation for node, either `add` or `del`
	Op *NodeOperationType `json:"op,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeOperation instantiates a new NodeOperation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeOperation() *NodeOperation {
	this := NodeOperation{}
	return &this
}

// NewNodeOperationWithDefaults instantiates a new NodeOperation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeOperationWithDefaults() *NodeOperation {
	this := NodeOperation{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NodeOperation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeOperation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NodeOperation) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NodeOperation) SetName(v string) {
	o.Name = &v
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *NodeOperation) GetOp() NodeOperationType {
	if o == nil || o.Op == nil {
		var ret NodeOperationType
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeOperation) GetOpOk() (*NodeOperationType, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *NodeOperation) HasOp() bool {
	return o != nil && o.Op != nil
}

// SetOp gets a reference to the given NodeOperationType and assigns it to the Op field.
func (o *NodeOperation) SetOp(v NodeOperationType) {
	o.Op = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeOperation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string            `json:"name,omitempty"`
		Op   *NodeOperationType `json:"op,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "op"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Op != nil && !all.Op.IsValid() {
		hasInvalidField = true
	} else {
		o.Op = all.Op
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
