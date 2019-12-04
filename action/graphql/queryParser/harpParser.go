package queryParser

import (
	"errors"
	"hcc/piccolo/data"
	"hcc/piccolo/http"
	"strconv"
)

func Subnet(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var subnetData data.SubnetData
	query := "query { subnet(uuid: \"" + uuid + "\") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"

	return http.DoHTTPRequest("harp", true, "SubnetData", subnetData, query)
}

func ListSubnet(args map[string]interface{}) (interface{}, error) {
	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	nextServer, nextServerOk := args["next_server"].(string)
	nameServer, nameServerOk := args["name_server"].(string)
	domainName, domainNameOk := args["domain_name"].(string)
	serverUUID, serverUUIDOk := args["sever_uuid"].(string)
	leaderNodeUUID, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	os, osOk := args["os"].(string)
	subnetName, subnetNameOk := args["subnet_name"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
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
		arguments += "name_server:\"" + nameServer + "\","
	}
	if domainNameOk {
		arguments += "domain_name:\"" + domainName + "\","
	}
	if serverUUIDOk {
		arguments += "sever_uuid:\"" + serverUUID + "\","
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

	var listServerData data.ListServerData
	query := "query { list_subnet(" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"

	return http.DoHTTPRequest("harp", true, "ListServerData", listServerData, query)
}

func AllSubnet(args map[string]interface{}) (interface{}, error) {
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	var query string

	if !rowOk && !pageOk {
		query = "query { all_subnet { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"
	} else if rowOk && pageOk {
		query = "query { all_subnet(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
			") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	var allServerData data.AllServerData

	return http.DoHTTPRequest("harp", true, "AllServerData", allServerData, query)
}

func NumSubnet() (interface{}, error) {
	var numSubnetData data.NumSubnetData
	query := "query { num_subnet { number } }"

	return http.DoHTTPRequest("harp", true, "NumSubnetData", numSubnetData, query)
}

func AdaptiveIP(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var adaptiveIPData data.AdaptiveIPData
	query := "query { adaptiveip(uuid: \"" + uuid + "\") { uuid network_address netmask gateway start_ip_address end_ip_address created_at} }"

	return http.DoHTTPRequest("harp", true, "AdaptiveIPData", adaptiveIPData, query)
}

func ListAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	networkAddress, networkAddressOk := args["network_address"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	startIPaddress, startIPaddressOk := args["start_ip_address"].(string)
	endIPaddress, endIPaddressOk := args["end_ip_address"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	if networkAddressOk {
		arguments += "network_address:\"" + networkAddress + "\","
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

	var listAdaptiveIPData data.ListAdaptiveIPData
	query := "query { list_adaptiveip(" + arguments + ") { uuid network_address netmask gateway start_ip_address end_ip_address created_at} }"

	return http.DoHTTPRequest("harp", true, "ListAdaptiveIPData", listAdaptiveIPData, query)
}

func AllAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	var query string

	if !rowOk && !pageOk {
		query = "query { all_adaptiveip { uuid network_address netmask gateway start_ip_address end_ip_address created_at} }"
	} else if rowOk && pageOk {
		query = "query { all_adaptiveip(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
			") { uuid network_address netmask gateway start_ip_address end_ip_address created_at} }"
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	var allAdaptiveIPData data.AllAdaptiveIPData

	return http.DoHTTPRequest("harp", true, "AllAdaptiveIPData", allAdaptiveIPData, query)
}

func NumAdaptiveIP() (interface{}, error) {
	var numAdaptiveIPData data.NumAdaptiveIPData
	query := "query { num_adaptiveip { number } }"

	return http.DoHTTPRequest("harp", true, "NumAdaptiveIPData", numAdaptiveIPData, query)
}

func AdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	adaptiveIPUUID, adaptiveIPUUIDOk := args["adaptiveip_uuid"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)

	if !adaptiveIPUUIDOk || !serverUUIDOk {
		return nil, errors.New("need adaptiveip_uuid and server_uuid arguments")
	}

	var adaptiveIPServerData data.AdaptiveIPServerData
	query := "query { adaptiveip_server(adaptiveip_uuid: \"" + adaptiveIPUUID + "\", server_uuid: \"" + serverUUID + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"

	return http.DoHTTPRequest("harp", true, "AdaptiveIPServerData", adaptiveIPServerData, query)
}

func ListAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	publicIP, publicIPOk := args["public_ip"].(string)
	privateIP, privateIPOk := args["private_ip"].(string)
	privateGateway, privateGatewayOk := args["private_gateway"].(string)

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	arguments += "server_uuid:\"" + serverUUID + "\","

	if publicIPOk {
		arguments += "public_ip:\"" + publicIP + "\","
	}
	if privateIPOk {
		arguments += "private_ip:\"" + privateIP + "\","
	}
	if privateGatewayOk {
		arguments += "private_gateway:\"" + privateGateway + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var listAdaptiveIPServerData data.ListAdaptiveIPServerData
	query := "query { list_adaptiveip_server(" + arguments + ") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"

	return http.DoHTTPRequest("harp", true, "ListAdaptiveIPServerData", listAdaptiveIPServerData, query)
}

func AllAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	var query string

	if !rowOk && !pageOk {
		query = "query { all_adaptiveip_server(server_uuid: \"" + serverUUID + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"
	} else if rowOk && pageOk {
		query = "query { all_adaptiveip_server(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
			", server_uuid: \"" + serverUUID + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	var allAdaptiveIPServerData data.AllAdaptiveIPServerData

	return http.DoHTTPRequest("harp", true, "AllAdaptiveIPServerData", allAdaptiveIPServerData, query)
}

func NumAdaptiveIPServer() (interface{}, error) {
	var numAdaptiveIPSeverData data.NumAdaptiveIPServerData
	query := "query { num_adaptiveip_server { number } }"

	return http.DoHTTPRequest("harp", true, "NumAdaptiveIPServerData", numAdaptiveIPSeverData, query)
}
