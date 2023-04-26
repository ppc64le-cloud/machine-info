# machine-info

This is a simple Go program that starts an HTTP server on port 8080 and provides an endpoint at /machine-info. When a request is made to this endpoint, the program executes the `uname -a` command and returns the output as the response.

## How to build the container image

```shell
$ docker buildx build --platform=linux/ppc64le,linux/amd64,linux/arm64 --push -t quay.io/mkumatag/machine-info:latest .
```

## Deploy this app

### Prerequsites
- Openshift cluster

### How to deploy/use
- Modify the following route to appropriete one
```shell
spec:
  host: machine-info.apps.het-hc-1.hypershift-ppc64le.com
```

to 
```shell
spec:
  host: <user supplied hostname, can be used from the existing routes>
```
- Deploy the application
```shell
$ oc apply -f deploy.yaml
```
