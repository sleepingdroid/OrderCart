package jwt_token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("your-secret-key")

func generateJWT(userID string, role string) (string, error) {
	// กำหนด Claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // หมดอายุใน 24 ชั่วโมง
	}

	// สร้าง Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ลงนาม Token ด้วย Secret Key
	return token.SignedString(secretKey)
}

func NewJWT(userID string, role string) (string, error) {
	token, err := generateJWT("123456", "admin")
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return "", err
	}
	return token, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบว่าใช้ Signing Method เดียวกัน
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}

func CheckJWT(tokenString string) bool {
	jwtToken, err := validateJWT(tokenString)
	if err != nil {
		return false
	}
	return jwtToken.Valid
}

// func parseJWT(tokenString string) {
// 	// ตรวจสอบ Token
// 	token, err := validateJWT(tokenString)
// 	if err != nil {
// 		fmt.Println("Invalid Token:", err)
// 		return
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		fmt.Println("Token is valid!")
// 		fmt.Println("User ID:", claims["user_id"])
// 		fmt.Println("Role:", claims["role"])
// 	} else {
// 		fmt.Println("Invalid Token Claims")
// 	}
// }
