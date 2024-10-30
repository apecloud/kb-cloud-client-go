// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type OpsRebuildInstanceRequestsItem struct {
	BackupName *string `json:"backupName,omitempty"`
	Instances []OpsRebuildInstanceInstanceParam `json:"instances"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewOpsRebuildInstanceRequestsItem instantiates a new OpsRebuildInstanceRequestsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsRebuildInstanceRequestsItem(instances []OpsRebuildInstanceInstanceParam) *OpsRebuildInstanceRequestsItem {
	this := OpsRebuildInstanceRequestsItem{}
	this.Instances = instances
	return &this
}

// NewOpsRebuildInstanceRequestsItemWithDefaults instantiates a new OpsRebuildInstanceRequestsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsRebuildInstanceRequestsItemWithDefaults() *OpsRebuildInstanceRequestsItem {
	this := OpsRebuildInstanceRequestsItem{}
	return &this
}
// GetBackupName returns the BackupName field value if set, zero value otherwise.
func (o *OpsRebuildInstanceRequestsItem) GetBackupName() string {
	if o == nil || o.BackupName == nil {
		var ret string
		return ret
	}
	return *o.BackupName
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsRebuildInstanceRequestsItem) GetBackupNameOk() (*string, bool) {
	if o == nil || o.BackupName == nil {
		return nil, false
	}
	return o.BackupName, true
}

// HasBackupName returns a boolean if a field has been set.
func (o *OpsRebuildInstanceRequestsItem) HasBackupName() bool {
	return o != nil && o.BackupName != nil
}

// SetBackupName gets a reference to the given string and assigns it to the BackupName field.
func (o *OpsRebuildInstanceRequestsItem) SetBackupName(v string) {
	o.BackupName = &v
}


// GetInstances returns the Instances field value.
func (o *OpsRebuildInstanceRequestsItem) GetInstances() []OpsRebuildInstanceInstanceParam {
	if o == nil {
		var ret []OpsRebuildInstanceInstanceParam
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *OpsRebuildInstanceRequestsItem) GetInstancesOk() (*[]OpsRebuildInstanceInstanceParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instances, true
}

// SetInstances sets field value.
func (o *OpsRebuildInstanceRequestsItem) SetInstances(v []OpsRebuildInstanceInstanceParam) {
	o.Instances = v
}



// MarshalJSON serializes the struct using spec logic.
func (o OpsRebuildInstanceRequestsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.BackupName != nil {
		toSerialize["backupName"] = o.BackupName
	}
	toSerialize["instances"] = o.Instances

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsRebuildInstanceRequestsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackupName *string `json:"backupName,omitempty"`
		Instances *[]OpsRebuildInstanceInstanceParam `json:"instances"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Instances == nil {
		return fmt.Errorf("required field instances missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "backupName", "instances",  })
	} else {
		return err
	}
	o.BackupName = all.BackupName
	o.Instances = *all.Instances

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
