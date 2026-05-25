// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentConversation struct {
	ConversationId string    `json:"conversationId"`
	OrgName        string    `json:"orgName"`
	ClusterName    *string   `json:"clusterName,omitempty"`
	Title          *string   `json:"title,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentConversation instantiates a new HermesAgentConversation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentConversation(conversationId string, orgName string, createdAt time.Time) *HermesAgentConversation {
	this := HermesAgentConversation{}
	this.ConversationId = conversationId
	this.OrgName = orgName
	this.CreatedAt = createdAt
	return &this
}

// NewHermesAgentConversationWithDefaults instantiates a new HermesAgentConversation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentConversationWithDefaults() *HermesAgentConversation {
	this := HermesAgentConversation{}
	return &this
}

// GetConversationId returns the ConversationId field value.
func (o *HermesAgentConversation) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentConversation) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *HermesAgentConversation) SetConversationId(v string) {
	o.ConversationId = v
}

// GetOrgName returns the OrgName field value.
func (o *HermesAgentConversation) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *HermesAgentConversation) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *HermesAgentConversation) SetOrgName(v string) {
	o.OrgName = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HermesAgentConversation) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentConversation) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HermesAgentConversation) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HermesAgentConversation) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *HermesAgentConversation) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentConversation) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *HermesAgentConversation) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *HermesAgentConversation) SetTitle(v string) {
	o.Title = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *HermesAgentConversation) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *HermesAgentConversation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *HermesAgentConversation) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentConversation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["orgName"] = o.OrgName
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentConversation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversationId *string    `json:"conversationId"`
		OrgName        *string    `json:"orgName"`
		ClusterName    *string    `json:"clusterName,omitempty"`
		Title          *string    `json:"title,omitempty"`
		CreatedAt      *time.Time `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conversationId", "orgName", "clusterName", "title", "createdAt"})
	} else {
		return err
	}
	o.ConversationId = *all.ConversationId
	o.OrgName = *all.OrgName
	o.ClusterName = all.ClusterName
	o.Title = all.Title
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
