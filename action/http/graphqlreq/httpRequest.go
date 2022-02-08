package graphqlreq

import (
	"encoding/json"
	"errors"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DoHTTPRequest : Send http request to other modules with GraphQL query string.
func DoHTTPRequest(moduleName string, needData bool, data interface{}, queryName string, query string) (interface{}, error) {
	client := &http.Client{Timeout: time.Duration(config.Timpani.RequestTimeoutMs) * time.Millisecond}
	url := "http://"
	switch moduleName {
	case "timpani":
		url += config.Timpani.ServerAddress + ":" + strconv.Itoa(int(config.Timpani.ServerPort))

	default:
		return nil, errors.New("unknown module name")
	}
	url += "/graphql?query=" + queryURLEncoder(query)
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

				switch queryName {
				case "mastersync":
					masterSync := data.(model.MasterSync)
					err = json.Unmarshal([]byte(result), &(masterSync))
					// fmt.Println("listNodeData: ", listNodeData)

					if err != nil {
						return nil, err
					}
					return masterSync, nil
				case "backup":
					backup := data.(model.Backup)
					err = json.Unmarshal([]byte(result), &(backup))
					// fmt.Println("listNodeData: ", listNodeData)

					if err != nil {
						return nil, err
					}
					return backup, nil
				case "backupschduler":
					backupSchduler := data.(model.BackupScheduler)
					err = json.Unmarshal([]byte(result), &(backupSchduler))
					// fmt.Println("listNodeData: ", listNodeData)

					if err != nil {
						return nil, err
					}
					return backupSchduler, nil

				default:
					return nil, errors.New("data is not supported for " + moduleName + " module")
				}
			}

			return result, nil
		}

		return nil, err
	}

	return nil, errors.New("http response returned error code")
}

func compatibleRFC3986Encode(str string) string {
	resultStr := str
	resultStr = strings.Replace(resultStr, "+", "%20", -1)
	return resultStr
}

func queryURLEncoder(queryString string) string {
	params := url.Values{
		"query_string": {queryString},
	}

	urlEncode := compatibleRFC3986Encode(params.Encode())
	urlEncode = urlEncode[len("query_string="):]

	return urlEncode
}
