{{ define "CustomBodyForAlerting" }}

{{ if gt (len .Alerts.Firing) 0 }}
**🚨 Firing Alerts 🚨**

{{ range .Alerts.Firing }}
**Alert Name:** {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" }}

- **Load Balancer:** {{ .Labels.LoadBalancer }}

  {{ range $k, $v := .Values }}
  - **Target Group / Metric Key:** {{ $k }}
    - **Unhealthy Host Count:** {{ $v }}
  {{ end }}

{{ end }}
{{ end }}
{{ end }}

{{ if gt (len .Alerts.Resolved) 0 }}
**✅ Resolved Alerts**

{{ range .Alerts.Resolved }}
**Alert Name:** {{ .Labels.alert_name }}

{{ if eq .Labels.alert_name "unhealthyhost" }}

- **Load Balancer:** {{ .Labels.LoadBalancer }}

  {{ range $k, $v := .Values }}
  - **Target Group / Metric Key:** {{ $k }}
    - **Resolved Unhealthy Host Count:** {{ $v }}
  {{ end }}

{{ end }}
{{ end }}
{{ end }}

{{ end }}


======================
{{ define "CustomBodyfor5XXAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**Firing Alerts**
{{ range .Alerts.Firing -}}
{{ if or (eq .Labels.alert_name "unhealthyhost") (eq .Labels.alert_name "High Request Count") (eq .Labels.alert_name "High Target Response Time") (eq .Labels.alert_name "5XX Error Rate") -}}
**Alert Details:**
- **Alert Name**: {{ .Labels.alert_name }}

{{ else if eq .Labels.alert_name "5XX Error Rate" }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Total 5XX Errors**: 5XX Errors are above 10% 

{{ end -}}
{{ end -}}
{{ end -}}


{{ else if gt (len .Alerts.Resolved) 0 -}}
**Resolved Alerts**
{{ range .Alerts.Resolved -}}
{{ if or (eq .Labels.alert_name "5XX Error Rate") -}}
**Alert Details:**

- **Alert Name**: {{ .Labels.alert_name }}
{{ else if eq .Labels.alert_name "5XX Error Rate" }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Total 5XX Errors**: - **Total 5XX Errors**: 5XX Errors are above 10% 

{{ end -}}
{{ end -}}
{{ end -}}

{{ else -}}
No active alerts.
{{ end -}}
{{ end -}}


