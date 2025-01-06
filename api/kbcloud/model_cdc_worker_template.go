// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CdcWorkerTemplate struct {
	UsingTool *CdcToolTemplate `json:"usingTool,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcWorkerTemplate instantiates a new CdcWorkerTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcWorkerTemplate() *CdcWorkerTemplate {
	this := CdcWorkerTemplate{}
	return &this
}

// NewCdcWorkerTemplateWithDefaults instantiates a new CdcWorkerTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcWorkerTemplateWithDefaults() *CdcWorkerTemplate {
	this := CdcWorkerTemplate{}
	return &this
}

// GetUsingTool returns the UsingTool field value if set, zero value otherwise.
func (o *CdcWorkerTemplate) GetUsingTool() CdcToolTemplate {
	if o == nil || o.UsingTool == nil {
		var ret CdcToolTemplate
		return ret
	}
	return *o.UsingTool
}

// GetUsingToolOk returns a tuple with the UsingTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcWorkerTemplate) GetUsingToolOk() (*CdcToolTemplate, bool) {
	if o == nil || o.UsingTool == nil {
		return nil, false
	}
	return o.UsingTool, true
}

// HasUsingTool returns a boolean if a field has been set.
func (o *CdcWorkerTemplate) HasUsingTool() bool {
	return o != nil && o.UsingTool != nil
}

// SetUsingTool gets a reference to the given CdcToolTemplate and assigns it to the UsingTool field.
func (o *CdcWorkerTemplate) SetUsingTool(v CdcToolTemplate) {
	o.UsingTool = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcWorkerTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.UsingTool != nil {
		toSerialize["usingTool"] = o.UsingTool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcWorkerTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UsingTool *CdcToolTemplate `json:"usingTool,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"usingTool"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.UsingTool != nil && all.UsingTool.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UsingTool = all.UsingTool

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
