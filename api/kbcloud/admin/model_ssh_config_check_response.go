// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

type SshConfigCheckResponse struct {
	Results []SshConfigCheckResult `json:"results"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSshConfigCheckResponse instantiates a new SshConfigCheckResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSshConfigCheckResponse(results []SshConfigCheckResult) *SshConfigCheckResponse {
	this := SshConfigCheckResponse{}
	this.Results = results
	return &this
}

// NewSshConfigCheckResponseWithDefaults instantiates a new SshConfigCheckResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSshConfigCheckResponseWithDefaults() *SshConfigCheckResponse {
	this := SshConfigCheckResponse{}
	return &this
}

// GetResults returns the Results field value.
func (o *SshConfigCheckResponse) GetResults() []SshConfigCheckResult {
	if o == nil {
		var ret []SshConfigCheckResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SshConfigCheckResponse) GetResultsOk() (*[]SshConfigCheckResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value.
func (o *SshConfigCheckResponse) SetResults(v []SshConfigCheckResult) {
	o.Results = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SshConfigCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SshConfigCheckResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Results *[]SshConfigCheckResult `json:"results"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Results == nil {
		return fmt.Errorf("required field results missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"results"})
	} else {
		return err
	}
	o.Results = *all.Results

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
