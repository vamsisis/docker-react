aws_ecs_cluster_name
label_values(aws_ecs_cluster_name)

aws_ecs_task_family
label_values({aws_ecs_cluster_name="$aws_ecs_cluster_name"}, aws_ecs_task_family)

aws_ecs_task_id
label_values({aws_ecs_cluster_name="$aws_ecs_cluster_name",aws_ecs_task_family="$aws_ecs_task_family"}, aws_ecs_task_id)

===================================================================================================================================


container_cpu_usage_system_Nanoseconds{
  aws_ecs_cluster_name="$aws_ecs_cluster_name",
  aws_ecs_docker_name=~".*",
  aws_ecs_task_id="$aws_ecs_task_id",
  aws_ecs_task_family="$aws_ecs_task_family",
  container_image_tag=~".*"
}

===========================================================

max by (aws_ecs_cluster_name, aws_ecs_docker_name, container_image_tag) (
    container_memory_usage_Bytes{
        aws_ecs_cluster_name="$aws_ecs_cluster_name", 
        aws_ecs_task_id="$aws_ecs_task_id", 
        aws_ecs_task_family="$aws_ecs_task_family"
    }
)

