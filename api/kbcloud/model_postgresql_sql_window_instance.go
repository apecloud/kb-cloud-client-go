// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PostgresqlSQLWindowInstance struct {
	InstanceId     *string `json:"instanceId,omitempty"`
	InstanceName   *string `json:"instanceName,omitempty"`
	LastObservedAt *string `json:"lastObservedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowInstance instantiates a new PostgresqlSQLWindowInstance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowInstance() *PostgresqlSQLWindowInstance {
	this := PostgresqlSQLWindowInstance{}
	return &this
}

// NewPostgresqlSQLWindowInstanceWithDefaults instantiates a new PostgresqlSQLWindowInstance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowInstanceWithDefaults() *PostgresqlSQLWindowInstance {
	this := PostgresqlSQLWindowInstance{}
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowInstance) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowInstance) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowInstance) HasInstanceId() bool {
	return o != nil && o.InstanceId != nil
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *PostgresqlSQLWindowInstance) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowInstance) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowInstance) HasInstanceName() bool {
	return o != nil && o.InstanceName != nil
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *PostgresqlSQLWindowInstance) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetLastObservedAt returns the LastObservedAt field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowInstance) GetLastObservedAt() string {
	if o == nil || o.LastObservedAt == nil {
		var ret string
		return ret
	}
	return *o.LastObservedAt
}

// GetLastObservedAtOk returns a tuple with the LastObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowInstance) GetLastObservedAtOk() (*string, bool) {
	if o == nil || o.LastObservedAt == nil {
		return nil, false
	}
	return o.LastObservedAt, true
}

// HasLastObservedAt returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowInstance) HasLastObservedAt() bool {
	return o != nil && o.LastObservedAt != nil
}

// SetLastObservedAt gets a reference to the given string and assigns it to the LastObservedAt field.
func (o *PostgresqlSQLWindowInstance) SetLastObservedAt(v string) {
	o.LastObservedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.InstanceId != nil {
		toSerialize["instanceId"] = o.InstanceId
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.LastObservedAt != nil {
		toSerialize["lastObservedAt"] = o.LastObservedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowInstance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceId     *string `json:"instanceId,omitempty"`
		InstanceName   *string `json:"instanceName,omitempty"`
		LastObservedAt *string `json:"lastObservedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceId", "instanceName", "lastObservedAt"})
	} else {
		return err
	}
	o.InstanceId = all.InstanceId
	o.InstanceName = all.InstanceName
	o.LastObservedAt = all.LastObservedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
