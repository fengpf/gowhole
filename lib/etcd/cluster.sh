#!/bin/sh

goreman -f Procfile start

# kill etcd2
$ goreman run stop etcd2
# restart etcd2
$ goreman run restart etcd2

# use API version 2
export ETCDCTL_API=2

etcdctl --output=table --endpoints=http://127.0.0.1:2379,http://127.0.0.1:22379,http://127.0.0.1:32379  member list

# use API version 3  

https://github.com/coreos/etcd/blob/master/Documentation/dev-guide/api_grpc_gateway.md

export ETCDCTL_API=3

etcdctl -w=table --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379 member list

etcdctl --write-out=table --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379  endpoint status

etcdctl --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379  endpoint health


etcdctl get --from-key ''

curl -iL http://localhost:2379/v3/kv/put   -X POST -d '{"key": "bmFtZQ0K", "value": "6Zi/55quDQo="}'

curl -iL http://localhost:2379/v3/kv/range   -X POST -d '{"key": "bmFtZQ0K", "range_end": "6Zi/55quDQo="}'