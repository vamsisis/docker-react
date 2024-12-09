
{{ define "HighTargetResponseTimeBody" -}}
**Alert Details:**
- Alert Name: High Target Response Time
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Target Response Time : {{ if .Values }}{{ range $refID, $value := .Values -}}{{ $value }}{{ end }}{{ else }}(Value unavailable){{ end }}
{{ if .Values }}
- Severity: {{ if ge (index .Values "B") 1.5 }}Critical (Urgent){{ else if ge (index .Values "B") 1.2 }}Critical{{ else if ge (index .Values "B") 0.8 }}Warning{{ else }}Info{{ end }}
{{ else }}
- Severity: (Severity unavailable)
{{ end }}
{{ end -}}

=======================================================================================
{{ define "CustomBodyforAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**Firing Alerts**
{{ range .Alerts.Firing -}}
**Alert Details:**
- Alert Name: {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current unhealthy host Count is : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High 5XX Errors" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current 5XX Errors are : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Request Count" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Request Count is : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Target Response Time : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}
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
- Current unhealthy host Count is : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High 5XX Errors" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current 5XX Errors are : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Request Count" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Request Count is : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- Load Balancer: {{ .Labels.LoadBalancer }}
- Current Target Response Time : {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}
- Severity: {{ if ge (index .Values "B") 1.5 }}Critical (Urgent){{ else if ge (index .Values "B") 1.2 }}Critical{{ else if ge (index .Values "B") 0.8 }}Warning{{ else }}Info{{ end }}

{{ else -}}
- Description: {{ .Annotations.description }}
{{ end -}}
{{ end -}}

{{ else -}}
No active alerts.
{{ end -}}
{{ end -}}
