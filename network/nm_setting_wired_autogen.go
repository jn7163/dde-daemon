// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingWiredKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_WIRED_PORT:
		t = ktypeString
	case NM_SETTING_WIRED_SPEED:
		t = ktypeUint32
	case NM_SETTING_WIRED_DUPLEX:
		t = ktypeString
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		t = ktypeBoolean
	case NM_SETTING_WIRED_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		t = ktypeArrayString
	case NM_SETTING_WIRED_MTU:
		t = ktypeUint32
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		t = ktypeArrayString
	case NM_SETTING_WIRED_S390_NETTYPE:
		t = ktypeString
	case NM_SETTING_WIRED_S390_OPTIONS:
		t = ktypeDictStringString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingWired(key string) bool {
	switch key {
	case NM_SETTING_WIRED_PORT:
		return true
	case NM_SETTING_WIRED_SPEED:
		return true
	case NM_SETTING_WIRED_DUPLEX:
		return true
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		return true
	case NM_SETTING_WIRED_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		return true
	case NM_SETTING_WIRED_MTU:
		return true
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		return true
	case NM_SETTING_WIRED_S390_NETTYPE:
		return true
	case NM_SETTING_WIRED_S390_OPTIONS:
		return true
	}
	return false
}

// Get key's default value
func getSettingWiredKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		Logger.Error("invalid key:", key)
	case NM_SETTING_WIRED_PORT:
		valueJSON = `""`
	case NM_SETTING_WIRED_SPEED:
		valueJSON = `0`
	case NM_SETTING_WIRED_DUPLEX:
		valueJSON = `""`
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		valueJSON = `true`
	case NM_SETTING_WIRED_MAC_ADDRESS:
		valueJSON = `""`
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		valueJSON = `""`
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		valueJSON = `null`
	case NM_SETTING_WIRED_MTU:
		valueJSON = `0`
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		valueJSON = `null`
	case NM_SETTING_WIRED_S390_NETTYPE:
		valueJSON = `""`
	case NM_SETTING_WIRED_S390_OPTIONS:
		valueJSON = `null`
	}
	return
}

// Get JSON value generally
func generalGetSettingWiredKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		Logger.Error("generalGetSettingWiredKeyJSON: invalide key", key)
	case NM_SETTING_WIRED_PORT:
		value = getSettingWiredPortJSON(data)
	case NM_SETTING_WIRED_SPEED:
		value = getSettingWiredSpeedJSON(data)
	case NM_SETTING_WIRED_DUPLEX:
		value = getSettingWiredDuplexJSON(data)
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		value = getSettingWiredAutoNegotiateJSON(data)
	case NM_SETTING_WIRED_MAC_ADDRESS:
		value = getSettingWiredMacAddressJSON(data)
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		value = getSettingWiredClonedMacAddressJSON(data)
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		value = getSettingWiredMacAddressBlacklistJSON(data)
	case NM_SETTING_WIRED_MTU:
		value = getSettingWiredMtuJSON(data)
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		value = getSettingWiredS390SubchannelsJSON(data)
	case NM_SETTING_WIRED_S390_NETTYPE:
		value = getSettingWiredS390NettypeJSON(data)
	case NM_SETTING_WIRED_S390_OPTIONS:
		value = getSettingWiredS390OptionsJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingWiredKeyJSON(data _ConnectionData, key, valueJSON string) {
	switch key {
	default:
		Logger.Error("generalSetSettingWiredKeyJSON: invalide key", key)
	case NM_SETTING_WIRED_PORT:
		setSettingWiredPortJSON(data, valueJSON)
	case NM_SETTING_WIRED_SPEED:
		setSettingWiredSpeedJSON(data, valueJSON)
	case NM_SETTING_WIRED_DUPLEX:
		setSettingWiredDuplexJSON(data, valueJSON)
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		setSettingWiredAutoNegotiateJSON(data, valueJSON)
	case NM_SETTING_WIRED_MAC_ADDRESS:
		setSettingWiredMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		setSettingWiredClonedMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		setSettingWiredMacAddressBlacklistJSON(data, valueJSON)
	case NM_SETTING_WIRED_MTU:
		setSettingWiredMtuJSON(data, valueJSON)
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		setSettingWiredS390SubchannelsJSON(data, valueJSON)
	case NM_SETTING_WIRED_S390_NETTYPE:
		setSettingWiredS390NettypeJSON(data, valueJSON)
	case NM_SETTING_WIRED_S390_OPTIONS:
		setSettingWiredS390OptionsJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingWiredPortExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT)
}
func isSettingWiredSpeedExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED)
}
func isSettingWiredDuplexExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX)
}
func isSettingWiredAutoNegotiateExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE)
}
func isSettingWiredMacAddressExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS)
}
func isSettingWiredClonedMacAddressExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS)
}
func isSettingWiredMacAddressBlacklistExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST)
}
func isSettingWiredMtuExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU)
}
func isSettingWiredS390SubchannelsExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS)
}
func isSettingWiredS390NettypeExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE)
}
func isSettingWiredS390OptionsExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS)
}

