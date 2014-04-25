// This file is automatically generated, please don't edit manully.
package main

// All virtual keys data
var virtualKeys = []VirtualKey{
	VirtualKey{NM_SETTING_VK_802_1X_ENABLE, ktypeBoolean, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_VK_NONE_RELATED_KEY, true, false},
	VirtualKey{NM_SETTING_VK_802_1X_EAP, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_EAP, true, false},
	VirtualKey{NM_SETTING_VK_802_1X_PAC_FILE, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_PAC_FILE, true, false},
	VirtualKey{NM_SETTING_VK_802_1X_CA_CERT, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_CA_CERT, true, false},
	VirtualKey{NM_SETTING_VK_802_1X_CLIENT_CERT, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_CLIENT_CERT, true, false},
	VirtualKey{NM_SETTING_VK_802_1X_PRIVATE_KEY, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_PRIVATE_KEY, true, false},
	VirtualKey{NM_SETTING_VK_CONNECTION_NO_PERMISSION, ktypeBoolean, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, true, false},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_ADDRESSES_ADDRESS, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, true, false},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_ADDRESSES_MASK, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, true, false},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_ADDRESSES_GATEWAY, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, true, true},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_DNS, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, true, false},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_ADDRESS, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_MASK, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_NEXTHOP, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_METRIC, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_ADDRESSES_ADDRESS, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, true, false},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_ADDRESSES_PREFIX, ktypeUint32, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, true, false},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_ADDRESSES_GATEWAY, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, true, true},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_DNS, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, true, false},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_ADDRESS, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_PREFIX, ktypeUint32, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_NEXTHOP, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_METRIC, ktypeUint32, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false},
	VirtualKey{NM_SETTING_VK_PPP_LCP_ECHO_ENABLE, ktypeBoolean, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE, true, false},
	VirtualKey{NM_SETTING_VK_VPN_L2TP_LCP_ECHO_ENABLE, ktypeBoolean, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE, true, false},
	VirtualKey{NM_SETTING_VK_WIRELESS_SECURITY_KEY_MGMT, ktypeString, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, true, false},
}

// Get JSON value generally
func generalGetVirtualKeyJSON(data _ConnectionData, field, key string) (valueJSON string) {
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_802_1X_ENABLE:
			return getSettingVk8021xEnableJSON(data)
		case NM_SETTING_VK_802_1X_EAP:
			return getSettingVk8021xEapJSON(data)
		case NM_SETTING_VK_802_1X_PAC_FILE:
			return getSettingVk8021xPacFileJSON(data)
		case NM_SETTING_VK_802_1X_CA_CERT:
			return getSettingVk8021xCaCertJSON(data)
		case NM_SETTING_VK_802_1X_CLIENT_CERT:
			return getSettingVk8021xClientCertJSON(data)
		case NM_SETTING_VK_802_1X_PRIVATE_KEY:
			return getSettingVk8021xPrivateKeyJSON(data)
		}
	case NM_SETTING_CONNECTION_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_CONNECTION_NO_PERMISSION:
			return getSettingVkConnectionNoPermissionJSON(data)
		}
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_ADDRESS:
			return getSettingVkIp4ConfigAddressesAddressJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_MASK:
			return getSettingVkIp4ConfigAddressesMaskJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_GATEWAY:
			return getSettingVkIp4ConfigAddressesGatewayJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_DNS:
			return getSettingVkIp4ConfigDnsJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_ADDRESS:
			return getSettingVkIp4ConfigRoutesAddressJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_MASK:
			return getSettingVkIp4ConfigRoutesMaskJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_NEXTHOP:
			return getSettingVkIp4ConfigRoutesNexthopJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_METRIC:
			return getSettingVkIp4ConfigRoutesMetricJSON(data)
		}
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_ADDRESS:
			return getSettingVkIp6ConfigAddressesAddressJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_PREFIX:
			return getSettingVkIp6ConfigAddressesPrefixJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_GATEWAY:
			return getSettingVkIp6ConfigAddressesGatewayJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_DNS:
			return getSettingVkIp6ConfigDnsJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_ADDRESS:
			return getSettingVkIp6ConfigRoutesAddressJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_PREFIX:
			return getSettingVkIp6ConfigRoutesPrefixJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_NEXTHOP:
			return getSettingVkIp6ConfigRoutesNexthopJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_METRIC:
			return getSettingVkIp6ConfigRoutesMetricJSON(data)
		}
	case NM_SETTING_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_PPP_LCP_ECHO_ENABLE:
			return getSettingVkPppLcpEchoEnableJSON(data)
		}
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_L2TP_LCP_ECHO_ENABLE:
			return getSettingVkVpnL2tpLcpEchoEnableJSON(data)
		}
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_WIRELESS_SECURITY_KEY_MGMT:
			return getSettingVkWirelessSecurityKeyMgmtJSON(data)
		}
	}
	Logger.Error("invalid virtual key:", field, key)
	return
}

