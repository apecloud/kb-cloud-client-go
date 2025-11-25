// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CreatePGExtensionResp struct {
	// OpsRequestName is the name of a KubeBlocks OpsRequest
	OpsRequestResult *OpsRequestName `json:"opsRequestResult,omitempty"`
	PgExtension      *PgExtension    `json:"PgExtension,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreatePGExtensionResp instantiates a new CreatePGExtensionResp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreatePGExtensionResp() *CreatePGExtensionResp {
	this := CreatePGExtensionResp{}
	return &this
}

// NewCreatePGExtensionRespWithDefaults instantiates a new CreatePGExtensionResp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreatePGExtensionRespWithDefaults() *CreatePGExtensionResp {
	this := CreatePGExtensionResp{}
	return &this
}

// GetOpsRequestResult returns the OpsRequestResult field value if set, zero value otherwise.
func (o *CreatePGExtensionResp) GetOpsRequestResult() OpsRequestName {
	if o == nil || o.OpsRequestResult == nil {
		var ret OpsRequestName
		return ret
	}
	return *o.OpsRequestResult
}

// GetOpsRequestResultOk returns a tuple with the OpsRequestResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePGExtensionResp) GetOpsRequestResultOk() (*OpsRequestName, bool) {
	if o == nil || o.OpsRequestResult == nil {
		return nil, false
	}
	return o.OpsRequestResult, true
}

// HasOpsRequestResult returns a boolean if a field has been set.
func (o *CreatePGExtensionResp) HasOpsRequestResult() bool {
	return o != nil && o.OpsRequestResult != nil
}

// SetOpsRequestResult gets a reference to the given OpsRequestName and assigns it to the OpsRequestResult field.
func (o *CreatePGExtensionResp) SetOpsRequestResult(v OpsRequestName) {
	o.OpsRequestResult = &v
}

// GetPgExtension returns the PgExtension field value if set, zero value otherwise.
func (o *CreatePGExtensionResp) GetPgExtension() PgExtension {
	if o == nil || o.PgExtension == nil {
		var ret PgExtension
		return ret
	}
	return *o.PgExtension
}

// GetPgExtensionOk returns a tuple with the PgExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePGExtensionResp) GetPgExtensionOk() (*PgExtension, bool) {
	if o == nil || o.PgExtension == nil {
		return nil, false
	}
	return o.PgExtension, true
}

// HasPgExtension returns a boolean if a field has been set.
func (o *CreatePGExtensionResp) HasPgExtension() bool {
	return o != nil && o.PgExtension != nil
}

// SetPgExtension gets a reference to the given PgExtension and assigns it to the PgExtension field.
func (o *CreatePGExtensionResp) SetPgExtension(v PgExtension) {
	o.PgExtension = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreatePGExtensionResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OpsRequestResult != nil {
		toSerialize["opsRequestResult"] = o.OpsRequestResult
	}
	if o.PgExtension != nil {
		toSerialize["PgExtension"] = o.PgExtension
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreatePGExtensionResp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpsRequestResult *OpsRequestName `json:"opsRequestResult,omitempty"`
		PgExtension      *PgExtension    `json:"PgExtension,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"opsRequestResult", "PgExtension"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OpsRequestResult != nil && all.OpsRequestResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OpsRequestResult = all.OpsRequestResult
	if all.PgExtension != nil && all.PgExtension.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PgExtension = all.PgExtension

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
