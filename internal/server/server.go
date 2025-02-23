package server

import (
	"context"
	"net/http"

	"github.com/osamikoyo/IM-basket/internal/data"
	"github.com/osamikoyo/IM-basket/internal/data/models"
	"github.com/osamikoyo/IM-basket/pkg/proto/pb"
)

type Server struct {
	pb.UnimplementedBasketServiceServer
    Storage *data.Storage
}

func (s *Server) AddProduct(_ context.Context,req *pb.AddProductReq) (*pb.Response, error){
    err := s.Storage.AddProduct(req.BasketID, models.ToModelsProduct(req.Product))
    if err != nil{
        return &pb.Response{
            Status: http.StatusInternalServerError,
            Error: err.Error(),
        }, nil
    }

    return &pb.Response{
        Status: http.StatusOK,
        Error: "",
    }, nil
}

func (s *Server) DeleteProduct(_ context.Context,req *pb.DeleteProductReq) (*pb.Response, error){
    err := s.Storage.DeleteProduct(req.BasketID, req.ProductID)
    if err != nil{
        return &pb.Response{
            Status: http.StatusInternalServerError,
            Error: err.Error(),
        }, nil
    }

    return &pb.Response{
        Status: http.StatusOK,
        Error: "",
    }, nil
}
func (s *Server) Get(_ context.Context,req *pb.GetBasketReq) (*pb.GetBasketResp, error){
    basket, err := s.Storage.Get(req.BasketID)
    if err != nil{
        return &pb.GetBasketResp{
            Response: &pb.Response{
                Status: http.StatusInternalServerError,
                Error: err.Error(),
            },
            Basket: nil,
        }, nil
    }

    return &pb.GetBasketResp{
        Response: &pb.Response{
            Status: http.StatusOK,
            Error: "",
        },
        Basket: models.ToPBbasket(*basket),
    }, nil
}
func (s *Server) New(_ context.Context,req *pb.NewBasketReq) (*pb.Response, error){}