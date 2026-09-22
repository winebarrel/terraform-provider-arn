# arn:aws:ec2:ap-northeast-1:111111111111:traffic-mirror-session/traffic-mirror-session-id
output "ec2_traffic_mirror_session" {
  value = provider::arn::ec2_traffic_mirror_session("traffic-mirror-session-id")
}
