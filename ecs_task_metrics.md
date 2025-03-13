# Task Metrics 

ecs_task_memory_reserved_Megabytes
ecs_task_memory_usage_Bytes
ecs_task_memory_usage_limit_Bytes
ecs_task_memory_usage_max_Bytes
ecs_task_memory_utilized_Megabytes
ecs_task_network_io_usage_rx_bytes_Bytes
ecs_task_network_io_usage_rx_dropped_Count
ecs_task_network_io_usage_rx_errors_Count
ecs_task_network_io_usage_rx_packets_Count
ecs_task_network_io_usage_tx_bytes_Bytes
ecs_task_network_io_usage_tx_dropped_Count
ecs_task_network_io_usage_tx_errors_Count
ecs_task_network_io_usage_tx_packets_Count
ecs_task_network_rate_rx_Bytes_per_Second
ecs_task_network_rate_tx_Bytes_per_Second
ecs_task_storage_read_bytes_Bytes
ecs_task_storage_write_bytes_Bytes

# Availabe Labels

aws_ecs_cluster_name
aws_ecs_container_created_at
aws_ecs_container_exit_code
aws_ecs_container_finished_at
aws_ecs_container_image_id
aws_ecs_container_know_status
aws_ecs_container_started_at
aws_ecs_docker_name
aws_ecs_launchtype
aws_ecs_task_arn
aws_ecs_task_family
aws_ecs_task_id
aws_ecs_task_known_status
aws_ecs_task_launch_type
aws_ecs_task_pull_started_at
aws_ecs_task_pull_stopped_at
aws_ecs_task_revision
aws_ecs_task_version

# Container metrics 

container_cpu_usage_kernelmode_Nanosecondscontainer_cpu_onlines_Countcontainer_cpu_reserved_Nonecontainer_cpu_utilized_Nonecontainer_network_rate_tx_Bytes_per_Secondcontainer_network_io_usage_rx_errors_Countcontainer_network_io_usage_rx_dropped_Countcontainer_network_io_usage_tx_dropped_Countcontainer_network_io_usage_rx_bytes_Bytescontainer_memory_usage_limit_Bytescontainer_memory_usage_Bytescontainer_network_io_usage_tx_packets_Countcontainer_storage_write_bytes_Bytescontainer_cpu_cores_Countcontainer_cpu_usage_vcpu_vCPUcontainer_cpu_usage_usermode_Nanosecondscontainer_memory_reserved_Megabytescontainer_cpu_usage_total_Nanosecondscontainer_memory_usage_max_Bytescontainer_network_io_usage_tx_bytes_Bytescontainer_memory_utilized_Megabytescontainer_network_io_usage_rx_packets_Countcontainer_duration_Secondscontainer_cpu_usage_system_Nanosecondscontainer_network_rate_rx_Bytes_per_Secondcontainer_storage_read_bytes_Bytes

100 * (container_memory_utilized_Megabytes{aws_ecs_cluster_name="$aws_ecs_cluster_name"} / 
       (container_memory_usage_limit_Bytes{aws_ecs_cluster_name="$aws_ecs_cluster_name"} / 1024 / 1024)) > 80

100 * (rate(container_cpu_usage_total_Nanoseconds{aws_ecs_cluster_name="$aws_ecs_cluster_name"}[1m]) / 
       (container_cpu_cores_Count{aws_ecs_cluster_name="$aws_ecs_cluster_name"} * 1e9)) > 80


aws ecs list-clusters --query "clusterArns" --output text | xargs -n1 -I {} sh -c 'echo "Cluster: {}"; aws ecs list-tasks --cluster {} --query "taskArns" --output text | grep . || echo "  No running tasks in this cluster"'


aws ecs list-clusters --query "clusterArns" --output text | xargs -n1 -I {} sh -c 'if [ -z "$(aws ecs list-tasks --cluster {} --query "taskArns" --output text)" ]; then echo {}; fi'


#!/bin/bash

# Specify the AWS region
AWS_REGION="your-region"

# Get the list of all ECS clusters
clusters=$(aws ecs list-clusters --region "$AWS_REGION" --query "clusterArns" --output text)

if [[ -z "$clusters" ]]; then
  echo "No ECS clusters found in region $AWS_REGION."
  exit 0
fi

echo "Checking ECS clusters in region $AWS_REGION..."

