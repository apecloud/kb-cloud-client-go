// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrgMember Org Member info

type OrgMember struct {
	// The display name of User. Read-Only
	DisplayName *string `json:"displayName,omitempty"`
	// The email of User. Required when create. Read-Only after create
	Email string `json:"email"`
	// The role of the User in the Org. Required
	Role OrgMemberRole `json:"role"`
	// The ID of User. Read-Only
	UserId string `json:"userId"`
	// Return true if the member is freezed in the organization
	Freezed bool `json:"freezed"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgMember instantiates a new OrgMember object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgMember(email string, role OrgMemberRole, userId string, freezed bool) *OrgMember {
	this := OrgMember{}
	this.Email = email
	this.Role = role
	this.UserId = userId
	this.Freezed = freezed
	return &this
}

// NewOrgMemberWithDefaults instantiates a new OrgMember object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgMemberWithDefaults() *OrgMember {
	this := OrgMember{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OrgMember) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OrgMember) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OrgMember) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value.
func (o *OrgMember) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OrgMember) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *OrgMember) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value.
func (o *OrgMember) GetRole() OrgMemberRole {
	if o == nil {
		var ret OrgMemberRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *OrgMember) GetRoleOk() (*OrgMemberRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *OrgMember) SetRole(v OrgMemberRole) {
	o.Role = v
}

// GetUserId returns the UserId field value.
func (o *OrgMember) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *OrgMember) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *OrgMember) SetUserId(v string) {
	o.UserId = v
}

// GetFreezed returns the Freezed field value.
func (o *OrgMember) GetFreezed() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Freezed
}

// GetFreezedOk returns a tuple with the Freezed field value
// and a boolean to check if the value has been set.
func (o *OrgMember) GetFreezedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Freezed, true
}

// SetFreezed sets field value.
func (o *OrgMember) SetFreezed(v bool) {
	o.Freezed = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["email"] = o.Email
	toSerialize["role"] = o.Role
	toSerialize["userId"] = o.UserId
	toSerialize["freezed"] = o.Freezed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgMember) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string        `json:"displayName,omitempty"`
		Email       *string        `json:"email"`
		Role        *OrgMemberRole `json:"role"`
		UserId      *string        `json:"userId"`
		Freezed     *bool          `json:"freezed"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field userId missing")
	}
	if all.Freezed == nil {
		return fmt.Errorf("required field freezed missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"displayName", "email", "role", "userId", "freezed"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Email = *all.Email
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}
	o.UserId = *all.UserId
	o.Freezed = *all.Freezed

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
