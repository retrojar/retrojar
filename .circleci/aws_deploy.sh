#!/usr/bin/env sh

aws configure set region $AWS_REGION
$(aws ecr get-login)

server_img=$(printf "%s.dkr.ecr.%s.amazonaws.com/server:%s" $AWS_ACCOUNT_ID $AWS_REGION $CIRCLE_SHA1)
client_img=$(printf "%s.dkr.ecr.%s.amazonaws.com/client:%s" $AWS_ACCOUNT_ID $AWS_REGION $CIRCLE_SHA1)

docker tag server $server_img
docker tag client $client_img
docker push $server_img
docker push $client_img

task_template='
[{
    "name": "server",
    "image": "%s",
    "essential": true,
    "memoryReservation": 64
},
{
    "name": "client",
    "image": "%s",
    "essential": true,
    "memoryReservation": 64,
    "portMappings": [
        {
            "containerPort": 80
        }
    ],
    "environment": [
        {
            "name": "API_HOST",
            "value": "https://api.retrojar.top"
        }
    ]
}]
'

task_def=$(printf "$task_template" $server_img $client_img)
echo "$task_def"

json=$(aws ecs register-task-definition --container-definitions "$task_def" --family "retrojar")
echo "$json"

revision=$(echo "$json" | grep -o '"revision": [0-9]*' | grep -Eo '[0-9]+')
echo "$revision"

aws ecs update-service --cluster "default" --service "retrojar" --task-definition "retrojar":"$revision"
