# setup-dev-input

```yaml
apiVersion: v1
items:
- apiVersion: cloud.appscode.com/v1alpha1
  kind: Credential
  metadata:
    creationTimestamp: null
    name: azure-credentials
  spec:
    name: ""
    ownerID: 0
    type: ""
  status: {}
- apiVersion: cloud.appscode.com/v1alpha1
  kind: Credential
  metadata:
    creationTimestamp: null
    name: azure-storage-credentials
  spec:
    name: ""
    ownerID: 0
    type: ""
  status: {}
- apiVersion: meta.k8s.appscode.com/v1alpha1
  kind: ClusterProfile
  metadata:
    creationTimestamp: null
    name: aws-dbaas
  spec:
    description: ""
    title: AWS DBaaS
kind: List
metadata: {}
```