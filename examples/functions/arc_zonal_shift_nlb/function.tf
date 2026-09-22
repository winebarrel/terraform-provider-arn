# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:loadbalancer/net/load-balancer-name/load-balancer-id
output "arc_zonal_shift_nlb" {
  value = provider::arn::arc_zonal_shift_nlb("load-balancer-name", "load-balancer-id")
}
