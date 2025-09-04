// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportSupportSourcesResponse Response for supported data sources list
type ImportSupportSourcesResponse struct {
	// Target engine type
	TargetEngine string `json:"targetEngine"`
	// List of supported data sources
	Sources []ImportSupportedSource `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportSupportSourcesResponse instantiates a new ImportSupportSourcesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportSupportSourcesResponse(targetEngine string, sources []ImportSupportedSource) *ImportSupportSourcesResponse {
	this := ImportSupportSourcesResponse{}
	this.TargetEngine = targetEngine
	this.Sources = sources
	return &this
}

// NewImportSupportSourcesResponseWithDefaults instantiates a new ImportSupportSourcesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportSupportSourcesResponseWithDefaults() *ImportSupportSourcesResponse {
	this := ImportSupportSourcesResponse{}
	return &this
}

// GetTargetEngine returns the TargetEngine field value.
func (o *ImportSupportSourcesResponse) GetTargetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetEngine
}

// GetTargetEngineOk returns a tuple with the TargetEngine field value
// and a boolean to check if the value has been set.
func (o *ImportSupportSourcesResponse) GetTargetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEngine, true
}

// SetTargetEngine sets field value.
func (o *ImportSupportSourcesResponse) SetTargetEngine(v string) {
	o.TargetEngine = v
}

// GetSources returns the Sources field value.
func (o *ImportSupportSourcesResponse) GetSources() []ImportSupportedSource {
	if o == nil {
		var ret []ImportSupportedSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ImportSupportSourcesResponse) GetSourcesOk() (*[]ImportSupportedSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ImportSupportSourcesResponse) SetSources(v []ImportSupportedSource) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportSupportSourcesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["targetEngine"] = o.TargetEngine
	toSerialize["sources"] = o.Sources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportSupportSourcesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TargetEngine *string                  `json:"targetEngine"`
		Sources      *[]ImportSupportedSource `json:"sources"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TargetEngine == nil {
		return fmt.Errorf("required field targetEngine missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"targetEngine", "sources"})
	} else {
		return err
	}
	o.TargetEngine = *all.TargetEngine
	o.Sources = *all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
