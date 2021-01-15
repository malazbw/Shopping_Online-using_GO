module blatt2-grp03

go 1.15

require (
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/logger/zerolog/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	github.com/micro/go-plugins/store/redis/v2 v2.9.1
	github.com/rs/zerolog v1.20.0
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.22.0
)

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
