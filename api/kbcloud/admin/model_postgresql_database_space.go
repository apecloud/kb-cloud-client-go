// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlDatabaseSpace struct {
	Name              *string `json:"name,omitempty"`
	SizeBytes         *int64  `json:"sizeBytes,omitempty"`
	NoHistoryYet      bool    `json:"noHistoryYet"`
	CannotProveGrowth bool    `json:"cannotProveGrowth"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlDatabaseSpace instantiates a new PostgresqlDatabaseSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlDatabaseSpace(noHistoryYet bool, cannotProveGrowth bool) *PostgresqlDatabaseSpace {
	this := PostgresqlDatabaseSpace{}
	this.NoHistoryYet = noHistoryYet
	this.CannotProveGrowth = cannotProveGrowth
	return &this
}

// NewPostgresqlDatabaseSpaceWithDefaults instantiates a new PostgresqlDatabaseSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlDatabaseSpaceWithDefaults() *PostgresqlDatabaseSpace {
	this := PostgresqlDatabaseSpace{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostgresqlDatabaseSpace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDatabaseSpace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostgresqlDatabaseSpace) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostgresqlDatabaseSpace) SetName(v string) {
	o.Name = &v
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise.
func (o *PostgresqlDatabaseSpace) GetSizeBytes() int64 {
	if o == nil || o.SizeBytes == nil {
		var ret int64
		return ret
	}
	return *o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDatabaseSpace) GetSizeBytesOk() (*int64, bool) {
	if o == nil || o.SizeBytes == nil {
		return nil, false
	}
	return o.SizeBytes, true
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *PostgresqlDatabaseSpace) HasSizeBytes() bool {
	return o != nil && o.SizeBytes != nil
}

// SetSizeBytes gets a reference to the given int64 and assigns it to the SizeBytes field.
func (o *PostgresqlDatabaseSpace) SetSizeBytes(v int64) {
	o.SizeBytes = &v
}

// GetNoHistoryYet returns the NoHistoryYet field value.
func (o *PostgresqlDatabaseSpace) GetNoHistoryYet() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NoHistoryYet
}

// GetNoHistoryYetOk returns a tuple with the NoHistoryYet field value
// and a boolean to check if the value has been set.
func (o *PostgresqlDatabaseSpace) GetNoHistoryYetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoHistoryYet, true
}

// SetNoHistoryYet sets field value.
func (o *PostgresqlDatabaseSpace) SetNoHistoryYet(v bool) {
	o.NoHistoryYet = v
}

// GetCannotProveGrowth returns the CannotProveGrowth field value.
func (o *PostgresqlDatabaseSpace) GetCannotProveGrowth() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CannotProveGrowth
}

// GetCannotProveGrowthOk returns a tuple with the CannotProveGrowth field value
// and a boolean to check if the value has been set.
func (o *PostgresqlDatabaseSpace) GetCannotProveGrowthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProveGrowth, true
}

// SetCannotProveGrowth sets field value.
func (o *PostgresqlDatabaseSpace) SetCannotProveGrowth(v bool) {
	o.CannotProveGrowth = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlDatabaseSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SizeBytes != nil {
		toSerialize["sizeBytes"] = o.SizeBytes
	}
	toSerialize["noHistoryYet"] = o.NoHistoryYet
	toSerialize["cannotProveGrowth"] = o.CannotProveGrowth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlDatabaseSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string `json:"name,omitempty"`
		SizeBytes         *int64  `json:"sizeBytes,omitempty"`
		NoHistoryYet      *bool   `json:"noHistoryYet"`
		CannotProveGrowth *bool   `json:"cannotProveGrowth"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NoHistoryYet == nil {
		return fmt.Errorf("required field noHistoryYet missing")
	}
	if all.CannotProveGrowth == nil {
		return fmt.Errorf("required field cannotProveGrowth missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "sizeBytes", "noHistoryYet", "cannotProveGrowth"})
	} else {
		return err
	}
	o.Name = all.Name
	o.SizeBytes = all.SizeBytes
	o.NoHistoryYet = *all.NoHistoryYet
	o.CannotProveGrowth = *all.CannotProveGrowth

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
