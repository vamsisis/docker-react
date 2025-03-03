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


"__inputs": [
    {
      "name": "DS_CLOUDWATCH",
      "label": "CloudWatch",
      "description": "",
      "type": "datasource",
      "pluginId": "cloudwatch",
      "pluginName": "CloudWatch"
    }
  ],
  "__elements": {},
  "annotations": {
    "list": [
      {
        "datasource": "${DS_CLOUDWATCH}",
        "enable": true,
        "name": "Deployment Events",
        "query": {
          "queryMode": "Annotations",
          "refId": "Annotations"
        }
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 1,
  "id": null,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "datasource": "${DS_CLOUDWATCH}",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "options": {
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "showThresholdLabels": false,
        "showThresholdMarkers": true
      },
      "pluginVersion": "9.2.2",
      "targets": [
        {
          "alias": "Container Instances",
          "datasource": "${DS_CLOUDWATCH}",
          "dimensions": {
            "ClusterName": "$ClusterName"
          },
          "expression": "",
          "id": "",
          "matchExact": true,
          "metricName": "ContainerInstanceCount",
          "namespace": "AWS/ECS",
          "period": "30",
          "refId": "A",
          "region": "default",
          "statistics": [
            "Average"
          ]
        }
      ],
      "title": "Container Instance Count",
      "type": "gauge"
    },
    {
      "datasource": "${DS_CLOUDWATCH}",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 0
      },
      "id": 4,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single"
        }
      },
      "targets": [
        {
          "alias": "Tasks",
          "datasource": "${DS_CLOUDWATCH}",
          "dimensions": {
            "ClusterName": "$ClusterName"
          },
          "expression": "",
          "id": "",
          "matchExact": true,
          "metricName": "TaskCount",
          "namespace": "AWS/ECS",
          "period": "30",
          "refId": "A",
          "region": "default",
          "statistics": [
            "Sum"
          ]
        }
      ],
      "title": "Total Tasks in Cluster",
      "type": "timeseries"
    },
    {
      "datasource": "${DS_CLOUDWATCH}",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 8
      },
      "id": 6,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single"
        }
      },
      "targets": [
        {
          "alias": "Running Tasks",
          "datasource": "${DS_CLOUDWATCH}",
          "dimensions": {
            "ClusterName": "$ClusterName",
            "ServiceName": "$ServiceName"
          },
          "expression": "",
          "id": "",
          "matchExact": true,
          "metricName": "RunningTaskCount",
          "namespace": "AWS/ECS",
          "period": "30",
          "refId": "A",
          "region": "default",
          "statistics": [
            "Average"
          ]
        }
      ],
      "title": "Running Tasks (Service Level)",
      "type": "timeseries"
    },
    {
      "collapsed": false,
      "gridPos": {
        "h": 1,
        "w": 24,
        "x": 0,
        "y": 16
      },
      "id": 12,
      "panels": [],
      "title": "Resource Utilization",
      "type": "row"
    },
    {
      "datasource": "${DS_CLOUDWATCH}",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 10,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 17
      },
      "id": 8,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "multi"
        }
      },
      "targets": [
        {
          "alias": "CPU Utilized",
          "datasource": "${DS_CLOUDWATCH}",
          "dimensions": {
            "ClusterName": "$ClusterName",
            "ServiceName": "$ServiceName"
          },
          "expression": "",
          "id": "",
          "matchExact": true,
          "metricName": "CpuUtilized",
          "namespace": "AWS/ECS",
          "period": "30",
          "refId": "A",
          "region": "default",
          "statistics": [
            "Average"
          ]
        },
        {
          "alias": "CPU Reserved",
          "datasource": "${DS_CLOUDWATCH}",
          "dimensions": {
            "ClusterName": "$ClusterName",
            "ServiceName": "$ServiceName"
          },
          "expression": "",
          "id": "",
          "matchExact": true,
          "metricName": "CpuReserved",
          "namespace": "AWS/ECS",
          "period": "30",
          "refId": "B",
          "region": "default",
          "statistics": [
            "Average"
          ]
        }
      ],
      "title": "CPU Utilization",
      "type": "timeseries"
    },
    {
      "datasource": "${DS_CLOUDWATCH}",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 10,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 2,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "never",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 17
      },
      "id": 10,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "multi"
        }
      },
      "targets": [
        {
          "alias": "Memory Utilized",
          "datasource": "${DS_CLOUDWATCH}",
          "dimensions": {
            "ClusterName": "$ClusterName",
            "ServiceName": "$ServiceName"
          },
          "expression": "",
          "id": "",
          "matchExact": true,
          "metricName": "MemoryUtilized",
          "namespace": "AWS/ECS",
          "period": "30",
          "refId": "A",
          "region": "default",
          "statistics": [
            "Average"
          ]
        },
        {
          "alias": "Memory Reserved",
          "datasource": "${DS_CLOUDWATCH}",
          "dimensions": {
            "ClusterName": "$ClusterName",
            "ServiceName": "$ServiceName"
          },
          "expression": "",
          "id": "",
          "matchExact": true,
          "metricName": "MemoryReserved",
          "namespace": "AWS/ECS",
          "period": "30",
          "refId": "B",
          "region": "default",
          "statistics": [
            "Average"
          ]
        }
      ],
      "title": "Memory Utilization",
      "type": "timeseries"
    }
  ],
  "refresh": "30s",
  "schemaVersion": 37,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": [
      {
        "current": {
          "selected": false,
          "text": "default",
          "value": "default"
        },
        "hide": 0,
        "includeAll": false,
        "label": "Region",
        "multi": false,
        "name": "region",
        "options": [],
        "query": "default",
        "refresh": 1,
        "skipUrlSync": false,
        "type": "datasource"
      },
      {
        "definition": "dimension_values(default,AWS/ECS,ClusterName,ClusterName)",
        "hide": 0,
        "includeAll": false,
        "label": "Cluster Name",
        "multi": false,
        "name": "ClusterName",
        "options": [],
        "query": {
          "query": "dimension_values(default,AWS/ECS,ClusterName,ClusterName)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 1,
        "skipUrlSync": false,
        "type": "query"
      },
      {
        "definition": "dimension_values(default,AWS/ECS,ClusterName,ServiceName,ClusterName=$ClusterName)",
        "hide": 0,
        "includeAll": false,
        "label": "Service Name",
        "multi": false,
        "name": "ServiceName",
        "options": [],
        "query": {
          "query": "dimension_values(default,AWS/ECS,ClusterName,ServiceName,ClusterName=$ClusterName)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 1,
        "skipUrlSync": false,
        "type": "query"
      }
    ]
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "ECS Cluster Monitoring",
  "version": 1,
  "weekStart": ""
}
