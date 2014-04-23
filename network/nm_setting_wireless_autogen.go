// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingWirelessKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_WIRELESS_SSID:
		t = ktypeWrapperString
	case NM_SETTING_WIRELESS_MODE:
		t = ktypeString
	case NM_SETTING_WIRELESS_BAND:
		t = ktypeString
	case NM_SETTING_WIRELESS_CHANNEL:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_BSSID:
		t = ktypeWrapperString
	case NM_SETTING_WIRELESS_RATE:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_TX_POWER:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_MTU:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SEC:
		t = ktypeString
	case NM_SETTING_WIRELESS_HIDDEN:
		t = ktypeBoolean
	}
	return
}

// Check is key in current setting field
func isKeyInSettingWireless(key string) bool {
	switch key {
	case NM_SETTING_WIRELESS_SSID:
		return true
	case NM_SETTING_WIRELESS_MODE:
		return true
	case NM_SETTING_WIRELESS_BAND:
		return true
	case NM_SETTING_WIRELESS_CHANNEL:
		return true
	case NM_SETTING_WIRELESS_BSSID:
		return true
	case NM_SETTING_WIRELESS_RATE:
		return true
	case NM_SETTING_WIRELESS_TX_POWER:
		return true
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		return true
	case NM_SETTING_WIRELESS_MTU:
		return true
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		return true
	case NM_SETTING_WIRELESS_SEC:
		return true
	case NM_SETTING_WIRELESS_HIDDEN:
		return true
	}
	return false
}

// Get key's default value
func getSettingWirelessKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		Logger.Error("invalid key:", key)
	case NM_SETTING_WIRELESS_SSID:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_MODE:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_BAND:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_CHANNEL:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_BSSID:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_RATE:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_TX_POWER:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		valueJSON = `null`
	case NM_SETTING_WIRELESS_MTU:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		valueJSON = `null`
	case NM_SETTING_WIRELESS_SEC:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_HIDDEN:
		valueJSON = `false`
	}
	return
}

// Get JSON value generally
func generalGetSettingWirelessKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		Logger.Error("generalGetSettingWirelessKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SSID:
		value = getSettingWirelessSsidJSON(data)
	case NM_SETTING_WIRELESS_MODE:
		value = getSettingWirelessModeJSON(data)
	case NM_SETTING_WIRELESS_BAND:
		value = getSettingWirelessBandJSON(data)
	case NM_SETTING_WIRELESS_CHANNEL:
		value = getSettingWirelessChannelJSON(data)
	case NM_SETTING_WIRELESS_BSSID:
		value = getSettingWirelessBssidJSON(data)
	case NM_SETTING_WIRELESS_RATE:
		value = getSettingWirelessRateJSON(data)
	case NM_SETTING_WIRELESS_TX_POWER:
		value = getSettingWirelessTxPowerJSON(data)
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		value = getSettingWirelessMacAddressJSON(data)
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		value = getSettingWirelessClonedMacAddressJSON(data)
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		value = getSettingWirelessMacAddressBlacklistJSON(data)
	case NM_SETTING_WIRELESS_MTU:
		value = getSettingWirelessMtuJSON(data)
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		value = getSettingWirelessSeenBssidsJSON(data)
	case NM_SETTING_WIRELESS_SEC:
		value = getSettingWirelessSecJSON(data)
	case NM_SETTING_WIRELESS_HIDDEN:
		value = getSettingWirelessHiddenJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingWirelessKeyJSON(data _ConnectionData, key, valueJSON string) {
	switch key {
	default:
		Logger.Error("generalSetSettingWirelessKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SSID:
		setSettingWirelessSsidJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MODE:
		setSettingWirelessModeJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_BAND:
		setSettingWirelessBandJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_CHANNEL:
		setSettingWirelessChannelJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_BSSID:
		setSettingWirelessBssidJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_RATE:
		setSettingWirelessRateJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_TX_POWER:
		setSettingWirelessTxPowerJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		setSettingWirelessMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		setSettingWirelessClonedMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		setSettingWirelessMacAddressBlacklistJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MTU:
		setSettingWirelessMtuJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		setSettingWirelessSeenBssidsJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SEC:
		setSettingWirelessSecJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_HIDDEN:
		setSettingWirelessHiddenJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingWirelessSsidExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID)
}
func isSettingWirelessModeExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE)
}
func isSettingWirelessBandExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND)
}
func isSettingWirelessChannelExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL)
}
func isSettingWirelessBssidExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID)
}
func isSettingWirelessRateExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE)
}
func isSettingWirelessTxPowerExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER)
}
func isSettingWirelessMacAddressExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS)
}
func isSettingWirelessClonedMacAddressExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS)
}
func isSettingWirelessMacAddressBlacklistExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST)
}
func isSettingWirelessMtuExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU)
}
func isSettingWirelessSeenBssidsExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS)
}
func isSettingWirelessSecExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC)
}
func isSettingWirelessHiddenExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN)
}

