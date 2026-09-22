# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:loadbalancer/app/load-balancer-name/load-balancer-id
output "elasticloadbalancing_loadbalancer_app" {
  value = provider::arn::elasticloadbalancing_loadbalancer_app("load-balancer-name", "load-balancer-id")
}
