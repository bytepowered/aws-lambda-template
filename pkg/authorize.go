package pkg

import (
    "errors"
    "fmt"
    "github.com/golang-jwt/jwt/v5"
    "log"
    "os"
    "time"
)

func AuthorizeSecretKey() []byte {
    return []byte(os.Getenv("AUTH_JWT_SECRET"))
}

func AuthorizeSignedToken(sub string, exp int64, roles []string) (string, error) {
    token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
        Subject:   sub,
        ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
        Audience:  roles,
    }).SignedString(AuthorizeSecretKey())
    if err != nil {
        return "", fmt.Errorf("generte jwt token %w", err)
    }
    return token, nil
}

func AuthorizeParseToken(tokenStr string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return AuthorizeSecretKey(), nil
    })
    if token.Valid {
        return token, nil
    }
    if errors.Is(err, jwt.ErrTokenMalformed) {
        log.Printf("JWT invalid token: %s", tokenStr)
    } else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
        log.Printf("JWT invalid sign: %s", tokenStr)
    } else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
        log.Printf("JWT invalid expire: %s", tokenStr)
    } else {
        log.Printf("JWT unknown error, token: %s, error: %s", tokenStr, err)
    }
    return nil, err
}
