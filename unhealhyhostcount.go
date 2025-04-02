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

{{ define "CustomBodyForAlerting" }}

{{ if gt (len .Alerts.Firing) 0 }}
**🚨 Unhealthy Targets Detected**

{{ range .Alerts.Firing }}
- **Alert Name**: {{ .Labels.alert_name }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Target Group**: {{ .Labels.TargetGroup }}
- **Unhealthy Host Count**: {{ index .Values "Value" }}

{{ end }}
{{ else }}
No firing alerts found.
{{ end }}

{{ end }}


