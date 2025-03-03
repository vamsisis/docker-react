ClusterName

MemoryReserved
StorageReadBytes
EphemeralStorageReserved
StorageWriteBytes
TaskCount
NetworkRxBytes
EphemeralStorageUtilized
ServiceCount
CpuReserved
NetworkTxBytes
MemoryUtilized
CpuUtilized
ContainerInstanceCount


ClusterName,TaskDefinitionFamily

EphemeralStorageUtilized
StorageWriteBytes
NetworkTxBytes
EphemeralStorageReserved
StorageReadBytes
CpuUtilized
NetworkRxBytes
MemoryReserved
CpuReserved
MemoryUtilized

ClusterName,SericeName

CpuUtilized - completed 
DesiredTaskCount - completed 
TaskSetCount
EphemeralStorageReserved
NetworkRxBytes
DeploymentCount
EphemeralStorageUtilized
NetworkTxBytes
PendingTaskCount - completed
RunningTaskCount - completed 
MemoryUtilized
StorageReadBytes
CpuReserved
StorageWriteBytes
MemoryReserved

# Variables 
ClusterName

ServiceName
Dimesions - clusterName

tasks
taskDefinitionFamily clusterName



{
  "dashboard": {
    "id": null,
    "uid": null,
    "title": "ECS Container Insights Dashboard",
    "tags": ["ECS", "ContainerInsights", "CloudWatch"],
    "timezone": "browser",
    "schemaVersion": 36,
    "version": 0,
    "panels": [
      {
        "title": "CPU Utilization",
        "type": "graph",
        "targets": [
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "CpuUtilized",
            "dimensions": {
              "ClusterName": "$ClusterName",
              "ServiceName": "$ServiceName"
            },
            "stat": "Average",
            "period": 60
          }
        ]
      },
      {
        "title": "Memory Utilization",
        "type": "graph",
        "targets": [
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "MemoryUtilized",
            "dimensions": {
              "ClusterName": "$ClusterName",
              "ServiceName": "$ServiceName"
            },
            "stat": "Average",
            "period": 60
          }
        ]
      },
      {
        "title": "Network Traffic",
        "type": "graph",
        "targets": [
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "NetworkRxBytes",
            "dimensions": {
              "ClusterName": "$ClusterName"
            },
            "stat": "Sum",
            "period": 60
          },
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "NetworkTxBytes",
            "dimensions": {
              "ClusterName": "$ClusterName"
            },
            "stat": "Sum",
            "period": 60
          }
        ]
      },
      {
        "title": "Storage Metrics",
        "type": "graph",
        "targets": [
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "StorageReadBytes",
            "dimensions": {
              "ClusterName": "$ClusterName"
            },
            "stat": "Sum",
            "period": 60
          },
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "StorageWriteBytes",
            "dimensions": {
              "ClusterName": "$ClusterName"
            },
            "stat": "Sum",
            "period": 60
          }
        ]
      },
      {
        "title": "Task and Service Monitoring",
        "type": "table",
        "targets": [
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "RunningTaskCount",
            "dimensions": {
              "ClusterName": "$ClusterName"
            },
            "stat": "Maximum",
            "period": 60
          },
          {
            "namespace": "ECS/ContainerInsights",
            "metricName": "PendingTaskCount",
            "dimensions": {
              "ClusterName": "$ClusterName"
            },
            "stat": "Maximum",
            "period": 60
          }
        ]
      }
    ],
    "templating": {
      "list": [
        {
          "name": "ClusterName",
          "type": "query",
          "query": "ECS/ContainerInsights",
          "refresh": 1
        },
        {
          "name": "ServiceName",
          "type": "query",
          "query": "ECS/ContainerInsights",
          "refresh": 1
        }
      ]
    }
  }
}


