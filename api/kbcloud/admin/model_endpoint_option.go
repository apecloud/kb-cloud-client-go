// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EndpointOption struct {
	Title      LocalizedDescription `json:"title"`
	Component  string               `json:"component"`
	PortName   string               `json:"portName"`
	Port       int32                `json:"port"`
	Protocol   string               `json:"protocol"`
	TargetPort string               `json:"targetPort"`
	Type       []string             `json:"type"`
	// whether the endpoint supports system use, such as health check, dms, databases & accounts management etc.
	SupportsSystemUse *bool `json:"supportsSystemUse,omitempty"`
	// whether the engine supports readonly endpoint
	SupportsReadonly *bool `json:"supportsReadonly,omitempty"`
	// service name pattern, e.g. ClusterName-ComponentName or .ClusterName`
	ServicePattern *EngineOptionsServicePattern `json:"servicePattern,omitempty"`
	// ServiceName regular expression
	ServiceNameRegex *string `json:"serviceNameRegex,omitempty"`
	// service suffix, defined in ComponentDefinition
	ServiceName *string `json:"serviceName,omitempty"`
	// selector of k8s service
	Selector map[string]string `json:"selector,omitempty"`
	// whether the endpoint follows the network mode of the component
	FollowNetworkMode *bool `json:"followNetworkMode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEndpointOption instantiates a new EndpointOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEndpointOption(title LocalizedDescription, component string, portName string, port int32, protocol string, targetPort string, typeVar []string) *EndpointOption {
	this := EndpointOption{}
	this.Title = title
	this.Component = component
	this.PortName = portName
	this.Port = port
	this.Protocol = protocol
	this.TargetPort = targetPort
	this.Type = typeVar
	var servicePattern EngineOptionsServicePattern = EngineOptionsServicePatternClusterComponent
	this.ServicePattern = &servicePattern
	var followNetworkMode bool = false
	this.FollowNetworkMode = &followNetworkMode
	return &this
}

// NewEndpointOptionWithDefaults instantiates a new EndpointOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEndpointOptionWithDefaults() *EndpointOption {
	this := EndpointOption{}
	var servicePattern EngineOptionsServicePattern = EngineOptionsServicePatternClusterComponent
	this.ServicePattern = &servicePattern
	var followNetworkMode bool = false
	this.FollowNetworkMode = &followNetworkMode
	return &this
}

// GetTitle returns the Title field value.
func (o *EndpointOption) GetTitle() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetTitleOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *EndpointOption) SetTitle(v LocalizedDescription) {
	o.Title = v
}

// GetComponent returns the Component field value.
func (o *EndpointOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *EndpointOption) SetComponent(v string) {
	o.Component = v
}

// GetPortName returns the PortName field value.
func (o *EndpointOption) GetPortName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PortName
}

// GetPortNameOk returns a tuple with the PortName field value
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetPortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortName, true
}

// SetPortName sets field value.
func (o *EndpointOption) SetPortName(v string) {
	o.PortName = v
}

// GetPort returns the Port field value.
func (o *EndpointOption) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value.
func (o *EndpointOption) SetPort(v int32) {
	o.Port = v
}

// GetProtocol returns the Protocol field value.
func (o *EndpointOption) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value.
func (o *EndpointOption) SetProtocol(v string) {
	o.Protocol = v
}

// GetTargetPort returns the TargetPort field value.
func (o *EndpointOption) GetTargetPort() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetTargetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPort, true
}

// SetTargetPort sets field value.
func (o *EndpointOption) SetTargetPort(v string) {
	o.TargetPort = v
}

// GetType returns the Type field value.
func (o *EndpointOption) GetType() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetTypeOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EndpointOption) SetType(v []string) {
	o.Type = v
}

// GetSupportsSystemUse returns the SupportsSystemUse field value if set, zero value otherwise.
func (o *EndpointOption) GetSupportsSystemUse() bool {
	if o == nil || o.SupportsSystemUse == nil {
		var ret bool
		return ret
	}
	return *o.SupportsSystemUse
}

// GetSupportsSystemUseOk returns a tuple with the SupportsSystemUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetSupportsSystemUseOk() (*bool, bool) {
	if o == nil || o.SupportsSystemUse == nil {
		return nil, false
	}
	return o.SupportsSystemUse, true
}

// HasSupportsSystemUse returns a boolean if a field has been set.
func (o *EndpointOption) HasSupportsSystemUse() bool {
	return o != nil && o.SupportsSystemUse != nil
}

// SetSupportsSystemUse gets a reference to the given bool and assigns it to the SupportsSystemUse field.
func (o *EndpointOption) SetSupportsSystemUse(v bool) {
	o.SupportsSystemUse = &v
}

// GetSupportsReadonly returns the SupportsReadonly field value if set, zero value otherwise.
func (o *EndpointOption) GetSupportsReadonly() bool {
	if o == nil || o.SupportsReadonly == nil {
		var ret bool
		return ret
	}
	return *o.SupportsReadonly
}

// GetSupportsReadonlyOk returns a tuple with the SupportsReadonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetSupportsReadonlyOk() (*bool, bool) {
	if o == nil || o.SupportsReadonly == nil {
		return nil, false
	}
	return o.SupportsReadonly, true
}

// HasSupportsReadonly returns a boolean if a field has been set.
func (o *EndpointOption) HasSupportsReadonly() bool {
	return o != nil && o.SupportsReadonly != nil
}

// SetSupportsReadonly gets a reference to the given bool and assigns it to the SupportsReadonly field.
func (o *EndpointOption) SetSupportsReadonly(v bool) {
	o.SupportsReadonly = &v
}

// GetServicePattern returns the ServicePattern field value if set, zero value otherwise.
func (o *EndpointOption) GetServicePattern() EngineOptionsServicePattern {
	if o == nil || o.ServicePattern == nil {
		var ret EngineOptionsServicePattern
		return ret
	}
	return *o.ServicePattern
}

// GetServicePatternOk returns a tuple with the ServicePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetServicePatternOk() (*EngineOptionsServicePattern, bool) {
	if o == nil || o.ServicePattern == nil {
		return nil, false
	}
	return o.ServicePattern, true
}

// HasServicePattern returns a boolean if a field has been set.
func (o *EndpointOption) HasServicePattern() bool {
	return o != nil && o.ServicePattern != nil
}

// SetServicePattern gets a reference to the given EngineOptionsServicePattern and assigns it to the ServicePattern field.
func (o *EndpointOption) SetServicePattern(v EngineOptionsServicePattern) {
	o.ServicePattern = &v
}

// GetServiceNameRegex returns the ServiceNameRegex field value if set, zero value otherwise.
func (o *EndpointOption) GetServiceNameRegex() string {
	if o == nil || o.ServiceNameRegex == nil {
		var ret string
		return ret
	}
	return *o.ServiceNameRegex
}

// GetServiceNameRegexOk returns a tuple with the ServiceNameRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetServiceNameRegexOk() (*string, bool) {
	if o == nil || o.ServiceNameRegex == nil {
		return nil, false
	}
	return o.ServiceNameRegex, true
}

// HasServiceNameRegex returns a boolean if a field has been set.
func (o *EndpointOption) HasServiceNameRegex() bool {
	return o != nil && o.ServiceNameRegex != nil
}

// SetServiceNameRegex gets a reference to the given string and assigns it to the ServiceNameRegex field.
func (o *EndpointOption) SetServiceNameRegex(v string) {
	o.ServiceNameRegex = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *EndpointOption) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *EndpointOption) HasServiceName() bool {
	return o != nil && o.ServiceName != nil
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *EndpointOption) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *EndpointOption) GetSelector() map[string]string {
	if o == nil || o.Selector == nil {
		var ret map[string]string
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetSelectorOk() (*map[string]string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return &o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *EndpointOption) HasSelector() bool {
	return o != nil && o.Selector != nil
}

// SetSelector gets a reference to the given map[string]string and assigns it to the Selector field.
func (o *EndpointOption) SetSelector(v map[string]string) {
	o.Selector = v
}

// GetFollowNetworkMode returns the FollowNetworkMode field value if set, zero value otherwise.
func (o *EndpointOption) GetFollowNetworkMode() bool {
	if o == nil || o.FollowNetworkMode == nil {
		var ret bool
		return ret
	}
	return *o.FollowNetworkMode
}

// GetFollowNetworkModeOk returns a tuple with the FollowNetworkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOption) GetFollowNetworkModeOk() (*bool, bool) {
	if o == nil || o.FollowNetworkMode == nil {
		return nil, false
	}
	return o.FollowNetworkMode, true
}

// HasFollowNetworkMode returns a boolean if a field has been set.
func (o *EndpointOption) HasFollowNetworkMode() bool {
	return o != nil && o.FollowNetworkMode != nil
}

// SetFollowNetworkMode gets a reference to the given bool and assigns it to the FollowNetworkMode field.
func (o *EndpointOption) SetFollowNetworkMode(v bool) {
	o.FollowNetworkMode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EndpointOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["title"] = o.Title
	toSerialize["component"] = o.Component
	toSerialize["portName"] = o.PortName
	toSerialize["port"] = o.Port
	toSerialize["protocol"] = o.Protocol
	toSerialize["targetPort"] = o.TargetPort
	toSerialize["type"] = o.Type
	if o.SupportsSystemUse != nil {
		toSerialize["supportsSystemUse"] = o.SupportsSystemUse
	}
	if o.SupportsReadonly != nil {
		toSerialize["supportsReadonly"] = o.SupportsReadonly
	}
	if o.ServicePattern != nil {
		toSerialize["servicePattern"] = o.ServicePattern
	}
	if o.ServiceNameRegex != nil {
		toSerialize["serviceNameRegex"] = o.ServiceNameRegex
	}
	if o.ServiceName != nil {
		toSerialize["serviceName"] = o.ServiceName
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.FollowNetworkMode != nil {
		toSerialize["followNetworkMode"] = o.FollowNetworkMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EndpointOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title             *LocalizedDescription        `json:"title"`
		Component         *string                      `json:"component"`
		PortName          *string                      `json:"portName"`
		Port              *int32                       `json:"port"`
		Protocol          *string                      `json:"protocol"`
		TargetPort        *string                      `json:"targetPort"`
		Type              *[]string                    `json:"type"`
		SupportsSystemUse *bool                        `json:"supportsSystemUse,omitempty"`
		SupportsReadonly  *bool                        `json:"supportsReadonly,omitempty"`
		ServicePattern    *EngineOptionsServicePattern `json:"servicePattern,omitempty"`
		ServiceNameRegex  *string                      `json:"serviceNameRegex,omitempty"`
		ServiceName       *string                      `json:"serviceName,omitempty"`
		Selector          map[string]string            `json:"selector,omitempty"`
		FollowNetworkMode *bool                        `json:"followNetworkMode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.PortName == nil {
		return fmt.Errorf("required field portName missing")
	}
	if all.Port == nil {
		return fmt.Errorf("required field port missing")
	}
	if all.Protocol == nil {
		return fmt.Errorf("required field protocol missing")
	}
	if all.TargetPort == nil {
		return fmt.Errorf("required field targetPort missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"title", "component", "portName", "port", "protocol", "targetPort", "type", "supportsSystemUse", "supportsReadonly", "servicePattern", "serviceNameRegex", "serviceName", "selector", "followNetworkMode"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = *all.Title
	o.Component = *all.Component
	o.PortName = *all.PortName
	o.Port = *all.Port
	o.Protocol = *all.Protocol
	o.TargetPort = *all.TargetPort
	o.Type = *all.Type
	o.SupportsSystemUse = all.SupportsSystemUse
	o.SupportsReadonly = all.SupportsReadonly
	if all.ServicePattern != nil && !all.ServicePattern.IsValid() {
		hasInvalidField = true
	} else {
		o.ServicePattern = all.ServicePattern
	}
	o.ServiceNameRegex = all.ServiceNameRegex
	o.ServiceName = all.ServiceName
	o.Selector = all.Selector
	o.FollowNetworkMode = all.FollowNetworkMode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
