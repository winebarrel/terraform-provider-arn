# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:loadbalancer/net/load-balancer-name/load-balancer-id
output "elasticloadbalancing_loadbalancer_net" {
  value = provider::arn::elasticloadbalancing_loadbalancer_net("load-balancer-name", "load-balancer-id")
}
