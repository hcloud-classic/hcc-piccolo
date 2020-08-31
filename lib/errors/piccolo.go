package errors

const (
	PiccoloInternalInitFail              = piccolo + internal + initFail
	PiccoloInternalConnectionFail  = piccolo + internal + connectionFail

	PiccoloGrpcRequestError        = piccolo + grpc + requestError

	PiccoloGraphQLArgumentError       = piccolo + graphql + argumentError
	PiccoloGraphQLLoginFailed  = piccolo + graphql + loginFailed
	PiccoloGraphQLTokenExpired = piccolo + graphql + tokenExpired
)
