resource "helm_release" "prometheus" {
  namespace = kubernetes_namespace_v1.statuz.id
  repository = "https://prometheus-community.github.io/helm-charts"
  chart = "prometheus"
  name  = "prometheus"
  version                    = "25.2.0"
  values = [
    <<-EOT
      serverFiles:
        rule_files:
          - /etc/config/recording_rules.yml
          - /etc/config/alerting_rules.yml
        prometheus.yml:
          scrape_configs:
            - job_name: prometheus
              honor_timestamps: true
              scrape_interval: 15s
              scrape_timeout: 10s
              metrics_path: /metrics
              scheme: http
              static_configs:
                - targets:
                    - localhost:9090
            - job_name: watcher
              honor_timestamps: true
              scrape_interval: 15s
              scrape_timeout: 10s
              metrics_path: /metrics
              scheme: http
              static_configs:
                - targets:
                    - watcher:8080
    EOT
  ]
}