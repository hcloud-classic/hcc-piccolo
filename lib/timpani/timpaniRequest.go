package timpani

import (
	"bytes"
	"errors"

	//"crypto/tls"
	"encoding/json"
	"hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// NormalRebootNotification : Notify normal reboot to timpani module
func NormalRebootNotification(nodeUUID string) (string, error) {
	// If timpani uses HTTPS
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Timeout: time.Duration(config.Timpani.RequestTimeoutMs) * time.Millisecond}

	resetType := normalRebootNotificationRequestType{NodeUUID: nodeUUID, BootKind: "Normal"}
	jsonBytes, err := json.Marshal(resetType)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "http://"+config.Timpani.ServerAddress+"/v1/hcloud/action/boot", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	// If timpani uses basic auth
	// req.SetBasicAuth(config.Timpani.Username, config.Timpani.Password)

	for i := 0; i < int(config.Timpani.RequestRetry); i++ {
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
			if err != nil {
				logger.Logger.Println(err)
			} else {
				_ = resp.Body.Close()
				logger.Logger.Println("NormalRebootNotification(): nodeUUID=" + nodeUUID + " http response returned error code " + strconv.Itoa(resp.StatusCode))
			}
			logger.Logger.Println("NormalRebootNotification(): nodeUUID=" + nodeUUID + " Retrying request " + strconv.Itoa(i+1) + "/" + strconv.Itoa(int(config.Timpani.RequestRetry)))
			continue
		} else {
			if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
				// Check response
				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					_ = resp.Body.Close()
					return "", err
				}

				var response normalRebootNotificationResponseType
				err = json.Unmarshal(respBody, &response)
				if err != nil {
					return "", err
				}

				_ = resp.Body.Close()

				if response.Result != "0000" {
					return "", errors.New(response.ResultMessage)
				}

				return response.ResultMessage, nil
			}

			_ = resp.Body.Close()
			return "", err
		}
	}

	return "", errors.New("NormalRebootNotification(): nodeUUID=" + nodeUUID + " retry count exceeded")
}
