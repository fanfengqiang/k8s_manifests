[root@node1 fluentd]# helm install --name fluentd  --namespace logs -f ./values.yaml ./fluentd-1.9.0.tgz 
NAME:   fluentd
LAST DEPLOYED: Tue Jul  2 10:06:22 2019
NAMESPACE: logs
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME     DATA  AGE
fluentd  4     2s

==> v1/PersistentVolumeClaim
NAME     STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
fluentd  Bound   pvc-90ffd494-24f6-49f6-85ce-84c9f59798b4  10Gi      RWO           ceph          2s

==> v1/Pod(related)
NAME                    READY  STATUS   RESTARTS  AGE
fluentd-c66bdbc7-tqkwm  0/1    Pending  0         1s

==> v1/Service
NAME     TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)    AGE
fluentd  ClusterIP  10.100.131.222  <none>       24220/TCP  2s

==> v1beta1/Deployment
NAME     READY  UP-TO-DATE  AVAILABLE  AGE
fluentd  0/1    1           0          1s

==> v1beta1/Ingress
NAME     HOSTS              ADDRESS  PORTS  AGE
fluentd  fluentd.5ik8s.com  80       1s


NOTES:
To verify that Fluentd Elasticsearch has started, run:

  kubectl --namespace=logs get all -l "app=fluentd,release=fluentd"

THIS APPLICATION CAPTURES ALL CONSOLE OUTPUT AND FORWARDS IT TO Elasticsearch. Anything that might be identifying,
including things like IP addresses, container images, and object names will NOT be anonymized.

