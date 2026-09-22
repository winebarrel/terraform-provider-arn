# arn:aws:ec2:ap-northeast-1:111111111111:traffic-mirror-filter/traffic-mirror-filter-id
output "ec2_traffic_mirror_filter" {
  value = provider::arn::ec2_traffic_mirror_filter("traffic-mirror-filter-id")
}
