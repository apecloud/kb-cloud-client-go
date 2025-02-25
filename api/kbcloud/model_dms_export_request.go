// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

type DmsExportRequest struct {
	// the database of the table or view
	Database string `json:"database,omitempty"`
	// the target table or view name
	Table string `json:"table,omitempty"`
	// return limited number of data
	Limit int32 `json:"limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExportRequest instantiates a new DmsExportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExportRequest() *DmsExportRequest {
	this := DmsExportRequest{}
	this.Database = ""
	this.Limit = 200
	return &this
}
