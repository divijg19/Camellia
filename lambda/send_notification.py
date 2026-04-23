import os
import boto3
from botocore.exceptions import ClientError


def lambda_handler(event, context):
    """Publish a message to SNS.

    Reads `SNS_TOPIC_ARN` and optional `AWS_REGION` from environment variables.
    Returns a JSON-like dict with status and message id or error details.
    """
    topic_arn = os.environ.get(
        "SNS_TOPIC_ARN", "arn:aws:sns:region:account-id:your-topic"
    )
    region = os.environ.get("AWS_REGION")

    try:
        sns = boto3.client("sns", region_name=region) if region else boto3.client("sns")

        response = sns.publish(
            TopicArn=topic_arn,
            Message=event.get("message", "Hello from Camellia!"),
            Subject="Camellia Notification",
        )

        return {"statusCode": 200, "messageId": response.get("MessageId")}

    except ClientError as e:
        return {"statusCode": 500, "error": str(e)}
    except Exception as e:
        return {"statusCode": 500, "error": str(e)}
