#!/bin/sh

set -e
echo "STUB_SERVER: INFO: 00001: setup started."

# configure callback url
CALLBACK_URL="http://localhost:8065/hooks/xxx"

# login to mattermost
docker-compose exec chat \
mmctl auth login http://localhost:8065/ --name local-server --username mmadmin --password mm123admin

# create channel
docker-compose exec chat \
mmctl channel create --team devops --name chat-ops --display_name "ChatOps"

# get channel id
CHANNEL_ID=$(docker-compose exec chat mmctl channel search chat-ops 2>/dev/null | grep -oP '(?<=Channel ID :).+')

# create incoming-webhook
docker-compose exec chat \
mmctl webhook create-incoming \
--display-name "ChatOps IncomingWebhook" \
--description "Stub Server" \
--user mmadmin \
--channel $CHANNEL_ID

# create outgoing-webhook
docker-compose exec chat \
mmctl webhook create-outgoing \
--team devops \
--channel $CHANNEL_ID \
--user mmadmin \
--display-name "ChatOps OutgoingWebhook" \
--description "Stub Server" \
--trigger-word "#build" \
--trigger-word "#test" \
--url $CALLBACK_URL \
--content-type "application/json"

# setup completed
echo "STUB_SERVER: INFO: 00002: setup completed."