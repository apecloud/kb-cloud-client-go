// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CreatePgExtensionResp struct {
	// OpsRequestName is the name of a KubeBlocks OpsRequest response
	OpsRequestResult *OpsRequestName           `json:"opsRequestResult,omitempty"`
	Items            []PgCreateExtensionResult `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreatePgExtensionResp instantiates a new CreatePgExtensionResp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreatePgExtensionResp(items []PgCreateExtensionResult) *CreatePgExtensionResp {
	this := CreatePgExtensionResp{}
	this.Items = items
	return &this
}

// NewCreatePgExtensionRespWithDefaults instantiates a new CreatePgExtensionResp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreatePgExtensionRespWithDefaults() *CreatePgExtensionResp {
	this := CreatePgExtensionResp{}
	return &this
}

// GetOpsRequestResult returns the OpsRequestResult field value if set, zero value otherwise.
func (o *CreatePgExtensionResp) GetOpsRequestResult() OpsRequestName {
	if o == nil || o.OpsRequestResult == nil {
		var ret OpsRequestName
		return ret
	}
	return *o.OpsRequestResult
}

// GetOpsRequestResultOk returns a tuple with the OpsRequestResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePgExtensionResp) GetOpsRequestResultOk() (*OpsRequestName, bool) {
	if o == nil || o.OpsRequestResult == nil {
		return nil, false
	}
	return o.OpsRequestResult, true
}

// HasOpsRequestResult returns a boolean if a field has been set.
func (o *CreatePgExtensionResp) HasOpsRequestResult() bool {
	return o != nil && o.OpsRequestResult != nil
}

// SetOpsRequestResult gets a reference to the given OpsRequestName and assigns it to the OpsRequestResult field.
func (o *CreatePgExtensionResp) SetOpsRequestResult(v OpsRequestName) {
	o.OpsRequestResult = &v
}

// GetItems returns the Items field value.
func (o *CreatePgExtensionResp) GetItems() []PgCreateExtensionResult {
	if o == nil {
		var ret []PgCreateExtensionResult
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CreatePgExtensionResp) GetItemsOk() (*[]PgCreateExtensionResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *CreatePgExtensionResp) SetItems(v []PgCreateExtensionResult) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreatePgExtensionResp) MarshalJSON() ([]byte, error) {
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
func (o *CreatePgExtensionResp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpsRequestResult *OpsRequestName            `json:"opsRequestResult,omitempty"`
		Items            *[]PgCreateExtensionResult `json:"items"`
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
