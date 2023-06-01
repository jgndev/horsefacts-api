# horsefacts-api

A REST API written in Go for retrieving facts about horses.

## Endpoints

`/api/facts`

`/api/breeds`

`/api/breeds/{id}`


### Create the HorseFacts table

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

### Create the HorseBreeds table

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


### Seeding the database

```bash
go build -o seed seed/seed.go
seed/seed

```

### Deployment to Elastic Beanstalk

```
eb init
eb create horsefacts-api-env
eb use horsefacts-api-env
eb create --platform docker --region us-east-1 --application horsefacts-api-docker --image jgndev/horsefacts-api:latest -t docker
eb setenv AWS_REGION=us-east-2 AWS_CLIENT_ID=<YOUR_CLIENT_ID> AWS_CLIENT_SECRET=<YOUR_CLIENT_SECRET> FACTS_TABLE=<YOUR_FACTS_TABLE> BREEDS_TABLE=<YOUR_BREEDS_TABLE>
eb deploy
```
