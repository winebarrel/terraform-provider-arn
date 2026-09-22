# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:loadbalancer/gwy/load-balancer-name/load-balancer-id
output "elasticloadbalancing_loadbalancer_gwy" {
  value = provider::arn::elasticloadbalancing_loadbalancer_gwy("load-balancer-name", "load-balancer-id")
}