// Ensure field and key exists and not empty
func ensureFieldSettingWiredExists(data _ConnectionData, errs FieldKeyErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_WIRED_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_WIRED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_WIRED_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_WIRED_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_WIRED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_WIRED_SETTING_NAME))
	}
}
func ensureSettingWiredPortNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredPortExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredPort(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredSpeedNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredSpeedExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWiredDuplexNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredDuplexExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredDuplex(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredAutoNegotiateNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredAutoNegotiateExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWiredMacAddressNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredClonedMacAddressNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredClonedMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredClonedMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredMacAddressBlacklistNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredMacAddressBlacklistExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredMacAddressBlacklist(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredMtuNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredMtuExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWiredS390SubchannelsNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredS390SubchannelsExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredS390Subchannels(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredS390NettypeNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredS390NettypeExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredS390Nettype(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredS390OptionsNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWiredS390OptionsExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredS390Options(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingWiredPort(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT).(string)
	return
}
func getSettingWiredSpeed(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED).(uint32)
	return
}
func getSettingWiredDuplex(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX).(string)
	return
}
func getSettingWiredAutoNegotiate(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE).(bool)
	return
}
func getSettingWiredMacAddress(data _ConnectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS).([]byte)
	return
}
func getSettingWiredClonedMacAddress(data _ConnectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS).([]byte)
	return
}
func getSettingWiredMacAddressBlacklist(data _ConnectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST).([]string)
	return
}
func getSettingWiredMtu(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU).(uint32)
	return
}
func getSettingWiredS390Subchannels(data _ConnectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS).([]string)
	return
}
func getSettingWiredS390Nettype(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE).(string)
	return
}
func getSettingWiredS390Options(data _ConnectionData) (value map[string]string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS).(map[string]string)
	return
}

// Setter
func setSettingWiredPort(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, value)
}
func setSettingWiredSpeed(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, value)
}
func setSettingWiredDuplex(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, value)
}
func setSettingWiredAutoNegotiate(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, value)
}
func setSettingWiredMacAddress(data _ConnectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, value)
}
func setSettingWiredClonedMacAddress(data _ConnectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, value)
}
func setSettingWiredMacAddressBlacklist(data _ConnectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, value)
}
func setSettingWiredMtu(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, value)
}
func setSettingWiredS390Subchannels(data _ConnectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, value)
}
func setSettingWiredS390Nettype(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, value)
}
func setSettingWiredS390Options(data _ConnectionData, value map[string]string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, value)
}

// JSON Getter
func getSettingWiredPortJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, getSettingWiredKeyType(NM_SETTING_WIRED_PORT))
	return
}
func getSettingWiredSpeedJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, getSettingWiredKeyType(NM_SETTING_WIRED_SPEED))
	return
}
func getSettingWiredDuplexJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, getSettingWiredKeyType(NM_SETTING_WIRED_DUPLEX))
	return
}
func getSettingWiredAutoNegotiateJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, getSettingWiredKeyType(NM_SETTING_WIRED_AUTO_NEGOTIATE))
	return
}
func getSettingWiredMacAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS))
	return
}
func getSettingWiredClonedMacAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, getSettingWiredKeyType(NM_SETTING_WIRED_CLONED_MAC_ADDRESS))
	return
}
func getSettingWiredMacAddressBlacklistJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST))
	return
}
func getSettingWiredMtuJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, getSettingWiredKeyType(NM_SETTING_WIRED_MTU))
	return
}
func getSettingWiredS390SubchannelsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, getSettingWiredKeyType(NM_SETTING_WIRED_S390_SUBCHANNELS))
	return
}
func getSettingWiredS390NettypeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, getSettingWiredKeyType(NM_SETTING_WIRED_S390_NETTYPE))
	return
}
func getSettingWiredS390OptionsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, getSettingWiredKeyType(NM_SETTING_WIRED_S390_OPTIONS))
	return
}

// JSON Setter
func setSettingWiredPortJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_PORT))
}
func setSettingWiredSpeedJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_SPEED))
}
func setSettingWiredDuplexJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_DUPLEX))
}
func setSettingWiredAutoNegotiateJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_AUTO_NEGOTIATE))
}
func setSettingWiredMacAddressJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS))
}
func setSettingWiredClonedMacAddressJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_CLONED_MAC_ADDRESS))
}
func setSettingWiredMacAddressBlacklistJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST))
}
func setSettingWiredMtuJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_MTU))
}
func setSettingWiredS390SubchannelsJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_S390_SUBCHANNELS))
}
func setSettingWiredS390NettypeJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_S390_NETTYPE))
}
func setSettingWiredS390OptionsJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_S390_OPTIONS))
}

// Remover
func removeSettingWiredPort(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT)
}
func removeSettingWiredSpeed(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED)
}
func removeSettingWiredDuplex(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX)
}
func removeSettingWiredAutoNegotiate(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE)
}
func removeSettingWiredMacAddress(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS)
}
func removeSettingWiredClonedMacAddress(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS)
}
func removeSettingWiredMacAddressBlacklist(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST)
}
func removeSettingWiredMtu(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU)
}
func removeSettingWiredS390Subchannels(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS)
}
func removeSettingWiredS390Nettype(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE)
}
func removeSettingWiredS390Options(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS)
}
