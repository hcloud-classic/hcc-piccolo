package queryParser

import (
	"errors"
	"hcc/piccolo/data"
	"hcc/piccolo/http"
	"strconv"
)

func PowerStateNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var powerStateNodeData data.PowerStateNodeData
	query := "query { power_state_node(uuid: \"" + uuid + "\") }"

	return http.DoHTTPRequest("flute", true, "PowerStateNodeData", powerStateNodeData, query)
}

func Node(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var nodeData data.NodeData
	query := "query { node(uuid: \"" + uuid + "\") { uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"

	return http.DoHTTPRequest("flute", true, "NodeData", nodeData, query)
}

func ListNode(args map[string]interface{}) (interface{}, error) {
	var noLimit bool

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAdr, pxeMacAdrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	description, descriptionOk := args["description"].(string)
	active, activeOk := args["active"].(string)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	if !rowOk && !pageOk {
		noLimit = true
	} else if rowOk && pageOk {
		noLimit = false
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	var arguments string

	if noLimit {
		arguments = ""
	} else {
		arguments = "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	}

	if serverUUIDOk {
		arguments += "server_uuid:\"" + serverUUID + "\","
	}
	if bmcMacAddrOk {
		arguments += "bmc_mac_addr:\"" + bmcMacAddr + "\","
	}
	if bmcIPOk {
		arguments += "bmc_ip:\"" + bmcIP + "\","
	}
	if pxeMacAdrOk {
		arguments += "pxe_mac_addr:\"" + pxeMacAdr + "\","
	}
	if statusOk {
		arguments += "status:\"" + status + "\","
	}
	if cpuCoresOk {
		arguments += "cpu_cores:" + strconv.Itoa(cpuCores) + "\","
	}
	if memoryOk {
		arguments += "memory:" + strconv.Itoa(memory) + "\","
	}
	if descriptionOk {
		arguments += "description:\"" + description + "\","
	}
	if activeOk {
		arguments += "active:\"" + active + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var listNodeData data.ListNodeData
	query := "query { list_node(" + arguments + ") { uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"

	return http.DoHTTPRequest("flute", true, "ListNodeData", listNodeData, query)
}

func AllNode(args map[string]interface{}) (interface{}, error) {
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	active, activeOk := args["active"].(int)
	var query string

	if !rowOk && !pageOk {
		query = "query { all_node { uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"
		if activeOk {
			query = "query { all_node(active:" + strconv.Itoa(active) + ") { uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"
		}
	} else if rowOk && pageOk {
		query = "query { all_node(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
			") { uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"
		if activeOk {
			query = "query { all_node(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
				", active:" + strconv.Itoa(active) + ") { uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"
		}
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	var allNodeData data.AllNodeData

	return http.DoHTTPRequest("flute", true, "AllNodeData", allNodeData, query)
}

func NumNode() (interface{}, error) {
	var numNodeData data.NumNodeData
	query := "query { num_node { number } }"

	return http.DoHTTPRequest("flute", true, "NumNodeData", numNodeData, query)
}

func NodeDetail(args map[string]interface{}) (interface{}, error) {
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)

	if !nodeUUIDOk {
		return nil, errors.New("need a node_uuid argument")
	}

	var nodeDetailData data.NodeDetailData
	query := "query { detail_node(node_uuid: \"" + nodeUUID + "\") { node_uuid cpu_model cpu_processors cpu_threads } }"

	return http.DoHTTPRequest("flute", true, "NodeDetailData", nodeDetailData, query)
}
