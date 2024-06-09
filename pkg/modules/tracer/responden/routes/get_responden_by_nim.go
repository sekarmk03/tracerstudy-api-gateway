package routes

import (
	"context"
	"net/http"
	"tracerstudy-api-gateway/pkg/pb"
	"tracerstudy-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func GetRespondenByNim(ctx *gin.Context, cr pb.RespondenServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))
	
	nim := ctx.Param("nim")

	res, err := cr.GetRespondenByNim(grpcCtx, &pb.GetRespondenByNimRequest{
		Nim: nim,
	})

	if err != nil {
		errResp := utils.GetGrpcError(err)
		if errResp.Code == http.StatusNotFound {
			// create responden
			_, err := cr.CreateResponden(grpcCtx, &pb.CreateRespondenRequest{
				Nim: nim,
			})

			if err != nil {
				errResp := utils.GetGrpcError(err)
				ctx.AbortWithStatusJSON(
					errResp.Code,
					errResp,
				)
				return
			}

			resUpdate, err := cr.UpdateRespondenFromSiak(grpcCtx, &pb.UpdateRespondenFromSiakRequest{
				Nim: nim,
			})

			if err != nil {
				errResp := utils.GetGrpcError(err)
				ctx.AbortWithStatusJSON(
					errResp.Code,
					errResp,
				)
				return
			}

			ctx.JSON(http.StatusOK, &resUpdate)
		} else {
			ctx.AbortWithStatusJSON(
				errResp.Code,
				errResp,
			)
			return
		}
	}

	if res.Data.StatusUpdate == 0 {
		resUpdate, err := cr.UpdateRespondenFromSiak(grpcCtx, &pb.UpdateRespondenFromSiakRequest{
			Nim: nim,
		})

		if err != nil {
			errResp := utils.GetGrpcError(err)
			ctx.AbortWithStatusJSON(
				errResp.Code,
				errResp,
			)
			return
		}

		ctx.JSON(http.StatusOK, &resUpdate)
	}

	ctx.JSON(http.StatusOK, &res)
}
