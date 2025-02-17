// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DatabaseOptionCreateOptionsItem struct {
	EnableCharset *bool `json:"enableCharset,omitempty"`
	EnableCollate *bool `json:"enableCollate,omitempty"`
	EnableCtype   *bool `json:"enableCtype,omitempty"`
	// 适用的数据库模式, 用于区分特殊的数据库模式, 如oceanbase的mysql模式
	//
	Mode *string `json:"mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseOptionCreateOptionsItem instantiates a new DatabaseOptionCreateOptionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseOptionCreateOptionsItem() *DatabaseOptionCreateOptionsItem {
	this := DatabaseOptionCreateOptionsItem{}
	return &this
}

// NewDatabaseOptionCreateOptionsItemWithDefaults instantiates a new DatabaseOptionCreateOptionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseOptionCreateOptionsItemWithDefaults() *DatabaseOptionCreateOptionsItem {
	this := DatabaseOptionCreateOptionsItem{}
	return &this
}

// GetEnableCharset returns the EnableCharset field value if set, zero value otherwise.
func (o *DatabaseOptionCreateOptionsItem) GetEnableCharset() bool {
	if o == nil || o.EnableCharset == nil {
		var ret bool
		return ret
	}
	return *o.EnableCharset
}

// GetEnableCharsetOk returns a tuple with the EnableCharset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptionCreateOptionsItem) GetEnableCharsetOk() (*bool, bool) {
	if o == nil || o.EnableCharset == nil {
		return nil, false
	}
	return o.EnableCharset, true
}

// HasEnableCharset returns a boolean if a field has been set.
func (o *DatabaseOptionCreateOptionsItem) HasEnableCharset() bool {
	return o != nil && o.EnableCharset != nil
}

// SetEnableCharset gets a reference to the given bool and assigns it to the EnableCharset field.
func (o *DatabaseOptionCreateOptionsItem) SetEnableCharset(v bool) {
	o.EnableCharset = &v
}

// GetEnableCollate returns the EnableCollate field value if set, zero value otherwise.
func (o *DatabaseOptionCreateOptionsItem) GetEnableCollate() bool {
	if o == nil || o.EnableCollate == nil {
		var ret bool
		return ret
	}
	return *o.EnableCollate
}

// GetEnableCollateOk returns a tuple with the EnableCollate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptionCreateOptionsItem) GetEnableCollateOk() (*bool, bool) {
	if o == nil || o.EnableCollate == nil {
		return nil, false
	}
	return o.EnableCollate, true
}

// HasEnableCollate returns a boolean if a field has been set.
func (o *DatabaseOptionCreateOptionsItem) HasEnableCollate() bool {
	return o != nil && o.EnableCollate != nil
}

// SetEnableCollate gets a reference to the given bool and assigns it to the EnableCollate field.
func (o *DatabaseOptionCreateOptionsItem) SetEnableCollate(v bool) {
	o.EnableCollate = &v
}

// GetEnableCtype returns the EnableCtype field value if set, zero value otherwise.
func (o *DatabaseOptionCreateOptionsItem) GetEnableCtype() bool {
	if o == nil || o.EnableCtype == nil {
		var ret bool
		return ret
	}
	return *o.EnableCtype
}

// GetEnableCtypeOk returns a tuple with the EnableCtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptionCreateOptionsItem) GetEnableCtypeOk() (*bool, bool) {
	if o == nil || o.EnableCtype == nil {
		return nil, false
	}
	return o.EnableCtype, true
}

// HasEnableCtype returns a boolean if a field has been set.
func (o *DatabaseOptionCreateOptionsItem) HasEnableCtype() bool {
	return o != nil && o.EnableCtype != nil
}

// SetEnableCtype gets a reference to the given bool and assigns it to the EnableCtype field.
func (o *DatabaseOptionCreateOptionsItem) SetEnableCtype(v bool) {
	o.EnableCtype = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DatabaseOptionCreateOptionsItem) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseOptionCreateOptionsItem) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DatabaseOptionCreateOptionsItem) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DatabaseOptionCreateOptionsItem) SetMode(v string) {
	o.Mode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseOptionCreateOptionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnableCharset != nil {
		toSerialize["enableCharset"] = o.EnableCharset
	}
	if o.EnableCollate != nil {
		toSerialize["enableCollate"] = o.EnableCollate
	}
	if o.EnableCtype != nil {
		toSerialize["enableCtype"] = o.EnableCtype
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabaseOptionCreateOptionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnableCharset *bool   `json:"enableCharset,omitempty"`
		EnableCollate *bool   `json:"enableCollate,omitempty"`
		EnableCtype   *bool   `json:"enableCtype,omitempty"`
		Mode          *string `json:"mode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enableCharset", "enableCollate", "enableCtype", "mode"})
	} else {
		return err
	}
	o.EnableCharset = all.EnableCharset
	o.EnableCollate = all.EnableCollate
	o.EnableCtype = all.EnableCtype
	o.Mode = all.Mode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
