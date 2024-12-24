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

args:

```bash
- --kubeconfig(-c): Specify kubeconfig;
- --output(-o): export result to file.
```
