package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	demomod "moul.io/adapterkit-module-demo"

	"github.com/bwmarrin/discordgo"
)

var (
	token string

	client *demomod.DemomodSvcClient

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "sum",
			Description: "sum two numbers",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "a",
					Description: "first number",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "b",
					Description: "second number",
					Required:    true,
				},
			},
		},
		{
			Name:        "say-hello",
			Description: "say hello to someone",
		},
		{
			Name:        "echo-stream",
			Description: "setup echo-stream",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"sum": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var A int64
			var B int64
			for _, option := range i.ApplicationCommandData().Options {
				switch option.Name {
				case "a":
					A = option.IntValue()
				case "b":
					B = option.IntValue()
				}
			}

			res, err := SumAction(*client, A, B)
			if err != nil {
				panic(err)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("%d", res),
				},
			})
		},
		"say-hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			res, err := SayHelloAction(*client)
			if err != nil {
				panic(err)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: res,
				},
			})
		},
		"echo-stream": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if echoStream.isOpen == true {
				echoStream.isOpen = false
				echoStream.sendChan <- ""
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "echo stream closed",
					},
				})
				return
			}
			err := EchoStreamAction(*client, &echoStream.sendChan, &echoStream.recvChan)
			if err != nil {
				panic(err)
			}

			echoStream.isOpen = true
			echoStream.id = i.ChannelID
			go func() {
				for {
					select {
					case msg := <-echoStream.recvChan:
						s.ChannelMessageSend(echoStream.id, msg)
					}
				}
			}()
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "echo stream opened",
				},
			})
			if err != nil {
				panic(err)
				return
			}
		},
	}

	echoStream = struct {
		isOpen   bool
		id       string
		sendChan chan string
		recvChan chan string
	}{
		isOpen:   false,
		id:       "",
		sendChan: make(chan string),
		recvChan: make(chan string),
	}
)

func init() {
	c, err := getClient()
	if err != nil {
		panic(err)
	}
	client = &c

	flag.StringVar(&token, "token", "", "discord bot token")
	flag.Parse()
}

func main() {
	if err := discord(); err != nil {
		panic(err)
	}

	return
}

func discord() error {
	if token == "" {
		return fmt.Errorf("missing TOKEN env var")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	err = dg.Open()
	if err != nil {
		return err
	}
	defer dg.Close()

	fmt.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		fmt.Println("Created command: ", cmd.Name)
		registeredCommands[i] = cmd
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if echoStream.isOpen == true {
			echoStream.sendChan <- m.Content
		}
	})

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}
