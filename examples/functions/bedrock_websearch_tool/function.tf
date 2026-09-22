# arn:aws:bedrock-websearch:ap-northeast-1:aws:tool/tool-name
output "bedrock_websearch_tool" {
  value = provider::arn::bedrock_websearch_tool("tool-name")
}
