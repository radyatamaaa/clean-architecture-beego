package grpc

import (
	"clean-architecture-beego/internal/domain"
	_productUsecaseMock "clean-architecture-beego/internal/product/mocks"
	"context"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log"
	"net"
	"testing"
)
const bufSize = 1024 * 1024

var (
	lis *bufconn.Listener
	productUsecaseMock  *_productUsecaseMock.Usecase
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	productUsecaseMock = new(_productUsecaseMock.Usecase)

	productService := NewProductService(productUsecaseMock)
	RegisterProductServiceServer(s, productService)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestProductService_GetProducts(t *testing.T) {
	//resultMock
	mockProduct := []domain.Product{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	for i := 0; i < 1; i++ {
		t.Run("TestProductService_integration_GetProducts-success",func(t *testing.T) {
			t.Parallel()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &GetProductsParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Limit:         "1",
				Offset:        "10",
			}
			client := NewProductServiceClient(connection)
			resp, err := client.GetProducts(context.Background(), req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})

		t.Run("TestProductService_mocks_GetProducts-success",func(t *testing.T) {
			t.Parallel()

			productUsecaseMock.On("GetProducts", mock.Anything,
				mock.AnythingOfType("int"),
				mock.AnythingOfType("int")).
				Return(mockProduct, nil)

			ctx := context.Background()
			var connection *grpc.ClientConn
			connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &GetProductsParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Limit:         "1",
				Offset:        "10",
			}
			client := NewProductServiceClient(connection)
			resp, err := client.GetProducts(ctx, req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}

}
