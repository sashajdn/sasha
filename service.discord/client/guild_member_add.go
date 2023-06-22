package client

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/monzo/slog"

	discordproto "github.com/sashajdn/sasha/service.discord/proto"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

func init() {
	registerGuildMemberAddHandler("create-internal-account", createInternalAccount)
}

func createInternalAccount(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	slog.Info(ctx, "Guild member add event received: %s %s", u.User.ID, u.User.Username)

	errParams := map[string]string{
		"user_id": u.User.ID,
		"name":    u.User.Username,
	}

	// Create account.
	if _, err := (&tradeaccountproto.CreateAccountRequest{
		UserId:   u.User.ID,
		Username: u.User.Username,
		Email:    u.User.Email,
	}).Send(ctx).Response(); err != nil {
		// Best effort for now.
		slog.Error(ctx, "Failed to create a users account on guild member add event: %+v: Error: %v", errParams, err)
		return
	}

	// Read account.
	rsp, err := (&tradeaccountproto.ReadAccountRequest{
		UserId: u.User.ID,
	}).Send(ctx).Response()
	if err != nil {
		slog.Error(ctx, "Failed to read users account: %+v, Error: %s", errParams, err)
		return
	}

	// Build welcome message.
	welcomeMsg := `
:wave: Hello <@%s>, welcome to the Swallowtail Crypto Group :dove:

I have created you an account on arrival, so you can get started right away.
Please ask in the support channels for any help!

Account:
	`
	formattedWelcomeMsg := fmt.Sprintf(welcomeMsg, u.User.ID)

	account := rsp.GetAccount()
	accountMsg := `
UserID:             %s
Email:              %s
Username:           %s
Created:            %s
Is Futures Member:  %v
Primary Venue:      %s
	`
	formattedAccountMsg := fmt.Sprintf(
		accountMsg,
		account.GetUserId(),
		account.GetEmail(),
		account.GetUsername(),
		account.GetCreated().AsTime(),
		account.GetIsFuturesMember(),
		account.GetPrimaryVenue(),
	)

	// Send welcome message.
	idempotencyKey := fmt.Sprintf("guildmemberaddwelcome-%s-%s", u.User.ID, time.Now().UTC().Truncate(10*time.Minute))
	errParams["idempotency_key"] = idempotencyKey

	if _, err := (&discordproto.SendMsgToPrivateChannelRequest{
		UserId:         u.User.ID,
		SenderId:       "discordsystem:event_handler",
		IdempotencyKey: idempotencyKey,
		Content:        fmt.Sprintf("%s```%s```", formattedWelcomeMsg, formattedAccountMsg),
	}).Send(ctx).Response(); err != nil {
		slog.Error(ctx, "Failed to send welcome message on guild member add: %+v, Error: %s", errParams, err)
	}
}
