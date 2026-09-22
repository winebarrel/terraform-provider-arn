# arn:aws:codeguru-profiler:ap-northeast-1:111111111111:profilingGroup/profiling-group-name
output "codeguru_profiler_profiling_group" {
  value = provider::arn::codeguru_profiler_profiling_group("profiling-group-name")
}
