# arn:aws:mgh:ap-northeast-1:111111111111:connection-arn
output "mgh_connection_resource" {
  value = provider::arn::mgh_connection_resource("connection-arn")
}
