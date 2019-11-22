package mutationParser

import (
	"errors"
	"hcc/piccolo/data"
	"hcc/piccolo/http"
)

func checkSubnetArgsEach(args map[string]interface{}) bool {
	_, networkIPOk := args["network_ip"].(string)
	_, netmaskOk := args["netmask"].(string)
	_, gatewayOk := args["gateway"].(string)
	_, nextServerOk := args["next_server"].(string)
	_, nameServerOk := args["name_server"].(string)
	_, domainNameOk := args["domain_name"].(string)
	_, serverUUIDOk := args["server_uuid"].(string)
	_, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	_, osOk := args["os"].(string)
	_, subnetNameOk := args["subnet_name"].(string)

	return networkIPOk || netmaskOk || gatewayOk || nextServerOk || nameServerOk || domainNameOk || serverUUIDOk || leaderNodeUUIDOk || osOk || subnetNameOk
}

func checkSubnetArgsAll(args map[string]interface{}) bool {
	_, networkIPOk := args["network_ip"].(string)
	_, netmaskOk := args["netmask"].(string)
	_, gatewayOk := args["gateway"].(string)
	_, nextServerOk := args["next_server"].(string)
	_, nameServerOk := args["name_server"].(string)
	_, domainNameOk := args["domain_name"].(string)
	_, serverUUIDOk := args["server_uuid"].(string)
	_, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	_, osOk := args["os"].(string)
	_, subnetNameOk := args["subnet_name"].(string)

	return networkIPOk && netmaskOk && gatewayOk && nextServerOk && nameServerOk && domainNameOk && serverUUIDOk && leaderNodeUUIDOk && osOk && subnetNameOk
}

func CreateSubnet(args map[string]interface{}) (interface{}, error) {
	if !checkSubnetArgsAll(args) {
		return nil, errors.New("check needed arguments (network_ip, netmask, gateway, next_server, name_server, domain_name, server_uuid, leader_node_uuid, os, subnet_name)")
	}

	networkIP, _ := args["network_ip"].(string)
	netmask, _ := args["netmask"].(string)
	gateway, _ := args["gateway"].(string)
	nextServer, _ := args["next_server"].(string)
	nameServer, _ := args["name_server"].(string)
	domainName, _ := args["domain_name"].(string)
	serverUUID, _ := args["server_uuid"].(string)
	leaderNodeUUID, _ := args["leader_node_uuid"].(string)
	os, _ := args["os"].(string)
	subnetName, _ := args["subnet_name"].(string)

	var createSubnetData data.CreateSubnetData
	query := "mutation _ { create_subnet(network_ip: \"" + networkIP + "\", netmask: \"" + netmask + "\", gateway: \"" +
		gateway + "\", next_server: \"" + nextServer + "\", name_server: \"" + nameServer + "\", domain_name: \"" +
		domainName + "\", server_uuid: \"" + serverUUID + "\", leader_node_uuid: \"" + leaderNodeUUID + "\", os: \"" +
		os + "\", subnet_name: \"" + subnetName + "\") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	return http.DoHTTPRequest("harp", true, "CreateSubnetData", createSubnetData, query)
}

func UpdateSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if checkSubnetArgsEach(args) {
		return nil, errors.New("need some arguments")
	}

	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	nextServer, nextServerOk := args["next_server"].(string)
	nameServer, nameServerOk := args["name_server"].(string)
	domainName, domainNameOk := args["domain_name"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	leaderNodeUUID, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	os, osOk := args["os"].(string)
	subnetName, subnetNameOk := args["subnet_name"].(string)

	arguments := "uuid:\"" + requestedUUID + "\""
	if networkIPOk {
		arguments += "network_ip:\"" + networkIP + "\","
	}
	if netmaskOk {
		arguments += "netmask:\"" + netmask + "\","
	}
	if gatewayOk {
		arguments += "gateway:\"" + gateway + "\","
	}
	if nextServerOk {
		arguments += "next_server:\"" + nextServer + "\","
	}
	if nameServerOk {
		arguments += "next_server:\"" + nameServer + "\","
	}
	if domainNameOk {
		arguments += "domain_name:\"" + domainName + "\","
	}
	if serverUUIDOk {
		arguments += "server_uuid:\"" + serverUUID + "\","
	}
	if leaderNodeUUIDOk {
		arguments += "leader_node_uuid:\"" + leaderNodeUUID + "\","
	}
	if osOk {
		arguments += "os:\"" + os + "\","
	}
	if subnetNameOk {
		arguments += "subnet_name:\"" + subnetName + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var updateSubnetData data.UpdateSubnetData
	query := "mutation _ { update_subnet(" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	return http.DoHTTPRequest("harp", true, "UpdateSubnetData", updateSubnetData, query)
}

func DeleteSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var deleteSubnetData data.DeleteSubnetData
	query := "mutation _ { delete_subnet(uuid:\"" + requestedUUID + "\") { uuid } }"

	return http.DoHTTPRequest("harp", true, "DeleteSubnetData", deleteSubnetData, query)
}

func CreateDHCPDConf(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	node_uuids, node_uuidsOk := args["node_uuids"].(string)
	if !subnetUUIDOk || !node_uuidsOk {
		return nil, errors.New("need subnet_uuid and node_uuids arguments")
	}

	query := "mutation _ { create_dhcpd_conf(subnet_uuid: \"" + subnetUUID + "\", node_uuids: \"" + node_uuids + "\") }"

	return http.DoHTTPRequest("harp", false, "", nil, query)
}

