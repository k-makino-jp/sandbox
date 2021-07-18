#!/bin/sh

set -e
echo "STUB_SERVER: INFO: 00001: setup started."

# configure callback url
CALLBACK_URL="http://localhost:8065/hooks/xxx"

# login to mattermost
mmctl auth login http://localhost:8065/ --name local-server --username mmadmin --password mm123admin

# create channel
mmctl channel create --team devops --name chat-ops --display_name "ChatOps"

# get channel id
CHANNEL_ID=$(mmctl channel search chat-ops 2>/dev/null | grep -oP '(?<=Channel ID :).+')

# create incoming-webhook
mmctl webhook create-incoming \
--channel $CHANNEL_ID \
--user mmadmin \
--display-name "ChatOps IncomingWebhook" \
--description "Stub Server"

# create outgoing-webhook
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