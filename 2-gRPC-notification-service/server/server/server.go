package server

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/40471/go-grpc-examples/2-gRPC-notification-service/proto"
	"github.com/40471/go-grpc-examples/2-gRPC-notification-service/pubsub"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
	subscribers map[string]map[string]map[string]chan *pb.Notification
	mu          sync.Mutex
}

func NewServer() *Server {
	return &Server{
		subscribers: make(map[string]map[string]map[string]chan *pb.Notification),
	}
}

func (s *Server) Subscribe(req *pb.SubscriptionRequest, stream pb.NotificationService_SubscribeServer) error {
	s.mu.Lock()
	if s.subscribers[req.EventType] == nil {
		s.subscribers[req.EventType] = make(map[string]map[string]chan *pb.Notification)
	}
	if s.subscribers[req.EventType][req.Filter] == nil {
		s.subscribers[req.EventType][req.Filter] = make(map[string]chan *pb.Notification)
	}
	notifications := make(chan *pb.Notification, 10)
	s.subscribers[req.EventType][req.Filter][req.UserId] = notifications
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.subscribers[req.EventType][req.Filter], req.UserId)
		s.mu.Unlock()
	}()

	for notification := range notifications {
		if err := stream.Send(notification); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) PublishEvent(ctx context.Context, req *pb.PublishEventRequest) (*pb.PublishEventResponse, error) {
	log.Printf("Publishing event: EventType=%s, Filter=%s, Message=%s", req.EventType, req.Filter, req.Message)
	pubsub.Publish(req.EventType, req.Filter, req.Message)
	return &pb.PublishEventResponse{Status: "Event published successfully"}, nil
}

func (s *Server) StartConsumer(eventType, filter string) {
	pubsub.Consume(eventType, filter, func(msg string) {
		s.mu.Lock()
		defer s.mu.Unlock()

		notification := &pb.Notification{
			EventType: eventType,
			Filter:    filter,
			Message:   msg,
			Timestamp: time.Now().Unix(),
		}

		log.Printf("Distributing notification to subscribers: EventType=%s, Filter=%s", eventType, filter)
		for _, ch := range s.subscribers[eventType][filter] {
			ch <- notification
		}
	})
}
