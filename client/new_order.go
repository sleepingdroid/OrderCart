package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/sleepingdroid/OrderCart/generated/orderpb"
	jwt_token "github.com/sleepingdroid/OrderCart/jwt_token"

	"google.golang.org/grpc"
)

func main() {
	// ตั้งค่า gRPC connection ไปยัง server

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // ใช้ WithTransportCredentials สำหรับ production
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	// ข้อมูล Payload ที่ต้องการส่ง
	order := map[string]interface{}{
		"order_id": "123",
		"user_id":  "456",
		"status":   "pending",
	}

	// แปลง Payload เป็น JSON string
	payloadBytes, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Failed to marshal payload: %v", err)
	}
	payload := string(payloadBytes)
	token, err := jwt_token.NewJWT("456", "user")
	if err != nil {
		log.Fatalf("Failed to generate JWT token: %v", err)
	}

	// สร้าง NewOrderRequest
	req := &pb.NewOrderRequest{
		Payload: payload,
		Jwt:     token,
	}

	// ตั้ง context พร้อม timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// เรียกใช้ NewOrder ผ่าน gRPC
	res, err := client.NewOrder(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}

	// แสดงผลลัพธ์จาก server
	log.Printf("Response: Status=%s, Message=%s", res.Status, res.Message)
}
