# RancherRedeployAction

### 一般用法

```yaml
- name: Restart
  uses: MultiMx/RancherRedeployAction@v3.0
  with:
    backend: 'https://rancher.example.domain/v3/'
    token: ${{ secrets.CATTLE_BEARER_TOKEN }}
    project: 'local:p-qgid4'
    namespace: 'control'
    workload: 'worker'
```

### 等待工作负载 100% 可用

```yaml
- name: Restart
  uses: MultiMx/RancherRedeployAction@v3.0
  with:
    backend: 'https://rancher.example.domain/v3/'
    token: ${{ secrets.CATTLE_BEARER_TOKEN }}
    project: 'local:p-qgid4'
    namespace: 'control'
    workload: 'worker'
    wait: true
```