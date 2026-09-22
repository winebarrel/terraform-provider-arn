# arn:aws:qdeveloper:ap-northeast-1:111111111111:plugin/identifier
output "q_plugin" {
  value = provider::arn::q_plugin("identifier")
}
