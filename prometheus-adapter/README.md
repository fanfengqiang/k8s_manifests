[root@node1 prometheus-adapter]# helm install --name prometheus-adapter --namespace monitor -f ./values.yaml ./prometheus-adapter-v0.5.0.tgz 
NAME:   prometheus-adapter
LAST DEPLOYED: Tue Jul  2 22:36:33 2019
NAMESPACE: monitor
STATUS: DEPLOYED

RESOURCES:
==> v1/ClusterRole
NAME                                 AGE
prometheus-adapter-external-metrics  2s
prometheus-adapter-resource-reader   2s
prometheus-adapter-server-resources  2s

==> v1/ClusterRoleBinding
NAME                                                AGE
prometheus-adapter-hpa-controller                   2s
prometheus-adapter-hpa-controller-external-metrics  2s
prometheus-adapter-resource-reader                  2s
prometheus-adapter:system:auth-delegator            2s

==> v1/ConfigMap
NAME                DATA  AGE
prometheus-adapter  1     2s

==> v1/Deployment
NAME                READY  UP-TO-DATE  AVAILABLE  AGE
prometheus-adapter  0/1    1           0          1s

==> v1/Pod(related)
NAME                                READY  STATUS             RESTARTS  AGE
prometheus-adapter-9bb84b4dd-qntnk  0/1    ContainerCreating  0         1s

==> v1/RoleBinding
NAME                            AGE
prometheus-adapter-auth-reader  2s

==> v1/Service
NAME                TYPE       CLUSTER-IP   EXTERNAL-IP  PORT(S)  AGE
prometheus-adapter  ClusterIP  10.98.2.193  <none>       443/TCP  1s

==> v1/ServiceAccount
NAME                SECRETS  AGE
prometheus-adapter  1        2s

==> v1beta1/APIService
NAME                             AGE
v1beta1.custom.metrics.k8s.io    1s
v1beta1.external.metrics.k8s.io  1s


NOTES:
prometheus-adapter has been deployed.
In a few minutes you should be able to list metrics using the following command:

  kubectl get --raw /apis/custom.metrics.k8s.io/v1beta1


