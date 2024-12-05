// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupRepoStatsList backupRepoStatsList is a list of backup repo stats
type BackupRepoStatsList struct {
	// Items is the list of backup repo stats in the list
	Items   []BackupRepoStats `json:"items"`
	Current *BackupRepoStats  `json:"current,omitempty"`
	// PageResult info
	PageResult *PageResult `json:"pageResult,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepoStatsList instantiates a new BackupRepoStatsList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepoStatsList(items []BackupRepoStats) *BackupRepoStatsList {
	this := BackupRepoStatsList{}
	this.Items = items
	return &this
}

// NewBackupRepoStatsListWithDefaults instantiates a new BackupRepoStatsList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoStatsListWithDefaults() *BackupRepoStatsList {
	this := BackupRepoStatsList{}
	return &this
}

// GetItems returns the Items field value.
func (o *BackupRepoStatsList) GetItems() []BackupRepoStats {
	if o == nil {
		var ret []BackupRepoStats
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *BackupRepoStatsList) GetItemsOk() (*[]BackupRepoStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *BackupRepoStatsList) SetItems(v []BackupRepoStats) {
	o.Items = v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *BackupRepoStatsList) GetCurrent() BackupRepoStats {
	if o == nil || o.Current == nil {
		var ret BackupRepoStats
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStatsList) GetCurrentOk() (*BackupRepoStats, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *BackupRepoStatsList) HasCurrent() bool {
	return o != nil && o.Current != nil
}

// SetCurrent gets a reference to the given BackupRepoStats and assigns it to the Current field.
func (o *BackupRepoStatsList) SetCurrent(v BackupRepoStats) {
	o.Current = &v
}

// GetPageResult returns the PageResult field value if set, zero value otherwise.
func (o *BackupRepoStatsList) GetPageResult() PageResult {
	if o == nil || o.PageResult == nil {
		var ret PageResult
		return ret
	}
	return *o.PageResult
}

// GetPageResultOk returns a tuple with the PageResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStatsList) GetPageResultOk() (*PageResult, bool) {
	if o == nil || o.PageResult == nil {
		return nil, false
	}
	return o.PageResult, true
}

// HasPageResult returns a boolean if a field has been set.
func (o *BackupRepoStatsList) HasPageResult() bool {
	return o != nil && o.PageResult != nil
}

// SetPageResult gets a reference to the given PageResult and assigns it to the PageResult field.
func (o *BackupRepoStatsList) SetPageResult(v PageResult) {
	o.PageResult = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepoStatsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.PageResult != nil {
		toSerialize["pageResult"] = o.PageResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepoStatsList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]BackupRepoStats `json:"items"`
		Current    *BackupRepoStats   `json:"current,omitempty"`
		PageResult *PageResult        `json:"pageResult,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "current", "pageResult"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = *all.Items
	if all.Current != nil && all.Current.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Current = all.Current
	if all.PageResult != nil && all.PageResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PageResult = all.PageResult

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
