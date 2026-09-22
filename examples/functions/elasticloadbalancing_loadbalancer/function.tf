# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:loadbalancer/load-balancer-name
output "elasticloadbalancing_loadbalancer" {
  value = provider::arn::elasticloadbalancing_loadbalancer("load-balancer-name")
}
