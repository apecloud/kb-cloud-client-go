// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	_io "io"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterLicense struct {
	// NODESCRIPTION Id
	Id *string `json:"id,omitempty"`
	// NODESCRIPTION Name
	Name *string `json:"name,omitempty"`
	// NODESCRIPTION ExpiredAt
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	// NODESCRIPTION Key
	Key *_io.Reader `json:"key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterLicense instantiates a new ClusterLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterLicense() *ClusterLicense {
	this := ClusterLicense{}
	return &this
}

// NewClusterLicenseWithDefaults instantiates a new ClusterLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterLicenseWithDefaults() *ClusterLicense {
	this := ClusterLicense{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterLicense) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLicense) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterLicense) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterLicense) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterLicense) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLicense) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterLicense) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterLicense) SetName(v string) {
	o.Name = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *ClusterLicense) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLicense) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *ClusterLicense) HasExpiredAt() bool {
	return o != nil && o.ExpiredAt != nil
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *ClusterLicense) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ClusterLicense) GetKey() _io.Reader {
	if o == nil || o.Key == nil {
		var ret _io.Reader
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLicense) GetKeyOk() (*_io.Reader, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ClusterLicense) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given _io.Reader and assigns it to the Key field.
func (o *ClusterLicense) SetKey(v _io.Reader) {
	o.Key = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExpiredAt != nil {
		if o.ExpiredAt.Nanosecond() == 0 {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string     `json:"id,omitempty"`
		Name      *string     `json:"name,omitempty"`
		ExpiredAt *time.Time  `json:"expiredAt,omitempty"`
		Key       *_io.Reader `json:"key,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "expiredAt", "key"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.ExpiredAt = all.ExpiredAt
	o.Key = all.Key

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
