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


import boto3

def check_clusters(region):
    client = boto3.client("ecs", region_name=region)
    
    try:
        # Get all clusters
        clusters = client.list_clusters()["clusterArns"]
        if not clusters:
            print(f"No ECS clusters found in region {region}.")
            return
        
        for cluster in clusters:
            print(f"Checking cluster: {cluster}")
            
            # Get running tasks
            tasks = client.list_tasks(cluster=cluster)["taskArns"]
            if not tasks:
                print(f"  No running tasks in cluster {cluster}.")
            else:
                print(f"  Running tasks found in cluster {cluster}:")
                for task in tasks:
                    print(f"    {task}")
    except Exception as e:
        print(f"Error: {e}")

# Replace with your AWS region
region = "your-region"
check_clusters(region)
