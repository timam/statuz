resource "helm_release" "grafana" {
  namespace        = kubernetes_namespace_v1.statuz.id
  repository       = "https://grafana.github.io/helm-charts"
  chart            = "grafana"
  name             = "grafana"
  version          = "6.61.1"
  values = [
    <<-EOT
      grafana.ini:
        auth.anonymous:
          enabled: true
          org_role: Admin
      datasources:
        datasources.yaml:
          apiVersion: 1
          datasources:
          - name: Prometheus
            type: prometheus
            url: http://prometheus-server:80
    EOT
  ]
}