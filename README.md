```
go get github.com/aws/aws-lambda-go/events
go get github.com/aws/aws-lambda-go/lambda
```

Create the HorseFacts table

```
aws dynamodb create-table \
  --table-name HorseFacts \
  --attribute-definitions \
    AttributeName=ID,AttributeType=N \
  --key-schema \
    AttributeName=ID,KeyType=HASH \
  --provisioned-throughput \
    ReadCapacityUnits=5,WriteCapacityUnits=5
```

Create the HorseBreeds table

```
aws dynamodb create-table \
  --table-name HorseBreeds \
  --attribute-definitions \
    AttributeName=ID,AttributeType=N \
  --key-schema \
    AttributeName=ID,KeyType=HASH \
  --provisioned-throughput \
    ReadCapacityUnits=5,WriteCapacityUnits=5
```
go get github.com/spf13/viper


```bash
go build -o seed seed/seed.go
seed/seed

```
