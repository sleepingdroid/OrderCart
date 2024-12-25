package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"

	pb "github.com/sleepingdroid/OrderCart/generated/orderpb"
	jwt_token "github.com/sleepingdroid/OrderCart/jwt_token"

	"google.golang.org/grpc"
	_ "modernc.org/sqlite" // ใช้ SQLite driver ที่ไม่ใช้ cgo
)

// สร้าง struct สำหรับ implement gRPC service
type server struct {
	pb.UnimplementedOrderServiceServer // Embed สำหรับ method ที่ไม่ต้อง implement
	db                                 *sql.DB
}

// Implement method สำหรับ GetOrder ตามที่กำหนดใน .proto
func (s *server) GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	if !jwt_token.CheckJWT(request.GetJwt()) {
		return nil, fmt.Errorf("JWT Invalid")
	}
	// fetch all order
	return &pb.GetOrderResponse{}, nil
}

func (s *server) NewOrder(ctx context.Context, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error) {
	if !jwt_token.CheckJWT(req.GetJwt()) {
		return &pb.NewOrderResponse{
			Status:  "error",
			Message: "Invalid JWT Token in payload",
		}, nil
	}

	var order *pb.Order
	err := json.Unmarshal([]byte(req.Payload), &order)
	if err != nil {
		return &pb.NewOrderResponse{
			Status:  "error",
			Message: "Invalid JSON in payload",
		}, nil
	}

	// เพิ่มคำสั่งใหม่ในฐานข้อมูล (ดูขั้นตอนถัดไป)
	fmt.Printf("Creating new order: %+v\n", order)

	// เชื่อมต่อกับฐานข้อมูล
	db, err := sql.Open("sqlite", "db/cart.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// บันทึกคำสั่งใหม่ในฐานข้อมูล
	results, err := insertOrder(db, order)
	if err != nil {
		log.Println("err:", err)
		return &pb.NewOrderResponse{
			Status:  "error",
			Message: "Failed to create order",
		}, nil
	}
	log.Println("results:", results)

	return &pb.NewOrderResponse{
		Status:  "success",
		Message: fmt.Sprintf("Order %s created successfully", order.Id),
	}, nil
}

// ฟังก์ชันสำหรับเพิ่มคำสั่งใหม่ในฐานข้อมูล
func insertOrder(db *sql.DB, order *pb.Order) (sql.Result, error) {
	query := `
        INSERT INTO orders (id, user_id, total, status, created_at, updated_at)
        VALUES (?, ?, ?, ?, datetime('now'), datetime('now'));
    `
	results, err := db.Exec(order.Id, query, order.UserId, 0, order.Status)
	return results, err
}

func main() {
	// สร้าง listener สำหรับการเชื่อมต่อ gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// สร้าง gRPC server
	s := grpc.NewServer()

	// ลงทะเบียน service ที่สร้างขึ้น
	pb.RegisterOrderServiceServer(s, &server{})

	fmt.Println("Server started at :50051")
	// เริ่มต้น server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
