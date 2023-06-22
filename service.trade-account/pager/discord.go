package pager

import (
	"context"
	"fmt"
	"time"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/libraries/util"
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

type discordPager struct{}

func init() {
	register(tradeaccountproto.PagerType_DISCORD.String(), &discordPager{})
}

func (d *discordPager) Page(ctx context.Context, userID, msg string) error {
	hashedContent := util.Sha256Hash(msg)
	if _, err := (&discordproto.SendMsgToPrivateChannelRequest{
		UserId:   userID,
		Content:  msg,
		SenderId: "system:s.account:pager",
		// Idempotent on channel, message & day.
		IdempotencyKey: fmt.Sprintf("%s-%s-%s", userID, hashedContent[:8], time.Now().UTC().Truncate(24*time.Hour)),
	}).Send(ctx).Response(); err != nil {
		return gerrors.Augment(err, "failed_to_page_user", map[string]string{
			"user_id": userID,
		})
	}

	return nil
}
