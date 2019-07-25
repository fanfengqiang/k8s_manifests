[root@node1 grafana]# helm install --name grafana --namespace monitor -f ./values.yaml ./grafana-3.3.4.tgz 
NAME:   grafana
LAST DEPLOYED: Tue Jul  2 22:52:49 2019
NAMESPACE: monitor
STATUS: DEPLOYED

RESOURCES:
==> v1/ClusterRole
NAME                 AGE
grafana-clusterrole  2s

==> v1/ClusterRoleBinding
NAME                        AGE
grafana-clusterrolebinding  2s

==> v1/ConfigMap
NAME          DATA  AGE
grafana       1     2s
grafana-test  1     2s

==> v1/PersistentVolumeClaim
NAME     STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
grafana  Bound   pvc-b38d2f05-a3ba-4856-9f4d-2b8e25d6223c  2Gi       RWO           ceph          2s

==> v1/Pod(related)
NAME                      READY  STATUS    RESTARTS  AGE
grafana-85f5c7df69-spqw7  0/1    Init:0/1  0         1s

==> v1/Secret
NAME     TYPE    DATA  AGE
grafana  Opaque  3     2s

==> v1/Service
NAME     TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)  AGE
grafana  ClusterIP  10.103.210.22  <none>       80/TCP   1s

==> v1/ServiceAccount
NAME     SECRETS  AGE
grafana  1        2s

==> v1beta1/Ingress
NAME     HOSTS              ADDRESS  PORTS  AGE
grafana  grafana.5ik8s.com  80       1s

==> v1beta1/PodSecurityPolicy
NAME     PRIV   CAPS      SELINUX   RUNASUSER  FSGROUP   SUPGROUP  READONLYROOTFS  VOLUMES
grafana  false  RunAsAny  RunAsAny  RunAsAny   RunAsAny  false     configMap,emptyDir,projected,secret,downwardAPI,persistentVolumeClaim

==> v1beta1/Role
NAME     AGE
grafana  2s

==> v1beta1/RoleBinding
NAME     AGE
grafana  2s

==> v1beta2/Deployment
NAME     READY  UP-TO-DATE  AVAILABLE  AGE
grafana  0/1    1           0          1s


NOTES:
1. Get your 'admin' user password by running:

   kubectl get secret --namespace monitor grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

2. The Grafana server can be accessed via port 80 on the following DNS name from within your cluster:

   grafana.monitor.svc.cluster.local

   From outside the cluster, the server URL(s) are:
     http://grafana.5ik8s.com


3. Login with the password from step 1 and the username: admin

