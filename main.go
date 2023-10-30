package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// get message from stdin: `echo "hello world!" | go run .`
	reader := bufio.NewReader(os.Stdin)
	msg, _ := io.ReadAll(reader)

	// start a channel for exiting
	done := make(chan bool, 1)

	// add a handler for when discord is ready
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.Ready) {
		channel, err := s.UserChannelCreate(USER)

		if err != nil {
			fmt.Println("channel not created with User:", USER)
			panic("UserChannelCreated Failed" + err.Error())
		}

		_, err = s.ChannelMessageSend(channel.ID, string(msg))

		if err != nil {
			panic("ChannelMessageSend Failed" + err.Error())
		}

		done <- true
	})

	// actually connect to discord
	if err = discord.Open(); err != nil {
		panic("An error occurred while trying to open bot connection: " + err.Error())
	}

	// add handler for CTRL+C exiting
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

		// wait 2 seconds before we give the message to press CTRL-C
		time.Sleep(time.Second * 2)

		fmt.Println("Press CTRL-C to exit...")

		<-sc
		done <- true
	}()

	<-done
	fmt.Println("done.")
}
