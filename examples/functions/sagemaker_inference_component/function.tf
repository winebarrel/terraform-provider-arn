# arn:aws:sagemaker:ap-northeast-1:111111111111:inference-component/inference-component-name
output "sagemaker_inference_component" {
  value = provider::arn::sagemaker_inference_component("inference-component-name")
}
