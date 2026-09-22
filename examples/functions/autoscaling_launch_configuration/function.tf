# arn:aws:autoscaling:ap-northeast-1:111111111111:launchConfiguration:id:launchConfigurationName/launch-configuration-name
output "autoscaling_launch_configuration" {
  value = provider::arn::autoscaling_launch_configuration("id", "launch-configuration-name")
}
