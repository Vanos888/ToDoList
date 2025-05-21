package interceptor

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

const signKey = "dsdasadasdasdasdsad23232324sdsdaasd#@$%65dsdasfgrgAASWFSDGaSFFHGEFSDSWAWSGTER435FDG"

const authHeader = "authorization"

func UnaryJWTServerInterceptor() grpc.UnaryServerInterceptor {
	whiteList := map[string]struct{}{
		"/api.AuthApi/Login":       {},
		"/api.AuthApi/Register":    {},
		"/api.AuthApi/ConfirmCode": {},
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println(info.FullMethod)
		if _, ok := whiteList[info.FullMethod]; ok {
			return handler(ctx, req)
		}

		tokenFromReq, err := AuthFromMD(ctx)
		if err != nil {
			return nil, err
		}

		fmt.Println(tokenFromReq)

		token, err := jwt.ParseWithClaims(
			tokenFromReq,
			&jwt.StandardClaims{},
			func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("unexpected token signing method")
				}

				return []byte(signKey), nil
			},
		)

		if err != nil {
			return nil, fmt.Errorf("invalid token: %w", err)
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			return nil, fmt.Errorf("invalid token claims")
		}

		ctxWithUserID := context.WithValue(ctx, "user_id", claims.Id)

		return handler(ctxWithUserID, req)
	}
}

func AuthFromMD(ctx context.Context) (string, error) {
	tokenStr := metadata.ValueFromIncomingContext(ctx, authHeader)
	if len(tokenStr) == 0 {
		return "", fmt.Errorf("invalid token")
	}

	_, token, _ := strings.Cut(tokenStr[0], " ")

	return token, nil
}
