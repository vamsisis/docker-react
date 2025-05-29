import json
import urllib.request

def lambda_handler(event, context):
    # Replace with your xMatters inbound integration URL
    xmatters_url = 'https://your-xmatters-url.example.com'

    # Basic payload to test xMatters
    payload = {
        "message": "Test message from AWS Lambda",
        "source": "lambda-function",
        "timestamp": context.aws_request_id  # Optional: use AWS request ID as a unique ref
    }

    data = json.dumps(payload).encode('utf-8')
    headers = {
        'Content-Type': 'application/json'
    }

    try:
        req = urllib.request.Request(xmatters_url, data=data, headers=headers)
        with urllib.request.urlopen(req) as response:
            result = response.read().decode('utf-8')
        print(f"xMatters response: {result}")
        return {
            'statusCode': 200,
            'body': 'Payload sent successfully to xMatters.'
        }
    except Exception as e:
        print(f"Error sending to xMatters: {e}")
        return {
            'statusCode': 500,
            'body': 'Failed to send payload to xMatters.'
        }
