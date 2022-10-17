package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "sum",
			Description: "sum two numbers",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "A",
					Description: "first number",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "B",
					Description: "second number",
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"sum": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var A int64
			var B int64
			for _, option := range i.ApplicationCommandData().Options {
				switch option.Name {
				case "A":
					A = option.IntValue()
				case "B":
					B = option.IntValue()
				}
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("%d", A+B),
				},
			})
		},
	}
)

func main() {

}

func discord() error {
	token := os.Getenv("TOKEN")
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

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}
