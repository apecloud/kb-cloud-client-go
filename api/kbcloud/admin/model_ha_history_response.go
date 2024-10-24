// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// HaHistoryResponse hahistory is the payload to get ha history of a KubeBlocks cluster
// NODESCRIPTION HaHistoryResponse
//
// Deprecated: This model is deprecated.
type HaHistoryResponse struct {
	// NODESCRIPTION ComponentName
	ComponentName *string `json:"componentName,omitempty"`
	// NODESCRIPTION Records
	Records []HaHistoryResponseRecordsItem `json:"records,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHaHistoryResponse instantiates a new HaHistoryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHaHistoryResponse() *HaHistoryResponse {
	this := HaHistoryResponse{}
	return &this
}

// NewHaHistoryResponseWithDefaults instantiates a new HaHistoryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHaHistoryResponseWithDefaults() *HaHistoryResponse {
	this := HaHistoryResponse{}
	return &this
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *HaHistoryResponse) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponse) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *HaHistoryResponse) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *HaHistoryResponse) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *HaHistoryResponse) GetRecords() []HaHistoryResponseRecordsItem {
	if o == nil || o.Records == nil {
		var ret []HaHistoryResponseRecordsItem
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponse) GetRecordsOk() (*[]HaHistoryResponseRecordsItem, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return &o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *HaHistoryResponse) HasRecords() bool {
	return o != nil && o.Records != nil
}

// SetRecords gets a reference to the given []HaHistoryResponseRecordsItem and assigns it to the Records field.
func (o *HaHistoryResponse) SetRecords(v []HaHistoryResponseRecordsItem) {
	o.Records = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HaHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HaHistoryResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentName *string                        `json:"componentName,omitempty"`
		Records       []HaHistoryResponseRecordsItem `json:"records,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentName", "records"})
	} else {
		return err
	}
	o.ComponentName = all.ComponentName
	o.Records = all.Records

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
