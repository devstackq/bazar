package server

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/devstackq/bazar/internal/auth"
	"github.com/devstackq/bazar/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/devstackq/bazar/db"

	v1 "github.com/devstackq/bazar/internal/gallery/delivery/http"

	_ "github.com/lib/pq"
)

//refactor all structure; use gin, run in docker

type App struct {
	// grpc        grpc.Server
	authUseCase auth.UseCase
	cfg    *config.Config
	db     *sql.DB
	router *gin.Engine
	Logger *logrus.Logger
}

// interface {Signup, Signin}; stuct Grpc - own realize; struct http - own realize, grpcServer
//singletone - prepare app, connect layers with interface; init app

//cfg *config.Config
func NewApp(cfg *config.Config) (*App, error) {
	if cfg == nil {
		return nil, fmt.Errorf("errro app struct")
	}
	return &App{cfg: cfg}, nil
}

func (a *App) Initialize() {

	gin.SetMode(a.cfg.App.Mode)

	a.router = gin.New()
	a.Logger = logrus.New()

	a.router.Use(gin.Recovery())
	a.router.Use(cors.New(cors.Config{
		AllowOrigins:     a.cfg.App.Cors.AllowOrigins,
		MaxAge:           30 * time.Second,
		AllowMethods:     a.cfg.App.Cors.AllowMethods,
		AllowHeaders:     a.cfg.App.Cors.AllowHeaders,
		ExposeHeaders:    a.cfg.App.Cors.ExposeHeaders,
		AllowCredentials: a.cfg.App.Cors.AllowCredentials,
		AllowWildcard:    true,
	}))
	//psql case
	// storage2 := db.NewPostgresStorage("postgres", "password", "localhost:", "5432", "testdb") os.LookUp
	// dbSql, err := storage2.InitPostgresDb()
	// repoSql := psql.NewUserRepository(dbSql)

	//mongo case
	// storage := db.NewMongoStorage("mongo", "", "mongo", "27017", "testdb") // os.LookUp
	// dbMongo, err := storage.InitMongoDb()

	//2 variant db, method fabric
	// mongoObject := db.NewDbObject("mongodb", viper.GetString("mongo.username"), viper.GetString("mongo.password"), viper.GetString("mongo.host"), viper.GetString("mongo.port"), viper.GetString("mongo.dbName"), viper.GetString("mongo.user_collection"))
	// // db.SetConfig("", "", viper.GetString("mongo.uri"), "27017", "users")
	// db, err := mongoObject.InitDb()
	// if err != nil {
	// 	log.Println(err)
	a.setComponents()
	// 	return nil
	// }
	// repo := mongoRepo.NewUserRepository(db.(*mongo.Database), viper.GetString("mongo.user_collection"))
	
	sqlObject := db.NewDbObject("postgresql", a.cfg.DB.Username, a.cfg.DB.Password, a.cfg.DB.Host, a.cfg.DB.Port,  a.cfg.DB.DBName )
	db, err := sqlObject.InitDb()
	if err != nil {
		log.Println(err)
		return
	}
	a.db = db.(*sql.DB)
	
	log.Println("init db")

	// repo := psql.NewUserRepository(db.(*sql.DB))

	// log.Print(repoMongo, "init repo")
	// return &App{
	// 	authUseCase: usecase.NewAuthUseCase(repo, []byte(viper.GetString("auth.hash_salt")), []byte(viper.GetString("auth.secret_key")), viper.GetDuration("auth.token_ttl")),
	// 	// httpServer:  server.(http.Server),
	// }
	a.setComponents()

}

  func (a *App) Run(ctx context.Context) {
	//func() - connect client
	var frontend fs.FS = os.DirFS("../../client/public")
	httpFS := http.FS(frontend)
	fileServer := http.FileServer(httpFS)
	serveIndex := serveFileContents("index.html", httpFS)
	http.Handle("/", intercept404(fileServer, serveIndex))

	srv := http.Server{
		Addr:           a.cfg.App.Port,
		Handler:        a.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    a.cfg.App.ReadTimeout,
		WriteTimeout:   a.cfg.App.WriteTimeout,
	}
	go func() {
		a.Logger.Info("starting web server on port: ", a.cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Logger.Fatal(err.Error())
		}
	}()

	<-ctx.Done()

	a.Logger.Info("shutting down web server...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second) // todo: change time context;
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.Logger.Fatal("application forced to shutdown: ", err.Error())
	}
	a.Logger.Info("application exiting")
}

func (a *App) setComponents() {
	apiVersion := a.router.Group("/v1")
	v1.SetGalleryEndpoints(a.cfg, a.db, a.Logger, apiVersion)
	// v1.SetAuthEndpoints(a.cfg, a.db, a.Logger, apiVersion)
}


func serveFileContents(file string, files http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  // Restrict only to instances where the browser is looking for an HTML file
	  if !strings.Contains(r.Header.Get("Accept"), "text/html") {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 not found")
  
		return
	  }
	  // Open the file and return its contents using http.ServeContent
	  index, err := files.Open(file)
	  if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not found", file)
  
		return
	  }
  
	  fi, err := index.Stat()
	  if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not found", file)
  
		return
	  }
  
	  w.Header().Set("Content-Type", "text/html; charset=utf-8")
	  http.ServeContent(w, r, fi.Name(), fi.ModTime(), index)
	}
  }

  type hookedResponseWriter struct {
	http.ResponseWriter
	got404 bool
  }
  
  func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == http.StatusNotFound {
	  // Don't actually write the 404 header, just set a flag.
	  hrw.got404 = true
	} else {
	  hrw.ResponseWriter.WriteHeader(status)
	}
  }
  
  func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.got404 {
	  // No-op, but pretend that we wrote len(p) bytes to the writer.
	  return len(p), nil
	}
	return hrw.ResponseWriter.Write(p)
  }
  
  func intercept404(handler, on404 http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  hookedWriter := &hookedResponseWriter{ResponseWriter: w}
	  handler.ServeHTTP(hookedWriter, r)
  
	  if hookedWriter.got404 {
		on404.ServeHTTP(w, r)
	  }
	})
  }