# arn:aws:lightsail:ap-northeast-1:111111111111:LoadBalancer/id
output "lightsail_load_balancer" {
  value = provider::arn::lightsail_load_balancer("id")
}
