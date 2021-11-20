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

... A context element in a kubeconfig file is used to group access parameters under a convenient name. Each context has three parameters: cluster, namespace, and user. [k8s Doc: Context](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#context)
... A kubeconfig file will be created at `$HOME/.kube/config`.

Q: Why using Consul instead of just Envoy
A: Consul provides service discovery, mTLS, and access control (intention). Since we use Consul for the production service mesh, this allows us to treat the local cluster as another cluster in the mesh.

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
