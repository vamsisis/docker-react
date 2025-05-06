#!/bin/bash

TARGET_TAG="v.35.9.3"

for cluster in $(aws ecs list-clusters --query "clusterArns" --output text); do
  task_arns=$(aws ecs list-tasks --cluster "$cluster" --desired-status RUNNING --query "taskArns" --output text)

  if [ -z "$task_arns" ]; then
    echo "No running tasks in Cluster: $cluster"
    continue
  fi

  otel_found_in_cluster="no"

  # Describe all tasks in one call
  task_descriptions=$(aws ecs describe-tasks --cluster "$cluster" --tasks $task_arns \
    --query 'tasks[].containers[].[taskArn, name, image]' --output text)

  # Loop through each line: taskArn containerName containerImage
  while read -r task_arn container_name container_image; do
    if [ "$container_name" = "otel-container" ]; then
      otel_found_in_cluster="yes"

      image_tag="${container_image##*:}"
      if [ "$image_tag" = "$TARGET_TAG" ]; then
        echo "otel-container FOUND | Cluster: $cluster | Task: $task_arn | Image Tag: $image_tag"
      fi
    fi
  done <<< "$task_descriptions"

  if [ "$otel_found_in_cluster" = "no" ]; then
    echo "otel-container NOT FOUND in Cluster: $cluster"
  fi
done