// Set JSON value generally
func generalSetVirtualKeyJSON(data _ConnectionData, field, key string, valueJSON string) {
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_802_1X_ENABLE:
			logicSetSettingVk8021xEnableJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_EAP:
			logicSetSettingVk8021xEapJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_PAC_FILE:
			logicSetSettingVk8021xPacFileJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_CA_CERT:
			logicSetSettingVk8021xCaCertJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_CLIENT_CERT:
			logicSetSettingVk8021xClientCertJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_PRIVATE_KEY:
			logicSetSettingVk8021xPrivateKeyJSON(data, valueJSON)
			return
		}
	case NM_SETTING_CONNECTION_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_CONNECTION_NO_PERMISSION:
			logicSetSettingVkConnectionNoPermissionJSON(data, valueJSON)
			return
		}
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_ADDRESS:
			logicSetSettingVkIp4ConfigAddressesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_MASK:
			logicSetSettingVkIp4ConfigAddressesMaskJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_GATEWAY:
			logicSetSettingVkIp4ConfigAddressesGatewayJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_DNS:
			logicSetSettingVkIp4ConfigDnsJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_ADDRESS:
			logicSetSettingVkIp4ConfigRoutesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_MASK:
			logicSetSettingVkIp4ConfigRoutesMaskJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_NEXTHOP:
			logicSetSettingVkIp4ConfigRoutesNexthopJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_METRIC:
			logicSetSettingVkIp4ConfigRoutesMetricJSON(data, valueJSON)
			return
		}
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_ADDRESS:
			logicSetSettingVkIp6ConfigAddressesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_PREFIX:
			logicSetSettingVkIp6ConfigAddressesPrefixJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_GATEWAY:
			logicSetSettingVkIp6ConfigAddressesGatewayJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_DNS:
			logicSetSettingVkIp6ConfigDnsJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_ADDRESS:
			logicSetSettingVkIp6ConfigRoutesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_PREFIX:
			logicSetSettingVkIp6ConfigRoutesPrefixJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_NEXTHOP:
			logicSetSettingVkIp6ConfigRoutesNexthopJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_METRIC:
			logicSetSettingVkIp6ConfigRoutesMetricJSON(data, valueJSON)
			return
		}
	case NM_SETTING_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_PPP_LCP_ECHO_ENABLE:
			logicSetSettingVkPppLcpEchoEnableJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_L2TP_LCP_ECHO_ENABLE:
			logicSetSettingVkVpnL2tpLcpEchoEnableJSON(data, valueJSON)
			return
		}
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_WIRELESS_SECURITY_KEY_MGMT:
			logicSetSettingVkWirelessSecurityKeyMgmtJSON(data, valueJSON)
			return
		}
	}
	Logger.Error("invalid virtual key:", field, key)
	return
}

// JSON getter
func getSettingVk8021xEnableJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xEnable(data))
	return
}
func getSettingVk8021xEapJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xEap(data))
	return
}
func getSettingVk8021xPacFileJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xPacFile(data))
	return
}
func getSettingVk8021xCaCertJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xCaCert(data))
	return
}
func getSettingVk8021xClientCertJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xClientCert(data))
	return
}
func getSettingVk8021xPrivateKeyJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xPrivateKey(data))
	return
}
func getSettingVkConnectionNoPermissionJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkConnectionNoPermission(data))
	return
}
func getSettingVkIp4ConfigAddressesAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigAddressesAddress(data))
	return
}
func getSettingVkIp4ConfigAddressesMaskJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigAddressesMask(data))
	return
}
func getSettingVkIp4ConfigAddressesGatewayJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigAddressesGateway(data))
	return
}
func getSettingVkIp4ConfigDnsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigDns(data))
	return
}
func getSettingVkIp4ConfigRoutesAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesAddress(data))
	return
}
func getSettingVkIp4ConfigRoutesMaskJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesMask(data))
	return
}
func getSettingVkIp4ConfigRoutesNexthopJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesNexthop(data))
	return
}
func getSettingVkIp4ConfigRoutesMetricJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesMetric(data))
	return
}
func getSettingVkIp6ConfigAddressesAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigAddressesAddress(data))
	return
}
func getSettingVkIp6ConfigAddressesPrefixJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigAddressesPrefix(data))
	return
}
func getSettingVkIp6ConfigAddressesGatewayJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigAddressesGateway(data))
	return
}
func getSettingVkIp6ConfigDnsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigDns(data))
	return
}
func getSettingVkIp6ConfigRoutesAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesAddress(data))
	return
}
func getSettingVkIp6ConfigRoutesPrefixJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesPrefix(data))
	return
}
func getSettingVkIp6ConfigRoutesNexthopJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesNexthop(data))
	return
}
func getSettingVkIp6ConfigRoutesMetricJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesMetric(data))
	return
}
func getSettingVkPppLcpEchoEnableJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkPppLcpEchoEnable(data))
	return
}
func getSettingVkVpnL2tpLcpEchoEnableJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnL2tpLcpEchoEnable(data))
	return
}
func getSettingVkWirelessSecurityKeyMgmtJSON(data _ConnectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkWirelessSecurityKeyMgmt(data))
	return
}

