# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:membership/membership-id/ml-input-channel/resource-id
output "cleanrooms_ml_ml_input_channel" {
  value = provider::arn::cleanrooms_ml_ml_input_channel("membership-id", "resource-id")
}
