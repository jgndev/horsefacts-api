```
go get github.com/aws/aws-lambda-go/events
go get github.com/aws/aws-lambda-go/lambda
```

## ChatGPT:

3.5 needed a lot of correction and prompting and made
mistakes after being asked to correct them (e.g. duplicates)

4.0 was better but stopped generating frequently. It did
however continue where it left off when clicking the 
'continue generating' button.

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
