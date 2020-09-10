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
	case "cello":
		timeout = time.Duration(config.Cello.RequestTimeoutMs)
		url += config.Cello.ServerAddress + ":" + strconv.Itoa(int(config.Cello.ServerPort))
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
