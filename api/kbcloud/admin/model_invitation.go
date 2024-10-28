// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Invitation Invitation info

type Invitation struct {
	// User has accepted or not
	Accepted bool `json:"accepted"`
	// the created time of the invitation
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The email of the invitee
	Email string `json:"email"`
	// Expire time of this invitation
	ExpireTime time.Time `json:"expireTime"`
	// The ID for the invitation
	Id string `json:"id"`
	// Last email send time
	LastEmailTime time.Time `json:"lastEmailTime"`
	// The name of the organization, for output only
	OrgName string `json:"orgName"`
	// The name of the role
	RoleName string `json:"roleName"`
	// InvitationSender is the User who send the Invitation
	Sender *InvitationSender `json:"sender,omitempty"`
	// [internal, not for caller display] token of invitation, base64 encoded string including invitation and org ID
	Token string `json:"token"`
	// the updated time of the invitation
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInvitation instantiates a new Invitation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInvitation(accepted bool, email string, expireTime time.Time, id string, lastEmailTime time.Time, orgName string, roleName string, token string) *Invitation {
	this := Invitation{}
	this.Accepted = accepted
	this.Email = email
	this.ExpireTime = expireTime
	this.Id = id
	this.LastEmailTime = lastEmailTime
	this.OrgName = orgName
	this.RoleName = roleName
	this.Token = token
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetAccepted returns the Accepted field value.
func (o *Invitation) GetAccepted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetAcceptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accepted, true
}

// SetAccepted sets field value.
func (o *Invitation) SetAccepted(v bool) {
	o.Accepted = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Invitation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Invitation) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Invitation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value.
func (o *Invitation) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *Invitation) SetEmail(v string) {
	o.Email = v
}

// GetExpireTime returns the ExpireTime field value.
func (o *Invitation) GetExpireTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetExpireTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpireTime, true
}

// SetExpireTime sets field value.
func (o *Invitation) SetExpireTime(v time.Time) {
	o.ExpireTime = v
}

// GetId returns the Id field value.
func (o *Invitation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Invitation) SetId(v string) {
	o.Id = v
}

// GetLastEmailTime returns the LastEmailTime field value.
func (o *Invitation) GetLastEmailTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastEmailTime
}

// GetLastEmailTimeOk returns a tuple with the LastEmailTime field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetLastEmailTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEmailTime, true
}

// SetLastEmailTime sets field value.
func (o *Invitation) SetLastEmailTime(v time.Time) {
	o.LastEmailTime = v
}

// GetOrgName returns the OrgName field value.
func (o *Invitation) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *Invitation) SetOrgName(v string) {
	o.OrgName = v
}

// GetRoleName returns the RoleName field value.
func (o *Invitation) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value.
func (o *Invitation) SetRoleName(v string) {
	o.RoleName = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *Invitation) GetSender() InvitationSender {
	if o == nil || o.Sender == nil {
		var ret InvitationSender
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetSenderOk() (*InvitationSender, bool) {
	if o == nil || o.Sender == nil {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *Invitation) HasSender() bool {
	return o != nil && o.Sender != nil
}

// SetSender gets a reference to the given InvitationSender and assigns it to the Sender field.
func (o *Invitation) SetSender(v InvitationSender) {
	o.Sender = &v
}

// GetToken returns the Token field value.
func (o *Invitation) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *Invitation) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value.
func (o *Invitation) SetToken(v string) {
	o.Token = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invitation) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invitation) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Invitation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["accepted"] = o.Accepted
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["email"] = o.Email
	if o.ExpireTime.Nanosecond() == 0 {
		toSerialize["expireTime"] = o.ExpireTime.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expireTime"] = o.ExpireTime.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["id"] = o.Id
	if o.LastEmailTime.Nanosecond() == 0 {
		toSerialize["lastEmailTime"] = o.LastEmailTime.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["lastEmailTime"] = o.LastEmailTime.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["orgName"] = o.OrgName
	toSerialize["roleName"] = o.RoleName
	if o.Sender != nil {
		toSerialize["sender"] = o.Sender
	}
	toSerialize["token"] = o.Token
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
func (o *Invitation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accepted      *bool             `json:"accepted"`
		CreatedAt     *time.Time        `json:"createdAt,omitempty"`
		Email         *string           `json:"email"`
		ExpireTime    *time.Time        `json:"expireTime"`
		Id            *string           `json:"id"`
		LastEmailTime *time.Time        `json:"lastEmailTime"`
		OrgName       *string           `json:"orgName"`
		RoleName      *string           `json:"roleName"`
		Sender        *InvitationSender `json:"sender,omitempty"`
		Token         *string           `json:"token"`
		UpdatedAt     *time.Time        `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Accepted == nil {
		return fmt.Errorf("required field accepted missing")
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.ExpireTime == nil {
		return fmt.Errorf("required field expireTime missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.LastEmailTime == nil {
		return fmt.Errorf("required field lastEmailTime missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.RoleName == nil {
		return fmt.Errorf("required field roleName missing")
	}
	if all.Token == nil {
		return fmt.Errorf("required field token missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"accepted", "createdAt", "email", "expireTime", "id", "lastEmailTime", "orgName", "roleName", "sender", "token", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Accepted = *all.Accepted
	o.CreatedAt = all.CreatedAt
	o.Email = *all.Email
	o.ExpireTime = *all.ExpireTime
	o.Id = *all.Id
	o.LastEmailTime = *all.LastEmailTime
	o.OrgName = *all.OrgName
	o.RoleName = *all.RoleName
	if all.Sender != nil && all.Sender.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sender = all.Sender
	o.Token = *all.Token
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
