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

aws_ecs_docker_name{aws_ecs_cluster_name="$aws_ecs_cluster_name", aws_ecs_task_id="$aws_ecs_task_id", aws_ecs_task_family="$aws_ecs_task_family"} 
or 
aws_ecs_container_image_id{aws_ecs_cluster_name="$aws_ecs_cluster_name", aws_ecs_task_id="$aws_ecs_task_id", aws_ecs_task_family="$aws_ecs_task_family"}

label_values({aws_ecs_cluster_name="$aws_ecs_cluster_name", aws_ecs_task_id="$aws_ecs_task_id"}, container_image_tag)
