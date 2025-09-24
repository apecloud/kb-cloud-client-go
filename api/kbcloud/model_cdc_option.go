// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CdcOption struct {
	Versions []string           `json:"versions,omitempty"`
	Source   *CdcSettings       `json:"source,omitempty"`
	Sink     *CdcSettings       `json:"sink,omitempty"`
	Worker   *CdcWorkerTemplate `json:"worker,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcOption instantiates a new CdcOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcOption() *CdcOption {
	this := CdcOption{}
	return &this
}

// NewCdcOptionWithDefaults instantiates a new CdcOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcOptionWithDefaults() *CdcOption {
	this := CdcOption{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *CdcOption) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcOption) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *CdcOption) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *CdcOption) SetVersions(v []string) {
	o.Versions = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CdcOption) GetSource() CdcSettings {
	if o == nil || o.Source == nil {
		var ret CdcSettings
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcOption) GetSourceOk() (*CdcSettings, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CdcOption) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given CdcSettings and assigns it to the Source field.
func (o *CdcOption) SetSource(v CdcSettings) {
	o.Source = &v
}

// GetSink returns the Sink field value if set, zero value otherwise.
func (o *CdcOption) GetSink() CdcSettings {
	if o == nil || o.Sink == nil {
		var ret CdcSettings
		return ret
	}
	return *o.Sink
}

// GetSinkOk returns a tuple with the Sink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcOption) GetSinkOk() (*CdcSettings, bool) {
	if o == nil || o.Sink == nil {
		return nil, false
	}
	return o.Sink, true
}

// HasSink returns a boolean if a field has been set.
func (o *CdcOption) HasSink() bool {
	return o != nil && o.Sink != nil
}

// SetSink gets a reference to the given CdcSettings and assigns it to the Sink field.
func (o *CdcOption) SetSink(v CdcSettings) {
	o.Sink = &v
}

// GetWorker returns the Worker field value if set, zero value otherwise.
func (o *CdcOption) GetWorker() CdcWorkerTemplate {
	if o == nil || o.Worker == nil {
		var ret CdcWorkerTemplate
		return ret
	}
	return *o.Worker
}

// GetWorkerOk returns a tuple with the Worker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcOption) GetWorkerOk() (*CdcWorkerTemplate, bool) {
	if o == nil || o.Worker == nil {
		return nil, false
	}
	return o.Worker, true
}

// HasWorker returns a boolean if a field has been set.
func (o *CdcOption) HasWorker() bool {
	return o != nil && o.Worker != nil
}

// SetWorker gets a reference to the given CdcWorkerTemplate and assigns it to the Worker field.
func (o *CdcOption) SetWorker(v CdcWorkerTemplate) {
	o.Worker = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Sink != nil {
		toSerialize["sink"] = o.Sink
	}
	if o.Worker != nil {
		toSerialize["worker"] = o.Worker
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Versions []string           `json:"versions,omitempty"`
		Source   *CdcSettings       `json:"source,omitempty"`
		Sink     *CdcSettings       `json:"sink,omitempty"`
		Worker   *CdcWorkerTemplate `json:"worker,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"versions", "source", "sink", "worker"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Versions = all.Versions
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source
	if all.Sink != nil && all.Sink.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sink = all.Sink
	if all.Worker != nil && all.Worker.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Worker = all.Worker

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
