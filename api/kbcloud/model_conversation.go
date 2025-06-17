// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/google/uuid"
)

type Conversation struct {
	// Conversation ID
	Id *uuid.UUID `json:"id,omitempty"`
	// User ID associated with the conversation
	UserId *uuid.UUID `json:"userId,omitempty"`
	// Name of the cluster associated with the conversation
	ClusterName *string `json:"clusterName,omitempty"`
	// Name of the organization associated with the conversation
	OrgName *string `json:"orgName,omitempty"`
	// Optional title for the conversation
	Title *string `json:"title,omitempty"`
	// Timestamp of when the conversation was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp of when the conversation was last updated
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConversation instantiates a new Conversation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConversation() *Conversation {
	this := Conversation{}
	return &this
}

// NewConversationWithDefaults instantiates a new Conversation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConversationWithDefaults() *Conversation {
	this := Conversation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Conversation) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Conversation) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *Conversation) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Conversation) GetUserId() uuid.UUID {
	if o == nil || o.UserId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Conversation) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given uuid.UUID and assigns it to the UserId field.
func (o *Conversation) SetUserId(v uuid.UUID) {
	o.UserId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *Conversation) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *Conversation) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *Conversation) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Conversation) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Conversation) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Conversation) SetOrgName(v string) {
	o.OrgName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Conversation) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Conversation) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Conversation) SetTitle(v string) {
	o.Title = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Conversation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Conversation) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Conversation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Conversation) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conversation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Conversation) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Conversation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Conversation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Conversation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *uuid.UUID `json:"id,omitempty"`
		UserId      *uuid.UUID `json:"userId,omitempty"`
		ClusterName *string    `json:"clusterName,omitempty"`
		OrgName     *string    `json:"orgName,omitempty"`
		Title       *string    `json:"title,omitempty"`
		CreatedAt   *time.Time `json:"createdAt,omitempty"`
		UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "userId", "clusterName", "orgName", "title", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Id = all.Id
	o.UserId = all.UserId
	o.ClusterName = all.ClusterName
	o.OrgName = all.OrgName
	o.Title = all.Title
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
