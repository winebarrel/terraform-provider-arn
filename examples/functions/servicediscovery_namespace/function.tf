# arn:aws:servicediscovery:ap-northeast-1:111111111111:namespace/namespace-id
output "servicediscovery_namespace" {
  value = provider::arn::servicediscovery_namespace("namespace-id")
}
