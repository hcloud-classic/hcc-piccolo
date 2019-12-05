package mutationParser

import (
	"errors"
	"hcc/piccolo/data"
	"hcc/piccolo/http"
	"strconv"
)

func checkServerArgsEach(args map[string]interface{}) bool {
	_, subnetUUIDOk := args["subnet_uuid"].(string)
	_, osOk := args["os"].(string)
	_, serverNameOk := args["server_name"].(string)
	_, serverDescOk := args["server_desc"].(string)
	_, cpuOk := args["cpu"].(int)
	_, memoryOk := args["memory"].(int)
	_, diskSizeOk := args["disk_size"].(int)
	_, statusOk := args["status"].(string)
	_, userUUIDOk := args["user_uuid"].(string)

	return subnetUUIDOk || osOk || serverNameOk || serverDescOk || cpuOk || memoryOk || diskSizeOk || statusOk || userUUIDOk
}

func checkServerArgsAll(args map[string]interface{}) bool {
	_, subnetUUIDOk := args["subnet_uuid"].(string)
	_, osOk := args["os"].(string)
	_, serverNameOk := args["server_name"].(string)
	_, serverDescOk := args["server_desc"].(string)
	_, cpuOk := args["cpu"].(int)
	_, memoryOk := args["memory"].(int)
	_, diskSizeOk := args["disk_size"].(int)
	_, userUUIDOk := args["user_uuid"].(string)
	_, nrNodeOk := args["nr_node"].(int)

	return subnetUUIDOk && osOk && serverNameOk && serverDescOk && cpuOk && memoryOk && diskSizeOk && userUUIDOk && nrNodeOk
}

func CreateServer(args map[string]interface{}) (interface{}, error) {
	if !checkServerArgsAll(args) {
		return nil, errors.New("check needed arguments (subnet_uuid, os, server_name, server_desc, cpu, memory, disk_size, user_uuid, nr_node)")
	}

	subnetUUID, _ := args["subnet_uuid"].(string)
	os, _ := args["os"].(string)
	serverName, _ := args["server_name"].(string)
	serverDesc, _ := args["server_desc"].(string)
	cpu, _ := args["cpu"].(int)
	memory, _ := args["memory"].(int)
	diskSize, _ := args["disk_size"].(int)
	userUUID, _ := args["user_uuid"].(string)
	nrNode, _ := args["nr_node"].(int)

	var createServerData data.CreateServerData
	query := "mutation _ { create_server(subnet_uuid: \"" + subnetUUID + "\", os: \"" + os + "\", server_name: \"" +
		serverName + "\", server_desc: \"" + serverDesc + "\", cpu: " + strconv.Itoa(cpu) + ", memory: " +
		strconv.Itoa(memory) + ", disk_size: " + strconv.Itoa(diskSize) + ", user_uuid: \"" +
		userUUID + "\", nr_node: " + strconv.Itoa(nrNode) + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size user_uuid } }"

	return http.DoHTTPRequest("violin", true, "CreateServerData", createServerData, query)
}

func UpdateServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if checkServerArgsEach(args) {
		return nil, errors.New("need some arguments")
	}

	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOk := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	status, statusOk := args["status"].(string)
	userUUID, userUUIDOk := args["user_uuid"].(string)

	arguments := "uuid:\"" + requestedUUID + "\""
	if subnetUUIDOk {
		arguments += "subnet_uuid:\"" + subnetUUID + "\","
	}
	if osOk {
		arguments += "os:\"" + os + "\","
	}
	if serverNameOk {
		arguments += "server_name:\"" + serverName + "\","
	}
	if serverDescOk {
		arguments += "server_desc:\"" + serverDesc + "\","
	}
	if cpuOk {
		arguments += "cpu:" + strconv.Itoa(cpu) + ","
	}
	if memoryOk {
		arguments += "memory:" + strconv.Itoa(memory) + "\","
	}
	if diskSizeOk {
		arguments += "disk_size:" + strconv.Itoa(diskSize) + "\","
	}
	if statusOk {
		arguments += "status:\"" + status + "\","
	}
	if userUUIDOk {
		arguments += "user_uuid:\"" + userUUID + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var updateServerData data.UpdateServerData
	query := "mutation _ { update_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	return http.DoHTTPRequest("violin", true, "UpdateServerData", updateServerData, query)
}

func DeleteServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	status, statusOk := args["status"].(string)
	if !requestedUUIDOk || !statusOk {
		return nil, errors.New("need a uuid ans status argument")
	}

	var deleteServerData data.DeleteServerData
	query := "mutation _ { delete_server(uuid:\"" + requestedUUID + "\", status: \"" + status + "\") { uuid } }"

	return http.DoHTTPRequest("violin", true, "DeleteServerData", deleteServerData, query)
}

func CreateServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)

	if !serverUUIDOk || !nodeUUIDOk {
		return nil, errors.New("need server_uuid and node_uuid arguments")
	}

	var createServerNodeData data.CreateServerNodeData
	query := "mutation _ { create_server_node(server_uuid: \"" + serverUUID + "\", node_uuid: \"" + nodeUUID +
		"\") { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "CreateServerNodeData", createServerNodeData, query)
}

func DeleteServerNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var deleteServerNodeData data.DeleteServerNodeData
	query := "mutation _ { delete_server_node(uuid:\"" + requestedUUID + "\") { uuid } }"

	return http.DoHTTPRequest("violin", true, "DeleteServerNodeData", deleteServerNodeData, query)
}