// Ensure field and key exists and not empty
func ensureFieldSettingWirelessExists(data _ConnectionData, errs FieldKeyErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_WIRELESS_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_WIRELESS_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_WIRELESS_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_WIRELESS_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_WIRELESS_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_WIRELESS_SETTING_NAME))
	}
}
func ensureSettingWirelessSsidNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessSsidExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSsid(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessModeNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessModeExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessMode(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessBandNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessBandExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessBand(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessChannelNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessChannelExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessBssidNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessBssidExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessBssid(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessRateNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessRateExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessTxPowerNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessTxPowerExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessMacAddressNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessClonedMacAddressNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessClonedMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessClonedMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessMacAddressBlacklistNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessMacAddressBlacklistExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessMacAddressBlacklist(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessMtuNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessMtuExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessSeenBssidsNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessSeenBssidsExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSeenBssids(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessSecExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSec(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessHiddenNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingWirelessHiddenExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingWirelessSsid(data _ConnectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID).([]byte)
	return
}
func getSettingWirelessMode(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE).(string)
	return
}
func getSettingWirelessBand(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND).(string)
	return
}
func getSettingWirelessChannel(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL).(uint32)
	return
}
func getSettingWirelessBssid(data _ConnectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID).([]byte)
	return
}
func getSettingWirelessRate(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE).(uint32)
	return
}
func getSettingWirelessTxPower(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER).(uint32)
	return
}
func getSettingWirelessMacAddress(data _ConnectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS).([]byte)
	return
}
func getSettingWirelessClonedMacAddress(data _ConnectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS).([]byte)
	return
}
func getSettingWirelessMacAddressBlacklist(data _ConnectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST).([]string)
	return
}
func getSettingWirelessMtu(data _ConnectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU).(uint32)
	return
}
func getSettingWirelessSeenBssids(data _ConnectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS).([]string)
	return
}
func getSettingWirelessSec(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC).(string)
	return
}
func getSettingWirelessHidden(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN).(bool)
	return
}

// Setter
func setSettingWirelessSsid(data _ConnectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, value)
}
func setSettingWirelessMode(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, value)
}
func setSettingWirelessBand(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, value)
}
func setSettingWirelessChannel(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, value)
}
func setSettingWirelessBssid(data _ConnectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, value)
}
func setSettingWirelessRate(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, value)
}
func setSettingWirelessTxPower(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, value)
}
func setSettingWirelessMacAddress(data _ConnectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, value)
}
func setSettingWirelessClonedMacAddress(data _ConnectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, value)
}
func setSettingWirelessMacAddressBlacklist(data _ConnectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, value)
}
func setSettingWirelessMtu(data _ConnectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, value)
}
func setSettingWirelessSeenBssids(data _ConnectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, value)
}
func setSettingWirelessSec(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, value)
}
func setSettingWirelessHidden(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, value)
}

// JSON Getter
func getSettingWirelessSsidJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SSID))
	return
}
func getSettingWirelessModeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MODE))
	return
}
func getSettingWirelessBandJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BAND))
	return
}
func getSettingWirelessChannelJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CHANNEL))
	return
}
func getSettingWirelessBssidJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BSSID))
	return
}
func getSettingWirelessRateJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, getSettingWirelessKeyType(NM_SETTING_WIRELESS_RATE))
	return
}
func getSettingWirelessTxPowerJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, getSettingWirelessKeyType(NM_SETTING_WIRELESS_TX_POWER))
	return
}
func getSettingWirelessMacAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS))
	return
}
func getSettingWirelessClonedMacAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS))
	return
}
func getSettingWirelessMacAddressBlacklistJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST))
	return
}
func getSettingWirelessMtuJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MTU))
	return
}
func getSettingWirelessSeenBssidsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEEN_BSSIDS))
	return
}
func getSettingWirelessSecJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEC))
	return
}
func getSettingWirelessHiddenJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, getSettingWirelessKeyType(NM_SETTING_WIRELESS_HIDDEN))
	return
}

// JSON Setter
func setSettingWirelessSsidJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SSID))
}
func setSettingWirelessModeJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MODE))
}
func setSettingWirelessBandJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BAND))
}
func setSettingWirelessChannelJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CHANNEL))
}
func setSettingWirelessBssidJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BSSID))
}
func setSettingWirelessRateJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_RATE))
}
func setSettingWirelessTxPowerJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_TX_POWER))
}
func setSettingWirelessMacAddressJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS))
}
func setSettingWirelessClonedMacAddressJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS))
}
func setSettingWirelessMacAddressBlacklistJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST))
}
func setSettingWirelessMtuJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MTU))
}
func setSettingWirelessSeenBssidsJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEEN_BSSIDS))
}
func setSettingWirelessSecJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEC))
}
func setSettingWirelessHiddenJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_HIDDEN))
}

// Remover
func removeSettingWirelessSsid(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID)
}
func removeSettingWirelessMode(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE)
}
func removeSettingWirelessBand(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND)
}
func removeSettingWirelessChannel(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL)
}
func removeSettingWirelessBssid(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID)
}
func removeSettingWirelessRate(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE)
}
func removeSettingWirelessTxPower(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER)
}
func removeSettingWirelessMacAddress(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS)
}
func removeSettingWirelessClonedMacAddress(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS)
}
func removeSettingWirelessMacAddressBlacklist(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST)
}
func removeSettingWirelessMtu(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU)
}
func removeSettingWirelessSeenBssids(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS)
}
func removeSettingWirelessSec(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC)
}
func removeSettingWirelessHidden(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN)
}
