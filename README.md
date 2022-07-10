# EC2Search
CLI tool to help you search EC2 instances with AWS CLI


## Caveats

1. You need to have  the AWS CLI installed.
2. Have the aws cli configured `aws configure` 
3. try some simple commands like `aws ec2 describe-instances` to verify that you have necassary permissions to list the instances.
4. If you have the named profiles you can pass it as an option to search for the instance using specific profile ( otherwise default would be taken)
5. EC2 Search uses the aws cli tool, its just a wrapper, EC2 Search does not directly connect to AWS or use your security credentials anyway.



## How to use it.

Download the right binary or download the source code or build it by running the following commands


```

git clone https://github.com/AKSarav/EC2Search.git

# If you want to use the default profile 
go run . <SEARCH_STRING>

# If you have multiple named profiles
go run . <SEARCH_STRING> <AWS CLI PROFILE>

```


