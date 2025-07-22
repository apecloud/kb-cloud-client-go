// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ServiceDescriptor serviceDescriptor that will be used in serviceRef. The field definition is in line with kubeblocks.
type ServiceDescriptor struct {
	Host     *string `json:"host,omitempty"`
	Port     *string `json:"port,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDescriptor instantiates a new ServiceDescriptor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDescriptor() *ServiceDescriptor {
	this := ServiceDescriptor{}
	return &this
}

// NewServiceDescriptorWithDefaults instantiates a new ServiceDescriptor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDescriptorWithDefaults() *ServiceDescriptor {
	this := ServiceDescriptor{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ServiceDescriptor) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDescriptor) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ServiceDescriptor) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ServiceDescriptor) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ServiceDescriptor) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDescriptor) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ServiceDescriptor) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *ServiceDescriptor) SetPort(v string) {
	o.Port = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ServiceDescriptor) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDescriptor) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ServiceDescriptor) HasEndpoint() bool {
	return o != nil && o.Endpoint != nil
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ServiceDescriptor) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ServiceDescriptor) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDescriptor) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ServiceDescriptor) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ServiceDescriptor) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ServiceDescriptor) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDescriptor) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ServiceDescriptor) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ServiceDescriptor) SetPassword(v string) {
	o.Password = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Host     *string `json:"host,omitempty"`
		Port     *string `json:"port,omitempty"`
		Endpoint *string `json:"endpoint,omitempty"`
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"host", "port", "endpoint", "username", "password"})
	} else {
		return err
	}
	o.Host = all.Host
	o.Port = all.Port
	o.Endpoint = all.Endpoint
	o.Username = all.Username
	o.Password = all.Password

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
