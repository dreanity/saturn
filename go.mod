module github.com/dreanity/saturn

go 1.16

require (
	github.com/cosmos/cosmos-proto v1.0.0-alpha7
	github.com/cosmos/cosmos-sdk v0.45.4
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/drand/drand v1.3.2
	github.com/drand/kyber v1.1.11
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ignite/cli v0.22.2
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.2
	github.com/tendermint/spn v0.2.1-0.20220610090138-44b136f042c4
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/genproto v0.0.0-20220728213248-dd149ef739b9
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
