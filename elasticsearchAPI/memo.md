# Research

## Downward API

### Gettable parameters

* [official doc](https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/#capabilities-of-the-downward-api)

### Using environment variables

* [official doc](https://kubernetes.io/ja/docs/tasks/inject-data-application/environment-variable-expose-pod-information/#use-container-fields-as-values-for-environment-variables)

* manifest

~~~
apiVersion: v1
kind: Pod
metadata:
  name: busybox
spec:
  containers:
    - name: busybox-container
      image: busybox
      command: [ "sh", "-c", "sleep 300"]
      resources:
        limits:
          cpu: "250m"
      env:
        - name: MY_CPU_LIMIT
          valueFrom:
            resourceFieldRef:
              containerName: busybox-container
              resource: limits.cpu
              divisor: 1m
  restartPolicy: Never
~~~

* log

~~~
$ kubectl apply -f pod.yaml
$ kubectl exec -ti busybox -- /bin/sh
# echo $MY_CPU_LIMIT
250
~~~

### Using file

* [official doc](https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/#store-container-fields)

* manifest

~~~
apiVersion: v1
kind: Pod
metadata:
  name: busybox
spec:
  containers:
    - name: busybox-container
      image: busybox
      command: ["sh", "-c", "sleep 300"]
      resources:
        limits:
          cpu: "1"
      volumeMounts:
        - name: limitvcpu
          mountPath: /etc/podinfo
  volumes:
    - name: limitvcpu
      downwardAPI:
        items:
          - path: "cpu_limit"
            resourceFieldRef:
              containerName: busybox-container
              resource: limits.cpu
              divisor: 1m
~~~

* log

~~~
$ kubectl apply -f pod.yaml
$ kubectl exec -ti busybox -- cat /etc/podinfo/cpu_limit
1000
~~~

