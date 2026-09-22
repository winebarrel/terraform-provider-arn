# arn:aws:health:*::event/service/event-type-code/*
output "health_event" {
  value = provider::arn::health_event("service", "event-type-code")
}
