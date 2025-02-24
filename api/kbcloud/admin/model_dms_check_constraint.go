// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsCheckConstraint struct {
	// The name of the check constraint
	Name *string `json:"name,omitempty"`
	// The expression for the check constraint
	Expression *string `json:"expression,omitempty"`
	// Whether the check constraint is enforced
	Enforced *bool `json:"enforced,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsCheckConstraint instantiates a new DmsCheckConstraint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsCheckConstraint() *DmsCheckConstraint {
	this := DmsCheckConstraint{}
	return &this
}

// NewDmsCheckConstraintWithDefaults instantiates a new DmsCheckConstraint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsCheckConstraintWithDefaults() *DmsCheckConstraint {
	this := DmsCheckConstraint{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DmsCheckConstraint) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsCheckConstraint) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DmsCheckConstraint) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DmsCheckConstraint) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *DmsCheckConstraint) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsCheckConstraint) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *DmsCheckConstraint) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *DmsCheckConstraint) SetExpression(v string) {
	o.Expression = &v
}

// GetEnforced returns the Enforced field value if set, zero value otherwise.
func (o *DmsCheckConstraint) GetEnforced() bool {
	if o == nil || o.Enforced == nil {
		var ret bool
		return ret
	}
	return *o.Enforced
}

// GetEnforcedOk returns a tuple with the Enforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsCheckConstraint) GetEnforcedOk() (*bool, bool) {
	if o == nil || o.Enforced == nil {
		return nil, false
	}
	return o.Enforced, true
}

// HasEnforced returns a boolean if a field has been set.
func (o *DmsCheckConstraint) HasEnforced() bool {
	return o != nil && o.Enforced != nil
}

// SetEnforced gets a reference to the given bool and assigns it to the Enforced field.
func (o *DmsCheckConstraint) SetEnforced(v bool) {
	o.Enforced = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsCheckConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Enforced != nil {
		toSerialize["enforced"] = o.Enforced
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsCheckConstraint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string `json:"name,omitempty"`
		Expression *string `json:"expression,omitempty"`
		Enforced   *bool   `json:"enforced,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "expression", "enforced"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Expression = all.Expression
	o.Enforced = all.Enforced

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
