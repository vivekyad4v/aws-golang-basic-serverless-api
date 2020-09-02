你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# aws-golang-serverless-basic-api

##### CICD Golang serverless API on AWS using SAM

##### This tutorial creates a Cloudformation stack, API gateway & two Lambda functions(Hello World & Stores) for demo purpose.

##### Make sure your CLI session has authorized access to AWS account.

#### 1. Export necessary variables
``` 
    export ORG_ID=foo
    export ENVIRON=uat
    export PROJECT_NAME=play-with-stores
    
```
You can change these variables while deploying it using CICD tools like Codepipeline, Jenkins, TravisCI etc.

#### 2. Deploy locally

```
    make clean build configure run-local
```

#### 3. Deploy on your AWS account (Docker must be pre-installed)

```
    make clean build configure package validate deploy describe outputs
```

You will get the API Gateway URL in the outputs something as below -    
 Ex - https://hahft1bb2c.execute-api.ap-south-1.amazonaws.com/uat

Check your APIs on above URL using curl -

```
    curl https://hahft1bb2c.execute-api.ap-south-1.amazonaws.com/uat/hello
    curl https://hahft1bb2c.execute-api.ap-south-1.amazonaws.com/uat/stores/store?type=grocery
    curl https://hahft1bb2c.execute-api.ap-south-1.amazonaws.com/uat/stores/store?type=fruit
    curl https://hahft1bb2c.execute-api.ap-south-1.amazonaws.com/uat/stores/store?type=all
    curl https://hahft1bb2c.execute-api.ap-south-1.amazonaws.com/uat/stores/store?type=doesnotexist
    curl https://hahft1bb2c.execute-api.ap-south-1.amazonaws.com/uat/whateverthrowerror
    
```

Note - You can get more details in Cloudwatch, Lambda, API Gateway console.


#### 4. Destroy everything

```
    make clean destroy 
```


