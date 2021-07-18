
# Webhooks

## Configure incoming and outgoing webhooks with script

* Please execute below command.

~~~
docker-compose exec chat /opt/stub_server/script/setup.sh
~~~

## Incoming Webhooks

### Configuration

* `[Integrations] > [Incoming Webhooks] > [Add Incoming Webhook]`
* Please configure below key-value.
  * Title: ChatOps
  * Description: Stub Server
  * Channel: Town Square
  * Lock to this channel: yes

### Tesing

* Please execute below command.

~~~
curl -i -X POST -H 'Content-Type: application/json' -d '{"text": "Hello, this is some text\nThis is more text. :tada:"}' http://{your-mattermost-site}/hooks/xxx-generatedkey-xxx
~~~

* expect

~~~
Hello, this is some text
This is more text. :tada:
~~~

## Outgoing Webhooks

### Configuration

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

### Tesing

* Please submit below message at Town Square.

~~~
#test this is sample text.
~~~

* expect

~~~
#test this is sample text.
~~~

