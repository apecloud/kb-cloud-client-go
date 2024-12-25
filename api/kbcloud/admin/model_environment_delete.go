// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// EnvironmentDelete Environment deletion option
type EnvironmentDelete struct {
	// clean up resources in the cloud (only valid if creation is done by role ARN)
	CleanCloudResources *bool `json:"cleanCloudResources,omitempty"`
	// The policy to clean cloud resources, either `Delete` or `Retain`
	Minio CloudResourceCleanPolicy `json:"minio"`
	// The policy to clean cloud resources, either `Delete` or `Retain`
	VictoriaMetrics CloudResourceCleanPolicy `json:"victoriaMetrics"`
	// The policy to clean cloud resources, either `Delete` or `Retain`
	Loki CloudResourceCleanPolicy `json:"loki"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentDelete instantiates a new EnvironmentDelete object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentDelete(minio CloudResourceCleanPolicy, victoriaMetrics CloudResourceCleanPolicy, loki CloudResourceCleanPolicy) *EnvironmentDelete {
	this := EnvironmentDelete{}
	this.Minio = minio
	this.VictoriaMetrics = victoriaMetrics
	this.Loki = loki
	return &this
}

// NewEnvironmentDeleteWithDefaults instantiates a new EnvironmentDelete object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentDeleteWithDefaults() *EnvironmentDelete {
	this := EnvironmentDelete{}
	var minio CloudResourceCleanPolicy = CloudResourceCleanPolicyRetain
	this.Minio = minio
	var victoriaMetrics CloudResourceCleanPolicy = CloudResourceCleanPolicyRetain
	this.VictoriaMetrics = victoriaMetrics
	var loki CloudResourceCleanPolicy = CloudResourceCleanPolicyRetain
	this.Loki = loki
	return &this
}

// GetCleanCloudResources returns the CleanCloudResources field value if set, zero value otherwise.
func (o *EnvironmentDelete) GetCleanCloudResources() bool {
	if o == nil || o.CleanCloudResources == nil {
		var ret bool
		return ret
	}
	return *o.CleanCloudResources
}

// GetCleanCloudResourcesOk returns a tuple with the CleanCloudResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetCleanCloudResourcesOk() (*bool, bool) {
	if o == nil || o.CleanCloudResources == nil {
		return nil, false
	}
	return o.CleanCloudResources, true
}

// HasCleanCloudResources returns a boolean if a field has been set.
func (o *EnvironmentDelete) HasCleanCloudResources() bool {
	return o != nil && o.CleanCloudResources != nil
}

// SetCleanCloudResources gets a reference to the given bool and assigns it to the CleanCloudResources field.
func (o *EnvironmentDelete) SetCleanCloudResources(v bool) {
	o.CleanCloudResources = &v
}

// GetMinio returns the Minio field value.
func (o *EnvironmentDelete) GetMinio() CloudResourceCleanPolicy {
	if o == nil {
		var ret CloudResourceCleanPolicy
		return ret
	}
	return o.Minio
}

// GetMinioOk returns a tuple with the Minio field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetMinioOk() (*CloudResourceCleanPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minio, true
}

// SetMinio sets field value.
func (o *EnvironmentDelete) SetMinio(v CloudResourceCleanPolicy) {
	o.Minio = v
}

// GetVictoriaMetrics returns the VictoriaMetrics field value.
func (o *EnvironmentDelete) GetVictoriaMetrics() CloudResourceCleanPolicy {
	if o == nil {
		var ret CloudResourceCleanPolicy
		return ret
	}
	return o.VictoriaMetrics
}

// GetVictoriaMetricsOk returns a tuple with the VictoriaMetrics field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetVictoriaMetricsOk() (*CloudResourceCleanPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VictoriaMetrics, true
}

// SetVictoriaMetrics sets field value.
func (o *EnvironmentDelete) SetVictoriaMetrics(v CloudResourceCleanPolicy) {
	o.VictoriaMetrics = v
}

// GetLoki returns the Loki field value.
func (o *EnvironmentDelete) GetLoki() CloudResourceCleanPolicy {
	if o == nil {
		var ret CloudResourceCleanPolicy
		return ret
	}
	return o.Loki
}

// GetLokiOk returns a tuple with the Loki field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetLokiOk() (*CloudResourceCleanPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loki, true
}

// SetLoki sets field value.
func (o *EnvironmentDelete) SetLoki(v CloudResourceCleanPolicy) {
	o.Loki = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CleanCloudResources != nil {
		toSerialize["cleanCloudResources"] = o.CleanCloudResources
	}
	toSerialize["minio"] = o.Minio
	toSerialize["victoriaMetrics"] = o.VictoriaMetrics
	toSerialize["loki"] = o.Loki

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentDelete) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CleanCloudResources *bool                     `json:"cleanCloudResources,omitempty"`
		Minio               *CloudResourceCleanPolicy `json:"minio"`
		VictoriaMetrics     *CloudResourceCleanPolicy `json:"victoriaMetrics"`
		Loki                *CloudResourceCleanPolicy `json:"loki"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Minio == nil {
		return fmt.Errorf("required field minio missing")
	}
	if all.VictoriaMetrics == nil {
		return fmt.Errorf("required field victoriaMetrics missing")
	}
	if all.Loki == nil {
		return fmt.Errorf("required field loki missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cleanCloudResources", "minio", "victoriaMetrics", "loki"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CleanCloudResources = all.CleanCloudResources
	if !all.Minio.IsValid() {
		hasInvalidField = true
	} else {
		o.Minio = *all.Minio
	}
	if !all.VictoriaMetrics.IsValid() {
		hasInvalidField = true
	} else {
		o.VictoriaMetrics = *all.VictoriaMetrics
	}
	if !all.Loki.IsValid() {
		hasInvalidField = true
	} else {
		o.Loki = *all.Loki
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
