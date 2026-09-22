# arn:aws:sagemaker:ap-northeast-1:111111111111:inference-experiment/inference-experiment-name
output "sagemaker_inference_experiment" {
  value = provider::arn::sagemaker_inference_experiment("inference-experiment-name")
}
