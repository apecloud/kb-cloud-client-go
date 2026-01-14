// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Dns struct {
	// the name of the DNS service
	ServiceName string `json:"serviceName"`
	// the namespace of the DNS service
	ServiceNamespace string `json:"serviceNamespace"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDns instantiates a new Dns object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDns(serviceName string, serviceNamespace string) *Dns {
	this := Dns{}
	this.ServiceName = serviceName
	this.ServiceNamespace = serviceNamespace
	return &this
}

// NewDnsWithDefaults instantiates a new Dns object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDnsWithDefaults() *Dns {
	this := Dns{}
	return &this
}

// GetServiceName returns the ServiceName field value.
func (o *Dns) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *Dns) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *Dns) SetServiceName(v string) {
	o.ServiceName = v
}

// GetServiceNamespace returns the ServiceNamespace field value.
func (o *Dns) GetServiceNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceNamespace
}

// GetServiceNamespaceOk returns a tuple with the ServiceNamespace field value
// and a boolean to check if the value has been set.
func (o *Dns) GetServiceNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceNamespace, true
}

// SetServiceNamespace sets field value.
func (o *Dns) SetServiceNamespace(v string) {
	o.ServiceNamespace = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Dns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["serviceNamespace"] = o.ServiceNamespace

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Dns) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServiceName      *string `json:"serviceName"`
		ServiceNamespace *string `json:"serviceNamespace"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field serviceName missing")
	}
	if all.ServiceNamespace == nil {
		return fmt.Errorf("required field serviceNamespace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"serviceName", "serviceNamespace"})
	} else {
		return err
	}
	o.ServiceName = *all.ServiceName
	o.ServiceNamespace = *all.ServiceNamespace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
