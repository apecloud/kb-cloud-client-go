// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KingbaseSessionList struct {
	Items       []KingbaseSession        `json:"items"`
	Limit       int64                    `json:"limit"`
	Source      KingbaseDiagnosticSource `json:"source"`
	Warnings    []string                 `json:"warnings"`
	CollectedAt string                   `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKingbaseSessionList instantiates a new KingbaseSessionList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKingbaseSessionList(items []KingbaseSession, limit int64, source KingbaseDiagnosticSource, warnings []string, collectedAt string) *KingbaseSessionList {
	this := KingbaseSessionList{}
	this.Items = items
	this.Limit = limit
	this.Source = source
	this.Warnings = warnings
	this.CollectedAt = collectedAt
	return &this
}

// NewKingbaseSessionListWithDefaults instantiates a new KingbaseSessionList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKingbaseSessionListWithDefaults() *KingbaseSessionList {
	this := KingbaseSessionList{}
	return &this
}

// GetItems returns the Items field value.
func (o *KingbaseSessionList) GetItems() []KingbaseSession {
	if o == nil {
		var ret []KingbaseSession
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionList) GetItemsOk() (*[]KingbaseSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *KingbaseSessionList) SetItems(v []KingbaseSession) {
	o.Items = v
}

// GetLimit returns the Limit field value.
func (o *KingbaseSessionList) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionList) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *KingbaseSessionList) SetLimit(v int64) {
	o.Limit = v
}

// GetSource returns the Source field value.
func (o *KingbaseSessionList) GetSource() KingbaseDiagnosticSource {
	if o == nil {
		var ret KingbaseDiagnosticSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionList) GetSourceOk() (*KingbaseDiagnosticSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *KingbaseSessionList) SetSource(v KingbaseDiagnosticSource) {
	o.Source = v
}

// GetWarnings returns the Warnings field value.
func (o *KingbaseSessionList) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionList) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *KingbaseSessionList) SetWarnings(v []string) {
	o.Warnings = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *KingbaseSessionList) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionList) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *KingbaseSessionList) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o KingbaseSessionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["limit"] = o.Limit
	toSerialize["source"] = o.Source
	toSerialize["warnings"] = o.Warnings
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KingbaseSessionList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items       *[]KingbaseSession        `json:"items"`
		Limit       *int64                    `json:"limit"`
		Source      *KingbaseDiagnosticSource `json:"source"`
		Warnings    *[]string                 `json:"warnings"`
		CollectedAt *string                   `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "limit", "source", "warnings", "collectedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = *all.Items
	o.Limit = *all.Limit
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	o.Warnings = *all.Warnings
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
