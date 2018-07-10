#!/bin/sh

goreman -f Procfile start

# kill etcd2
$ goreman run stop etcd2
# restart etcd2
$ goreman run restart etcd2

# use API version 2
etcdctl --output=table --endpoints=http://127.0.0.1:2379 member list

# use API version 3
export ETCDCTL_API=3

etcdctl -w=table --endpoints=http://127.0.0.1:2379 member list