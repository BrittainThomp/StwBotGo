package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotID string

func main() {

	token := "YOUR TOKEN GOES HERE"
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = user.ID
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

	dg.AddHandler(fkstw)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	<-make(chan struct{})
	return
}

//message responses
func fkstw(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "FK STW" {
		s.ChannelMessageSend(m.ChannelID, "Your feedback is appreciated. Thank you "+m.Author.Username+"!")
	}

	if m.Content == "fk stw" {
		s.ChannelMessageSend(m.ChannelID, "I would prefer you be more passionate please, "+m.Author.Username+"!")
	}

	if m.Author.ID == "USER ID GOES HERE" {
		s.ChannelMessageSend(m.ChannelID, "You're doing great, bud!")
	}
}

//have not tested this implementation yet
func greetings(s *discordgo.Session, a *discordgo.GuildMemberAdd) {
	s.ChannelMessageSend("CHANNEL ID GOES HERE", "Weclome! Feel free to to kick things off with a 'fk stw' in chat!")
}
