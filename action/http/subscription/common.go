package subscription

import (
	"context"
	"github.com/functionalfoundry/graphqlws"
	graphqlgo "github.com/graphql-go/graphql"
	"hcc/piccolo/action/graphql"
	piccoloConfig "hcc/piccolo/lib/config"
	"hcc/piccolo/lib/logger"
	"time"
)

func graphqlCommonSubscription(conn graphqlws.Connection,
	opID string,
	data *graphqlws.StartMessagePayload) {
	ctx := context.Background()

	for true {
		opCancelReadLock.Lock()
		if isOpStopped(conn, opID) {
			opCancelReadLock.Unlock()
			return
		}
		opCancelReadLock.Unlock()

		params := graphqlgo.Params{
			Schema:         graphql.Schema,
			RequestString:  data.Query,
			VariableValues: data.Variables,
			OperationName:  data.OperationName,
			Context:        ctx,
		}
		//logger.Logger.Println("query", data.query)
		//logger.Logger.Println("data.Variables", data.Variables)

		result := graphqlgo.Do(params)

		graphqlData := graphqlws.DataMessagePayload{
			Data:   result.Data,
			Errors: graphqlws.ErrorsFromGraphQLErrors(result.Errors),
		}
		conn.SendData(opID, &graphqlData)
		if graphqlData.Errors != nil {
			logger.Logger.Println("graphqlCommonSubscription(): Query: ", data.Query, " Errors: ", graphqlData.Errors)
		}

		time.Sleep(time.Millisecond * time.Duration(piccoloConfig.GraphQL.SubscriptionInterval))
	}
}