func checkAdaptiveIPArgsEach(args map[string]interface{}) bool {
	_, networkAddressOk := args["network_address"].(string)
	_, netmaskOk := args["netmask"].(string)
	_, gatewayOk := args["gateway"].(string)
	_, startIPaddressOk := args["start_ip_address"].(string)
	_, endIPaddressOk := args["end_ip_address"].(string)

	return networkAddressOk || netmaskOk || gatewayOk || startIPaddressOk || endIPaddressOk
}

func checkAdaptiveIPArgsAll(args map[string]interface{}) bool {
	_, networkAddressOk := args["network_address"].(string)
	_, netmaskOk := args["netmask"].(string)
	_, gatewayOk := args["gateway"].(string)
	_, startIPaddressOk := args["start_ip_address"].(string)
	_, endIPaddressOk := args["end_ip_address"].(string)

	return networkAddressOk && netmaskOk && gatewayOk && startIPaddressOk && endIPaddressOk
}

func CreateAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	if !checkAdaptiveIPArgsAll(args) {
		return nil, errors.New("check needed arguments (network_ip, netmask, gateway, next_server, name_server, domain_name, server_uuid, leader_node_uuid, os, adaptiveip_name)")
	}

	networkAddress, _ := args["network_address"].(string)
	netmask, _ := args["netmask"].(string)
	gateway, _ := args["gateway"].(string)
	startIPaddressOk, _ := args["start_ip_address"].(string)
	endIPaddressOk, _ := args["end_ip_address"].(string)

	var createAdaptiveIPData data.CreateAdaptiveIPData
	query := "mutation _ { create_adaptiveip(network_address: \"" + networkAddress + "\", netmask: \"" + netmask + "\", gateway: \"" +
		gateway + "\", start_ip_address: \"" + startIPaddressOk + "\", end_ip_address: \"" + endIPaddressOk + "\") { uuid network_address netmask gateway start_ip_address end_ip_address } }"

	return http.DoHTTPRequest("harp", true, "CreateAdaptiveIPData", createAdaptiveIPData, query)
}

func UpdateAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if checkAdaptiveIPArgsEach(args) {
		return nil, errors.New("need some arguments")
	}

	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	startIPaddress, startIPaddressOk := args["start_ip_address"].(string)
	endIPaddress, endIPaddressOk := args["end_ip_address"].(string)

	arguments := "uuid:\"" + requestedUUID + "\""
	if networkIPOk {
		arguments += "network_ip:\"" + networkIP + "\","
	}
	if netmaskOk {
		arguments += "netmask:\"" + netmask + "\","
	}
	if gatewayOk {
		arguments += "gateway:\"" + gateway + "\","
	}
	if startIPaddressOk {
		arguments += "start_ip_address:\"" + startIPaddress + "\","
	}
	if endIPaddressOk {
		arguments += "end_ip_address:\"" + endIPaddress + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var updateAdaptiveIPData data.UpdateAdaptiveIPData
	query := "mutation _ { update_adaptiveip(" + arguments + ") { uuid network_address netmask gateway start_ip_address end_ip_address } }"

	return http.DoHTTPRequest("harp", true, "UpdateAdaptiveIPData", updateAdaptiveIPData, query)
}

func DeleteAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var deleteAdaptiveIPData data.DeleteAdaptiveIPData
	query := "mutation _ { delete_adaptiveip(uuid:\"" + requestedUUID + "\") { uuid } }"

	return http.DoHTTPRequest("harp", true, "DeleteAdaptiveIPData", deleteAdaptiveIPData, query)
}

func checkAdaptiveIPServerArgsAll(args map[string]interface{}) bool {
	_, adaptiveIPUUIDOk := args["adaptiveip_uuid"].(string)
	_, serverUUIDOk := args["server_uuid"].(string)
	_, publicIPOk := args["public_ip"].(string)

	return adaptiveIPUUIDOk && serverUUIDOk && publicIPOk
}

func CreateAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	if !checkAdaptiveIPServerArgsAll(args) {
		return nil, errors.New("check needed arguments (adaptiveip_uuid, server_uuid, public_ip)")
	}

	adaptiveIPUUID, _ := args["adaptiveip_uuid"].(string)
	serverUUID, _ := args["server_uuid"].(string)
	publicIP, _ := args["public_ip"].(string)

	var createAdaptiveIPServerData data.CreateAdaptiveIPServerData
	query := "mutation _ { create_adaptiveip_server(adaptiveip_uuid: \"" + adaptiveIPUUID + "\", server_uuid: \"" + serverUUID + "\", public_ip: \"" +
		publicIP + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway } }"

	return http.DoHTTPRequest("harp", true, "CreateAdaptiveIPServerData", createAdaptiveIPServerData, query)
}

func DeleteAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["server_uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	var deleteAdaptiveIPServerData data.DeleteAdaptiveIPServerData
	query := "mutation _ { delete_adaptiveip_server(server_uuid:\"" + requestedUUID + "\") { server_uuid } }"

	return http.DoHTTPRequest("harp", true, "DeleteAdaptiveIPServerData", deleteAdaptiveIPServerData, query)
}
