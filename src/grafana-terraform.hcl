{{ define "CustomBodyfor5XXAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**🚨 Firing 5XX Alerts 🚨**

{{ range .Alerts.Firing -}}
{{ if eq .Labels.alert_name "High 5XX Errors" -}}
**Alert Details:**
- **Load Balancer**: {{ .Labels.LoadBalancer }}

{{ $totalRequests := float64 (index .Values "total_requests") }}
{{ $errorCount := float64 (index .Values "5xx_errors") }}

{{ $errorRate := printf "%.2f" (mul (div $errorCount $totalRequests) 100) }}

{{ if gt (float64 $errorRate) 10.0 -}}
- **Current 5XX Error Rate**: 🚨 {{ $errorRate }}% (Above Threshold)
- **Severity**: Critical (Error rate > 10%)
{{ else -}}
- **Current 5XX Error Rate**: ✅ {{ $errorRate }}% (Within Threshold)
- **Severity**: Warning (Error rate <= 10%)
{{ end }}
{{ end -}}
{{ end -}}

{{ else if gt (len .Alerts.Resolved) 0 -}}
**✅ Resolved 5XX Alerts ✅**

{{ range .Alerts.Resolved -}}
{{ if eq .Labels.alert_name "High 5XX Errors" -}}
**Alert Details:**
- **Load Balancer**: {{ .Labels.LoadBalancer }}

{{ $totalRequests := float64 (index .Values "total_requests") }}
{{ $errorCount := float64 (index .Values "5xx_errors") }}

{{ $errorRate := printf "%.2f" (mul (div $errorCount $totalRequests) 100) }}

- **Resolved 5XX Error Rate**: {{ $errorRate }}%
- **Status**: Resolved ✅
{{ end -}}
{{ end -}}

{{ else -}}
No active 5XX alerts at the moment.
{{ end -}}
{{ end -}}
