// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NetworkConfig Configuration of networking for this environment
// NODESCRIPTION NetworkConfig
//
// Deprecated: This model is deprecated.
type NetworkConfig struct {
	// Enable node port service for this environment
	NodePortEnabled *bool `json:"nodePortEnabled,omitempty"`
	// Enable load balancer service for this environment
	LbEnabled *bool `json:"lbEnabled,omitempty"`
	// Enable the Internet load balancer service for this environment
	InternetLbEnabled *bool `json:"internetLBEnabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkConfig instantiates a new NetworkConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkConfig() *NetworkConfig {
	this := NetworkConfig{}
	var nodePortEnabled bool = true
	this.NodePortEnabled = &nodePortEnabled
	var lbEnabled bool = true
	this.LbEnabled = &lbEnabled
	var internetLbEnabled bool = true
	this.InternetLbEnabled = &internetLbEnabled
	return &this
}

// NewNetworkConfigWithDefaults instantiates a new NetworkConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkConfigWithDefaults() *NetworkConfig {
	this := NetworkConfig{}
	var nodePortEnabled bool = true
	this.NodePortEnabled = &nodePortEnabled
	var lbEnabled bool = true
	this.LbEnabled = &lbEnabled
	var internetLbEnabled bool = true
	this.InternetLbEnabled = &internetLbEnabled
	return &this
}

// GetNodePortEnabled returns the NodePortEnabled field value if set, zero value otherwise.
func (o *NetworkConfig) GetNodePortEnabled() bool {
	if o == nil || o.NodePortEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NodePortEnabled
}

// GetNodePortEnabledOk returns a tuple with the NodePortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetNodePortEnabledOk() (*bool, bool) {
	if o == nil || o.NodePortEnabled == nil {
		return nil, false
	}
	return o.NodePortEnabled, true
}

// HasNodePortEnabled returns a boolean if a field has been set.
func (o *NetworkConfig) HasNodePortEnabled() bool {
	return o != nil && o.NodePortEnabled != nil
}

// SetNodePortEnabled gets a reference to the given bool and assigns it to the NodePortEnabled field.
func (o *NetworkConfig) SetNodePortEnabled(v bool) {
	o.NodePortEnabled = &v
}

// GetLbEnabled returns the LbEnabled field value if set, zero value otherwise.
func (o *NetworkConfig) GetLbEnabled() bool {
	if o == nil || o.LbEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LbEnabled
}

// GetLbEnabledOk returns a tuple with the LbEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetLbEnabledOk() (*bool, bool) {
	if o == nil || o.LbEnabled == nil {
		return nil, false
	}
	return o.LbEnabled, true
}

// HasLbEnabled returns a boolean if a field has been set.
func (o *NetworkConfig) HasLbEnabled() bool {
	return o != nil && o.LbEnabled != nil
}

// SetLbEnabled gets a reference to the given bool and assigns it to the LbEnabled field.
func (o *NetworkConfig) SetLbEnabled(v bool) {
	o.LbEnabled = &v
}

// GetInternetLbEnabled returns the InternetLbEnabled field value if set, zero value otherwise.
func (o *NetworkConfig) GetInternetLbEnabled() bool {
	if o == nil || o.InternetLbEnabled == nil {
		var ret bool
		return ret
	}
	return *o.InternetLbEnabled
}

// GetInternetLbEnabledOk returns a tuple with the InternetLbEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetInternetLbEnabledOk() (*bool, bool) {
	if o == nil || o.InternetLbEnabled == nil {
		return nil, false
	}
	return o.InternetLbEnabled, true
}

// HasInternetLbEnabled returns a boolean if a field has been set.
func (o *NetworkConfig) HasInternetLbEnabled() bool {
	return o != nil && o.InternetLbEnabled != nil
}

// SetInternetLbEnabled gets a reference to the given bool and assigns it to the InternetLbEnabled field.
func (o *NetworkConfig) SetInternetLbEnabled(v bool) {
	o.InternetLbEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodePortEnabled != nil {
		toSerialize["nodePortEnabled"] = o.NodePortEnabled
	}
	if o.LbEnabled != nil {
		toSerialize["lbEnabled"] = o.LbEnabled
	}
	if o.InternetLbEnabled != nil {
		toSerialize["internetLBEnabled"] = o.InternetLbEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodePortEnabled   *bool `json:"nodePortEnabled,omitempty"`
		LbEnabled         *bool `json:"lbEnabled,omitempty"`
		InternetLbEnabled *bool `json:"internetLBEnabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodePortEnabled", "lbEnabled", "internetLBEnabled"})
	} else {
		return err
	}
	o.NodePortEnabled = all.NodePortEnabled
	o.LbEnabled = all.LbEnabled
	o.InternetLbEnabled = all.InternetLbEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
