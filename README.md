# go-REST-simple-product

A simple REST API to manage products in a database.

This repo is intended for learning DevOps.

* A containirized REST API with a DB and integration tests
* CI/CD
* Deployed on a cloud

## Tools

* DB: PostgreSQL
* Routing: [gorilla/mux](https://github.com/gorilla/mux)
* Migrations: [migrate](https://github.com/golang-migrate/migrate)
* CI/CD: [circleci](https://circleci.com)
* Infrastructure as code: Terraform

## API Spec

[Postman Collection](https://documenter.getpostman.com/view/13097698/TVRoYmdb)

## Run

You need to have `golang-migrate` installed or use a tool of your choice.

```bash
brew install golang-migrate
```

```bash
# cd to the project dir
make run
```

## Test

You need to have `go` and `migrate` installed.

```bash
# clean the env first
make stop

# run tests
make test
```

## Env Vars

For production you need to set the following:

POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB

## CI/CD

### Integration

Workflow at circleci makes sure that the app builds successfully.

TODO: Use circleci docker orb to run `make test`

### Delivery

Workflow uses the `circleci/aws-ecr` orb to build and push the docker image to AWS/ECR.

Currently, the task is managed by terraform, each time we want to have another container live, we need to go to ECR, copy the URI and change the task definition in terraform.

For continuous delivery there is [this](https://stackoverflow.com/questions/64642858/continuous-deployment-and-delivery-on-ecs-fargate-with-circleci-and-terraform) option if I want to have terraform in sync. Alternatively, use circleci orb `aws/ecs` to manage the task definition and feed it with the latest image.
