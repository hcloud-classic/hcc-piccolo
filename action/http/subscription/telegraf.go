package subscription

import (
	"context"
	"fmt"
	"github.com/functionalfoundry/graphqlws"
	graphqlgo "github.com/graphql-go/graphql"
	"hcc/piccolo/action/graphql"
	piccoloConfig "hcc/piccolo/lib/config"
	"strings"
	"time"
)

func telegrafSubscriptionQueryTimeChange(query string, newTime string) string {
	newQuery := ""

	query = strings.Replace(query, "\t", "", -1)
	query = strings.Replace(query, "\n", "", -1)

	querySlice := strings.Split(query, ",")
	for i := range querySlice {
		if strings.Contains(querySlice[i], "time") {
			if strings.Contains(querySlice[i], "$time") {
				continue
			}
			querySlice[i] = strings.Replace(querySlice[i], " ", "", -1)
			s := strings.Split(querySlice[i], ":")
			if len(s) == 2 {
				querySlice[i] = "time: " + "\"" + newTime + "\""
			}
		}
		newQuery += querySlice[i] + ","
	}
	if strings.Contains(newQuery, "$uuid") &&
		!strings.Contains(newQuery, "$time") {
		newQuery = strings.Replace(newQuery, "$uuid: String!", "$time: String!, $uuid: String!", -1)
		newQuery = strings.Replace(newQuery, "uuid: $uuid", "time: $time, uuid: $uuid", -1)
	}

	//logger.Logger.Println("newQuery", newQuery)

	return newQuery
}

func telegrafSubscriptionGetNewTime(dataStr string) string {
	newTime := ""

	if strings.Contains(dataStr, "values") {
		dataStr = strings.TrimSpace(dataStr)
		dataStr = strings.Replace(dataStr, " ", "", -1)
		dataStr = strings.Replace(dataStr, "\t", "", -1)
		dataStr = strings.Replace(dataStr, "\n", "", -1)
		dataStr = strings.Replace(dataStr, "\\", "", -1)
		dataStr = strings.Replace(dataStr, "\"", "", -1)
		dataStr = strings.Replace(dataStr, "[", "", -1)
		dataStr = strings.Replace(dataStr, "]", "", -1)
		slices := strings.Split(dataStr, ":")
		length := len(slices)
		for i := range slices {
			if strings.Contains(slices[i], "values") && length >= i+2 {
				s := strings.Split(slices[i+1], ",")
				newTime = s[0]
				break
			}
		}
	}

	//logger.Logger.Println("newTime", newTime)

	return newTime
}

func telegrafSubscription(conn graphqlws.Connection,
	opID string,
	data *graphqlws.StartMessagePayload,
	newTime *string) {
	ctx := context.Background()

	for true {
		query := data.Query

		if *newTime != "" {
			query = telegrafSubscriptionQueryTimeChange(data.Query, *newTime)
			data.Variables["time"] = *newTime
		}

		params := graphqlgo.Params{
			Schema:         graphql.Schema,
			RequestString:  query,
			VariableValues: data.Variables,
			OperationName:  data.OperationName,
			Context:        ctx,
		}
		//logger.Logger.Println("query", query)
		//logger.Logger.Println("goroutineData.Variables", data.Variables)
		result := graphqlgo.Do(params)

		dataStr := fmt.Sprintf("%v", result.Data)
		*newTime = telegrafSubscriptionGetNewTime(dataStr)

		graphqlData := graphqlws.DataMessagePayload{
			Data:   result.Data,
			Errors: graphqlws.ErrorsFromGraphQLErrors(result.Errors),
		}
		conn.SendData(opID, &graphqlData)
		if graphqlData.Errors != nil {
			//logger.Logger.Println("subscription websocket Error: ", graphqlData.Errors)
		}

		if isConnectionClosed(conn.ID()) {
			return
		}

		time.Sleep(time.Millisecond * time.Duration(piccoloConfig.GraphQL.SubscriptionInterval))
	}
}
