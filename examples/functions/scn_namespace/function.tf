# arn:aws:scn:ap-northeast-1:111111111111:instance/instance-id/namespaces/namespace
output "scn_namespace" {
  value = provider::arn::scn_namespace("instance-id", "namespace")
}