# Loop through each cluster and check for running tasks
for cluster in $clusters; do
  # Get the list of running tasks in the cluster
  tasks=$(aws ecs list-tasks --cluster "$cluster" --region "$AWS_REGION" --query "taskArns" --output text)
  
  # If no tasks are running, print the cluster name
  if [[ -z "$tasks" ]]; then
    echo "Cluster with no running tasks: $cluster"
  fi
done



aws_ecs_container_image_id * on(aws_ecs_task_id) group_right() aws_ecs_task_known_status{aws_ecs_task_known_status="RUNNING"}

aws_ecs_cluster_name="ewc-svcgr",container_image_name="raas-docker-snapshot-local.docker.fis.dev/com/fis/raas/monitoring/aws-otel-collector",container_image_tag="v.35.9.5"}

aws_ecs_container_image_id{
    aws_ecs_cluster_name="ewc-svcgr",
    container_image_name="raas-docker-snapshot-local.docker.fis.dev/com/fis/raas/monitoring/aws-otel-collector",
    container_image_tag="v.35.9.5"
} 
* on(aws_ecs_task_id) group_right() aws_ecs_task_known_status{aws_ecs_task_known_status="RUNNING"}
=========================================================

{{ define "rdshighcpuusage" -}} 
{{ if gt (len .Alerts.Firing) 0 -}} 🚨 High CPU Utilization Alert 🚨

{{ range .Alerts.Firing -}} 
{{ if eq .Labels.alert_name "High CPU Utilization" -}} 
Alert Details:

Instance: {{ .Labels.Instance }}
Current CPU Utilization: 🚨 {{ index .Values "E" }}% (Above Threshold) 
{{ end -}} 
{{ end -}}

{{ else if gt (len .Alerts.Resolved) 0 -}} ✅ CPU Utilization Resolved ✅

{{ range .Alerts.Resolved -}} 
{{ if eq .Labels.alert_name "High CPU Utilization" -}} 
Alert Details:

Instance: {{ .Labels.Instance }}
Resolved CPU Utilization: {{ index .Values "E" }}%
Status: Resolved ✅ 
{{ end -}} 
{{ end -}}

{{ else -}} No active CPU utilization alerts at the moment. 
{{ end -}} 
{{ end -}}

(sum by(api) (http_server_requests_seconds_count{status=~"2.."}))
/
(sum by(api) (http_server_requests_seconds_count))
* 100


(sum by(api) (http_server_requests_seconds_count{status=~"4..|5.."}))
/
(sum by(api) (http_server_requests_seconds_count))
* 100


histogram_quantile(0.95, sum(rate(http_server_requests_seconds_bucket[5m])) by (le, api))


sum by(api) (rate(http_server_requests_seconds_count[5m]))


(sum by(api) (http_server_requests_seconds_count{status=~"5.."})) > 10

=================================================================

aws_ecs_cluster_name{aws_ecs_cluster_name=~"$cluster"}
  * on(aws_ecs_task_id) group_right() aws_ecs_container_image_id
  * on(aws_ecs_task_id) group_right() container_image_name{container_image_name=~"$image_name"}
  * on(aws_ecs_task_id) group_right() container_image_tag{container_image_tag=~"$image_tag"}

=========================================================================

aws_ecs_task_known_status{aws_ecs_task_known_status="RUNNING", aws_ecs_cluster_name=~"$cluster"}
  * on(aws_ecs_task_id) group_right()
  label_replace(aws_ecs_container_image_id, "container_image_id", "$1", "aws_ecs_container_image_id", "(.*)")
  * on(aws_ecs_task_id) group_right()
  label_replace(container_image_name, "image_name", "$1", "container_image_name", "(.*)")
  * on(aws_ecs_task_id) group_right()
  label_replace(container_image_tag, "image_tag", "$1", "container_image_tag", "(.*)")

======================================

groups:
  - name: ecs_oom_alerts
    rules:
      - alert: ECSOOMAlert
        expr: (rate(ecs_task_memory_usage_max_Bytes[5m]) / ecs_task_memory_usage_limit_Bytes) * 100 > 90
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "ECS Task OOM Alert"
          description: "Task {{ $labels.aws_ecs_task_id }} in Cluster {{ $labels.aws_ecs_cluster_name }} is running out of memory."
