// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsTableOptions struct {
	// The storage engine for the table
	Engine *string `json:"engine,omitempty"`
	// The character set for the table
	Charset *string `json:"charset,omitempty"`
	// The collation for the table
	Collation *string `json:"collation,omitempty"`
	// The type of the table
	Type *string `json:"type,omitempty"`
	// Other options for the table
	OtherOptions common.NullableList[string] `json:"otherOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsTableOptions instantiates a new DmsTableOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsTableOptions() *DmsTableOptions {
	this := DmsTableOptions{}
	return &this
}

// NewDmsTableOptionsWithDefaults instantiates a new DmsTableOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsTableOptionsWithDefaults() *DmsTableOptions {
	this := DmsTableOptions{}
	return &this
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *DmsTableOptions) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableOptions) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *DmsTableOptions) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *DmsTableOptions) SetEngine(v string) {
	o.Engine = &v
}

// GetCharset returns the Charset field value if set, zero value otherwise.
func (o *DmsTableOptions) GetCharset() string {
	if o == nil || o.Charset == nil {
		var ret string
		return ret
	}
	return *o.Charset
}

// GetCharsetOk returns a tuple with the Charset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableOptions) GetCharsetOk() (*string, bool) {
	if o == nil || o.Charset == nil {
		return nil, false
	}
	return o.Charset, true
}

// HasCharset returns a boolean if a field has been set.
func (o *DmsTableOptions) HasCharset() bool {
	return o != nil && o.Charset != nil
}

// SetCharset gets a reference to the given string and assigns it to the Charset field.
func (o *DmsTableOptions) SetCharset(v string) {
	o.Charset = &v
}

// GetCollation returns the Collation field value if set, zero value otherwise.
func (o *DmsTableOptions) GetCollation() string {
	if o == nil || o.Collation == nil {
		var ret string
		return ret
	}
	return *o.Collation
}

// GetCollationOk returns a tuple with the Collation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableOptions) GetCollationOk() (*string, bool) {
	if o == nil || o.Collation == nil {
		return nil, false
	}
	return o.Collation, true
}

// HasCollation returns a boolean if a field has been set.
func (o *DmsTableOptions) HasCollation() bool {
	return o != nil && o.Collation != nil
}

// SetCollation gets a reference to the given string and assigns it to the Collation field.
func (o *DmsTableOptions) SetCollation(v string) {
	o.Collation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DmsTableOptions) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableOptions) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DmsTableOptions) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DmsTableOptions) SetType(v string) {
	o.Type = &v
}

// GetOtherOptions returns the OtherOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsTableOptions) GetOtherOptions() []string {
	if o == nil || o.OtherOptions.Get() == nil {
		var ret []string
		return ret
	}
	return *o.OtherOptions.Get()
}

// GetOtherOptionsOk returns a tuple with the OtherOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsTableOptions) GetOtherOptionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherOptions.Get(), o.OtherOptions.IsSet()
}

// HasOtherOptions returns a boolean if a field has been set.
func (o *DmsTableOptions) HasOtherOptions() bool {
	return o != nil && o.OtherOptions.IsSet()
}

// SetOtherOptions gets a reference to the given common.NullableList[string] and assigns it to the OtherOptions field.
func (o *DmsTableOptions) SetOtherOptions(v []string) {
	o.OtherOptions.Set(&v)
}

// SetOtherOptionsNil sets the value for OtherOptions to be an explicit nil.
func (o *DmsTableOptions) SetOtherOptionsNil() {
	o.OtherOptions.Set(nil)
}

// UnsetOtherOptions ensures that no value is present for OtherOptions, not even an explicit nil.
func (o *DmsTableOptions) UnsetOtherOptions() {
	o.OtherOptions.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsTableOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Charset != nil {
		toSerialize["charset"] = o.Charset
	}
	if o.Collation != nil {
		toSerialize["collation"] = o.Collation
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.OtherOptions.IsSet() {
		toSerialize["otherOptions"] = o.OtherOptions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsTableOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine       *string                     `json:"engine,omitempty"`
		Charset      *string                     `json:"charset,omitempty"`
		Collation    *string                     `json:"collation,omitempty"`
		Type         *string                     `json:"type,omitempty"`
		OtherOptions common.NullableList[string] `json:"otherOptions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "charset", "collation", "type", "otherOptions"})
	} else {
		return err
	}
	o.Engine = all.Engine
	o.Charset = all.Charset
	o.Collation = all.Collation
	o.Type = all.Type
	o.OtherOptions = all.OtherOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
