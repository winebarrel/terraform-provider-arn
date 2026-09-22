# arn:aws:codeconnections:ap-northeast-1:111111111111:host/host-id
output "codeconnections_host" {
  value = provider::arn::codeconnections_host("host-id")
}
