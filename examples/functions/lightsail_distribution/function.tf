# arn:aws:lightsail:ap-northeast-1:111111111111:Distribution/id
output "lightsail_distribution" {
  value = provider::arn::lightsail_distribution("id")
}
