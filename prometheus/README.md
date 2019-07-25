[root@node1 prometheus]# helm install --name prometheus --namespace monitor -f ./values.yaml ./prometheus-8.10.3.tgz
NAME:   prometheus
LAST DEPLOYED: Tue Jul  2 22:34:30 2019
NAMESPACE: monitor
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                     DATA  AGE
prometheus-alertmanager  1     24s
prometheus-server        3     24s

==> v1/PersistentVolumeClaim
NAME                     STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
prometheus-alertmanager  Bound   pvc-ccb897d2-428a-4fd7-b3ef-765ab9912142  2Gi       RWO           ceph          24s
prometheus-server        Bound   pvc-26f5b539-94e3-4246-a1cc-e0e8bafaefb1  8Gi       RWO           ceph          23s

==> v1/Pod(related)
NAME                                            READY  STATUS             RESTARTS  AGE
prometheus-alertmanager-5f49f7bd68-26vbt        0/2    ContainerCreating  0         19s
prometheus-kube-state-metrics-64f766f4cd-hz4wq  1/1    Running            0         19s
prometheus-node-exporter-9w2cm                  1/1    Running            0         19s
prometheus-node-exporter-bljrs                  1/1    Running            0         19s
prometheus-node-exporter-xcwns                  1/1    Running            0         18s
prometheus-node-exporter-zbcx2                  0/1    ContainerCreating  0         19s
prometheus-pushgateway-7bfc59b475-pbzjr         0/1    Running            0         18s
prometheus-server-58794c66fd-bgvxl              0/2    Init:0/1           0         18s

==> v1/Service
NAME                           TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)   AGE
prometheus-alertmanager        ClusterIP  10.101.224.80   <none>       80/TCP    21s
prometheus-kube-state-metrics  ClusterIP  None            <none>       80/TCP    21s
prometheus-node-exporter       ClusterIP  None            <none>       9100/TCP  20s
prometheus-pushgateway         ClusterIP  10.104.232.255  <none>       9091/TCP  20s
prometheus-server              ClusterIP  10.98.255.77    <none>       80/TCP    19s

==> v1/ServiceAccount
NAME                           SECRETS  AGE
prometheus-alertmanager        1        23s
prometheus-kube-state-metrics  1        23s
prometheus-node-exporter       1        23s
prometheus-pushgateway         1        23s
prometheus-server              1        23s

==> v1beta1/ClusterRole
NAME                           AGE
prometheus-kube-state-metrics  22s
prometheus-server              22s

==> v1beta1/ClusterRoleBinding
NAME                           AGE
prometheus-kube-state-metrics  22s
prometheus-server              22s

==> v1beta1/DaemonSet
NAME                      DESIRED  CURRENT  READY  UP-TO-DATE  AVAILABLE  NODE SELECTOR  AGE
prometheus-node-exporter  4        4        3      4           3          <none>         19s

==> v1beta1/Deployment
NAME                           READY  UP-TO-DATE  AVAILABLE  AGE
prometheus-alertmanager        0/1    1           0          19s
prometheus-kube-state-metrics  1/1    1           1          19s
prometheus-pushgateway         0/1    1           0          19s
prometheus-server              0/1    1           0          19s

==> v1beta1/Ingress
NAME               HOSTS                 ADDRESS  PORTS  AGE
prometheus-server  prometheus.5ik8s.com  80       18s


NOTES:
The Prometheus server can be accessed via port 80 on the following DNS name from within your cluster:
prometheus-server.monitor.svc.cluster.local

From outside the cluster, the server URL(s) are:
http://prometheus.5ik8s.com


The Prometheus alertmanager can be accessed via port 80 on the following DNS name from within your cluster:
prometheus-alertmanager.monitor.svc.cluster.local


Get the Alertmanager URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace monitor -l "app=prometheus,component=alertmanager" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace monitor port-forward $POD_NAME 9093


The Prometheus PushGateway can be accessed via port 9091 on the following DNS name from within your cluster:
prometheus-pushgateway.monitor.svc.cluster.local


Get the PushGateway URL by running these commands in the same shell:
  export POD_NAME=$(kubectl get pods --namespace monitor -l "app=prometheus,component=pushgateway" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace monitor port-forward $POD_NAME 9091

For more information on running Prometheus, visit:
https://prometheus.io/

