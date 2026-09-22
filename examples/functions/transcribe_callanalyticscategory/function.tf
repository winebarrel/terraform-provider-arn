# arn:aws:transcribe:ap-northeast-1:111111111111:analytics-category/category-name
output "transcribe_callanalyticscategory" {
  value = provider::arn::transcribe_callanalyticscategory("category-name")
}
