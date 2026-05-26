// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentRegenerateReportRequest struct {
	Style           *string  `json:"style,omitempty"`
	IncludeSections []string `json:"includeSections,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRegenerateReportRequest instantiates a new AiAgentRegenerateReportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRegenerateReportRequest() *AiAgentRegenerateReportRequest {
	this := AiAgentRegenerateReportRequest{}
	return &this
}

// NewAiAgentRegenerateReportRequestWithDefaults instantiates a new AiAgentRegenerateReportRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRegenerateReportRequestWithDefaults() *AiAgentRegenerateReportRequest {
	this := AiAgentRegenerateReportRequest{}
	return &this
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *AiAgentRegenerateReportRequest) GetStyle() string {
	if o == nil || o.Style == nil {
		var ret string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRegenerateReportRequest) GetStyleOk() (*string, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *AiAgentRegenerateReportRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given string and assigns it to the Style field.
func (o *AiAgentRegenerateReportRequest) SetStyle(v string) {
	o.Style = &v
}

// GetIncludeSections returns the IncludeSections field value if set, zero value otherwise.
func (o *AiAgentRegenerateReportRequest) GetIncludeSections() []string {
	if o == nil || o.IncludeSections == nil {
		var ret []string
		return ret
	}
	return o.IncludeSections
}

// GetIncludeSectionsOk returns a tuple with the IncludeSections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRegenerateReportRequest) GetIncludeSectionsOk() (*[]string, bool) {
	if o == nil || o.IncludeSections == nil {
		return nil, false
	}
	return &o.IncludeSections, true
}

// HasIncludeSections returns a boolean if a field has been set.
func (o *AiAgentRegenerateReportRequest) HasIncludeSections() bool {
	return o != nil && o.IncludeSections != nil
}

// SetIncludeSections gets a reference to the given []string and assigns it to the IncludeSections field.
func (o *AiAgentRegenerateReportRequest) SetIncludeSections(v []string) {
	o.IncludeSections = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRegenerateReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.IncludeSections != nil {
		toSerialize["includeSections"] = o.IncludeSections
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentRegenerateReportRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Style           *string  `json:"style,omitempty"`
		IncludeSections []string `json:"includeSections,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"style", "includeSections"})
	} else {
		return err
	}
	o.Style = all.Style
	o.IncludeSections = all.IncludeSections

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
