package queryParser

import (
	"errors"
	"hcc/piccolo/data"
	"hcc/piccolo/http"
	"strconv"
)

func Server(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var serverData data.ServerData
	query := "query { server(uuid: \"" + uuid + "\") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "ServerData", serverData, query)
}

func ListServer(args map[string]interface{}) (interface{}, error) {
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOk := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	status, statusOk := args["status"].(string)
	userUUID, userUUIDOk := args["user_uuid"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
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

	var listServerData data.ListServerData
	query := "query { list_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	return http.DoHTTPRequest("violin", true, "ListServerData", listServerData, query)
}

func AllServer(args map[string]interface{}) (interface{}, error) {
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	var query string

	if !rowOk && !pageOk {
		query = "query { all_server { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid created_at } }"
	} else if rowOk && pageOk {
		query = "query { all_server(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
			") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid created_at } }"
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	var allServerData data.AllServerData

	return http.DoHTTPRequest("violin", true, "AllServerData", allServerData, query)
}

func NumServer() (interface{}, error) {
	var numServerData data.NumServerData
	query := "query { num_server { number } }"

	return http.DoHTTPRequest("violin", true, "NumServerData", numServerData, query)
}

func ServerNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var serverNodeData data.ServerNodeData
	query := "query { server_node(uuid: \"" + uuid + "\") { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "ServerNodeData", serverNodeData, query)
}

func ListServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	var listServerNodeData data.ListServerNodeData
	query := "query { list_server_node(server_uuid: \"" + serverUUID + "\") { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "ListServerNodeData", listServerNodeData, query)
}

func AllServerNode() (interface{}, error) {
	var allServerNodeData data.AllServerNodeData
	query := "query { all_server_node { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "AllServerNodeData", allServerNodeData, query)
}
