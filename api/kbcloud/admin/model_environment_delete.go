// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentDelete Environment deletion option
type EnvironmentDelete struct {
	// clean up resources in the cloud (only valid if creation is done by role ARN)
	CleanCloudResources *bool `json:"cleanCloudResources,omitempty"`
	// The policy to clean cloud resources, either `Delete` or `Retain`
	Backup CleanCloudResourcePolicy `json:"backup"`
	// The policy to clean cloud resources, either `Delete` or `Retain`
	Minio CleanCloudResourcePolicy `json:"minio"`
	// The policy to clean cloud resources, either `Delete` or `Retain`
	VictoriaMetrics CleanCloudResourcePolicy `json:"victoriaMetrics"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentDelete instantiates a new EnvironmentDelete object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentDelete(backup CleanCloudResourcePolicy, minio CleanCloudResourcePolicy, victoriaMetrics CleanCloudResourcePolicy) *EnvironmentDelete {
	this := EnvironmentDelete{}
	this.Backup = backup
	this.Minio = minio
	this.VictoriaMetrics = victoriaMetrics
	return &this
}

// NewEnvironmentDeleteWithDefaults instantiates a new EnvironmentDelete object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentDeleteWithDefaults() *EnvironmentDelete {
	this := EnvironmentDelete{}
	var backup CleanCloudResourcePolicy = CLEANCLOUDRESOURCEPOLICY_RETAIN
	this.Backup = backup
	var minio CleanCloudResourcePolicy = CLEANCLOUDRESOURCEPOLICY_RETAIN
	this.Minio = minio
	var victoriaMetrics CleanCloudResourcePolicy = CLEANCLOUDRESOURCEPOLICY_RETAIN
	this.VictoriaMetrics = victoriaMetrics
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

// GetBackup returns the Backup field value.
func (o *EnvironmentDelete) GetBackup() CleanCloudResourcePolicy {
	if o == nil {
		var ret CleanCloudResourcePolicy
		return ret
	}
	return o.Backup
}

// GetBackupOk returns a tuple with the Backup field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetBackupOk() (*CleanCloudResourcePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Backup, true
}

// SetBackup sets field value.
func (o *EnvironmentDelete) SetBackup(v CleanCloudResourcePolicy) {
	o.Backup = v
}

// GetMinio returns the Minio field value.
func (o *EnvironmentDelete) GetMinio() CleanCloudResourcePolicy {
	if o == nil {
		var ret CleanCloudResourcePolicy
		return ret
	}
	return o.Minio
}

// GetMinioOk returns a tuple with the Minio field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetMinioOk() (*CleanCloudResourcePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minio, true
}

// SetMinio sets field value.
func (o *EnvironmentDelete) SetMinio(v CleanCloudResourcePolicy) {
	o.Minio = v
}

// GetVictoriaMetrics returns the VictoriaMetrics field value.
func (o *EnvironmentDelete) GetVictoriaMetrics() CleanCloudResourcePolicy {
	if o == nil {
		var ret CleanCloudResourcePolicy
		return ret
	}
	return o.VictoriaMetrics
}

// GetVictoriaMetricsOk returns a tuple with the VictoriaMetrics field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDelete) GetVictoriaMetricsOk() (*CleanCloudResourcePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VictoriaMetrics, true
}

// SetVictoriaMetrics sets field value.
func (o *EnvironmentDelete) SetVictoriaMetrics(v CleanCloudResourcePolicy) {
	o.VictoriaMetrics = v
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
	toSerialize["backup"] = o.Backup
	toSerialize["minio"] = o.Minio
	toSerialize["victoriaMetrics"] = o.VictoriaMetrics

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentDelete) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CleanCloudResources *bool                     `json:"cleanCloudResources,omitempty"`
		Backup              *CleanCloudResourcePolicy `json:"backup"`
		Minio               *CleanCloudResourcePolicy `json:"minio"`
		VictoriaMetrics     *CleanCloudResourcePolicy `json:"victoriaMetrics"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Backup == nil {
		return fmt.Errorf("required field backup missing")
	}
	if all.Minio == nil {
		return fmt.Errorf("required field minio missing")
	}
	if all.VictoriaMetrics == nil {
		return fmt.Errorf("required field victoriaMetrics missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cleanCloudResources", "backup", "minio", "victoriaMetrics"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CleanCloudResources = all.CleanCloudResources
	if !all.Backup.IsValid() {
		hasInvalidField = true
	} else {
		o.Backup = *all.Backup
	}
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
