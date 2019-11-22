package http

import (
	"encoding/json"
	"errors"
	hccGatewayData "hcc/piccolo/data"
	"hcc/piccolo/lib/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// DoHTTPRequest : Send http request to other modules with GraphQL query string.
func DoHTTPRequest(moduleName string, needData bool, dataType string, data interface{}, query string) (interface{}, error) {
	var timeout time.Duration
	var url = "http://"
	switch moduleName {
	case "flute":
		timeout = time.Duration(config.Flute.RequestTimeoutMs)
		url += config.Flute.ServerAddress + ":" + strconv.Itoa(int(config.Flute.ServerPort))
		break
	case "harp":
		timeout = time.Duration(config.Harp.RequestTimeoutMs)
		url += config.Harp.ServerAddress + ":" + strconv.Itoa(int(config.Harp.ServerPort))
		break
	case "cello":
		timeout = time.Duration(config.Cello.RequestTimeoutMs)
		url += config.Cello.ServerAddress + ":" + strconv.Itoa(int(config.Cello.ServerPort))
		break
	case "violin":
		timeout = time.Duration(config.Violin.RequestTimeoutMs)
		url += config.Violin.ServerAddress + ":" + strconv.Itoa(int(config.Violin.ServerPort))
		break
	case "piano":
		timeout = time.Duration(config.Piano.RequestTimeoutMs)
		url += config.Piano.ServerAddress + ":" + strconv.Itoa(int(config.Piano.ServerPort))
		break
	default:
		return nil, errors.New("unknown module name")
	}
	url += "/graphql?query=" + queryURLEncoder(query)

	client := &http.Client{Timeout: timeout * time.Millisecond}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// Check response
		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			result := string(respBody)

			if strings.Contains(result, "errors") {
				return nil, errors.New(result)
			}

			if needData {
				if data == nil {
					return nil, errors.New("needData marked as true but data is nil")
				}

				switch dataType {
				case "NodeData":
					nodeData := data.(hccGatewayData.NodeData)
					err = json.Unmarshal([]byte(result), &nodeData)
					if err != nil {
						return nil, err
					}

					return nodeData.Data.Node, nil
				case "AllNodeData":
					allNodeData := data.(hccGatewayData.AllNodeData)
					err = json.Unmarshal([]byte(result), &allNodeData)
					if err != nil {
						return nil, err
					}

					return allNodeData.Data.AllNode, nil
				case "NumNodeData":
					numNodeData := data.(hccGatewayData.NumNodeData)
					err = json.Unmarshal([]byte(result), &numNodeData)
					if err != nil {
						return nil, err
					}

					return numNodeData.Data.NumNode, nil
				case "NodeDetailData":
					nodeDetailData := data.(hccGatewayData.NodeDetailData)
					err = json.Unmarshal([]byte(result), &nodeDetailData)
					if err != nil {
						return nil, err
					}

					return nodeDetailData.Data.NodeDetail, nil
				case "OnNodeData":
					onNodeData := data.(hccGatewayData.OnNodeData)
					err = json.Unmarshal([]byte(result), &onNodeData)
					if err != nil {
						return nil, err
					}

					return onNodeData.Data.Result, nil
				case "CreateNodeData":
					createNodeData := data.(hccGatewayData.CreateNodeData)
					err = json.Unmarshal([]byte(result), &createNodeData)
					if err != nil {
						return nil, err
					}

					return createNodeData.Data.Node, nil
				case "UpdateNodeData":
					updateNodeData := data.(hccGatewayData.UpdateNodeData)
					err = json.Unmarshal([]byte(result), &updateNodeData)
					if err != nil {
						return nil, err
					}

					return updateNodeData.Data.Node, nil
				case "DeleteNodeData":
					deleteNodeData := data.(hccGatewayData.DeleteNodeData)
					err = json.Unmarshal([]byte(result), &deleteNodeData)
					if err != nil {
						return nil, err
					}

					return deleteNodeData.Data.Node, nil
				case "CreateNodeDetailData":
					createNodeDetailData := data.(hccGatewayData.CreateNodeDetailData)
					err = json.Unmarshal([]byte(result), &createNodeDetailData)
					if err != nil {
						return nil, err
					}

					return createNodeDetailData.Data.NodeDetail, nil
				case "DeleteNodeDetailData":
					deleteNodeDetailData := data.(hccGatewayData.DeleteNodeDetailData)
					err = json.Unmarshal([]byte(result), &deleteNodeDetailData)
					if err != nil {
						return nil, err
					}

					return deleteNodeDetailData.Data.NodeDetail, nil
				case "ServerData":
					serverData := data.(hccGatewayData.ServerData)
					err = json.Unmarshal([]byte(result), &serverData)
					if err != nil {
						return nil, err
					}

					return serverData.Data.Server, nil
				case "ListServerData":
					listServerData := data.(hccGatewayData.ListServerData)
					err = json.Unmarshal([]byte(result), &listServerData)
					if err != nil {
						return nil, err
					}

					return listServerData.Data.ListServer, nil
				case "AllServerData":
					allServerData := data.(hccGatewayData.AllServerData)
					err = json.Unmarshal([]byte(result), &allServerData)
					if err != nil {
						return nil, err
					}

					return allServerData.Data.AllServer, nil
				case "NumServerData":
					numServerData := data.(hccGatewayData.NumServerData)
					err = json.Unmarshal([]byte(result), &numServerData)
					if err != nil {
						return nil, err
					}

					return numServerData.Data.NumServer, nil
				case "CreateServerData":
					createServerData := data.(hccGatewayData.CreateServerData)
					err = json.Unmarshal([]byte(result), &createServerData)
					if err != nil {
						return nil, err
					}

					return createServerData.Data.Server, nil
				case "UpdateServerData":
					updateServerData := data.(hccGatewayData.UpdateServerData)
					err = json.Unmarshal([]byte(result), &updateServerData)
					if err != nil {
						return nil, err
					}

					return updateServerData.Data.Server, nil
				case "DeleteServerData":
					deleteServerData := data.(hccGatewayData.DeleteServerData)
					err = json.Unmarshal([]byte(result), &deleteServerData)
					if err != nil {
						return nil, err
					}

					return deleteServerData.Data.Server, nil
				case "CreateServerNodeData":
					createServerNodeData := data.(hccGatewayData.CreateServerNodeData)
					err = json.Unmarshal([]byte(result), &createServerNodeData)
					if err != nil {
						return nil, err
					}

					return createServerNodeData.Data.Server, nil
				case "DeleteServerNodeData":
					deleteServerNodeData := data.(hccGatewayData.DeleteServerNodeData)
					err = json.Unmarshal([]byte(result), &deleteServerNodeData)
					if err != nil {
						return nil, err
					}

					return deleteServerNodeData.Data.Server, nil
				case "ServerNodeData":
					serverNodeData := data.(hccGatewayData.ServerNodeData)
					err = json.Unmarshal([]byte(result), &serverNodeData)
					if err != nil {
						return nil, err
					}

					return serverNodeData.Data.ServerNode, nil
				case "ListServerNodeData":
					listServerNodeData := data.(hccGatewayData.ListServerNodeData)
					err = json.Unmarshal([]byte(result), &listServerNodeData)
					if err != nil {
						return nil, err
					}

					return listServerNodeData.Data.ListServerNode, nil
				case "AllServerNodeData":
					allServerNodeData := data.(hccGatewayData.AllServerNodeData)
					err = json.Unmarshal([]byte(result), &allServerNodeData)
					if err != nil {
						return nil, err
					}

					return allServerNodeData.Data.AllServerNode, nil
				case "SubnetData":
					subnetData := data.(hccGatewayData.SubnetData)
					err = json.Unmarshal([]byte(result), &subnetData)
					if err != nil {
						return nil, err
					}

					return subnetData.Data.Subnet, nil
				case "ListSubnetData":
					listSubnetData := data.(hccGatewayData.ListSubnetData)
					err = json.Unmarshal([]byte(result), &listSubnetData)
					if err != nil {
						return nil, err
					}

					return listSubnetData.Data.ListSubnet, nil
				case "AllSubnetData":
					allSubnetData := data.(hccGatewayData.AllSubnetData)
					err = json.Unmarshal([]byte(result), &allSubnetData)
					if err != nil {
						return nil, err
					}

					return allSubnetData.Data.AllSubnet, nil
				case "NumSubnetData":
					numSubnetData := data.(hccGatewayData.NumSubnetData)
					err = json.Unmarshal([]byte(result), &numSubnetData)
					if err != nil {
						return nil, err
					}

					return numSubnetData.Data.NumSubnet, nil
				case "CreateSubnetData":
					createSubnetData := data.(hccGatewayData.CreateSubnetData)
					err = json.Unmarshal([]byte(result), &createSubnetData)
					if err != nil {
						return nil, err
					}

					return createSubnetData.Data.Subnet, nil
				case "UpdateSubnetData":
					updateSubnetData := data.(hccGatewayData.UpdateSubnetData)
					err = json.Unmarshal([]byte(result), &updateSubnetData)
					if err != nil {
						return nil, err
					}

					return updateSubnetData.Data.Subnet, nil
				case "DeleteSubnetData":
					deleteSubnetData := data.(hccGatewayData.DeleteSubnetData)
					err = json.Unmarshal([]byte(result), &deleteSubnetData)
					if err != nil {
						return nil, err
					}

					return deleteSubnetData.Data.Subnet, nil
				case "CreateDHCPDConfData":
					createDHCPDConfData := data.(hccGatewayData.CreateDHCPDConfData)
					err = json.Unmarshal([]byte(result), &createDHCPDConfData)
					if err != nil {
						return nil, err
					}

					return createDHCPDConfData.Data.Result, nil
				case "AdaptiveIPData":
					adaptiveIPData := data.(hccGatewayData.AdaptiveIPData)
					err = json.Unmarshal([]byte(result), &adaptiveIPData)
					if err != nil {
						return nil, err
					}

					return adaptiveIPData.Data.AdaptiveIP, nil
				case "ListAdaptiveIPData":
					listAdaptiveIPData := data.(hccGatewayData.ListAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &listAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return listAdaptiveIPData.Data.ListAdaptiveIP, nil
				case "AllAdaptiveIPData":
					allAdaptiveIPData := data.(hccGatewayData.AllAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &allAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return allAdaptiveIPData.Data.AllAdaptiveIP, nil
				case "NumAdaptiveIPData":
					numAdaptiveIPData := data.(hccGatewayData.NumAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &numAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return numAdaptiveIPData.Data.NumAdaptiveIP, nil
				case "CreateAdaptiveIPData":
					createAdaptiveIPData := data.(hccGatewayData.CreateAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &createAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return createAdaptiveIPData.Data.AdaptiveIP, nil
				case "UpdateAdaptiveIPData":
					updateAdaptiveIPData := data.(hccGatewayData.UpdateAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &updateAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return updateAdaptiveIPData.Data.AdaptiveIP, nil
				case "DeleteAdaptiveIPData":
					deleteAdaptiveIPData := data.(hccGatewayData.DeleteAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &deleteAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return deleteAdaptiveIPData.Data.AdaptiveIP, nil
				case "AdaptiveIPServerData":
					adaptiveIPServerData := data.(hccGatewayData.AdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &adaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return adaptiveIPServerData.Data.AdaptiveIPServer, nil
				case "ListAdaptiveIPServerData":
					listAdaptiveIPServerData := data.(hccGatewayData.ListAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &listAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return listAdaptiveIPServerData.Data.ListAdaptiveIPServer, nil
				case "AllAdaptiveIPServerData":
					allAdaptiveIPServerData := data.(hccGatewayData.AllAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &allAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return allAdaptiveIPServerData.Data.AllAdaptiveIPServer, nil
				case "NumAdaptiveIPServerData":
					numAdaptiveIPServerData := data.(hccGatewayData.NumAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &numAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return numAdaptiveIPServerData.Data.NumAdaptiveIPServer, nil
				case "CreateAdaptiveIPServerData":
					createAdaptiveIPServerData := data.(hccGatewayData.CreateAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &createAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return createAdaptiveIPServerData.Data.AdaptiveIPServer, nil
				case "DeleteAdaptiveIPServerData":
					deleteAdaptiveIPServerData := data.(hccGatewayData.DeleteAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &deleteAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return deleteAdaptiveIPServerData.Data.AdaptiveIPServer, nil
				case "VolumeData":
					volumeData := data.(hccGatewayData.VolumeData)
					err = json.Unmarshal([]byte(result), &volumeData)
					if err != nil {
						return nil, err
					}

					return volumeData.Data.Volume, nil
				case "ListVolumeData":
					listVolumeData := data.(hccGatewayData.ListVolumeData)
					err = json.Unmarshal([]byte(result), &listVolumeData)
					if err != nil {
						return nil, err
					}

					return listVolumeData.Data.ListVolume, nil
				case "AllVolumeData":
					allVolumeData := data.(hccGatewayData.AllVolumeData)
					err = json.Unmarshal([]byte(result), &allVolumeData)
					if err != nil {
						return nil, err
					}

					return allVolumeData.Data.AllVolume, nil
				case "NumVolumeData":
					numVolumeData := data.(hccGatewayData.NumVolumeData)
					err = json.Unmarshal([]byte(result), &numVolumeData)
					if err != nil {
						return nil, err
					}

					return numVolumeData.Data.NumVolume, nil
				case "VolumeAttatchmentData":
					volumeAttatchmentData := data.(hccGatewayData.VolumeAttatchmentData)
					err = json.Unmarshal([]byte(result), &volumeAttatchmentData)
					if err != nil {
						return nil, err
					}

					return volumeAttatchmentData.Data.VolumeAttatchment, nil
				case "ListVolumeAttatchmentData":
					listVolumeAttatchmentData := data.(hccGatewayData.ListVolumeAttatchmentData)
					err = json.Unmarshal([]byte(result), &listVolumeAttatchmentData)
					if err != nil {
						return nil, err
					}

					return listVolumeAttatchmentData.Data.ListVolumeAttatchment, nil
				case "AllVolumeAttatchmentData":
					allVolumeAttatchmentData := data.(hccGatewayData.AllVolumeAttatchmentData)
					err = json.Unmarshal([]byte(result), &allVolumeAttatchmentData)
					if err != nil {
						return nil, err
					}

					return allVolumeAttatchmentData.Data.AllVolumeAttatchment, nil
				case "CreateVolumeData":
					createVolumeData := data.(hccGatewayData.CreateVolumeData)
					err = json.Unmarshal([]byte(result), &createVolumeData)
					if err != nil {
						return nil, err
					}

					return createVolumeData.Data.Volume, nil
				case "UpdateVolumeData":
					updateVolumeData := data.(hccGatewayData.UpdateVolumeData)
					err = json.Unmarshal([]byte(result), &updateVolumeData)
					if err != nil {
						return nil, err
					}

					return updateVolumeData.Data.Volume, nil
				case "DeleteVolumeData":
					deleteVolumeData := data.(hccGatewayData.DeleteVolumeData)
					err = json.Unmarshal([]byte(result), &deleteVolumeData)
					if err != nil {
						return nil, err
					}

					return deleteVolumeData.Data.Volume, nil
				case "CreateVolumeAttatchmentData":
					createVolumeAttatchmentData := data.(hccGatewayData.CreateVolumeAttatchmentData)
					err = json.Unmarshal([]byte(result), &createVolumeAttatchmentData)
					if err != nil {
						return nil, err
					}

					return createVolumeAttatchmentData.Data.VolumeAttachment, nil
				case "UpdateVolumeAttatchmentData":
					updateVolumeAttatchmentData := data.(hccGatewayData.UpdateVolumeAttatchmentData)
					err = json.Unmarshal([]byte(result), &updateVolumeAttatchmentData)
					if err != nil {
						return nil, err
					}

					return updateVolumeAttatchmentData.Data.VolumeAttachment, nil
				case "DeleteVolumeAttatchmentData":
					deleteVolumeAttatchmentData := data.(hccGatewayData.DeleteVolumeAttatchmentData)
					err = json.Unmarshal([]byte(result), &deleteVolumeAttatchmentData)
					if err != nil {
						return nil, err
					}

					return deleteVolumeAttatchmentData.Data.VolumeAttachment, nil
				case "TelegrafData":
					telegrafData := data.(hccGatewayData.TelegrafData)
					err = json.Unmarshal([]byte(result), &telegrafData)
					if err != nil {
						return nil, err
					}

					return telegrafData.Data.Telegraf, nil
				default:
					return nil, errors.New("unknown data type")
				}
			}

			return result, nil
		}

		return nil, err
	}

	return nil, errors.New("http response returned error code")
}