// Logic JSON setter
func logicSetSettingVk8021xEnableJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	logicSetSettingVk8021xEnable(data, value)
}
func logicSetSettingVk8021xEapJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVk8021xEap(data, value)
}
func logicSetSettingVk8021xPacFileJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVk8021xPacFile(data, value)
}
func logicSetSettingVk8021xCaCertJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVk8021xCaCert(data, value)
}
func logicSetSettingVk8021xClientCertJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVk8021xClientCert(data, value)
}
func logicSetSettingVk8021xPrivateKeyJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVk8021xPrivateKey(data, value)
}
func logicSetSettingVkConnectionNoPermissionJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	logicSetSettingVkConnectionNoPermission(data, value)
}
func logicSetSettingVkIp4ConfigAddressesAddressJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigAddressesAddress(data, value)
}
func logicSetSettingVkIp4ConfigAddressesMaskJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigAddressesMask(data, value)
}
func logicSetSettingVkIp4ConfigAddressesGatewayJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigAddressesGateway(data, value)
}
func logicSetSettingVkIp4ConfigDnsJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigDns(data, value)
}
func logicSetSettingVkIp4ConfigRoutesAddressJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigRoutesAddress(data, value)
}
func logicSetSettingVkIp4ConfigRoutesMaskJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigRoutesMask(data, value)
}
func logicSetSettingVkIp4ConfigRoutesNexthopJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigRoutesNexthop(data, value)
}
func logicSetSettingVkIp4ConfigRoutesMetricJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp4ConfigRoutesMetric(data, value)
}
func logicSetSettingVkIp6ConfigAddressesAddressJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp6ConfigAddressesAddress(data, value)
}
func logicSetSettingVkIp6ConfigAddressesPrefixJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueUint32(valueJSON)
	logicSetSettingVkIp6ConfigAddressesPrefix(data, value)
}
func logicSetSettingVkIp6ConfigAddressesGatewayJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp6ConfigAddressesGateway(data, value)
}
func logicSetSettingVkIp6ConfigDnsJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp6ConfigDns(data, value)
}
func logicSetSettingVkIp6ConfigRoutesAddressJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp6ConfigRoutesAddress(data, value)
}
func logicSetSettingVkIp6ConfigRoutesPrefixJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueUint32(valueJSON)
	logicSetSettingVkIp6ConfigRoutesPrefix(data, value)
}
func logicSetSettingVkIp6ConfigRoutesNexthopJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkIp6ConfigRoutesNexthop(data, value)
}
func logicSetSettingVkIp6ConfigRoutesMetricJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueUint32(valueJSON)
	logicSetSettingVkIp6ConfigRoutesMetric(data, value)
}
func logicSetSettingVkPppLcpEchoEnableJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	logicSetSettingVkPppLcpEchoEnable(data, value)
}
func logicSetSettingVkVpnL2tpLcpEchoEnableJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	logicSetSettingVkVpnL2tpLcpEchoEnable(data, value)
}
func logicSetSettingVkWirelessSecurityKeyMgmtJSON(data _ConnectionData, valueJSON string) {
	value, _ := jsonToKeyValueString(valueJSON)
	logicSetSettingVkWirelessSecurityKeyMgmt(data, value)
}