package bot

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type TenguBot struct {
	log *zap.Logger

	session *discordgo.Session
}

func NewTenguBot(logger *zap.Logger, token string) *TenguBot {
	log := logger.Named("TenguBot")

	session, err := discordgo.New(token)
	if err != nil {
		logger.Fatal("Failed to create bot", zap.Error(err))
		return nil
	}

	return &TenguBot{
		log:     log,
		session: session,
	}
}

func (b *TenguBot) Start() {
	log := b.log.Named("Start")

	log.Info("Start bot")

	err := b.session.Open()
	if err != nil {
		log.Fatal("Failed to open session")
		return
	}
	defer b.session.Close()

	log.Info("Bot successfully start")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
