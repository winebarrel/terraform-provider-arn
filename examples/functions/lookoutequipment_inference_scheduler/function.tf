# arn:aws:lookoutequipment:ap-northeast-1:111111111111:inference-scheduler/inference-scheduler-name/inference-scheduler-id
output "lookoutequipment_inference_scheduler" {
  value = provider::arn::lookoutequipment_inference_scheduler("inference-scheduler-name", "inference-scheduler-id")
}
