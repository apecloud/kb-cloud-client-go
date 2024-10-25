// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION Kubeconfig
type Kubeconfig struct {
	// The base64 encoded kubeconfig file contents to connect to this Kubernetes
	Kubeconfig string `json:"kubeconfig"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKubeconfig instantiates a new Kubeconfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKubeconfig(kubeconfig string) *Kubeconfig {
	this := Kubeconfig{}
	this.Kubeconfig = kubeconfig
	return &this
}

// NewKubeconfigWithDefaults instantiates a new Kubeconfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKubeconfigWithDefaults() *Kubeconfig {
	this := Kubeconfig{}
	return &this
}

// GetKubeconfig returns the Kubeconfig field value.
func (o *Kubeconfig) GetKubeconfig() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Kubeconfig
}

// GetKubeconfigOk returns a tuple with the Kubeconfig field value
// and a boolean to check if the value has been set.
func (o *Kubeconfig) GetKubeconfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kubeconfig, true
}

// SetKubeconfig sets field value.
func (o *Kubeconfig) SetKubeconfig(v string) {
	o.Kubeconfig = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Kubeconfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["kubeconfig"] = o.Kubeconfig

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Kubeconfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kubeconfig *string `json:"kubeconfig"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Kubeconfig == nil {
		return fmt.Errorf("required field kubeconfig missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kubeconfig"})
	} else {
		return err
	}
	o.Kubeconfig = *all.Kubeconfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
