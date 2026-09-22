# arn:aws:elemental-inference:ap-northeast-1:111111111111:feed/id
output "elemental_inference_feed" {
  value = provider::arn::elemental_inference_feed("id")
}
