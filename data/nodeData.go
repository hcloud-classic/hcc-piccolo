package data

import "hcc/piccolo/model"

// Query

// NodeData : Data structure of node
type NodeData struct {
	Data struct {
		Node model.Node `json:"node"`
	} `json:"data"`
}

// ListNodeData : Data structure of list_node
type ListNodeData struct {
	Data struct {
		ListNode []model.Node `json:"list_node"`
	} `json:"data"`
}

// AllNodeData : Data structure of all_node
type AllNodeData struct {
	Data struct {
		AllNode []model.Node `json:"all_node"`
	} `json:"data"`
}

// NumNodeData : Data structure of num_node
type NumNodeData struct {
	Data struct {
		NumNode model.NodeNum `json:"num_node"`
	} `json:"data"`
}

// NodeDetailData : Data structure of detail_node
type NodeDetailData struct {
	Data struct {
		NodeDetail model.NodeDetail `json:"detail_node"`
	} `json:"data"`
}

// Mutation

// CreateNodeData : Data structure of create_node
type CreateNodeData struct {
	Data struct {
		Node model.Node `json:"create_node"`
	} `json:"data"`
}

// UpdateNodeData : Data structure of update_node
type UpdateNodeData struct {
	Data struct {
		Node model.Node `json:"update_node"`
	} `json:"data"`
}

// DeleteNodeData : Data structure of delete_node
type DeleteNodeData struct {
	Data struct {
		Node model.NodeDetail `json:"delete_node"`
	} `json:"data"`
}

// CreateNodeDetailData : Data structure of create_node_detail
type CreateNodeDetailData struct {
	Data struct {
		NodeDetail model.NodeDetail `json:"create_node_detail"`
	} `json:"data"`
}

// DeleteNodeDetailData : Data structure of delete_node_detail
type DeleteNodeDetailData struct {
	Data struct {
		NodeDetail model.Node `json:"delete_node_detail"`
	} `json:"data"`
}
