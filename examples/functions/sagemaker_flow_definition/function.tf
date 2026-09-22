# arn:aws:sagemaker:ap-northeast-1:111111111111:flow-definition/flow-definition-name
output "sagemaker_flow_definition" {
  value = provider::arn::sagemaker_flow_definition("flow-definition-name")
}
