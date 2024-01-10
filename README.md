# discord-notify

Send notifications to yourself

### Usage

Download:

```sh
VERSION=0.0.3
OS=linux
ARCH=amd64
curl -OL "https://github.com/bozdoz/discord-notify/releases/download/v${VERSION}/discord-notify-v${VERSION}-${OS}-${ARCH}.tar.gz"
```

Install:

```sh
tar -xvf discord-notify.tar.gz
```

Setup:

Setup environment variables for `NOTIFY_TOKEN` and `NOTIFY_USER` somehow.  Maybe with a `.env` file, adjacent to your binary.

```sh
NOTIFY_TOKEN=ASDF.asdf.asdf
NOTIFY_USER=9879879877987
```

Use:

Testing:

```sh
echo "Hello" | discord-notify
```

Printing a list of files/folders in a directory:

```sh
ls | discord-notify
```

Syntax Highlighting:

```sh
echo "const elephant: number = 123; // but really is it?"  | discord-notify --code=ts
```

Notifying when a Docker container exits:

```sh
docker ps -a --format "{{.Names}} | {{.Status}}" --filter status=exited | discord-notify --code=sh
```

For permissions:

1. Allow `Message Content Intent` in discord developers, under Bot: https://discord.com/developers/applications/
2. Visit a similar link, to add your bot to a server: # https://discord.com/api/oauth2/authorize?client_id=[YOUR_CLIENT_ID]&permissions=2112&scope=bot

How to get your USER ID? 

- Click your avatar in Discord and press "Copy User Id"
