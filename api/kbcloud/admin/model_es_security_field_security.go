// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ESSecurityFieldSecurity struct {
	Grant  []string `json:"grant,omitempty"`
	Except []string `json:"except,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityFieldSecurity instantiates a new ESSecurityFieldSecurity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityFieldSecurity() *ESSecurityFieldSecurity {
	this := ESSecurityFieldSecurity{}
	return &this
}

// NewESSecurityFieldSecurityWithDefaults instantiates a new ESSecurityFieldSecurity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityFieldSecurityWithDefaults() *ESSecurityFieldSecurity {
	this := ESSecurityFieldSecurity{}
	return &this
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *ESSecurityFieldSecurity) GetGrant() []string {
	if o == nil || o.Grant == nil {
		var ret []string
		return ret
	}
	return o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityFieldSecurity) GetGrantOk() (*[]string, bool) {
	if o == nil || o.Grant == nil {
		return nil, false
	}
	return &o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *ESSecurityFieldSecurity) HasGrant() bool {
	return o != nil && o.Grant != nil
}

// SetGrant gets a reference to the given []string and assigns it to the Grant field.
func (o *ESSecurityFieldSecurity) SetGrant(v []string) {
	o.Grant = v
}

// GetExcept returns the Except field value if set, zero value otherwise.
func (o *ESSecurityFieldSecurity) GetExcept() []string {
	if o == nil || o.Except == nil {
		var ret []string
		return ret
	}
	return o.Except
}

// GetExceptOk returns a tuple with the Except field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityFieldSecurity) GetExceptOk() (*[]string, bool) {
	if o == nil || o.Except == nil {
		return nil, false
	}
	return &o.Except, true
}

// HasExcept returns a boolean if a field has been set.
func (o *ESSecurityFieldSecurity) HasExcept() bool {
	return o != nil && o.Except != nil
}

// SetExcept gets a reference to the given []string and assigns it to the Except field.
func (o *ESSecurityFieldSecurity) SetExcept(v []string) {
	o.Except = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityFieldSecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Grant != nil {
		toSerialize["grant"] = o.Grant
	}
	if o.Except != nil {
		toSerialize["except"] = o.Except
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ESSecurityFieldSecurity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Grant  []string `json:"grant,omitempty"`
		Except []string `json:"except,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"grant", "except"})
	} else {
		return err
	}
	o.Grant = all.Grant
	o.Except = all.Except

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
