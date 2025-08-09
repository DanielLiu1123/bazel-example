package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NotificationService implements the notification gRPC service
type NotificationService struct {
	// UnimplementedNotificationServiceServer  // TODO: Add when proto is generated
}

// In-memory storage for demo purposes
var notifications = make(map[string]*Notification)
var notificationCounter = 0

// Notification represents a notification entity
type Notification struct {
	ID       string            `json:"id"`
	UserID   string            `json:"user_id"`
	Type     string            `json:"type"`
	Priority string            `json:"priority"`
	Title    string            `json:"title"`
	Content  string            `json:"content"`
	Metadata map[string]string `json:"metadata"`
	Sent     bool              `json:"sent"`
	SentAt   *time.Time        `json:"sent_at,omitempty"`
	CreateAt time.Time         `json:"created_at"`
}

func main() {
	// Create gRPC server
	s := grpc.NewServer()

	// Register notification service
	notificationService := &NotificationService{}
	// notification.RegisterNotificationServiceServer(s, notificationService)  // TODO: Add when proto is generated

	// Enable reflection for debugging
	reflection.Register(s)

	// Listen on port 9003
	lis, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start server in a goroutine
	go func() {
		log.Println("Notification service starting on port 9003...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Start background notification processor
	go startNotificationProcessor()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down notification service...")

	// Graceful shutdown
	s.GracefulStop()
	log.Println("Notification service stopped")
}

// SendNotification sends a notification
func (s *NotificationService) SendNotification(ctx context.Context, req interface{}) (interface{}, error) {
	log.Println("Sending notification")

	// TODO: Implement actual notification sending logic
	// For now, just create a mock notification
	notificationCounter++
	notification := &Notification{
		ID:       generateID(),
		UserID:   "user-123", // Extract from request
		Type:     "email",    // Extract from request
		Priority: "normal",   // Extract from request
		Title:    "Test Notification",
		Content:  "This is a test notification",
		Metadata: make(map[string]string),
		Sent:     false,
		CreateAt: time.Now(),
	}

	notifications[notification.ID] = notification

	// Queue for processing
	go processNotification(notification)

	log.Printf("Notification %s queued for processing", notification.ID)
	return notification, nil
}

// GetNotification retrieves a notification by ID
func (s *NotificationService) GetNotification(ctx context.Context, req interface{}) (interface{}, error) {
	// TODO: Extract ID from request
	id := "notification-123"
	log.Printf("Getting notification: %s", id)

	notification, exists := notifications[id]
	if !exists {
		return nil, nil // TODO: Return proper error
	}

	return notification, nil
}

// ListUserNotifications lists notifications for a user
func (s *NotificationService) ListUserNotifications(ctx context.Context, req interface{}) (interface{}, error) {
	// TODO: Extract user ID from request
	userID := "user-123"
	log.Printf("Listing notifications for user: %s", userID)

	var userNotifications []*Notification
	for _, notification := range notifications {
		if notification.UserID == userID {
			userNotifications = append(userNotifications, notification)
		}
	}

	return map[string]interface{}{
		"notifications": userNotifications,
		"total":         len(userNotifications),
	}, nil
}

// MarkAsRead marks a notification as read
func (s *NotificationService) MarkAsRead(ctx context.Context, req interface{}) (interface{}, error) {
	// TODO: Extract notification ID from request
	notificationID := "notification-123"
	log.Printf("Marking notification as read: %s", notificationID)

	notification, exists := notifications[notificationID]
	if !exists {
		return map[string]interface{}{"success": false}, nil
	}

	// TODO: Implement mark as read logic
	log.Printf("Notification %s marked as read", notificationID)

	return map[string]interface{}{
		"success":      true,
		"notification": notification,
	}, nil
}

// GetNotificationStats returns notification statistics
func (s *NotificationService) GetNotificationStats(ctx context.Context, req interface{}) (interface{}, error) {
	// TODO: Extract user ID from request
	userID := "user-123"
	log.Printf("Getting notification stats for user: %s", userID)

	totalCount := 0
	unreadCount := 0
	emailCount := 0
	smsCount := 0
	pushCount := 0

	for _, notification := range notifications {
		if notification.UserID == userID {
			totalCount++
			if !notification.Sent {
				unreadCount++
			}
			switch notification.Type {
			case "email":
				emailCount++
			case "sms":
				smsCount++
			case "push":
				pushCount++
			}
		}
	}

	return map[string]interface{}{
		"total_count":  totalCount,
		"unread_count": unreadCount,
		"email_count":  emailCount,
		"sms_count":    smsCount,
		"push_count":   pushCount,
	}, nil
}

// Background notification processor
func startNotificationProcessor() {
	log.Println("Starting notification processor...")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			processQueuedNotifications()
		}
	}
}

func processQueuedNotifications() {
	for _, notification := range notifications {
		if !notification.Sent {
			processNotification(notification)
		}
	}
}

func processNotification(notification *Notification) {
	log.Printf("Processing notification %s of type %s", notification.ID, notification.Type)

	// Simulate processing time
	time.Sleep(100 * time.Millisecond)

	// Mark as sent
	notification.Sent = true
	now := time.Now()
	notification.SentAt = &now

	log.Printf("Notification %s sent successfully", notification.ID)
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + string(rune(notificationCounter))
}
