# arn:aws:ts::aws:tool/tool-id
output "ts_tool" {
  value = provider::arn::ts_tool("tool-id")
}
