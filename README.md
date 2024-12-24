# krew plugin: crd2sample

## how to install

### Use [krew](https://krew.sigs.k8s.io/) (Windows / macOS / Linux)

```shell
kubectl krew index add crd2sample https://github.com/AGou-ops/kubectl-crd2sample.git
kubectl krew install crd2sample/crd2sample
kubectl crd2sample -h 
```

## how to use

```bash
kubectl crd2sample <CRD_NAME>
# example
kubectl crd2sample tenants.capsule.clastix.io
```
<details>
  <summary>Sample tenants</summary>
  
  ```yaml
  apiVersion: capsule.clastix.io/v1beta1
kind: Tenant
metadata:
  name: tenants-sample
spec:
  apiVersion: <string>
  kind: <string>
  metadata: {}
  spec:
    additionalRoleBindings:
    - clusterRoleName: <string>
      subjects:
      - apiGroup: <string>
        kind: <string>
        name: <string>
        namespace: <string>
    containerRegistries:
      allowed:
      - <string>
      allowedRegex: <string>
    imagePullPolicies:
    - <string>
    ingressOptions:
      allowedClasses:
        allowed:
        - <string>
        allowedRegex: <string>
      allowedHostnames:
        allowed:
        - <string>
        allowedRegex: <string>
      hostnameCollisionScope: <string>
    limitRanges:
      items:
      - limits:
        - default: {}
          defaultRequest: {}
          max: {}
          maxLimitRequestRatio: {}
          min: {}
          type: <string>
    namespaceOptions:
      additionalMetadata:
        annotations: {}
        labels: {}
      quota: 0
    networkPolicies:
      items:
      - egress:
        - ports:
          - endPort: 0
            port: <nil>
            protocol: <string>
          to:
          - ipBlock:
              cidr: <string>
              except:
              - <string>
            namespaceSelector:
              matchExpressions:
              - key: <string>
                operator: <string>
                values:
                - <string>
              matchLabels: {}
            podSelector:
              matchExpressions:
              - key: <string>
                operator: <string>
                values:
                - <string>
              matchLabels: {}
        ingress:
        - from:
          - ipBlock:
              cidr: <string>
              except:
              - <string>
            namespaceSelector:
              matchExpressions:
              - key: <string>
                operator: <string>
                values:
                - <string>
              matchLabels: {}
            podSelector:
              matchExpressions:
              - key: <string>
                operator: <string>
                values:
                - <string>
              matchLabels: {}
          ports:
          - endPort: 0
            port: <nil>
            protocol: <string>
        podSelector:
          matchExpressions:
          - key: <string>
            operator: <string>
            values:
            - <string>
          matchLabels: {}
        policyTypes:
        - <string>
    nodeSelector: {}
    owners:
    - kind: <string>
      name: <string>
      proxySettings:
      - kind: <string>
        operations:
        - <string>
    priorityClasses:
      allowed:
      - <string>
      allowedRegex: <string>
    resourceQuotas:
      items:
      - hard: {}
        scopeSelector:
          matchExpressions:
          - operator: <string>
            scopeName: <string>
            values:
            - <string>
        scopes:
        - <string>
      scope: <string>
    serviceOptions:
      additionalMetadata:
        annotations: {}
        labels: {}
      allowedServices:
        externalName: true
        loadBalancer: true
        nodePort: true
      externalIPs:
        allowed:
        - <string>
      forbiddenAnnotations:
        denied:
        - <string>
        deniedRegex: <string>
      forbiddenLabels:
        denied:
        - <string>
        deniedRegex: <string>
    storageClasses:
      allowed:
      - <string>
      allowedRegex: <string>
  status:
    namespaces:
    - <string>
    size: 0
    state: <string>
  ```

</details>

args:

```bash
- --kubeconfig(-c): Specify kubeconfig;
- --output(-o): export result to file.
```
