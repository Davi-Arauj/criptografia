package main

import (
	"criptografia/config"
	"criptografia/logger"
	"criptografia/config/database"
	"criptografia/interface/cadastros"	
	"log"
	"net/http"
	"net/http/pprof"
	"criptografia/middleware"
	"github.com/fvbock/endless"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"criptografia/validations"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	// db.Connection()
	// defer db.DB.Close()
	// Davi.ID = "1"
	// Davi.Document = "documento"
	// Davi.CreditCard = "cartão de debito"

	// rng := rand.Reader
	// cliente.GenerateKeypair()
	// cliente.Encrypt(Davi, rng)
	// cliente.Decrypt(rng)
	// cliente.Sign(rng)
	// cliente.Verify()

	// servidor := gin.New()

	// apiRoutes := servidor.Group("/api")
	// {
	// 	apiRoutes.POST("/cliente", func(c *gin.Context) {
	// 		err := clienteController.Save(ctx)
	// 		if err != nil {
	// 			ctx.json(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		}
	// 	})
	// }
	// servidor.Run(":8080")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		err  error
		logg *zap.Logger
	)

	config.CarregarConfiguracao()

	if logg, err = logger.SetupLogger(); err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logg.Sync() }()
	zap.ReplaceGlobals(logg)

	if err = database.AbrirConexao(); err != nil {
		zap.L().Fatal("Não foi possível conectar-se ao banco de dados", zap.Error(err))
	}
	defer database.FecharConexoes()

	validations.ConfigurarValidadores()
	// if err = middleware.InicializarVersionamento(); err != nil {
	// 	zap.L().Error("Não foi possível definir a versão do sistema", zap.Error(err))
	// }

	group := errgroup.Group{}
	group.Go(func() error {
		return endless.ListenAndServe(config.ObterConfiguracao().EnderecoInterno, internalRouter(logg))
	})
	group.Go(func() error {
		return endless.ListenAndServe(config.ObterConfiguracao().EnderecoExterno, externalRouter(logg))
	})

	// health.Init(middleware.Versao)

	if err = group.Wait(); err != nil {
		zap.L().Error("Erro ao inicializar aplicação", zap.Error(err))
	}
}

func externalRouter(logg *zap.Logger) http.Handler {
	r := gin.New()
	r.Use(
		middleware.IdentificadorRequisicao(),
		middleware.VersaoInfo(),
		middleware.GinZap(logg),
		ginzap.RecoveryWithZap(logg, true),
		// middleware.Autenticacao(),
	)
	v2 := r.Group("v2")
	internal := v2.Group("internal")
	// internal.Use(middleware.ModoAdministrador())
	pprofRouter(internal)
	cadastros.Router(v2.Group("cadastros"))

	//middleware.AutoPerm(r.Routes())

	return r
}

func internalRouter(logg *zap.Logger) http.Handler {
	r := gin.New()
	r.Use(
		middleware.IdentificadorRequisicao(),
		middleware.VersaoInfo(),
		middleware.GinZap(logg),
		ginzap.RecoveryWithZap(logg, true),
		middleware.Auditoria(),
	)

	api := r.Group("api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}

func pprofRouter(r *gin.RouterGroup) {
	prefixRouter := r.Group("debug/pprof")
	prefixRouter.GET("/", pprofHandler(pprof.Index))
	prefixRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
	prefixRouter.GET("/profile", pprofHandler(pprof.Profile))
	prefixRouter.POST("/symbol", pprofHandler(pprof.Symbol))
	prefixRouter.GET("/symbol", pprofHandler(pprof.Symbol))
	prefixRouter.GET("/trace", pprofHandler(pprof.Trace))
	prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
	prefixRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
	prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
	prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
	prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
	prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
