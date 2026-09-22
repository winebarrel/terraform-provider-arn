# arn:aws:sagemaker:ap-northeast-1:111111111111:human-loop/human-loop-name
output "sagemaker_human_loop" {
  value = provider::arn::sagemaker_human_loop("human-loop-name")
}
