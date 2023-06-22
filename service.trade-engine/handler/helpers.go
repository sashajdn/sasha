package handler

import (
	"strings"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-engine/domain"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

func isActorValid(actorID string) bool {
	switch actorID {
	case tradeengineproto.TradeEngineActorSatoshiSystem,
		tradeengineproto.TradeEngineActorManual,
		tradeengineproto.TradeEngineActorSatoshiCommand:
		return true
	default:
		return false
	}
}

func validateTradeStrategyParticipant(tradeStrategyParticipant *tradeengineproto.ExecuteTradeStrategyForParticipantRequest, tradeStrategy *domain.TradeStrategy) error {
	switch tradeStrategyParticipant.Venue {
	case tradeengineproto.VENUE_UNREQUIRED:
		// We don't need to validate anything at this point in time.
	default:
		if !isParticipantVenueAnAvailableVenue(tradeStrategyParticipant.Venue, tradeStrategy.TradeableVenues) {
			return gerrors.FailedPrecondition("invalid_trade_strategy_participant.venue_not_available", map[string]string{
				"participant_venue": tradeStrategyParticipant.Venue.String(),
				"available_venues":  strings.Join(tradeStrategy.TradeableVenues, ","),
			})
		}
	}

	switch {
	case tradeStrategyParticipant.Risk > 50:
		return gerrors.FailedPrecondition("invalid_trade_strategy_participant.risk_too_high", nil)
	case tradeStrategyParticipant.Size < 0 && tradeStrategyParticipant.Risk < 0:
		return gerrors.BadParam("bad_param.risk_or_size_cannot_be_less_than_zero", nil)
	case tradeStrategyParticipant.Size == 0 && tradeStrategyParticipant.Risk == 0:
		return gerrors.BadParam("bad_params.risk_and_size_cannot_be_zero", nil)
	}

	return nil
}

func isParticipantVenueAnAvailableVenue(participantVenue tradeengineproto.VENUE, availableVenues []string) bool {
	for _, av := range availableVenues {
		if strings.ToUpper(participantVenue.String()) == strings.ToUpper(av) {
			return true
		}
	}

	return false
}
