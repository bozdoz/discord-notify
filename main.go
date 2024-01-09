package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bwmarrin/discordgo"

	_ "github.com/joho/godotenv/autoload"
)

var TOKEN = os.Getenv("NOTIFY_TOKEN")
var USER = os.Getenv("NOTIFY_USER")

func main() {
	// creates Discord session
	discord, err := discordgo.New("Bot " + TOKEN)

	if err != nil {
		fmt.Println("Discord failed to run w/Token:", TOKEN[:5], "...", TOKEN[len(TOKEN)-5:])
		panic(err.Error())
	}

	defer discord.Close()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("data needs to be piped to stdin")
		fmt.Println("\nUsage:")
		fmt.Println("\techo 'hello' | docker-notify")
		os.Exit(1)
	}

	// get message from stdin: `echo "hello world!" | go run .`
	reader := bufio.NewReader(os.Stdin)

	msg, _ := io.ReadAll(reader)

	if len(msg) < 2 {
		fmt.Println("message is empty")
		fmt.Println("\nUsage:")
		fmt.Println("\techo 'hello' | docker-notify")
		os.Exit(1)
	}

	var code_format string

	flag.StringVar(&code_format, "code", "", "If passed, which code format the message should use")

	flag.Parse()

	// start a channel for exiting
	done := make(chan bool, 1)

	var dm string
	if len(code_format) > 0 {
		dm = fmt.Sprintf("```%s\n%s```", code_format, msg)
	} else {
		dm = string(msg)
	}

	// add a handler for when discord is ready
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.Ready) {
		channel, err := s.UserChannelCreate(USER)

		var channel_id string

		if err != nil {
			// TODO append errors to an array and output if there's an issue
			fmt.Println("channel not created with User: ", USER, " ", err.Error())
			// try setting the user id as the channel
			channel_id = USER
		} else {
			channel_id = channel.ID
		}

		// TODO: this has a limit of 4K characters
		_, err = s.ChannelMessageSend(channel_id, dm)

		if err != nil {
			panic("ChannelMessageSend Failed: " + err.Error())
		}

		done <- true
	})

	// actually connect to discord
	if err = discord.Open(); err != nil {
		panic("An error occurred while trying to open bot connection: " + err.Error())
	}

	<-done
	fmt.Println("done.")
}
