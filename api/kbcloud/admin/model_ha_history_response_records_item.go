// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HaHistoryResponseRecordsItem struct {
	StartAt    int32   `json:"StartAt"`
	EndAt      *int32  `json:"EndAt,omitempty"`
	OldPrimary *string `json:"OldPrimary,omitempty"`
	NewPrimary *string `json:"NewPrimary,omitempty"`
	Reason     *string `json:"Reason,omitempty"`
	UserId     *string `json:"UserID,omitempty"`
	UserName   *string `json:"UserName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHaHistoryResponseRecordsItem instantiates a new HaHistoryResponseRecordsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHaHistoryResponseRecordsItem(startAt int32) *HaHistoryResponseRecordsItem {
	this := HaHistoryResponseRecordsItem{}
	this.StartAt = startAt
	return &this
}

// NewHaHistoryResponseRecordsItemWithDefaults instantiates a new HaHistoryResponseRecordsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHaHistoryResponseRecordsItemWithDefaults() *HaHistoryResponseRecordsItem {
	this := HaHistoryResponseRecordsItem{}
	return &this
}

// GetStartAt returns the StartAt field value.
func (o *HaHistoryResponseRecordsItem) GetStartAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *HaHistoryResponseRecordsItem) GetStartAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value.
func (o *HaHistoryResponseRecordsItem) SetStartAt(v int32) {
	o.StartAt = v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *HaHistoryResponseRecordsItem) GetEndAt() int32 {
	if o == nil || o.EndAt == nil {
		var ret int32
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponseRecordsItem) GetEndAtOk() (*int32, bool) {
	if o == nil || o.EndAt == nil {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *HaHistoryResponseRecordsItem) HasEndAt() bool {
	return o != nil && o.EndAt != nil
}

// SetEndAt gets a reference to the given int32 and assigns it to the EndAt field.
func (o *HaHistoryResponseRecordsItem) SetEndAt(v int32) {
	o.EndAt = &v
}

// GetOldPrimary returns the OldPrimary field value if set, zero value otherwise.
func (o *HaHistoryResponseRecordsItem) GetOldPrimary() string {
	if o == nil || o.OldPrimary == nil {
		var ret string
		return ret
	}
	return *o.OldPrimary
}

// GetOldPrimaryOk returns a tuple with the OldPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponseRecordsItem) GetOldPrimaryOk() (*string, bool) {
	if o == nil || o.OldPrimary == nil {
		return nil, false
	}
	return o.OldPrimary, true
}

// HasOldPrimary returns a boolean if a field has been set.
func (o *HaHistoryResponseRecordsItem) HasOldPrimary() bool {
	return o != nil && o.OldPrimary != nil
}

// SetOldPrimary gets a reference to the given string and assigns it to the OldPrimary field.
func (o *HaHistoryResponseRecordsItem) SetOldPrimary(v string) {
	o.OldPrimary = &v
}

// GetNewPrimary returns the NewPrimary field value if set, zero value otherwise.
func (o *HaHistoryResponseRecordsItem) GetNewPrimary() string {
	if o == nil || o.NewPrimary == nil {
		var ret string
		return ret
	}
	return *o.NewPrimary
}

// GetNewPrimaryOk returns a tuple with the NewPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponseRecordsItem) GetNewPrimaryOk() (*string, bool) {
	if o == nil || o.NewPrimary == nil {
		return nil, false
	}
	return o.NewPrimary, true
}

// HasNewPrimary returns a boolean if a field has been set.
func (o *HaHistoryResponseRecordsItem) HasNewPrimary() bool {
	return o != nil && o.NewPrimary != nil
}

// SetNewPrimary gets a reference to the given string and assigns it to the NewPrimary field.
func (o *HaHistoryResponseRecordsItem) SetNewPrimary(v string) {
	o.NewPrimary = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *HaHistoryResponseRecordsItem) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponseRecordsItem) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *HaHistoryResponseRecordsItem) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *HaHistoryResponseRecordsItem) SetReason(v string) {
	o.Reason = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *HaHistoryResponseRecordsItem) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponseRecordsItem) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *HaHistoryResponseRecordsItem) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *HaHistoryResponseRecordsItem) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *HaHistoryResponseRecordsItem) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistoryResponseRecordsItem) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *HaHistoryResponseRecordsItem) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *HaHistoryResponseRecordsItem) SetUserName(v string) {
	o.UserName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HaHistoryResponseRecordsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["StartAt"] = o.StartAt
	if o.EndAt != nil {
		toSerialize["EndAt"] = o.EndAt
	}
	if o.OldPrimary != nil {
		toSerialize["OldPrimary"] = o.OldPrimary
	}
	if o.NewPrimary != nil {
		toSerialize["NewPrimary"] = o.NewPrimary
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.UserId != nil {
		toSerialize["UserID"] = o.UserId
	}
	if o.UserName != nil {
		toSerialize["UserName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HaHistoryResponseRecordsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartAt    *int32  `json:"StartAt"`
		EndAt      *int32  `json:"EndAt,omitempty"`
		OldPrimary *string `json:"OldPrimary,omitempty"`
		NewPrimary *string `json:"NewPrimary,omitempty"`
		Reason     *string `json:"Reason,omitempty"`
		UserId     *string `json:"UserID,omitempty"`
		UserName   *string `json:"UserName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.StartAt == nil {
		return fmt.Errorf("required field StartAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"StartAt", "EndAt", "OldPrimary", "NewPrimary", "Reason", "UserID", "UserName"})
	} else {
		return err
	}
	o.StartAt = *all.StartAt
	o.EndAt = all.EndAt
	o.OldPrimary = all.OldPrimary
	o.NewPrimary = all.NewPrimary
	o.Reason = all.Reason
	o.UserId = all.UserId
	o.UserName = all.UserName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
