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
- Deploy the application
```shell
$ oc apply -f deploy.yaml
```
- Get the route and curl the hosturl with /machine-info endpoint
```shell
$ oc get route machine-info-route

# replace the hostname based on the above output and execute it to test.
$ for i in `seq 1 10`; do curl -k https://machine-info-route-default.apps.het-hc-1.hypershift-ppc64le.com/machine-info; done
```
