package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"stockexchange/api/internal/types"
)

func HandleGrpcErrorToHttp(err error) *types.DetailMeta {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return &types.DetailMeta{
					Msg:    e.Message(),
					Status: http.StatusNotFound,
				}
			case codes.Internal:
				return &types.DetailMeta{
					Msg:    "内部错误",
					Status: http.StatusInternalServerError,
				}
			case codes.InvalidArgument:
				return &types.DetailMeta{
					Msg:    "参数错误",
					Status: http.StatusBadRequest,
				}
			case codes.Unavailable:
				return &types.DetailMeta{
					Msg:    "服务不可用",
					Status: http.StatusInternalServerError,
				}
			default:
				return &types.DetailMeta{
					Msg:    string(e.Code()),
					Status: http.StatusInternalServerError,
				}
			}
		}
	}
	return nil
}
