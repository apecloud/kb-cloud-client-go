// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BlueGreenDeploymentEngineOption blueGreenDeploymentEngine is the engine of a blue-green deployment.
type BlueGreenDeploymentEngineOption struct {
	// Whether the blue-green deployment is enabled.
	Enabled  *bool                             `json:"enabled,omitempty"`
	Accounts *BlueGreenDeploymentAccountOption `json:"accounts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentEngineOption instantiates a new BlueGreenDeploymentEngineOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentEngineOption() *BlueGreenDeploymentEngineOption {
	this := BlueGreenDeploymentEngineOption{}
	return &this
}

// NewBlueGreenDeploymentEngineOptionWithDefaults instantiates a new BlueGreenDeploymentEngineOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentEngineOptionWithDefaults() *BlueGreenDeploymentEngineOption {
	this := BlueGreenDeploymentEngineOption{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BlueGreenDeploymentEngineOption) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentEngineOption) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BlueGreenDeploymentEngineOption) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BlueGreenDeploymentEngineOption) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *BlueGreenDeploymentEngineOption) GetAccounts() BlueGreenDeploymentAccountOption {
	if o == nil || o.Accounts == nil {
		var ret BlueGreenDeploymentAccountOption
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentEngineOption) GetAccountsOk() (*BlueGreenDeploymentAccountOption, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *BlueGreenDeploymentEngineOption) HasAccounts() bool {
	return o != nil && o.Accounts != nil
}

// SetAccounts gets a reference to the given BlueGreenDeploymentAccountOption and assigns it to the Accounts field.
func (o *BlueGreenDeploymentEngineOption) SetAccounts(v BlueGreenDeploymentAccountOption) {
	o.Accounts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentEngineOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentEngineOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled  *bool                             `json:"enabled,omitempty"`
		Accounts *BlueGreenDeploymentAccountOption `json:"accounts,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "accounts"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.Accounts != nil && all.Accounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Accounts = all.Accounts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
