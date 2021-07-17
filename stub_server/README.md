# Reference

* [Local Machine Setup using Docker](https://docs.mattermost.com/install/setting-up-local-machine-using-docker.html)
* [Incoming Webhookds](https://docs.mattermost.com/developer/webhooks-incoming.html)

# Environment

* OS: Ubuntu 20.04 LTS

# Installed Software

* Docker
* Docker-Compose

# Installation

~~~
mkdir mattermost-data mysql
docker-compose up -d
~~~

# Access to mattermost

* Please access below url.

~~~
http://localhost:8065
~~~

# Incoming Webhooks

## Configuration

* `[Integrations] > [Incoming Webhooks] > [Add Incoming Webhook]`
* Please configure below key-value.
  * Title: ChatOps
  * Description: Stub Server
  * Channel: Town Square
  * Lock to this channel: yes

## Tesing

* Please execute below command.

~~~
curl -i -X POST -H 'Content-Type: application/json' -d '{"text": "Hello, this is some text\nThis is more text. :tada:"}' http://{your-mattermost-site}/hooks/xxx-generatedkey-xxx
~~~

* expect

~~~
Hello, this is some text
This is more text. :tada:
~~~

# Outgoing Webhooks

## Configuration

* `[Integrations] > [Outgoing Webhooks] > [Add Outgoing Webhook]`
* Please configure below key-value.
  * Title: ChatOps
  * Description: Stub Server
  * Content Type: application/json
  * Channel: Town Square
  * Trigger Words (One Per Line): `#test`
  * Trigger When: First word matches a trigger word exactly
  * Callback URLs (One Per Line): `http://localhost:8065/hooks/<xxx-generated-key>`
* `[System Console] > [ENVIRONMENT] > [Developer]`
* Please configure below key-value.
  * Allow untrusted internal connections to: localhost

## Tesing

* Please submit below message at Town Square.

~~~
#test this is sample text.
~~~

* expect

~~~
#test this is sample text.
~~~

# Uninstallation

~~~
docker-compose down
~~~

