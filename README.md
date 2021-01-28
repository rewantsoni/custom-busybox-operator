Custom BusyBox Operator
===

Create a custom operator which deploys a BusyBox image on kubernetes using a Deployment.

## Example Usage

**User creates custom resource yaml:**

**custom.yaml**


```yaml
apiVersion: store.example.com/v1beta1
kind: StoreFrnt
metadata:
  name: storefrnt-sample
spec:
  # Add fields here
  size: 1
  location: SAT
```

* **User deploys the yaml:**
```shell
$ kubectl create -f custom.yaml
storefrnt.store.example.com/storefrnt-sample created
```


* **Results:**
```shell
$ kubectl get deployment, pods
NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/storefrnt-sample   2/2     2            2           2m26s

NAME                                    READY   STATUS    RESTARTS   AGE
pod/storefrnt-sample-8676d76f87-g77xw   1/1     Running   0          2m26s
pod/storefrnt-sample-8676d76f87-ghsvz   1/1     Running   0          2m26s
```

## Env:
* Kubernetes on Minikube will be sufficient

### How to run:

```shell
make install
make run
```

### Deploy Sample Wordpress operator:
```shell
kubectl apply -f config/samples/store_v1beta1_storefrnt.yaml
```

### Check if operator is deployed:
```shell
kubectl get deployment,pods