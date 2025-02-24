round(sum by(exception) (increase(http_server_requests_seconds_count{exception!="None", job="$job", $_instance_label_name=~"$instance"}[$__range]))) > 0
