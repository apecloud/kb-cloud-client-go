// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ImportStringValidation Validation rules for string field type
type ImportStringValidation struct {
	// Minimum length
	MinLength *int32 `json:"minLength,omitempty"`
	// Maximum length
	MaxLength *int32 `json:"maxLength,omitempty"`
	// Regex pattern
	Pattern *string `json:"pattern,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportStringValidation instantiates a new ImportStringValidation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportStringValidation() *ImportStringValidation {
	this := ImportStringValidation{}
	return &this
}

// NewImportStringValidationWithDefaults instantiates a new ImportStringValidation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportStringValidationWithDefaults() *ImportStringValidation {
	this := ImportStringValidation{}
	return &this
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *ImportStringValidation) GetMinLength() int32 {
	if o == nil || o.MinLength == nil {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStringValidation) GetMinLengthOk() (*int32, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *ImportStringValidation) HasMinLength() bool {
	return o != nil && o.MinLength != nil
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *ImportStringValidation) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *ImportStringValidation) GetMaxLength() int32 {
	if o == nil || o.MaxLength == nil {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStringValidation) GetMaxLengthOk() (*int32, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *ImportStringValidation) HasMaxLength() bool {
	return o != nil && o.MaxLength != nil
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *ImportStringValidation) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *ImportStringValidation) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStringValidation) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *ImportStringValidation) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *ImportStringValidation) SetPattern(v string) {
	o.Pattern = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportStringValidation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportStringValidation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MinLength *int32  `json:"minLength,omitempty"`
		MaxLength *int32  `json:"maxLength,omitempty"`
		Pattern   *string `json:"pattern,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"minLength", "maxLength", "pattern"})
	} else {
		return err
	}
	o.MinLength = all.MinLength
	o.MaxLength = all.MaxLength
	o.Pattern = all.Pattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
