    - --oidc-issuer-url=https://dex.5ik8s.com
    - --oidc-client-id=loginapp
    - --oidc-ca-file=/etc/kubernetes/pki/ca-dex.crt
    - --oidc-username-claim=email
    - --oidc-groups-claim=groups


echo `curl -k https://192.168.65.91:6443/api/v1/namespaces/kube-public/configmaps/cluster-info | jq .data.kubeconfig` |grep certificate-authority-data|awk '{print $NF}'|base64 -d
