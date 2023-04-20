package handler

import (
	openaiproto "github.com/sashajdn/sasha/service.openai/proto"
)

type OpenAIService struct {
	*openaiproto.UnimplementedOpenaiServer
}
