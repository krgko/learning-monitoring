# EFK

EFK stack stand for Elasticsearch, Fluentd, Kibana which is nowaday popular monitoring stack

## Prerequisite

Should build docker images for `server` first

## Setting fluentd

- Build docker with plugin elasticsearch with config follow [this config](./fluentd/conf/fluent.conf)

  ```bash
  FROM fluent/fluentd:v0.12-debian
  RUN ["gem", "install", "fluent-plugin-elasticsearch", "--no-rdoc", "--no-ri", "--version", "1.9.2"]
  ```

- Start container and have fun ...

## Troubleshoot guide

1. For error

    ```bash
    [1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
    ```

    Fix by

    ```bash
    sudo sysctl -w vm.max_map_count=262144
    ```

2. Another problem can be found on [this link](https://stackoverflow.com/questions/55956645/docker-compose-yml-for-elasticsearch-7-0-1-and-kibana-7-0-1/55957214)
3. Fluentd cannot found elasticsearch - reconfig `fluent.conf` by change host to `learning_monitoring_elasticsearch` or correct one

## References

- [docker-logging-efk-compose](https://docs.fluentd.org/v/0.12/articles/docker-logging-efk-compose)
