# arn:aws:application-autoscaling:ap-northeast-1:111111111111:scalable-target/resource-id
output "application_autoscaling_scalable_target" {
  value = provider::arn::application_autoscaling_scalable_target("resource-id")
}
