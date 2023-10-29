package logic

import (
	"composeapp/db"
	"composeapp/internal/svc"
	"composeapp/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type MigrateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMigrateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MigrateLogic {
	return &MigrateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MigrateLogic) Migrate(req *types.MigrateReq) (resp *types.MigrateResp, err error) {

	//s := l.svcCtx.Client

	client := db.GetConnection()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println("failed creating schema resources: %v", err)
	}
	defer client.Close()
	return &types.MigrateResp{Status: "Migrate done"}, nil

}
