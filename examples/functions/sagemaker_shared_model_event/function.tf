# arn:aws:sagemaker:ap-northeast-1:111111111111:shared-model-event/event-id
output "sagemaker_shared_model_event" {
  value = provider::arn::sagemaker_shared_model_event("event-id")
}
