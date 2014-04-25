// This file is automatically generated, please don't edit manully.
package main

func generalIsKeyInSettingField(field, key string) bool {
	if isVirtualKey(field, key) {
		return true
	}
	switch field {
	default:
		Logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		return isKeyInSetting8021x(key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		return isKeyInSettingConnection(key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		return isKeyInSettingIp4Config(key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		return isKeyInSettingIp6Config(key)
	case NM_SETTING_PPP_SETTING_NAME:
		return isKeyInSettingPpp(key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		return isKeyInSettingPppoe(key)
	case NM_SETTING_VPN_SETTING_NAME:
		return isKeyInSettingVpn(key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		return isKeyInSettingVpnL2tp(key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		return isKeyInSettingVpnL2tpPpp(key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		return isKeyInSettingVpnL2tpIpsec(key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		return isKeyInSettingVpnOpenconnect(key)
	case NM_SETTING_WIRED_SETTING_NAME:
		return isKeyInSettingWired(key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		return isKeyInSettingWireless(key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		return isKeyInSettingWirelessSecurity(key)
	}
	return false
}

func generalGetSettingKeyType(field, key string) (t ktype) {
	if isVirtualKey(field, key) {
		t = getSettingVkKeyType(field, key)
		return
	}
	switch field {
	default:
		Logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		t = getSetting8021xKeyType(key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		t = getSettingConnectionKeyType(key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		t = getSettingIp4ConfigKeyType(key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		t = getSettingIp6ConfigKeyType(key)
	case NM_SETTING_PPP_SETTING_NAME:
		t = getSettingPppKeyType(key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		t = getSettingPppoeKeyType(key)
	case NM_SETTING_VPN_SETTING_NAME:
		t = getSettingVpnKeyType(key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		t = getSettingVpnL2tpKeyType(key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		t = getSettingVpnL2tpPppKeyType(key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		t = getSettingVpnL2tpIpsecKeyType(key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		t = getSettingVpnOpenconnectKeyType(key)
	case NM_SETTING_WIRED_SETTING_NAME:
		t = getSettingWiredKeyType(key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		t = getSettingWirelessKeyType(key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		t = getSettingWirelessSecurityKeyType(key)
	}
	return
}

func generalGetSettingAvailableKeys(data _ConnectionData, field string) (keys []string) {
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		keys = getSetting8021xAvailableKeys(data)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		keys = getSettingConnectionAvailableKeys(data)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		keys = getSettingIp4ConfigAvailableKeys(data)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		keys = getSettingIp6ConfigAvailableKeys(data)
	case NM_SETTING_PPP_SETTING_NAME:
		keys = getSettingPppAvailableKeys(data)
	case NM_SETTING_PPPOE_SETTING_NAME:
		keys = getSettingPppoeAvailableKeys(data)
	case NM_SETTING_VPN_SETTING_NAME:
		keys = getSettingVpnAvailableKeys(data)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		keys = getSettingVpnL2tpAvailableKeys(data)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		keys = getSettingVpnL2tpPppAvailableKeys(data)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		keys = getSettingVpnL2tpIpsecAvailableKeys(data)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		keys = getSettingVpnOpenconnectAvailableKeys(data)
	case NM_SETTING_WIRED_SETTING_NAME:
		keys = getSettingWiredAvailableKeys(data)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		keys = getSettingWirelessAvailableKeys(data)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		keys = getSettingWirelessSecurityAvailableKeys(data)
	}
	return
}

func generalGetSettingAvailableValues(data _ConnectionData, field, key string) (values []string, customizable bool) {
	if isVirtualKey(field, key) {
		values = generalGetSettingVkAvailableValues(data, field, key)
		return
	}
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		values, customizable = getSetting8021xAvailableValues(data, key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		values, customizable = getSettingConnectionAvailableValues(data, key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		values, customizable = getSettingIp4ConfigAvailableValues(data, key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		values, customizable = getSettingIp6ConfigAvailableValues(data, key)
	case NM_SETTING_PPP_SETTING_NAME:
		values, customizable = getSettingPppAvailableValues(data, key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		values, customizable = getSettingPppoeAvailableValues(data, key)
	case NM_SETTING_VPN_SETTING_NAME:
		values, customizable = getSettingVpnAvailableValues(data, key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		values, customizable = getSettingVpnL2tpAvailableValues(data, key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		values, customizable = getSettingVpnL2tpPppAvailableValues(data, key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		values, customizable = getSettingVpnL2tpIpsecAvailableValues(data, key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		values, customizable = getSettingVpnOpenconnectAvailableValues(data, key)
	case NM_SETTING_WIRED_SETTING_NAME:
		values, customizable = getSettingWiredAvailableValues(data, key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		values, customizable = getSettingWirelessAvailableValues(data, key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		values, customizable = getSettingWirelessSecurityAvailableValues(data, key)
	}
	return
}

func generalCheckSettingValues(data _ConnectionData, field string) (errs FieldKeyErrors) {
	switch field {
	default:
		Logger.Error("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		errs = checkSetting8021xValues(data)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		errs = checkSettingConnectionValues(data)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		errs = checkSettingIp4ConfigValues(data)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		errs = checkSettingIp6ConfigValues(data)
	case NM_SETTING_PPP_SETTING_NAME:
		errs = checkSettingPppValues(data)
	case NM_SETTING_PPPOE_SETTING_NAME:
		errs = checkSettingPppoeValues(data)
	case NM_SETTING_VPN_SETTING_NAME:
		errs = checkSettingVpnValues(data)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		errs = checkSettingVpnL2tpValues(data)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		errs = checkSettingVpnL2tpPppValues(data)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		errs = checkSettingVpnL2tpIpsecValues(data)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		errs = checkSettingVpnOpenconnectValues(data)
	case NM_SETTING_WIRED_SETTING_NAME:
		errs = checkSettingWiredValues(data)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		errs = checkSettingWirelessValues(data)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		errs = checkSettingWirelessSecurityValues(data)
	}
	return
}

func generalGetSettingKeyJSON(data _ConnectionData, field, key string) (valueJSON string) {
	if isVirtualKey(field, key) {
		valueJSON = generalGetVirtualKeyJSON(data, field, key)
		return
	}
	switch field {
	default:
		Logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		valueJSON = generalGetSetting8021xKeyJSON(data, key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		valueJSON = generalGetSettingConnectionKeyJSON(data, key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		valueJSON = generalGetSettingIp4ConfigKeyJSON(data, key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		valueJSON = generalGetSettingIp6ConfigKeyJSON(data, key)
	case NM_SETTING_PPP_SETTING_NAME:
		valueJSON = generalGetSettingPppKeyJSON(data, key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		valueJSON = generalGetSettingPppoeKeyJSON(data, key)
	case NM_SETTING_VPN_SETTING_NAME:
		valueJSON = generalGetSettingVpnKeyJSON(data, key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		valueJSON = generalGetSettingVpnL2tpKeyJSON(data, key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		valueJSON = generalGetSettingVpnL2tpPppKeyJSON(data, key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		valueJSON = generalGetSettingVpnL2tpIpsecKeyJSON(data, key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		valueJSON = generalGetSettingVpnOpenconnectKeyJSON(data, key)
	case NM_SETTING_WIRED_SETTING_NAME:
		valueJSON = generalGetSettingWiredKeyJSON(data, key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		valueJSON = generalGetSettingWirelessKeyJSON(data, key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		valueJSON = generalGetSettingWirelessSecurityKeyJSON(data, key)
	}
	return
}

func generalSetSettingKeyJSON(data _ConnectionData, field, key, valueJSON string) {
	if isVirtualKey(field, key) {
		generalSetVirtualKeyJSON(data, field, key, valueJSON)
		return
	}
	switch field {
	default:
		Logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		generalSetSetting8021xKeyJSON(data, key, valueJSON)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		generalSetSettingConnectionKeyJSON(data, key, valueJSON)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		generalSetSettingIp4ConfigKeyJSON(data, key, valueJSON)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		generalSetSettingIp6ConfigKeyJSON(data, key, valueJSON)
	case NM_SETTING_PPP_SETTING_NAME:
		generalSetSettingPppKeyJSON(data, key, valueJSON)
	case NM_SETTING_PPPOE_SETTING_NAME:
		generalSetSettingPppoeKeyJSON(data, key, valueJSON)
	case NM_SETTING_VPN_SETTING_NAME:
		generalSetSettingVpnKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		generalSetSettingVpnL2tpKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		generalSetSettingVpnL2tpPppKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		generalSetSettingVpnL2tpIpsecKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		generalSetSettingVpnOpenconnectKeyJSON(data, key, valueJSON)
	case NM_SETTING_WIRED_SETTING_NAME:
		generalSetSettingWiredKeyJSON(data, key, valueJSON)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		generalSetSettingWirelessKeyJSON(data, key, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		generalSetSettingWirelessSecurityKeyJSON(data, key, valueJSON)
	}
}

func getSettingKeyDefaultValueJSON(field, key string) (valueJSON string) {
	switch field {
	default:
		Logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		valueJSON = getSetting8021xKeyDefaultValueJSON(key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		valueJSON = getSettingConnectionKeyDefaultValueJSON(key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		valueJSON = getSettingIp4ConfigKeyDefaultValueJSON(key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		valueJSON = getSettingIp6ConfigKeyDefaultValueJSON(key)
	case NM_SETTING_PPP_SETTING_NAME:
		valueJSON = getSettingPppKeyDefaultValueJSON(key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		valueJSON = getSettingPppoeKeyDefaultValueJSON(key)
	case NM_SETTING_VPN_SETTING_NAME:
		valueJSON = getSettingVpnKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		valueJSON = getSettingVpnL2tpKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		valueJSON = getSettingVpnL2tpPppKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		valueJSON = getSettingVpnL2tpIpsecKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		valueJSON = getSettingVpnOpenconnectKeyDefaultValueJSON(key)
	case NM_SETTING_WIRED_SETTING_NAME:
		valueJSON = getSettingWiredKeyDefaultValueJSON(key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		valueJSON = getSettingWirelessKeyDefaultValueJSON(key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		valueJSON = getSettingWirelessSecurityKeyDefaultValueJSON(key)
	}
	return
}