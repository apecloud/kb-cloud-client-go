// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PgBatchCreateExtensionResp struct {
	// OpsRequestName is the name of a KubeBlocks OpsRequest response
	OpsRequestResult *OpsRequestName                `json:"opsRequestResult,omitempty"`
	Items            []PgBatchCreateExtensionResult `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPgBatchCreateExtensionResp instantiates a new PgBatchCreateExtensionResp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPgBatchCreateExtensionResp(items []PgBatchCreateExtensionResult) *PgBatchCreateExtensionResp {
	this := PgBatchCreateExtensionResp{}
	this.Items = items
	return &this
}

// NewPgBatchCreateExtensionRespWithDefaults instantiates a new PgBatchCreateExtensionResp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPgBatchCreateExtensionRespWithDefaults() *PgBatchCreateExtensionResp {
	this := PgBatchCreateExtensionResp{}
	return &this
}

// GetOpsRequestResult returns the OpsRequestResult field value if set, zero value otherwise.
func (o *PgBatchCreateExtensionResp) GetOpsRequestResult() OpsRequestName {
	if o == nil || o.OpsRequestResult == nil {
		var ret OpsRequestName
		return ret
	}
	return *o.OpsRequestResult
}

// GetOpsRequestResultOk returns a tuple with the OpsRequestResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionResp) GetOpsRequestResultOk() (*OpsRequestName, bool) {
	if o == nil || o.OpsRequestResult == nil {
		return nil, false
	}
	return o.OpsRequestResult, true
}

// HasOpsRequestResult returns a boolean if a field has been set.
func (o *PgBatchCreateExtensionResp) HasOpsRequestResult() bool {
	return o != nil && o.OpsRequestResult != nil
}

// SetOpsRequestResult gets a reference to the given OpsRequestName and assigns it to the OpsRequestResult field.
func (o *PgBatchCreateExtensionResp) SetOpsRequestResult(v OpsRequestName) {
	o.OpsRequestResult = &v
}

// GetItems returns the Items field value.
func (o *PgBatchCreateExtensionResp) GetItems() []PgBatchCreateExtensionResult {
	if o == nil {
		var ret []PgBatchCreateExtensionResult
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionResp) GetItemsOk() (*[]PgBatchCreateExtensionResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *PgBatchCreateExtensionResp) SetItems(v []PgBatchCreateExtensionResult) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PgBatchCreateExtensionResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OpsRequestResult != nil {
		toSerialize["opsRequestResult"] = o.OpsRequestResult
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PgBatchCreateExtensionResp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpsRequestResult *OpsRequestName                 `json:"opsRequestResult,omitempty"`
		Items            *[]PgBatchCreateExtensionResult `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"opsRequestResult", "items"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OpsRequestResult != nil && all.OpsRequestResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpsRequestResult = all.OpsRequestResult
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
