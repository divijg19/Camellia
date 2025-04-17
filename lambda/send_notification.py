import boto3

def lambda_handler(event, context):
    sns = boto3.client("sns")
    response = sns.publish(
        TopicArn="arn:aws:sns:region:account-id:your-topic",
        Message=event.get("message", "Hello from SkillShare!"),
        Subject="SkillShare Notification"
    )
    return {
        "statusCode": 200,
        "messageId": response["MessageId"]
    }
