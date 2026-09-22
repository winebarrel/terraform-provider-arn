# arn:aws:codeconnections:ap-northeast-1:111111111111:connection/connection-id
output "codeconnections_connection" {
  value = provider::arn::codeconnections_connection("connection-id")
}
