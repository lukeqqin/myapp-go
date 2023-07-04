package main

import (
	"context"
	"embed"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"myapp-go/internal/infrustructure/rk"
	_ "myapp-go/internal/infrustructure/rk"
	"myapp-go/internal/routes"
)

//go:embed boot.yaml
var embedFS embed.FS

func main() {
	rk.LoadMyEntry()
	boot := rkboot.NewBoot(rkboot.WithBootConfigPath("boot.yaml", &embedFS))
	// Bootstrap
	boot.Bootstrap(context.Background())
	rk.Init()
	ginEntry := rkgin.GetGinEntry("myapp")
	myapp := ginEntry.Router.Group("/myapp/api/v1")
	{
		genealogy := myapp.Group("/genealogy")
		{
			genealogy.POST("/assemble", routes.Assemble)
			genealogy.POST("/my", routes.MyGenealogy)
			genealogy.POST("/save", routes.SaveGenealogy)
			genealogy.POST("/members", routes.Members)
			genealogy.POST("/member/save", routes.SaveGenealogyMember)

		}
	}
	boot.WaitForShutdownSig(context.Background())

}
