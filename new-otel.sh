#!/bin/bash

TARGET_TAG="v.35.9.3"
OUTPUT_FILE="otel_report.csv"

# Write CSV headers
echo "ClusterName,TaskArn,ImageTag" > "$OUTPUT_FILE"

for cluster_arn in $(aws ecs list-clusters --query "clusterArns" --output text); do
  cluster_name=$(basename "$cluster_arn")

  task_arns=$(aws ecs list-tasks --cluster "$cluster_arn" --desired-status RUNNING --query "taskArns" --output text)
  if [ -z "$task_arns" ]; then
    # No running tasks, still report
    echo "$cluster_name,N/A,N/A" >> "$OUTPUT_FILE"
    continue
  fi

  otel_found_in_cluster="no"
  task_descriptions=$(aws ecs describe-tasks --cluster "$cluster_arn" --tasks $task_arns \
    --query 'tasks[].containers[].[taskArn, name, image]' --output text)

  while read -r task_arn container_name container_image; do
    if [ "$container_name" = "otel-container" ]; then
      otel_found_in_cluster="yes"
      image_tag="${container_image##*:}"

      if [ "$image_tag" = "$TARGET_TAG" ]; then
        echo "$cluster_name,$task_arn,$image_tag" >> "$OUTPUT_FILE"
      fi
    fi
  done <<< "$task_descriptions"

  if [ "$otel_found_in_cluster" = "no" ]; then
    echo "$cluster_name,NONE,NONE" >> "$OUTPUT_FILE"
  fi
done
