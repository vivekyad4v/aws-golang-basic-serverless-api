# golang-serverless-basic-api

CICD Golang serverless API on AWS using SAM

This tutorial creates a Cloudformation stack, API gateway & two Lambda functions(Hello World & Stores) for demo purpose.

Make sure your CLI session has authorized access to AWS account.

### 1. Export necessary variables
``` export ORG_ID=foo
    export ENVIRON=uat
    export PROJECT_NAME=play-with-stores
```
Note - You do not need to export any variable for local development. You only need to change these variables while deploying it using CICD tools like Codepipeline, Jenkins, TravisCI etc.

### 2. Deploy locally

```
    make clean build configure run-local
```

### 3. Deploy on your AWS account

```
    make clean build configure package validate deploy describe 
```
