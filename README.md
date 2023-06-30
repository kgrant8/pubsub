# Go Cloud Pub/Sub Application

This is a sample application demonstrating the usage of [Go Cloud Development Kit (CDK)](https://gocloud.dev/) to interact with various cloud-based message brokers. The application features both a publisher for sending messages to a topic and a subscriber for receiving messages from a topic.

## Features

- Publish messages to any supported cloud-based Pub/Sub service (like Google Cloud Pub/Sub, AWS SNS/SQS, Azure Service Bus) based on a URL.
- Consume messages from any supported cloud-based Pub/Sub service based on a URL.
- Leverage the Go CDK, a powerful library for developing cross-cloud Go applications.

## Usage

1. Clone the repository:

    ```sh
    git clone https://github.com/kgrant8/pubsub.git
    ```

2. Build the application:

    ```sh
    go build .
    ```

3. Run the publisher:

    ```sh
    ./app publisher --url=<pubsub-url> --message="Hello, World!"
    ```

    Replace `<pubsub-url>` with the URL of the cloud-based Pub/Sub topic you want to publish to.

4. Run the subscriber:

    ```sh
    ./app consumer --topic=<pubsub-url>
    ```

    Replace `<pubsub-url>` with the URL of the cloud-based Pub/Sub subscription you want to consume from.

## Authentication

The Go CDK library leverages the standard authentication mechanisms provided by each cloud provider's SDK. Here are the specific mechanisms for Google Cloud, AWS, and Azure:

### Google Cloud Pub/Sub

Google's Application Default Credentials are used. The following sources are checked in order for credentials:

1. `GOOGLE_APPLICATION_CREDENTIALS` environment variable. This should be the path to the service account key file.
2. Default service account associated with the application when running on Google Cloud services.
3. User account authenticated with the `gcloud` tool when running the application locally.

For more information, see the [Authentication Overview](https://cloud.google.com/docs/authentication/) in the Google Cloud Documentation.

### AWS SNS/SQS

The AWS SDK's default credential provider chain is utilized. It looks for credentials in this sequence:

1. `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `AWS_SESSION_TOKEN` environment variables.
2. AWS credentials file. By default, the file is located at `~/.aws/credentials`.
3. IAM role for the EC2 instance, if the application is running on an Amazon EC2 instance.

For more information, see [Specifying Credentials](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials) in the AWS SDK for Go Developer Guide.

### Azure Service Bus

Azure uses Azure Active Directory (Azure AD) for authentication. Here are the methods it supports:

1. Environment variables: `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`.
2. Managed identities for Azure resources when running the application on Azure services.

For more information, refer to the [Authentication with Azure](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-go-authorization) documentation.

### Other Supported Platforms

CDK supports many other platforms. Read about them [here](https://pkg.go.dev/gocloud.dev/pubsub?utm_source=godoc)

## Contributing

Contributions are very much appreciated. Please feel free to submit a Pull Request.
