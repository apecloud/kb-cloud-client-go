// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSpacePVC struct {
	// Current PVC used bytes from backend-owned metrics when available.
	CurrentUsedBytes  *int64 `json:"currentUsedBytes,omitempty"`
	NoHistoryYet      bool   `json:"noHistoryYet"`
	CannotProveGrowth bool   `json:"cannotProveGrowth"`
	Source            string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSpacePVC instantiates a new PostgresqlSpacePVC object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSpacePVC(noHistoryYet bool, cannotProveGrowth bool, source string) *PostgresqlSpacePVC {
	this := PostgresqlSpacePVC{}
	this.NoHistoryYet = noHistoryYet
	this.CannotProveGrowth = cannotProveGrowth
	this.Source = source
	return &this
}

// NewPostgresqlSpacePVCWithDefaults instantiates a new PostgresqlSpacePVC object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSpacePVCWithDefaults() *PostgresqlSpacePVC {
	this := PostgresqlSpacePVC{}
	return &this
}

// GetCurrentUsedBytes returns the CurrentUsedBytes field value if set, zero value otherwise.
func (o *PostgresqlSpacePVC) GetCurrentUsedBytes() int64 {
	if o == nil || o.CurrentUsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.CurrentUsedBytes
}

// GetCurrentUsedBytesOk returns a tuple with the CurrentUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSpacePVC) GetCurrentUsedBytesOk() (*int64, bool) {
	if o == nil || o.CurrentUsedBytes == nil {
		return nil, false
	}
	return o.CurrentUsedBytes, true
}

// HasCurrentUsedBytes returns a boolean if a field has been set.
func (o *PostgresqlSpacePVC) HasCurrentUsedBytes() bool {
	return o != nil && o.CurrentUsedBytes != nil
}

// SetCurrentUsedBytes gets a reference to the given int64 and assigns it to the CurrentUsedBytes field.
func (o *PostgresqlSpacePVC) SetCurrentUsedBytes(v int64) {
	o.CurrentUsedBytes = &v
}

// GetNoHistoryYet returns the NoHistoryYet field value.
func (o *PostgresqlSpacePVC) GetNoHistoryYet() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NoHistoryYet
}

// GetNoHistoryYetOk returns a tuple with the NoHistoryYet field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpacePVC) GetNoHistoryYetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoHistoryYet, true
}

// SetNoHistoryYet sets field value.
func (o *PostgresqlSpacePVC) SetNoHistoryYet(v bool) {
	o.NoHistoryYet = v
}

// GetCannotProveGrowth returns the CannotProveGrowth field value.
func (o *PostgresqlSpacePVC) GetCannotProveGrowth() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CannotProveGrowth
}

// GetCannotProveGrowthOk returns a tuple with the CannotProveGrowth field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpacePVC) GetCannotProveGrowthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProveGrowth, true
}

// SetCannotProveGrowth sets field value.
func (o *PostgresqlSpacePVC) SetCannotProveGrowth(v bool) {
	o.CannotProveGrowth = v
}

// GetSource returns the Source field value.
func (o *PostgresqlSpacePVC) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpacePVC) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *PostgresqlSpacePVC) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSpacePVC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CurrentUsedBytes != nil {
		toSerialize["currentUsedBytes"] = o.CurrentUsedBytes
	}
	toSerialize["noHistoryYet"] = o.NoHistoryYet
	toSerialize["cannotProveGrowth"] = o.CannotProveGrowth
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSpacePVC) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentUsedBytes  *int64  `json:"currentUsedBytes,omitempty"`
		NoHistoryYet      *bool   `json:"noHistoryYet"`
		CannotProveGrowth *bool   `json:"cannotProveGrowth"`
		Source            *string `json:"source"`
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
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"currentUsedBytes", "noHistoryYet", "cannotProveGrowth", "source"})
	} else {
		return err
	}
	o.CurrentUsedBytes = all.CurrentUsedBytes
	o.NoHistoryYet = *all.NoHistoryYet
	o.CannotProveGrowth = *all.CannotProveGrowth
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
