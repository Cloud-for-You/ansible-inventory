# ansible-inventory

API pro dynamicke ansible inventory různých IaaS providerů

### Build
```sh
go build -ldflags="-X main.Commit=$(git rev-parse HEAD)" -a -o ~/Documents/Kubernetes/_K8S_DEPLOY_CLUSTER/inventory/multipass ./inventory_from_multipass.go
```
