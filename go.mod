module hcc/piccolo

go 1.13

require (
	github.com/Terry-Mao/goconf v0.0.0-20161115082538-13cb73d70c44
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/websocket v1.4.2
	github.com/graphql-go/graphql v0.7.8
	github.com/graphql-go/handler v0.2.3
	github.com/hcloud-classic/hcc_errors v1.1.2
	github.com/hcloud-classic/pb v0.0.0
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20210110051926-789bb1bd4061 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210111234610-22ae2b108f89 // indirect
	google.golang.org/grpc v1.34.1
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/hcloud-classic/pb => ../pb
