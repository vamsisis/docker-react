#!/bin/bash

for cluster in $(aws ecs list-clusters --query "clusterArns" --output text); do
  otel_running="no"

  # list running task in the cluster
  task_arns=$(aws ecs list-tasks --cluster "$cluster" --desired-status RUNNING --query "taskArns" --output text)

  # if no tasks are found, skip to the next cluster
  if [ -z "$task_arns" ]; then
    continue
  fi

  # check if any task has the 'otel-container' running
  for task in $task_arns; do
    container_names=$(aws ecs describe-tasks --cluster "$cluster" --tasks "$task" --query "tasks[].containers[].name" --output text)


    if echo "$container_names" | grep -qw "otel-container"; then
      otel_running="yes"
      break
    fi
  done

  # output if 'otel-container' is not found

  if [ "$otel_running" = "no" ]; then
    echo "otel-container is NOT running in Cluster: $cluster"
  fi
done
