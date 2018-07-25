#!/usr/bin/env sh

# Login to AWS
aws configure set region $AWS_REGION
$(aws ecr get-login)
# Tag and push docker image
docker tag $DOCKER_IMAGE $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$DOCKER_IMAGE:$CIRCLE_SHA1
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$DOCKER_IMAGE:$CIRCLE_SHA1
# Create task for docker deploy
task_template='[
{
  "name": "%s",
  "image": "%s.dkr.ecr.%s.amazonaws.com/%s:%s",
  "essential": true,
  "memoryReservation": 64,
  "portMappings": [
    {
      "containerPort": %s
    }
  ]
}
]'
task_def=$(printf "$task_template" $TASK $AWS_ACCOUNT_ID $AWS_REGION $DOCKER_IMAGE $CIRCLE_SHA1 $PORT)
echo "$task_def"
# Register task definition
json=$(aws ecs register-task-definition --container-definitions "$task_def" --family "$FAMILY")
echo "$json"
# Grab revision # using regular bash and grep
revision=$(echo "$json" | grep -o '"revision": [0-9]*' | grep -Eo '[0-9]+')
echo "$revision"
# Deploy revision
aws ecs update-service --cluster "$CLUSTER" --service "$SERVICE" --task-definition "$FAMILY":"$revision"
