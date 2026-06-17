// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlWaitChainNode struct {
	Pid int64 `json:"pid"`
	// Role in the selected session wait graph. Known values are selected, blocking, blocked, and related.
	Role    string             `json:"role"`
	Session *PostgresqlSession `json:"session,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlWaitChainNode instantiates a new PostgresqlWaitChainNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlWaitChainNode(pid int64, role string) *PostgresqlWaitChainNode {
	this := PostgresqlWaitChainNode{}
	this.Pid = pid
	this.Role = role
	return &this
}

// NewPostgresqlWaitChainNodeWithDefaults instantiates a new PostgresqlWaitChainNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlWaitChainNodeWithDefaults() *PostgresqlWaitChainNode {
	this := PostgresqlWaitChainNode{}
	return &this
}

// GetPid returns the Pid field value.
func (o *PostgresqlWaitChainNode) GetPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitChainNode) GetPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value.
func (o *PostgresqlWaitChainNode) SetPid(v int64) {
	o.Pid = v
}

// GetRole returns the Role field value.
func (o *PostgresqlWaitChainNode) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitChainNode) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *PostgresqlWaitChainNode) SetRole(v string) {
	o.Role = v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *PostgresqlWaitChainNode) GetSession() PostgresqlSession {
	if o == nil || o.Session == nil {
		var ret PostgresqlSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitChainNode) GetSessionOk() (*PostgresqlSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *PostgresqlWaitChainNode) HasSession() bool {
	return o != nil && o.Session != nil
}

// SetSession gets a reference to the given PostgresqlSession and assigns it to the Session field.
func (o *PostgresqlWaitChainNode) SetSession(v PostgresqlSession) {
	o.Session = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlWaitChainNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pid"] = o.Pid
	toSerialize["role"] = o.Role
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlWaitChainNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pid     *int64             `json:"pid"`
		Role    *string            `json:"role"`
		Session *PostgresqlSession `json:"session,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Pid == nil {
		return fmt.Errorf("required field pid missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pid", "role", "session"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Pid = *all.Pid
	o.Role = *all.Role
	if all.Session != nil && all.Session.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Session = all.Session

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
