def lambda_handler(event, context):
    user = event.get("user")
    amount = event.get("amount")
    
    # Simulate payment processing
    return {
        "statusCode": 200,
        "body": f"Payment of ${amount} processed for {user}"
    }
