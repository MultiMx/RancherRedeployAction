# RancherRedeployAction

```yaml
- name: Restart
  uses: MultiMx/RancherRedeployAction@v2.1
  with:
    backend: 'https://rancher.example.domain/v3/'
    token: ${{ secrets.CATTLE_BEARER_TOKEN }}
    project: 'local:p-qgid4'
    namespace: 'control'
    workload: 'worker'
```