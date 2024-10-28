// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OrganizationItem Organization items.

type OrganizationItem struct {
	// The name of the organization.
	Name string `json:"name"`
	// The display name of the organization.
	DisplayName string `json:"displayName"`
	// The display name of the creator of the organization.
	Creator string `json:"creator"`
	// The ID of the creator of the organization
	CreatorId string `json:"creatorID"`
	// The total number of clusters in the organization.
	ClusterTotal int32 `json:"clusterTotal"`
	// The time the organization was created.
	CreatedAt time.Time `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrganizationItem instantiates a new OrganizationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrganizationItem(name string, displayName string, creator string, creatorId string, clusterTotal int32, createdAt time.Time) *OrganizationItem {
	this := OrganizationItem{}
	this.Name = name
	this.DisplayName = displayName
	this.Creator = creator
	this.CreatorId = creatorId
	this.ClusterTotal = clusterTotal
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationItemWithDefaults instantiates a new OrganizationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrganizationItemWithDefaults() *OrganizationItem {
	this := OrganizationItem{}
	return &this
}

// GetName returns the Name field value.
func (o *OrganizationItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OrganizationItem) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value.
func (o *OrganizationItem) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *OrganizationItem) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *OrganizationItem) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetCreator returns the Creator field value.
func (o *OrganizationItem) GetCreator() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *OrganizationItem) GetCreatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value.
func (o *OrganizationItem) SetCreator(v string) {
	o.Creator = v
}

// GetCreatorId returns the CreatorId field value.
func (o *OrganizationItem) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *OrganizationItem) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value.
func (o *OrganizationItem) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetClusterTotal returns the ClusterTotal field value.
func (o *OrganizationItem) GetClusterTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ClusterTotal
}

// GetClusterTotalOk returns a tuple with the ClusterTotal field value
// and a boolean to check if the value has been set.
func (o *OrganizationItem) GetClusterTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterTotal, true
}

// SetClusterTotal sets field value.
func (o *OrganizationItem) SetClusterTotal(v int32) {
	o.ClusterTotal = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OrganizationItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OrganizationItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrganizationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["creator"] = o.Creator
	toSerialize["creatorID"] = o.CreatorId
	toSerialize["clusterTotal"] = o.ClusterTotal
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
func (o *OrganizationItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string    `json:"name"`
		DisplayName  *string    `json:"displayName"`
		Creator      *string    `json:"creator"`
		CreatorId    *string    `json:"creatorID"`
		ClusterTotal *int32     `json:"clusterTotal"`
		CreatedAt    *time.Time `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.Creator == nil {
		return fmt.Errorf("required field creator missing")
	}
	if all.CreatorId == nil {
		return fmt.Errorf("required field creatorID missing")
	}
	if all.ClusterTotal == nil {
		return fmt.Errorf("required field clusterTotal missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "displayName", "creator", "creatorID", "clusterTotal", "createdAt"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.DisplayName = *all.DisplayName
	o.Creator = *all.Creator
	o.CreatorId = *all.CreatorId
	o.ClusterTotal = *all.ClusterTotal
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
