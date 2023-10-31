# discord-notify

Send notifications to yourself

### Usage

Download:

```sh
VERSION=0.0.1
OS=linux
ARCH=amd64
curl -OL "https://github.com/bozdoz/discord-notify/releases/download/v${VERSION}/discord-notify-v${VERSION}-${OS}-${ARCH}.tar.gz"
```

Install:

```sh
tar -xvf discord-notify.tar.gz
```

Setup:

Setup `NOTIFY_TOKEN` and `NOTIFY_USER` somehow.  Maybe with a `.env` file, adjacent to your binary.

```sh
NOTIFY_TOKEN=ASDF.asdf.asdf
NOTIFY_USER=9879879877987
```

Use:

```sh
echo "Hello" | discord-notify
```

```sh
ls | discord-notify
```

```sh
echo "const elephant: number = 123; // but really is it?"  | discord-notify --code=ts
```

For permissions:

1. Allow `Message Content Intent` in discord developers, under Bot: https://discord.com/developers/applications/
2. Visit a similar link, to add your bot to a server: # https://discord.com/api/oauth2/authorize?client_id=[YOUR_CLIENT_ID]&permissions=2112&scope=bot

How to get your USER ID? 

- Click your avatar in Discord and press "Copy User Id"