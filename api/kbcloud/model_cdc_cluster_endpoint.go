// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type CdcClusterEndpoint struct {
	Role         *string               `json:"role,omitempty"`
	Component    *string               `json:"component,omitempty"`
	PortName     *string               `json:"portName,omitempty"`
	Port         *int32                `json:"port,omitempty"`
	AuthDatabase common.NullableString `json:"authDatabase,omitempty"`
	NetworkType  *CdcNetworkType       `json:"networkType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcClusterEndpoint instantiates a new CdcClusterEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcClusterEndpoint() *CdcClusterEndpoint {
	this := CdcClusterEndpoint{}
	return &this
}

// NewCdcClusterEndpointWithDefaults instantiates a new CdcClusterEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcClusterEndpointWithDefaults() *CdcClusterEndpoint {
	this := CdcClusterEndpoint{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *CdcClusterEndpoint) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterEndpoint) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *CdcClusterEndpoint) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *CdcClusterEndpoint) SetRole(v string) {
	o.Role = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CdcClusterEndpoint) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterEndpoint) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CdcClusterEndpoint) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CdcClusterEndpoint) SetComponent(v string) {
	o.Component = &v
}

// GetPortName returns the PortName field value if set, zero value otherwise.
func (o *CdcClusterEndpoint) GetPortName() string {
	if o == nil || o.PortName == nil {
		var ret string
		return ret
	}
	return *o.PortName
}

// GetPortNameOk returns a tuple with the PortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterEndpoint) GetPortNameOk() (*string, bool) {
	if o == nil || o.PortName == nil {
		return nil, false
	}
	return o.PortName, true
}

// HasPortName returns a boolean if a field has been set.
func (o *CdcClusterEndpoint) HasPortName() bool {
	return o != nil && o.PortName != nil
}

// SetPortName gets a reference to the given string and assigns it to the PortName field.
func (o *CdcClusterEndpoint) SetPortName(v string) {
	o.PortName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CdcClusterEndpoint) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterEndpoint) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CdcClusterEndpoint) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *CdcClusterEndpoint) SetPort(v int32) {
	o.Port = &v
}

// GetAuthDatabase returns the AuthDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CdcClusterEndpoint) GetAuthDatabase() string {
	if o == nil || o.AuthDatabase.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthDatabase.Get()
}

// GetAuthDatabaseOk returns a tuple with the AuthDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CdcClusterEndpoint) GetAuthDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthDatabase.Get(), o.AuthDatabase.IsSet()
}

// HasAuthDatabase returns a boolean if a field has been set.
func (o *CdcClusterEndpoint) HasAuthDatabase() bool {
	return o != nil && o.AuthDatabase.IsSet()
}

// SetAuthDatabase gets a reference to the given common.NullableString and assigns it to the AuthDatabase field.
func (o *CdcClusterEndpoint) SetAuthDatabase(v string) {
	o.AuthDatabase.Set(&v)
}

// SetAuthDatabaseNil sets the value for AuthDatabase to be an explicit nil.
func (o *CdcClusterEndpoint) SetAuthDatabaseNil() {
	o.AuthDatabase.Set(nil)
}

// UnsetAuthDatabase ensures that no value is present for AuthDatabase, not even an explicit nil.
func (o *CdcClusterEndpoint) UnsetAuthDatabase() {
	o.AuthDatabase.Unset()
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *CdcClusterEndpoint) GetNetworkType() CdcNetworkType {
	if o == nil || o.NetworkType == nil {
		var ret CdcNetworkType
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterEndpoint) GetNetworkTypeOk() (*CdcNetworkType, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *CdcClusterEndpoint) HasNetworkType() bool {
	return o != nil && o.NetworkType != nil
}

// SetNetworkType gets a reference to the given CdcNetworkType and assigns it to the NetworkType field.
func (o *CdcClusterEndpoint) SetNetworkType(v CdcNetworkType) {
	o.NetworkType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcClusterEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.PortName != nil {
		toSerialize["portName"] = o.PortName
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.AuthDatabase.IsSet() {
		toSerialize["authDatabase"] = o.AuthDatabase.Get()
	}
	if o.NetworkType != nil {
		toSerialize["networkType"] = o.NetworkType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcClusterEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role         *string               `json:"role,omitempty"`
		Component    *string               `json:"component,omitempty"`
		PortName     *string               `json:"portName,omitempty"`
		Port         *int32                `json:"port,omitempty"`
		AuthDatabase common.NullableString `json:"authDatabase,omitempty"`
		NetworkType  *CdcNetworkType       `json:"networkType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "component", "portName", "port", "authDatabase", "networkType"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Role = all.Role
	o.Component = all.Component
	o.PortName = all.PortName
	o.Port = all.Port
	o.AuthDatabase = all.AuthDatabase
	if all.NetworkType != nil && !all.NetworkType.IsValid() {
		hasInvalidField = true
	} else {
		o.NetworkType = all.NetworkType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
