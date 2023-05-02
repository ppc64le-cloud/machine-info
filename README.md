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

Sample output for the above command:
```shell
# reached the x86 pod 6 times and ppc64le pod 4 times.
$ for i in `seq 1 10`; do curl -k https://machine-info-route-default.apps.het-hc-1.hypershift-ppc64le.com/machine-info; done
machine-info: Linux machine-info-7df6d58f94-q7bkw 4.18.0-372.41.1.el8_6.x86_64 #1 SMP Thu Jan 5 13:56:06 EST 2023 x86_64 GNU/Linux

machine-info: Linux machine-info-7df6d58f94-q7bkw 4.18.0-372.41.1.el8_6.x86_64 #1 SMP Thu Jan 5 13:56:06 EST 2023 x86_64 GNU/Linux

machine-info: Linux machine-info-7df6d58f94-96c9b 4.18.0-372.41.1.el8_6.ppc64le #1 SMP Thu Jan 5 13:49:25 EST 2023 ppc64le GNU/Linux

machine-info: Linux machine-info-7df6d58f94-q7bkw 4.18.0-372.41.1.el8_6.x86_64 #1 SMP Thu Jan 5 13:56:06 EST 2023 x86_64 GNU/Linux

machine-info: Linux machine-info-7df6d58f94-q7bkw 4.18.0-372.41.1.el8_6.x86_64 #1 SMP Thu Jan 5 13:56:06 EST 2023 x86_64 GNU/Linux

machine-info: Linux machine-info-7df6d58f94-96c9b 4.18.0-372.41.1.el8_6.ppc64le #1 SMP Thu Jan 5 13:49:25 EST 2023 ppc64le GNU/Linux

machine-info: Linux machine-info-7df6d58f94-96c9b 4.18.0-372.41.1.el8_6.ppc64le #1 SMP Thu Jan 5 13:49:25 EST 2023 ppc64le GNU/Linux

machine-info: Linux machine-info-7df6d58f94-96c9b 4.18.0-372.41.1.el8_6.ppc64le #1 SMP Thu Jan 5 13:49:25 EST 2023 ppc64le GNU/Linux

machine-info: Linux machine-info-7df6d58f94-q7bkw 4.18.0-372.41.1.el8_6.x86_64 #1 SMP Thu Jan 5 13:56:06 EST 2023 x86_64 GNU/Linux

machine-info: Linux machine-info-7df6d58f94-q7bkw 4.18.0-372.41.1.el8_6.x86_64 #1 SMP Thu Jan 5 13:56:06 EST 2023 x86_64 GNU/Linux
```
