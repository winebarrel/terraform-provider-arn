# arn:aws:ec2:ap-northeast-1:111111111111:traffic-mirror-target/traffic-mirror-target-id
output "ec2_traffic_mirror_target" {
  value = provider::arn::ec2_traffic_mirror_target("traffic-mirror-target-id")
}
