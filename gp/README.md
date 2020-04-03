# GP

Grafana and Prometheus is another stack for monitoring. **Does not link to API server yet**

## Start docker

```bash
docker-compose up -d
```

## Accessing

- [prometheus-web](http://localhost:9090)
- [grafana-web](http://localhost:9110) - with username: `admin` and password: `P@ssw0rd`

## Details

- `prometheus` is a system service monitoring. It collects matrics from configurated targets at given interval
- `node-exporter` is exporter to get machine matrics. You can select matric via dropdown `insert matric at cursor` on prometheus with select `Graph` tab display as graph
- `grafana` is graph display UI

## Linking grafana with prometheus

- Create data source: gear icon (configuration) > select prometheus > add URL as `http://learning_monitoring_prometheus:9090` > click `Save & Test` (Check that prometheus you set is a default one)
- New dashboard: select matric visualization > add query > select graph to display
- Done !!
