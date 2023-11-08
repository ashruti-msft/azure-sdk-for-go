//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeys

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BackupKeyResult.
func (b BackupKeyResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", b.Value, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BackupKeyResult.
func (b *BackupKeyResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = runtime.DecodeByteArray(string(val), &b.Value, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CreateKeyParameters.
func (c CreateKeyParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "crv", c.Curve)
	populate(objectMap, "attributes", c.KeyAttributes)
	populate(objectMap, "key_ops", c.KeyOps)
	populate(objectMap, "key_size", c.KeySize)
	populate(objectMap, "kty", c.Kty)
	populate(objectMap, "public_exponent", c.PublicExponent)
	populate(objectMap, "release_policy", c.ReleasePolicy)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateKeyParameters.
func (c *CreateKeyParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "crv":
			err = unpopulate(val, "Curve", &c.Curve)
			delete(rawMsg, key)
		case "attributes":
			err = unpopulate(val, "KeyAttributes", &c.KeyAttributes)
			delete(rawMsg, key)
		case "key_ops":
			err = unpopulate(val, "KeyOps", &c.KeyOps)
			delete(rawMsg, key)
		case "key_size":
			err = unpopulate(val, "KeySize", &c.KeySize)
			delete(rawMsg, key)
		case "kty":
			err = unpopulate(val, "Kty", &c.Kty)
			delete(rawMsg, key)
		case "public_exponent":
			err = unpopulate(val, "PublicExponent", &c.PublicExponent)
			delete(rawMsg, key)
		case "release_policy":
			err = unpopulate(val, "ReleasePolicy", &c.ReleasePolicy)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &c.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeletedKey.
func (d DeletedKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", d.Attributes)
	populateTimeUnix(objectMap, "deletedDate", d.DeletedDate)
	populate(objectMap, "key", d.Key)
	populate(objectMap, "managed", d.Managed)
	populate(objectMap, "recoveryId", d.RecoveryID)
	populate(objectMap, "release_policy", d.ReleasePolicy)
	populateTimeUnix(objectMap, "scheduledPurgeDate", d.ScheduledPurgeDate)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedKey.
func (d *DeletedKey) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &d.Attributes)
			delete(rawMsg, key)
		case "deletedDate":
			err = unpopulateTimeUnix(val, "DeletedDate", &d.DeletedDate)
			delete(rawMsg, key)
		case "key":
			err = unpopulate(val, "Key", &d.Key)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &d.Managed)
			delete(rawMsg, key)
		case "recoveryId":
			err = unpopulate(val, "RecoveryID", &d.RecoveryID)
			delete(rawMsg, key)
		case "release_policy":
			err = unpopulate(val, "ReleasePolicy", &d.ReleasePolicy)
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			err = unpopulateTimeUnix(val, "ScheduledPurgeDate", &d.ScheduledPurgeDate)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &d.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeletedKeyProperties.
func (d DeletedKeyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", d.Attributes)
	populateTimeUnix(objectMap, "deletedDate", d.DeletedDate)
	populate(objectMap, "kid", d.KID)
	populate(objectMap, "managed", d.Managed)
	populate(objectMap, "recoveryId", d.RecoveryID)
	populateTimeUnix(objectMap, "scheduledPurgeDate", d.ScheduledPurgeDate)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedKeyProperties.
func (d *DeletedKeyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &d.Attributes)
			delete(rawMsg, key)
		case "deletedDate":
			err = unpopulateTimeUnix(val, "DeletedDate", &d.DeletedDate)
			delete(rawMsg, key)
		case "kid":
			err = unpopulate(val, "KID", &d.KID)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &d.Managed)
			delete(rawMsg, key)
		case "recoveryId":
			err = unpopulate(val, "RecoveryID", &d.RecoveryID)
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			err = unpopulateTimeUnix(val, "ScheduledPurgeDate", &d.ScheduledPurgeDate)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &d.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeletedKeyPropertiesListResult.
func (d DeletedKeyPropertiesListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedKeyPropertiesListResult.
func (d *DeletedKeyPropertiesListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &d.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &d.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GetRandomBytesParameters.
func (g GetRandomBytesParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "count", g.Count)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetRandomBytesParameters.
func (g *GetRandomBytesParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, "Count", &g.Count)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImportKeyParameters.
func (i ImportKeyParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "Hsm", i.HSM)
	populate(objectMap, "key", i.Key)
	populate(objectMap, "attributes", i.KeyAttributes)
	populate(objectMap, "release_policy", i.ReleasePolicy)
	populate(objectMap, "tags", i.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImportKeyParameters.
func (i *ImportKeyParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "Hsm":
			err = unpopulate(val, "HSM", &i.HSM)
			delete(rawMsg, key)
		case "key":
			err = unpopulate(val, "Key", &i.Key)
			delete(rawMsg, key)
		case "attributes":
			err = unpopulate(val, "KeyAttributes", &i.KeyAttributes)
			delete(rawMsg, key)
		case "release_policy":
			err = unpopulate(val, "ReleasePolicy", &i.ReleasePolicy)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &i.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JSONWebKey.
func (j JSONWebKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "crv", j.Crv)
	populateByteArray(objectMap, "d", j.D, runtime.Base64URLFormat)
	populateByteArray(objectMap, "dp", j.DP, runtime.Base64URLFormat)
	populateByteArray(objectMap, "dq", j.DQ, runtime.Base64URLFormat)
	populateByteArray(objectMap, "e", j.E, runtime.Base64URLFormat)
	populateByteArray(objectMap, "k", j.K, runtime.Base64URLFormat)
	populate(objectMap, "kid", j.KID)
	populate(objectMap, "key_ops", j.KeyOps)
	populate(objectMap, "kty", j.Kty)
	populateByteArray(objectMap, "n", j.N, runtime.Base64URLFormat)
	populateByteArray(objectMap, "p", j.P, runtime.Base64URLFormat)
	populateByteArray(objectMap, "q", j.Q, runtime.Base64URLFormat)
	populateByteArray(objectMap, "qi", j.QI, runtime.Base64URLFormat)
	populateByteArray(objectMap, "key_hsm", j.T, runtime.Base64URLFormat)
	populateByteArray(objectMap, "x", j.X, runtime.Base64URLFormat)
	populateByteArray(objectMap, "y", j.Y, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JSONWebKey.
func (j *JSONWebKey) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "crv":
			err = unpopulate(val, "Crv", &j.Crv)
			delete(rawMsg, key)
		case "d":
			err = runtime.DecodeByteArray(string(val), &j.D, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "dp":
			err = runtime.DecodeByteArray(string(val), &j.DP, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "dq":
			err = runtime.DecodeByteArray(string(val), &j.DQ, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "e":
			err = runtime.DecodeByteArray(string(val), &j.E, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "k":
			err = runtime.DecodeByteArray(string(val), &j.K, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "kid":
			err = unpopulate(val, "KID", &j.KID)
			delete(rawMsg, key)
		case "key_ops":
			err = unpopulate(val, "KeyOps", &j.KeyOps)
			delete(rawMsg, key)
		case "kty":
			err = unpopulate(val, "Kty", &j.Kty)
			delete(rawMsg, key)
		case "n":
			err = runtime.DecodeByteArray(string(val), &j.N, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "p":
			err = runtime.DecodeByteArray(string(val), &j.P, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "q":
			err = runtime.DecodeByteArray(string(val), &j.Q, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "qi":
			err = runtime.DecodeByteArray(string(val), &j.QI, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "key_hsm":
			err = runtime.DecodeByteArray(string(val), &j.T, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "x":
			err = runtime.DecodeByteArray(string(val), &j.X, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "y":
			err = runtime.DecodeByteArray(string(val), &j.Y, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyAttributes.
func (k KeyAttributes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeUnix(objectMap, "created", k.Created)
	populate(objectMap, "enabled", k.Enabled)
	populateTimeUnix(objectMap, "exp", k.Expires)
	populate(objectMap, "exportable", k.Exportable)
	populate(objectMap, "hsmPlatform", k.HsmPlatform)
	populateTimeUnix(objectMap, "nbf", k.NotBefore)
	populate(objectMap, "recoverableDays", k.RecoverableDays)
	populate(objectMap, "recoveryLevel", k.RecoveryLevel)
	populateTimeUnix(objectMap, "updated", k.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyAttributes.
func (k *KeyAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulateTimeUnix(val, "Created", &k.Created)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, "Enabled", &k.Enabled)
			delete(rawMsg, key)
		case "exp":
			err = unpopulateTimeUnix(val, "Expires", &k.Expires)
			delete(rawMsg, key)
		case "exportable":
			err = unpopulate(val, "Exportable", &k.Exportable)
			delete(rawMsg, key)
		case "hsmPlatform":
			err = unpopulate(val, "HsmPlatform", &k.HsmPlatform)
			delete(rawMsg, key)
		case "nbf":
			err = unpopulateTimeUnix(val, "NotBefore", &k.NotBefore)
			delete(rawMsg, key)
		case "recoverableDays":
			err = unpopulate(val, "RecoverableDays", &k.RecoverableDays)
			delete(rawMsg, key)
		case "recoveryLevel":
			err = unpopulate(val, "RecoveryLevel", &k.RecoveryLevel)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeUnix(val, "Updated", &k.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyBundle.
func (k KeyBundle) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", k.Attributes)
	populate(objectMap, "key", k.Key)
	populate(objectMap, "managed", k.Managed)
	populate(objectMap, "release_policy", k.ReleasePolicy)
	populate(objectMap, "tags", k.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyBundle.
func (k *KeyBundle) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &k.Attributes)
			delete(rawMsg, key)
		case "key":
			err = unpopulate(val, "Key", &k.Key)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &k.Managed)
			delete(rawMsg, key)
		case "release_policy":
			err = unpopulate(val, "ReleasePolicy", &k.ReleasePolicy)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &k.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyOperationParameters.
func (k KeyOperationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "aad", k.AdditionalAuthenticatedData, runtime.Base64URLFormat)
	populate(objectMap, "alg", k.Algorithm)
	populateByteArray(objectMap, "tag", k.AuthenticationTag, runtime.Base64URLFormat)
	populateByteArray(objectMap, "iv", k.IV, runtime.Base64URLFormat)
	populateByteArray(objectMap, "value", k.Value, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyOperationParameters.
func (k *KeyOperationParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aad":
			err = runtime.DecodeByteArray(string(val), &k.AdditionalAuthenticatedData, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "alg":
			err = unpopulate(val, "Algorithm", &k.Algorithm)
			delete(rawMsg, key)
		case "tag":
			err = runtime.DecodeByteArray(string(val), &k.AuthenticationTag, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "iv":
			err = runtime.DecodeByteArray(string(val), &k.IV, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "value":
			err = runtime.DecodeByteArray(string(val), &k.Value, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyOperationResult.
func (k KeyOperationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "aad", k.AdditionalAuthenticatedData, runtime.Base64URLFormat)
	populateByteArray(objectMap, "tag", k.AuthenticationTag, runtime.Base64URLFormat)
	populateByteArray(objectMap, "iv", k.IV, runtime.Base64URLFormat)
	populate(objectMap, "kid", k.KID)
	populateByteArray(objectMap, "value", k.Result, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyOperationResult.
func (k *KeyOperationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aad":
			err = runtime.DecodeByteArray(string(val), &k.AdditionalAuthenticatedData, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "tag":
			err = runtime.DecodeByteArray(string(val), &k.AuthenticationTag, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "iv":
			err = runtime.DecodeByteArray(string(val), &k.IV, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "kid":
			err = unpopulate(val, "KID", &k.KID)
			delete(rawMsg, key)
		case "value":
			err = runtime.DecodeByteArray(string(val), &k.Result, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyProperties.
func (k KeyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", k.Attributes)
	populate(objectMap, "kid", k.KID)
	populate(objectMap, "managed", k.Managed)
	populate(objectMap, "tags", k.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyProperties.
func (k *KeyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &k.Attributes)
			delete(rawMsg, key)
		case "kid":
			err = unpopulate(val, "KID", &k.KID)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &k.Managed)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &k.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyPropertiesListResult.
func (k KeyPropertiesListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", k.NextLink)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyPropertiesListResult.
func (k *KeyPropertiesListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &k.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &k.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyReleasePolicy.
func (k KeyReleasePolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "contentType", k.ContentType)
	populateByteArray(objectMap, "data", k.EncodedPolicy, runtime.Base64URLFormat)
	populate(objectMap, "immutable", k.Immutable)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyReleasePolicy.
func (k *KeyReleasePolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, "ContentType", &k.ContentType)
			delete(rawMsg, key)
		case "data":
			err = runtime.DecodeByteArray(string(val), &k.EncodedPolicy, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "immutable":
			err = unpopulate(val, "Immutable", &k.Immutable)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyReleaseResult.
func (k KeyReleaseResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyReleaseResult.
func (k *KeyReleaseResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &k.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyRotationPolicy.
func (k KeyRotationPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", k.Attributes)
	populate(objectMap, "id", k.ID)
	populate(objectMap, "lifetimeActions", k.LifetimeActions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyRotationPolicy.
func (k *KeyRotationPolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &k.Attributes)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &k.ID)
			delete(rawMsg, key)
		case "lifetimeActions":
			err = unpopulate(val, "LifetimeActions", &k.LifetimeActions)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyRotationPolicyAttributes.
func (k KeyRotationPolicyAttributes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeUnix(objectMap, "created", k.Created)
	populate(objectMap, "expiryTime", k.ExpiryTime)
	populateTimeUnix(objectMap, "updated", k.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyRotationPolicyAttributes.
func (k *KeyRotationPolicyAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulateTimeUnix(val, "Created", &k.Created)
			delete(rawMsg, key)
		case "expiryTime":
			err = unpopulate(val, "ExpiryTime", &k.ExpiryTime)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeUnix(val, "Updated", &k.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyVerifyResult.
func (k KeyVerifyResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyVerifyResult.
func (k *KeyVerifyResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &k.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LifetimeAction.
func (l LifetimeAction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "action", l.Action)
	populate(objectMap, "trigger", l.Trigger)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LifetimeAction.
func (l *LifetimeAction) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "action":
			err = unpopulate(val, "Action", &l.Action)
			delete(rawMsg, key)
		case "trigger":
			err = unpopulate(val, "Trigger", &l.Trigger)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LifetimeActionTrigger.
func (l LifetimeActionTrigger) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "timeAfterCreate", l.TimeAfterCreate)
	populate(objectMap, "timeBeforeExpiry", l.TimeBeforeExpiry)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LifetimeActionTrigger.
func (l *LifetimeActionTrigger) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "timeAfterCreate":
			err = unpopulate(val, "TimeAfterCreate", &l.TimeAfterCreate)
			delete(rawMsg, key)
		case "timeBeforeExpiry":
			err = unpopulate(val, "TimeBeforeExpiry", &l.TimeBeforeExpiry)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LifetimeActionType.
func (l LifetimeActionType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LifetimeActionType.
func (l *LifetimeActionType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "type":
			err = unpopulate(val, "Type", &l.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RandomBytes.
func (r RandomBytes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", r.Value, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RandomBytes.
func (r *RandomBytes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = runtime.DecodeByteArray(string(val), &r.Value, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReleaseParameters.
func (r ReleaseParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enc", r.Algorithm)
	populate(objectMap, "nonce", r.Nonce)
	populate(objectMap, "target", r.TargetAttestationToken)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReleaseParameters.
func (r *ReleaseParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enc":
			err = unpopulate(val, "Algorithm", &r.Algorithm)
			delete(rawMsg, key)
		case "nonce":
			err = unpopulate(val, "Nonce", &r.Nonce)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "TargetAttestationToken", &r.TargetAttestationToken)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RestoreKeyParameters.
func (r RestoreKeyParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", r.KeyBackup, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoreKeyParameters.
func (r *RestoreKeyParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = runtime.DecodeByteArray(string(val), &r.KeyBackup, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SignParameters.
func (s SignParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "alg", s.Algorithm)
	populateByteArray(objectMap, "value", s.Value, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SignParameters.
func (s *SignParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "alg":
			err = unpopulate(val, "Algorithm", &s.Algorithm)
			delete(rawMsg, key)
		case "value":
			err = runtime.DecodeByteArray(string(val), &s.Value, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateKeyParameters.
func (u UpdateKeyParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", u.KeyAttributes)
	populate(objectMap, "key_ops", u.KeyOps)
	populate(objectMap, "release_policy", u.ReleasePolicy)
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateKeyParameters.
func (u *UpdateKeyParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "KeyAttributes", &u.KeyAttributes)
			delete(rawMsg, key)
		case "key_ops":
			err = unpopulate(val, "KeyOps", &u.KeyOps)
			delete(rawMsg, key)
		case "release_policy":
			err = unpopulate(val, "ReleasePolicy", &u.ReleasePolicy)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &u.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VerifyParameters.
func (v VerifyParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "alg", v.Algorithm)
	populateByteArray(objectMap, "digest", v.Digest, runtime.Base64URLFormat)
	populateByteArray(objectMap, "value", v.Signature, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VerifyParameters.
func (v *VerifyParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "alg":
			err = unpopulate(val, "Algorithm", &v.Algorithm)
			delete(rawMsg, key)
		case "digest":
			err = runtime.DecodeByteArray(string(val), &v.Digest, runtime.Base64URLFormat)
			delete(rawMsg, key)
		case "value":
			err = runtime.DecodeByteArray(string(val), &v.Signature, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateByteArray(m map[string]any, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
