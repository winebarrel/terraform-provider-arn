# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/retriever/retriever-id
output "qbusiness_retriever" {
  value = provider::arn::qbusiness_retriever("application-id", "retriever-id")
}
