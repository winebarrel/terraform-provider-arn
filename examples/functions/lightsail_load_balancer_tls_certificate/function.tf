# arn:aws:lightsail:ap-northeast-1:111111111111:LoadBalancerTlsCertificate/id
output "lightsail_load_balancer_tls_certificate" {
  value = provider::arn::lightsail_load_balancer_tls_certificate("id")
}
