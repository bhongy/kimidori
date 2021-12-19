# infrastructure

`infrastructure` is the directory contains infrastructure components (cluster, networking) to run the services.

## Set up Local Development Environment

Prerequisite:

- Docker Desktop
- [Kind](https://kind.sigs.k8s.io/)
- `kubectl`

Create a local kubenetes cluster with [kind](https://kind.sigs.k8s.io/).

> Docker Desktop is expected to be running on the local machine.

```sh
# kind create cluster --name kimidori.local
kind create cluster --config cluster-local.yaml 

# set kubectl to use the local k8s cluster
kubectl config use-context kind-kimidori.local
# or kubectl cluster-info --context kind-kimidori.local

# verify that the kind-kimidori-local cluster is used
kubectl config current-context
# or kubectl config view

# see all clusters you created with kind
kind get clusters

# delete the cluster
kind delete cluster --name kimidori.local
```

> There's an issue with the worker node bind mount on macOS. Need to uncheck "Use gRPC FUSE for file sharing" option in Docker Desktop.

TEMPORARY

```sh
# create a debug container to curl a service within a cluster
kubectl run curl -it --rm --image=curlimages/curl -- sh

# then curl a service using the name
curl example-svc

# or create the helper pod testcurl `testcurl.dev.yaml` (shut down after 10m)
kubectl apply -f testcurl.dev.yaml
kubectl exec -it testcurl -- sh
```

```sh
docker kill --signal=HUP <container_id>

curl -H "Authorization: Bearer <ACL_SecretID>" \
  http://127.0.0.1:8500/v1/agent/members
```

```sh
kubectl create secret generic ssh-key-secret --from-file=ssh-privatekey=/path/to/.ssh/id_rsa --from-file=ssh-publickey=/path/to/.ssh/id_rsa.pub

kubectl create secret generic ssh-key-secret \
  --from-file=ssh-privatekey="../secrets/kimidori.local" \
  --from-file=ssh-publickey="../secrets/kimidori.local.pub"
```

... A context element in a kubeconfig file is used to group access parameters under a convenient name. Each context has three parameters: cluster, namespace, and user. [k8s Doc: Context](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#context)
... A kubeconfig file will be created at `$HOME/.kube/config`.

Q: Why using Consul instead of just Envoy
A: Consul provides service discovery, mTLS, and access control (intention). Since we use Consul for the production service mesh, this allows us to treat the local cluster as another cluster in the mesh.

## Notes: Edge Ingress (local)

Edge Ingress (local) service manages the public traffic (internet to the cluster) when using the local (on laptop) cluster.

https://www.thoughtworks.com/en-us/insights/blog/building-service-mesh-envoy-0

https://github.com/turbinelabs/examples/tree/master/local-dev-kubernetes
https://github.com/turbinelabs/examples/blob/master/telepresence-houston/README.md
https://learn.hashicorp.com/tutorials/consul/kubernetes-kind

- connectivity: service discovery, load balancing
- communication resiliency: retries, timeouts, circuit breaking, and rate limiting
- security: mTLS
- observability

edge-ingress-dashboard: web ui to manage configuration

The virtual router has one or more virtual routes that adhere to the traffic policies and retry policies. (router can delegate to another router)

profile.app.local
profile.data.local
OR
profile.srv.local
profile.srv.local/api
