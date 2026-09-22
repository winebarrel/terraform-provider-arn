# arn:aws:iotevents:ap-northeast-1:111111111111:detectorModel/detector-model-name
output "iotevents_detector_model" {
  value = provider::arn::iotevents_detector_model("detector-model-name")
}
