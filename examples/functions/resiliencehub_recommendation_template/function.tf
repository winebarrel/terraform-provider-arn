# arn:aws:resiliencehub:ap-northeast-1:111111111111:recommendation-template/recommendation-template-id
output "resiliencehub_recommendation_template" {
  value = provider::arn::resiliencehub_recommendation_template("recommendation-template-id")
}
