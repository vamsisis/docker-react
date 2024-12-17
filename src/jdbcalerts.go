{{ define "custom_title" }}
  {{ if gt (len .Alerts.Firing) 0 }}🚨 Firing alert(s): JDBC Connection Usage Alert: {{ .CommonLabels.alert_name }} - {{ .CommonLabels.job }} - {{ .CommonLabels.value }}%{{ end }}
  {{ if gt (len .Alerts.Resolved) 0 }}✅ Resolved alert(s): JDBC Connection Usage Alert: {{ .CommonLabels.alert_name }} - {{ .CommonLabels.job }} - {{ .CommonLabels.value }}%{{ end }}
{{ end }}

=============================================================
JDBC Connection Usage: [ var='A' labels={http_scheme=http, instance=0.0.0.0:33501, job=rps-rem-services-svcgr, name=db2, net_host_port=33501, service_instance_id=0.0.0.0:33501, service_name=rps-rem-services-svcgr} value=4 ], [ var='C' labels={http_scheme=http, instance=0.0.0.0:33501, job=rps-rem-services-svcgr, name=db2, net_host_port=33501, service_instance_id=0.0.0.0:33501, service_name=rps-rem-services-svcgr} value=1 ]%
===============================================================================
{{ define "CustomBodyforAlerting" -}}
{{ if or (gt (len .Alerts.Firing) 0) (gt (len .Alerts.Resolved) 0) -}}

{{ if gt (len .Alerts.Firing) 0 -}}
**Firing Alerts**
{{ range .Alerts.Firing -}}
**Alert Details:**
- Alert Name: {{ .Labels.alert_name }}
- Job Name: {{ .Labels.job }}
- JDBC Connection Usage: {{ index .Annotations "usage" }}%

**Debug Information:**
- All Labels:
{{ range $key, $value := .Labels -}}
  - {{ $key }}: {{ $value }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ if gt (len .Alerts.Resolved) 0 -}}
**Resolved Alerts**
{{ range .Alerts.Resolved -}}
**Alert Details:**
- Alert Name: {{ .Labels.alert_name }}
- Job Name: {{ .Labels.job }}
- JDBC Connection Usage: {{ index .Annotations "usage" }}%

**Debug Information:**
- All Labels:
{{ range $key, $value := .Labels -}}
  - {{ $key }}: {{ $value }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ else -}}
No active alerts.
{{ end -}}
{{ end -}}

No active alerts.
{{ end -}}
{{ end -}}
