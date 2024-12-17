
{{ define "CustomBodyforAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**Firing Alerts**
{{ range .Alerts.Firing -}}
{{ if or (eq .Labels.alert_name "unhealthyhost") (eq .Labels.alert_name "High Request Count") (eq .Labels.alert_name "High Target Response Time") -}}
**Alert Details:**
- **Alert Name**: {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Current Unhealthy Host Count**: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Request Count" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Current Request Count**: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Current Target Response Time**: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}
- **Severity**: 
  {{ if ge (index .Values "B") 1.5 }}Critical (Urgent){{ else if ge (index .Values "B") 1.2 }}Critical{{ else if ge (index .Values "B") 0.8 }}Warning{{ else }}Info{{ end }}

{{ end -}}
{{ end -}}
{{ end -}}

{{ else if gt (len .Alerts.Resolved) 0 -}}
**Resolved Alerts**
{{ range .Alerts.Resolved -}}
{{ if or (eq .Labels.alert_name "unhealthyhost") (eq .Labels.alert_name "High Request Count") (eq .Labels.alert_name "High Target Response Time") -}}
**Alert Details:**
- **Alert Name**: {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Resolved Unhealthy Host Count**: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Request Count" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Resolved Request Count**: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Resolved Target Response Time**: {{ range $refID, $value := .Values -}}{{ $value }}{{ end }}
- **Severity**: 
  {{ if ge (index .Values "B") 1.5 }}Critical (Urgent){{ else if ge (index .Values "B") 1.2 }}Critical{{ else if ge (index .Values "B") 0.8 }}Warning{{ else }}Info{{ end }}

{{ end -}}
{{ end -}}
{{ end -}}

{{ else -}}
No active alerts.
{{ end -}}
{{ end -}}
