# arn:aws:autoscaling:ap-northeast-1:111111111111:autoScalingGroup:group-id:autoScalingGroupName/group-friendly-name
output "autoscaling_auto_scaling_group" {
  value = provider::arn::autoscaling_auto_scaling_group("group-id", "group-friendly-name")
}
