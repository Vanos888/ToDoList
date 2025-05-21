package auth

import (
	"ToDoList/internal/domain"
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"math/rand/v2"
	"time"
)

type IConfig interface {
	GetJWTSecret() string
}

type IAuthRepo interface { //Редис
	SaveTempUser(ctx context.Context, user domain.User, code int) error
	GetTempUser(ctx context.Context, tgName string) (domain.User, int, error)
}
type IUserRepo interface { //Постгресс
	CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error)
	GetUserByLogin(ctx context.Context, tgName string) (domain.User, error)
}
type ITelegramClient interface { //Телеграм код
	SendCode(ctx context.Context, tgName string, code int) error
}

type Service struct {
	authRepo       IAuthRepo
	userRepo       IUserRepo
	telegramClient ITelegramClient
	cfg            IConfig
}

func NewAuthService(
	authRepo IAuthRepo,
	userRepo IUserRepo,
	telegramClient ITelegramClient,
	cfg IConfig,
) *Service {
	return &Service{
		authRepo:       authRepo,
		userRepo:       userRepo,
		telegramClient: telegramClient,
		cfg:            cfg,
	}
}
func (s *Service) Register(ctx context.Context, user domain.User) error {
	//проверка на существование юзера, постгресс
	checkUser, err := s.userRepo.GetUserByLogin(ctx, user.TgName)
	if err != nil {
		return err
	}
	if !checkUser.IsEmpty() {
		return errors.New("user already exists")
	}
	//генерация кода
	code := randRange(100000, 999999)
	//перемешение юзера в редис

	if err := s.authRepo.SaveTempUser(ctx, user, code); err != nil {
		return err
	}

	//отправить код в тг
	if err := s.telegramClient.SendCode(ctx, user.TgName, code); err != nil {
		return err
	}

	return nil
}

func (s *Service) ConfirmCode(ctx context.Context, tgName string, code int) (uuid.UUID, error) {
	user, realCode, err := s.authRepo.GetTempUser(ctx, tgName) // получаем реальный код из редиса и юзера
	if err != nil {
		return uuid.Nil, err
	}
	if realCode != code { // сравниваем введенный юзером код с нашим кодом из редиса
		return uuid.Nil, errors.New("invalid code")
	}
	user.ID = uuid.New()
	return s.userRepo.CreateUser(ctx, user)
}

func (s *Service) Login(ctx context.Context, user domain.Login) (string, error) {

	currentUser, err := s.userRepo.GetUserByLogin(ctx, user.TgName)
	if err != nil {
		return "", err
	}

	if currentUser.PasswordHash != user.Password {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Subject:   user.TgName,
		Id:        currentUser.ID.String(),
	})

	return token.SignedString([]byte(s.cfg.GetJWTSecret()))
}

func randRange(min int, max int) int {
	return rand.IntN(max-min) + min
}
