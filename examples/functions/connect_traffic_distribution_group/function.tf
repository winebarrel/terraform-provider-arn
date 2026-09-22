# arn:aws:connect:ap-northeast-1:111111111111:traffic-distribution-group/traffic-distribution-group-id
output "connect_traffic_distribution_group" {
  value = provider::arn::connect_traffic_distribution_group("traffic-distribution-group-id")
}
