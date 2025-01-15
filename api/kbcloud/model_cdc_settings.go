// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CdcSettings struct {
	Config    *CdcClusterConfig   `json:"config,omitempty"`
	Account   *CdcClusterAccount  `json:"account,omitempty"`
	Endpoint  *CdcClusterEndpoint `json:"endpoint,omitempty"`
	Lifecycle *CdcLifecycle       `json:"lifecycle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcSettings instantiates a new CdcSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcSettings() *CdcSettings {
	this := CdcSettings{}
	return &this
}

// NewCdcSettingsWithDefaults instantiates a new CdcSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcSettingsWithDefaults() *CdcSettings {
	this := CdcSettings{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CdcSettings) GetConfig() CdcClusterConfig {
	if o == nil || o.Config == nil {
		var ret CdcClusterConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcSettings) GetConfigOk() (*CdcClusterConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CdcSettings) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given CdcClusterConfig and assigns it to the Config field.
func (o *CdcSettings) SetConfig(v CdcClusterConfig) {
	o.Config = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CdcSettings) GetAccount() CdcClusterAccount {
	if o == nil || o.Account == nil {
		var ret CdcClusterAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcSettings) GetAccountOk() (*CdcClusterAccount, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CdcSettings) HasAccount() bool {
	return o != nil && o.Account != nil
}

// SetAccount gets a reference to the given CdcClusterAccount and assigns it to the Account field.
func (o *CdcSettings) SetAccount(v CdcClusterAccount) {
	o.Account = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *CdcSettings) GetEndpoint() CdcClusterEndpoint {
	if o == nil || o.Endpoint == nil {
		var ret CdcClusterEndpoint
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcSettings) GetEndpointOk() (*CdcClusterEndpoint, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *CdcSettings) HasEndpoint() bool {
	return o != nil && o.Endpoint != nil
}

// SetEndpoint gets a reference to the given CdcClusterEndpoint and assigns it to the Endpoint field.
func (o *CdcSettings) SetEndpoint(v CdcClusterEndpoint) {
	o.Endpoint = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *CdcSettings) GetLifecycle() CdcLifecycle {
	if o == nil || o.Lifecycle == nil {
		var ret CdcLifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcSettings) GetLifecycleOk() (*CdcLifecycle, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *CdcSettings) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given CdcLifecycle and assigns it to the Lifecycle field.
func (o *CdcSettings) SetLifecycle(v CdcLifecycle) {
	o.Lifecycle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config    *CdcClusterConfig   `json:"config,omitempty"`
		Account   *CdcClusterAccount  `json:"account,omitempty"`
		Endpoint  *CdcClusterEndpoint `json:"endpoint,omitempty"`
		Lifecycle *CdcLifecycle       `json:"lifecycle,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"config", "account", "endpoint", "lifecycle"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Config != nil && all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = all.Config
	if all.Account != nil && all.Account.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Account = all.Account
	if all.Endpoint != nil && all.Endpoint.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Endpoint = all.Endpoint
	if all.Lifecycle != nil && all.Lifecycle.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Lifecycle = all.Lifecycle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
