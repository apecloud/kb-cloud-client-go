// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsExcludeConstraintExcludeItem struct {
	// The column(s) involved in the exclusion
	Columns *string `json:"columns,omitempty"`
	// The operator used in the exclusion
	Operator *string `json:"operator,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExcludeConstraintExcludeItem instantiates a new DmsExcludeConstraintExcludeItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExcludeConstraintExcludeItem() *DmsExcludeConstraintExcludeItem {
	this := DmsExcludeConstraintExcludeItem{}
	return &this
}

// NewDmsExcludeConstraintExcludeItemWithDefaults instantiates a new DmsExcludeConstraintExcludeItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExcludeConstraintExcludeItemWithDefaults() *DmsExcludeConstraintExcludeItem {
	this := DmsExcludeConstraintExcludeItem{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DmsExcludeConstraintExcludeItem) GetColumns() string {
	if o == nil || o.Columns == nil {
		var ret string
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExcludeConstraintExcludeItem) GetColumnsOk() (*string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DmsExcludeConstraintExcludeItem) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given string and assigns it to the Columns field.
func (o *DmsExcludeConstraintExcludeItem) SetColumns(v string) {
	o.Columns = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *DmsExcludeConstraintExcludeItem) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExcludeConstraintExcludeItem) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *DmsExcludeConstraintExcludeItem) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *DmsExcludeConstraintExcludeItem) SetOperator(v string) {
	o.Operator = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExcludeConstraintExcludeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExcludeConstraintExcludeItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns  *string `json:"columns,omitempty"`
		Operator *string `json:"operator,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"columns", "operator"})
	} else {
		return err
	}
	o.Columns = all.Columns
	o.Operator = all.Operator

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
