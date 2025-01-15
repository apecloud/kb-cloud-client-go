// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CdcLifecycleAction struct {
	Name        *string         `json:"name,omitempty"`
	SqlExecutor *CdcSqlExecutor `json:"sqlExecutor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcLifecycleAction instantiates a new CdcLifecycleAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcLifecycleAction() *CdcLifecycleAction {
	this := CdcLifecycleAction{}
	return &this
}

// NewCdcLifecycleActionWithDefaults instantiates a new CdcLifecycleAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcLifecycleActionWithDefaults() *CdcLifecycleAction {
	this := CdcLifecycleAction{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CdcLifecycleAction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcLifecycleAction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CdcLifecycleAction) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CdcLifecycleAction) SetName(v string) {
	o.Name = &v
}

// GetSqlExecutor returns the SqlExecutor field value if set, zero value otherwise.
func (o *CdcLifecycleAction) GetSqlExecutor() CdcSqlExecutor {
	if o == nil || o.SqlExecutor == nil {
		var ret CdcSqlExecutor
		return ret
	}
	return *o.SqlExecutor
}

// GetSqlExecutorOk returns a tuple with the SqlExecutor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcLifecycleAction) GetSqlExecutorOk() (*CdcSqlExecutor, bool) {
	if o == nil || o.SqlExecutor == nil {
		return nil, false
	}
	return o.SqlExecutor, true
}

// HasSqlExecutor returns a boolean if a field has been set.
func (o *CdcLifecycleAction) HasSqlExecutor() bool {
	return o != nil && o.SqlExecutor != nil
}

// SetSqlExecutor gets a reference to the given CdcSqlExecutor and assigns it to the SqlExecutor field.
func (o *CdcLifecycleAction) SetSqlExecutor(v CdcSqlExecutor) {
	o.SqlExecutor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcLifecycleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SqlExecutor != nil {
		toSerialize["sqlExecutor"] = o.SqlExecutor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcLifecycleAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string         `json:"name,omitempty"`
		SqlExecutor *CdcSqlExecutor `json:"sqlExecutor,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "sqlExecutor"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.SqlExecutor != nil && all.SqlExecutor.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SqlExecutor = all.SqlExecutor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
