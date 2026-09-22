# arn:aws:ts::111111111111:execution/user-id/tool-id/execution-id
output "ts_execution" {
  value = provider::arn::ts_execution("user-id", "tool-id", "execution-id")
}
