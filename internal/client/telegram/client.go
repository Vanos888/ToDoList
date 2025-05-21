package telegram

import "context"

type TelegramClient struct {
}

func NewTelegramClient() *TelegramClient {
	return &TelegramClient{}
}

func (c *TelegramClient) SendCode(ctx context.Context, tgName string, code int) error {
	return nil
}
