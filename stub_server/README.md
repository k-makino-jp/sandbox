# Reference

* [Local Machine Setup using Docker](https://docs.mattermost.com/install/setting-up-local-machine-using-docker.html)
* [Incoming Webhookds](https://docs.mattermost.com/developer/webhooks-incoming.html)

# Environment

* OS: Ubuntu 20.04 LTS

# Installed Software

* Docker
* Docker-Compose

# Installation

* Please configure OutGoing-WebHook CallBack URL written in `script/mm_setup.sh`.

~~~
CALLBACK_URL="http://localhost:8065/hooks/xxx"
~~~

* Please execute below commads.

~~~
docker-compose up -d
~~~

* Please access below url.

~~~
http://localhost:8065
~~~

* Please create below user.
  * email: <email>
  * username: mmadmin
  * password: mm123admin

* Please create below team.
  * teamname: devops

* Please execute below shell script.

~~~
docker-compose exec chat /opt/stub_server/script/mm_setup.sh
~~~

# Uninstallation

* Please execute below commands.

~~~
docker-compose down
rm -r mattermost-data/* mysql/*
~~~