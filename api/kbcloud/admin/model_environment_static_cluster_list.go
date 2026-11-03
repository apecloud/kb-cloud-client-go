// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentStaticClusterList struct {
	// The base64 encoded kubeconfig file contents to connect to this Kubernetes cluster.
	Kubeconfig string `json:"kubeconfig"`
	// Engine whose retained clusters should be listed. Candidates are discovered from the kubeconfig and existing cluster records. The caller must select one; the server does not choose a default.
	Engine StaticServiceClusterEngine `json:"engine"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStaticClusterList instantiates a new EnvironmentStaticClusterList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStaticClusterList(kubeconfig string, engine StaticServiceClusterEngine) *EnvironmentStaticClusterList {
	this := EnvironmentStaticClusterList{}
	this.Kubeconfig = kubeconfig
	this.Engine = engine
	return &this
}

// NewEnvironmentStaticClusterListWithDefaults instantiates a new EnvironmentStaticClusterList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStaticClusterListWithDefaults() *EnvironmentStaticClusterList {
	this := EnvironmentStaticClusterList{}
	return &this
}

// GetKubeconfig returns the Kubeconfig field value.
func (o *EnvironmentStaticClusterList) GetKubeconfig() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Kubeconfig
}

// GetKubeconfigOk returns a tuple with the Kubeconfig field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStaticClusterList) GetKubeconfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kubeconfig, true
}

// SetKubeconfig sets field value.
func (o *EnvironmentStaticClusterList) SetKubeconfig(v string) {
	o.Kubeconfig = v
}

// GetEngine returns the Engine field value.
func (o *EnvironmentStaticClusterList) GetEngine() StaticServiceClusterEngine {
	if o == nil {
		var ret StaticServiceClusterEngine
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStaticClusterList) GetEngineOk() (*StaticServiceClusterEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *EnvironmentStaticClusterList) SetEngine(v StaticServiceClusterEngine) {
	o.Engine = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStaticClusterList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["kubeconfig"] = o.Kubeconfig
	toSerialize["engine"] = o.Engine

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStaticClusterList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kubeconfig *string                     `json:"kubeconfig"`
		Engine     *StaticServiceClusterEngine `json:"engine"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Kubeconfig == nil {
		return fmt.Errorf("required field kubeconfig missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kubeconfig", "engine"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Kubeconfig = *all.Kubeconfig
	if !all.Engine.IsValid() {
		hasInvalidField = true
	} else {
		o.Engine = *all.Engine
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
