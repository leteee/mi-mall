module captcha-client

go 1.16

require (
	github.com/go-micro/plugins/v4/registry/consul v1.2.1
	go-micro.dev/v4 v4.9.0
	google.golang.org/protobuf v1.31.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
