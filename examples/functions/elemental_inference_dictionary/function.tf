# arn:aws:elemental-inference:ap-northeast-1:111111111111:dictionary/id
output "elemental_inference_dictionary" {
  value = provider::arn::elemental_inference_dictionary("id")
}
