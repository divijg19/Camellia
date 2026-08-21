Local lambda helper: how to install deps and test

1. Create a virtual environment and activate it:

```bash
uv venv
source .venv/bin/activate
```

2. Install dependencies:

```bash
uv sync --upgrade --all-extras
```

3. Run a quick import test:

```bash
python -c "import boto3; print('boto3', boto3.__version__)"
```

4. To run the handler locally:

```bash
python -c "from send_notification import lambda_handler; print(lambda_handler({'message':'hi'}, None))"
```

Note: AWS credentials and correct `SNS_TOPIC_ARN`/`AWS_REGION` environment variables are required to publish to SNS.
