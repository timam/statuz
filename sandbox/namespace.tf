resource "kubernetes_namespace_v1" "statuz" {
  metadata {
    name = "statuz"
  }
}