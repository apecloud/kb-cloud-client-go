// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION ViewInvolvedObjectsSelectorMatchExpressionsItem
type ViewInvolvedObjectsSelectorMatchExpressionsItem struct {
	// key is the label key that the selector applies to.
	Key *string `json:"key,omitempty"`
	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator *string `json:"operator,omitempty"`
	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.
	Values []string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewViewInvolvedObjectsSelectorMatchExpressionsItem instantiates a new ViewInvolvedObjectsSelectorMatchExpressionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewViewInvolvedObjectsSelectorMatchExpressionsItem() *ViewInvolvedObjectsSelectorMatchExpressionsItem {
	this := ViewInvolvedObjectsSelectorMatchExpressionsItem{}
	return &this
}

// NewViewInvolvedObjectsSelectorMatchExpressionsItemWithDefaults instantiates a new ViewInvolvedObjectsSelectorMatchExpressionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewViewInvolvedObjectsSelectorMatchExpressionsItemWithDefaults() *ViewInvolvedObjectsSelectorMatchExpressionsItem {
	this := ViewInvolvedObjectsSelectorMatchExpressionsItem{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) SetKey(v string) {
	o.Key = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) SetOperator(v string) {
	o.Operator = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ViewInvolvedObjectsSelectorMatchExpressionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
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
func (o *ViewInvolvedObjectsSelectorMatchExpressionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key      *string  `json:"key,omitempty"`
		Operator *string  `json:"operator,omitempty"`
		Values   []string `json:"values,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "operator", "values"})
	} else {
		return err
	}
	o.Key = all.Key
	o.Operator = all.Operator
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
