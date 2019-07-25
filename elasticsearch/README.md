[root@node1 elasticsearch]# helm install --name elasticsearch --namespace logs -f ./values.yaml ./elasticsearch-1.26.0.tgz 
NAME:   elasticsearch
LAST DEPLOYED: Tue Jul  2 09:27:22 2019
NAMESPACE: logs
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                DATA  AGE
elasticsearch       4     5s
elasticsearch-test  1     5s

==> v1/Pod(related)
NAME                                   READY  STATUS    RESTARTS  AGE
elasticsearch-client-6ddcc6f47f-8mfr8  0/1    Init:0/1  0         4s
elasticsearch-client-6ddcc6f47f-nm9rq  0/1    Init:0/1  0         4s
elasticsearch-data-0                   0/1    Pending   0         4s
elasticsearch-master-0                 0/1    Pending   0         4s

==> v1/Service
NAME                     TYPE       CLUSTER-IP   EXTERNAL-IP  PORT(S)   AGE
elasticsearch-client     ClusterIP  10.98.66.88  <none>       9200/TCP  5s
elasticsearch-discovery  ClusterIP  None         <none>       9300/TCP  5s

==> v1/ServiceAccount
NAME                  SECRETS  AGE
elasticsearch-client  1        5s
elasticsearch-data    1        5s
elasticsearch-master  1        5s

==> v1beta1/Deployment
NAME                  READY  UP-TO-DATE  AVAILABLE  AGE
elasticsearch-client  0/2    2           0          5s

==> v1beta1/Ingress
NAME                  HOSTS                    ADDRESS  PORTS  AGE
elasticsearch-client  elasticsearch.5ik8s.com  80       4s

==> v1beta1/StatefulSet
NAME                  READY  AGE
elasticsearch-data    0/2    5s
elasticsearch-master  0/3    4s


NOTES:
The elasticsearch cluster has been installed.

Elasticsearch can be accessed:

  * Within your cluster, at the following DNS name at port 9200:

    elasticsearch-client.logs.svc

  * From outside the cluster, run these commands in the same shell:

    export POD_NAME=$(kubectl get pods --namespace logs -l "app=elasticsearch,component=client,release=elasticsearch" -o jsonpath="{.items[0].metadata.name}")
    echo "Visit http://127.0.0.1:9200 to use Elasticsearch"
    kubectl port-forward --namespace logs $POD_NAME 9200:9200

