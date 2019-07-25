[root@node1 kibana]# helm install --name kibana --namespace logs -f ./values.yaml ./kibana-2.3.0.tgz 
NAME:   kibana
LAST DEPLOYED: Tue Jul  2 09:42:43 2019
NAMESPACE: logs
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                 DATA  AGE
kibana               1     3s
kibana-dashboards    1     3s
kibana-importscript  1     3s
kibana-test          1     2s

==> v1/PersistentVolumeClaim
NAME    STATUS  VOLUME                                    CAPACITY  ACCESS MODES  STORAGECLASS  AGE
kibana  Bound   pvc-14fd0dac-3d64-423c-b3fc-5596a442e09b  5Gi       RWO           ceph          2s

==> v1/Pod(related)
NAME                    READY  STATUS   RESTARTS  AGE
kibana-98764cbdc-ddb2w  0/1    Pending  0         1s

==> v1/Service
NAME    TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)  AGE
kibana  ClusterIP  10.111.150.194  <none>       443/TCP  2s

==> v1beta1/Deployment
NAME    READY  UP-TO-DATE  AVAILABLE  AGE
kibana  0/1    1           0          2s

==> v1beta1/Ingress
NAME    HOSTS             ADDRESS  PORTS  AGE
kibana  kibana.5ik8s.com  80       2s


NOTES:
To verify that kibana has started, run:

  kubectl --namespace=logs get pods -l "app=kibana"

Kibana can be accessed:

  * From outside the cluster, run these commands in the same shell:

    export POD_NAME=$(kubectl get pods --namespace logs -l "app=kibana,release=kibana" -o jsonpath="{.items[0].metadata.name}")
    echo "Visit http://127.0.0.1:5601 to use Kibana"
    kubectl port-forward --namespace logs $POD_NAME 5601:5601

