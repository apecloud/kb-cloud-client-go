// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DmsServiceAccountList MinIO service account access-key list
type DmsServiceAccountList struct {
	Items []DmsServiceAccount

	// UnparsedObject contains the raw value of the array if there was an error when deserializing into the struct
	UnparsedObject []interface{} `json:"-"`
}

// NewDmsServiceAccountList instantiates a new DmsServiceAccountList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsServiceAccountList() *DmsServiceAccountList {
	this := DmsServiceAccountList{}
	return &this
}

// NewDmsServiceAccountListWithDefaults instantiates a new DmsServiceAccountList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsServiceAccountListWithDefaults() *DmsServiceAccountList {
	this := DmsServiceAccountList{}
	return &this
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsServiceAccountList) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsServiceAccountList) UnmarshalJSON(bytes []byte) (err error) {
	if err = common.Unmarshal(bytes, &o.Items); err != nil {
		return err
	}

	if o.Items != nil && len(o.Items) > 0 {
		for _, v := range o.Items {
			if v.UnparsedObject != nil {
				return common.Unmarshal(bytes, &o.UnparsedObject)
			}
		}
	}

	return nil
}
