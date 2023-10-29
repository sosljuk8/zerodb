package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"composeapp/internal/svc"
	"composeapp/internal/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	//return &types.Response{Message: "Hello from zerodb " + req.Name}, nil

	type Category struct {
		Id          int
		ParentId    int
		Name        string
		SystemName  string
		Description string
	}

	db, err := sql.Open("mysql", "root:nopass@/test_db")
	defer db.Close()

	rows, err := db.Query("SELECT * FROM test_db.settings_categories")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	categories := []Category{}

	for rows.Next() {
		c := Category{}
		err := rows.Scan(&c.Id, &c.ParentId, &c.Name, &c.SystemName, &c.Description)
		if err != nil {
			fmt.Println(err)
			continue
		}
		categories = append(categories, c)
	}

	e, err := json.Marshal(categories)
	if err != nil {
		fmt.Println(err)
		return
	}

	return &types.Response{Message: string(e)}, nil

	//database = db

}
