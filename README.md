# Project structures

```
.
├── README.md                           <-- This instructions file
├── Makefile                            <-- Make to automate build
├── bin                                 <-- binary Go files
├── docs/                               <-- user documents
├── cmd/                                <-- main packages
│   └── custom                          <-- custom lambda function
├── scripts/                            <-- executable build/deployment scripts
├── pkg/                                <-- main packages for exposable business logic.
├── go.mod/                             <-- go module
├── go.sum                              <-- module hash
└── template.yaml                       <-- SAM template
```

## 要求　(Requirements)

* AWS CLI already configured with Administrator permission
* [AWS local CLI](https://github.com/localstack/awscli-local)
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process
- run as following in order:

```shell script
TMPDIR=/private$TMPDIR docker-compose up
```

- build lambda function
```makefile
make build-local
```

- init custom stack in localstak
```makefile
make setup-local
```

## Push message to queue
- Test push message to sqs queue

```shell script
./scripts/send-message-to-queue.sh "Test message"
```

## Cleanup
- kill daemon which you launched, and clean up build artifacts.

```shell
make clean
```

### Testing

- Put file to bucket
```shell script
awslocal s3 ls s3://custom-bucket
awslocal s3 ls s3://lambda-bucket
```

