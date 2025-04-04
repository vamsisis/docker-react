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


=============================================================
{map[%!f(string=Lo):%!f(string=ap)] 2.78}

{{ define "CustomBodyfor5XXAlerting" -}}
{{ if gt (len .Alerts.Firing) 0 -}}
**Firing 5XX Error Rate Alerts**
{{ range .Alerts.Firing -}}
{{ if eq .Labels.alert_name "5XX Error Rate" -}}
**Alert Details:**
- **Alert Name**: {{ .Labels.alert_name }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Total 5XX Errors**: 5XX Errors are above 10%
{{ end -}}
{{ end -}}

{{ else if gt (len .Alerts.Resolved) 0 -}}
**Resolved 5XX Error Rate Alerts**
{{ range .Alerts.Resolved -}}
{{ if eq .Labels.alert_name "5XX Error Rate" }}
**Alert Details:**
- **Alert Name**: {{ .Labels.alert_name }}
- **Load Balancer**: {{ .Labels.LoadBalancer }}
- **Total 5XX Errors**: 5XX Errors are above 10%
{{ end -}}
{{ end -}}

{{ else -}}
No active 5XX Error Rate alerts.
{{ end -}}
{{ end -}}

