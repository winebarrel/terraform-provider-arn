# arn:aws:bedrock::111111111111:system-tool/resource-id
output "bedrock_system_tool" {
  value = provider::arn::bedrock_system_tool("resource-id")
}
