// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ImportOption struct {
	// List of supported data sources for import
	SupportedSources []map[string]interface{} `json:"supportedSources,omitempty"`
	// Engine name for sink/target database
	SinkEngineName *string `json:"sinkEngineName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportOption instantiates a new ImportOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportOption() *ImportOption {
	this := ImportOption{}
	return &this
}

// NewImportOptionWithDefaults instantiates a new ImportOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportOptionWithDefaults() *ImportOption {
	this := ImportOption{}
	return &this
}

// GetSupportedSources returns the SupportedSources field value if set, zero value otherwise.
func (o *ImportOption) GetSupportedSources() []map[string]interface{} {
	if o == nil || o.SupportedSources == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.SupportedSources
}

// GetSupportedSourcesOk returns a tuple with the SupportedSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportOption) GetSupportedSourcesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.SupportedSources == nil {
		return nil, false
	}
	return &o.SupportedSources, true
}

// HasSupportedSources returns a boolean if a field has been set.
func (o *ImportOption) HasSupportedSources() bool {
	return o != nil && o.SupportedSources != nil
}

// SetSupportedSources gets a reference to the given []map[string]interface{} and assigns it to the SupportedSources field.
func (o *ImportOption) SetSupportedSources(v []map[string]interface{}) {
	o.SupportedSources = v
}

// GetSinkEngineName returns the SinkEngineName field value if set, zero value otherwise.
func (o *ImportOption) GetSinkEngineName() string {
	if o == nil || o.SinkEngineName == nil {
		var ret string
		return ret
	}
	return *o.SinkEngineName
}

// GetSinkEngineNameOk returns a tuple with the SinkEngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportOption) GetSinkEngineNameOk() (*string, bool) {
	if o == nil || o.SinkEngineName == nil {
		return nil, false
	}
	return o.SinkEngineName, true
}

// HasSinkEngineName returns a boolean if a field has been set.
func (o *ImportOption) HasSinkEngineName() bool {
	return o != nil && o.SinkEngineName != nil
}

// SetSinkEngineName gets a reference to the given string and assigns it to the SinkEngineName field.
func (o *ImportOption) SetSinkEngineName(v string) {
	o.SinkEngineName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SupportedSources != nil {
		toSerialize["supportedSources"] = o.SupportedSources
	}
	if o.SinkEngineName != nil {
		toSerialize["sinkEngineName"] = o.SinkEngineName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SupportedSources []map[string]interface{} `json:"supportedSources,omitempty"`
		SinkEngineName   *string                  `json:"sinkEngineName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supportedSources", "sinkEngineName"})
	} else {
		return err
	}
	o.SupportedSources = all.SupportedSources
	o.SinkEngineName = all.SinkEngineName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
