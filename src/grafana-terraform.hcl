{{ define "CustomBodyforAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**Firing Alerts**
{{ range .Alerts.Firing -}}
**Alert Details:**
- Alert Name: {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Unhealthy Host Count: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High 5XX Errors" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
{{ $totalRequests := index .Values "total_requests" }}
{{ $errorCount := index .Values "5xx_errors" }}
{{ $errorRate := (mul (div $errorCount $totalRequests) 100) | printf "%.2f" }}
{{ if gt (float64 $errorRate) 10.0 -}}
- Current 5XX Error Rate: {{ $errorRate }}% (Above Threshold)
{{ else -}}
- Current 5XX Error Rate: {{ $errorRate }}% (Within Threshold)
{{ end }}

{{ else if eq .Labels.alert_name "High Request Count" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Request Count: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Target Response Time: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}
- Severity: {{ if ge (index .Values "B") 1.5 }}Critical (Urgent){{ else if ge (index .Values "B") 1.2 }}Critical{{ else if ge (index .Values "B") 0.8 }}Warning{{ else }}Info{{ end }}

{{ else -}}
- Description: {{ .Annotations.description }}
{{ end -}}
{{ end -}}

{{ else if gt (len .Alerts.Resolved) 0 -}}
**Resolved Alerts**
{{ range .Alerts.Resolved -}}
**Alert Details:**
- Alert Name: {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Unhealthy Host Count: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High 5XX Errors" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
{{ $totalRequests := index .Values "total_requests" }}
{{ $errorCount := index .Values "5xx_errors" }}
{{ $errorRate := (mul (div $errorCount $totalRequests) 100) | printf "%.2f" }}
- Current 5XX Error Rate: {{ $errorRate }}%

{{ else if eq .Labels.alert_name "High Request Count" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Request Count: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Target Response Time: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}
- Severity: {{ if ge (index .Values "B") 1.5 }}Critical (Urgent){{ else if ge (index .Values "B") 1.2 }}Critical{{ else if ge (index .Values "B") 0.8 }}Warning{{ else }}Info{{ end }}

{{ else -}}
- Description: {{ .Annotations.description }}
{{ end -}}
{{ end -}}

{{ else -}}
No active alerts.
{{ end -}}
{{ end -}}
