#!/bin/bash

for cluster in $(aws ecs list-clusters --query "clusterArns" --output text); do
  task_arns=$(aws ecs list-tasks --cluster "$cluster" --desired-status RUNNING --query "taskArns" --output text)

  # Skip if no tasks
  if [ -z "$task_arns" ]; then
    echo "No running tasks in Cluster: $cluster"
    continue
  fi

  # Describe all tasks at once
  task_info=$(aws ecs describe-tasks --cluster "$cluster" --tasks $task_arns --query "tasks[].{TaskArn:taskArn,Containers:containers}" --output text)

  otel_found="no"

  while read -r line; do
    # Line will be like: taskArn containerName containerImage ...
    task_arn=$(echo "$line" | awk '{print $1}')
    container_name=$(echo "$line" | awk '{print $2}')
    container_image=$(echo "$line" | awk '{print $3}')

    if [ "$container_name" = "otel-container" ]; then
      echo "otel-container FOUND | Cluster: $cluster | Task: $task_arn | Image: $container_image"
      otel_found="yes"
    fi
  done <<< "$(aws ecs describe-tasks --cluster "$cluster" --tasks $task_arns \
    --query 'tasks[].containers[].[taskArn, name, image]' --output text)"

  if [ "$otel_found" != "yes" ]; then
    echo "otel-container NOT FOUND in Cluster: $cluster"
  fi
done
