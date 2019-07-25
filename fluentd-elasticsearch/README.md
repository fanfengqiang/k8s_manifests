[root@node1 fluentd-elasticsearch]# helm install --name fluentd-elasticsearch --namespace logs -f ./values.yaml ./fluentd-elasticsearch-2.0.7.tgz 
NAME:   fluentd-elasticsearch
LAST DEPLOYED: Tue Jul  2 10:15:39 2019
NAMESPACE: logs
STATUS: DEPLOYED

RESOURCES:
==> v1/ClusterRole
NAME                   AGE
fluentd-elasticsearch  5s

==> v1/ClusterRoleBinding
NAME                   AGE
fluentd-elasticsearch  5s

==> v1/ConfigMap
NAME                   DATA  AGE
fluentd-elasticsearch  6     5s

==> v1/DaemonSet
NAME                   DESIRED  CURRENT  READY  UP-TO-DATE  AVAILABLE  NODE SELECTOR  AGE
fluentd-elasticsearch  4        4        0      4           0          <none>         4s

==> v1/Pod(related)
NAME                         READY  STATUS             RESTARTS  AGE
fluentd-elasticsearch-4cbnt  0/1    ContainerCreating  0         4s
fluentd-elasticsearch-l6d2r  0/1    ContainerCreating  0         3s
fluentd-elasticsearch-nrv9t  0/1    ContainerCreating  0         3s
fluentd-elasticsearch-pq4tr  0/1    ContainerCreating  0         3s

==> v1/Service
NAME                   TYPE       CLUSTER-IP    EXTERNAL-IP  PORT(S)    AGE
fluentd-elasticsearch  ClusterIP  10.109.170.6  <none>       24231/TCP  4s

==> v1/ServiceAccount
NAME                   SECRETS  AGE
fluentd-elasticsearch  1        5s


NOTES:
1. To verify that Fluentd has started, run:

  kubectl --namespace=logs get pods -l "app.kubernetes.io/name=fluentd-elasticsearch,app.kubernetes.io/instance=fluentd-elasticsearch"

THIS APPLICATION CAPTURES ALL CONSOLE OUTPUT AND FORWARDS IT TO elasticsearch . Anything that might be identifying,
including things like IP addresses, container images, and object names will NOT be anonymized.
2. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace logs -l "app.kubernetes.io/name=fluentd-elasticsearch,app.kubernetes.io/instance=fluentd-elasticsearch" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl port-forward $POD_NAME 8080:80

