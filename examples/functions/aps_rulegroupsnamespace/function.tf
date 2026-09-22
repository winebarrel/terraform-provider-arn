# arn:aws:aps:ap-northeast-1:111111111111:rulegroupsnamespace/workspace-id/namespace
output "aps_rulegroupsnamespace" {
  value = provider::arn::aps_rulegroupsnamespace("workspace-id", "namespace")
}
