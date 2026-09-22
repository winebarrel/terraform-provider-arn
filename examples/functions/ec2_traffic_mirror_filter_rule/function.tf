# arn:aws:ec2:ap-northeast-1:111111111111:traffic-mirror-filter-rule/traffic-mirror-filter-rule-id
output "ec2_traffic_mirror_filter_rule" {
  value = provider::arn::ec2_traffic_mirror_filter_rule("traffic-mirror-filter-rule-id")
}
