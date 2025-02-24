round(sum by(exception) (increase(http_server_requests_seconds_count{exception!="None", job="$job", $_instance_label_name=~"$instance"}[$__range]))) > 0
=================================================================================

label_replace(
    sum by (method, uri, exception, status) (
        increase(
            http_server_requests_seconds_count{job="$job", $_instance_label_name=~"$instance", status=~"5..|400|499"}[$__range]
        )
    ) > 0, 
    'uri', 
    '$1/*$2', 
    'uri', 
    '(.*)\\/\\{.+\\}(.*)'
)
