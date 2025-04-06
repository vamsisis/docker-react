{{ define "CustomBodyforAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**Firing Alerts**
{{ range .Alerts.Firing -}}
{{ if or (eq .Labels.alert_name "5XX Error Rate") (eq .Labels.alert_name "High Request Count") (eq .Labels.alert_name "High Target Response Time") -}}
**Alert Details:**
- **Alert Name**: Alert: {{ .Labels.alert_name }} for {{ .Labels.LoadBalancer }}

{{ if eq .Labels.alert_name "5XX Error Rate" -}}
Description:
- Load Balancer: {{ .Labels.LoadBalancer }}
- Total 5XX Errors: {{ .Labels.errors }}
- Total Requests: {{ .Labels.total_requests }}
- Current 5XX Error Rate: {{ .Labels.error_rate }}%
- Threshold: >10%

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Current Target Response Time**: {{ .ValueString }}
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
- **Resolved Unhealthy Host Count**: {{ if .ValueString }}{{ .ValueString }}{{ else }}0{{ end }}

{{ else if eq .Labels.alert_name "High Request Count" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Resolved Request Count**: {{ if .ValueString }}{{ .ValueString }}{{ else }}0{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" -}}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Resolved Target Response Time**: {{ if .ValueString }}{{ .ValueString }}{{ else }}0{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ else -}}
No active alerts.
{{ end -}}
{{ end -}}
