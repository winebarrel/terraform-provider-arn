# arn:aws:mgn:ap-northeast-1:111111111111:source-server/source-server-id
output "mgn_source_server_resource" {
  value = provider::arn::mgn_source_server_resource("source-server-id")
}
