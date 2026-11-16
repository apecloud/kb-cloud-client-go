// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBBucketPermission struct {
	BucketId string `json:"bucketID"`
	Read     bool   `json:"read"`
	Write    bool   `json:"write"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBBucketPermission instantiates a new InfluxDBBucketPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBBucketPermission(bucketId string, read bool, write bool) *InfluxDBBucketPermission {
	this := InfluxDBBucketPermission{}
	this.BucketId = bucketId
	this.Read = read
	this.Write = write
	return &this
}

// NewInfluxDBBucketPermissionWithDefaults instantiates a new InfluxDBBucketPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBBucketPermissionWithDefaults() *InfluxDBBucketPermission {
	this := InfluxDBBucketPermission{}
	return &this
}

// GetBucketId returns the BucketId field value.
func (o *InfluxDBBucketPermission) GetBucketId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketPermission) GetBucketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketId, true
}

// SetBucketId sets field value.
func (o *InfluxDBBucketPermission) SetBucketId(v string) {
	o.BucketId = v
}

// GetRead returns the Read field value.
func (o *InfluxDBBucketPermission) GetRead() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Read
}

// GetReadOk returns a tuple with the Read field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketPermission) GetReadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Read, true
}

// SetRead sets field value.
func (o *InfluxDBBucketPermission) SetRead(v bool) {
	o.Read = v
}

// GetWrite returns the Write field value.
func (o *InfluxDBBucketPermission) GetWrite() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Write
}

// GetWriteOk returns a tuple with the Write field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucketPermission) GetWriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Write, true
}

// SetWrite sets field value.
func (o *InfluxDBBucketPermission) SetWrite(v bool) {
	o.Write = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBBucketPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["bucketID"] = o.BucketId
	toSerialize["read"] = o.Read
	toSerialize["write"] = o.Write

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBBucketPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketId *string `json:"bucketID"`
		Read     *bool   `json:"read"`
		Write    *bool   `json:"write"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.BucketId == nil {
		return fmt.Errorf("required field bucketID missing")
	}
	if all.Read == nil {
		return fmt.Errorf("required field read missing")
	}
	if all.Write == nil {
		return fmt.Errorf("required field write missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"bucketID", "read", "write"})
	} else {
		return err
	}
	o.BucketId = *all.BucketId
	o.Read = *all.Read
	o.Write = *all.Write

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
