# arn:aws:codecatalyst:ap-northeast-1:111111111111:/connections/connection-id
output "codecatalyst_connections" {
  value = provider::arn::codecatalyst_connections("connection-id")
}
