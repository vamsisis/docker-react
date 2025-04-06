{{ define "CustomBodyfor5XXAlerting" -}}

{{ if gt (len .Alerts.Firing) 0 }}
🔥 Firing Alerts
{{ range .Alerts.Firing }}
{{ if or (eq .Labels.alert_name "5XX Error Rate") (eq .Labels.alert_name "High Target Response Time") }}

Alert Name: {{ .Labels.alert_name }} for Load Balancer: {{ .Labels.LoadBalancer }}

{{ if eq .Labels.alert_name "5XX Error Rate" }}
Load Balancer: {{ .Labels.LoadBalancer }}
Total 5XX Errors: {{ .Labels.errors }}
Total Requests: {{ .Labels.total_requests }}
Current 5XX Error Rate: {{ .Labels.error_rate }}%
Threshold: >10%

{{ else if eq .Labels.alert_name "High Target Response Time" }}
Load Balancer: {{ .Labels.LoadBalancer }}
Target response time: {{ .Labels.response_time }}
Threshold: >1.5 seconds

{{ end }}
{{ end }}
{{ end }}

{{ else if gt (len .Alerts.Resolved) 0 }}
✅ Resolved Alerts
{{ range .Alerts.Resolved }}
{{ if or (eq .Labels.alert_name "5XX Error Rate") (eq .Labels.alert_name "High Target Response Time") }}

Alert Name: {{ .Labels.alert_name }} resolved for Load Balancer: {{ .Labels.LoadBalancer }}

{{ if eq .Labels.alert_name "5XX Error Rate" }}
Resolved 5XX Error Rate: {{ if .ValueString }}{{ .ValueString }}%{{ else }}0%{{ end }}

{{ else if eq .Labels.alert_name "High Target Response Time" }}
Resolved Target Response Time: {{ if .ValueString }}{{ .ValueString }}{{ else }}0s{{ end }}

{{ end }}
{{ end }}
{{ end }}

{{ else }}
No active or resolved alerts.
{{ end }}

{{ end }}



{{ define "loadbalancer" }}{{ if gt (len .Alerts.Firing) 0 }}[CRITICAL]{{ .CommonLabels.alert_name }} for {{ .CommonLabels.LoadBalancer }}{{ else if gt (len .Alerts.Resolved) 0 }}✅ Resolved alert(s):{{ .CommonLabels.alert_name }} for {{ .CommonLabels.LoadBalancer  }}{{ end }}{{ end }}
