package server

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/devstackq/go-clean/auth"
	authHttp "github.com/devstackq/go-clean/auth/delivery/http"
	"github.com/devstackq/go-clean/auth/repository/psql"
	"github.com/devstackq/go-clean/auth/usecase"

	"github.com/devstackq/go-clean/db"
	"github.com/devstackq/go-clean/transport"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

//refactor all structure; use gin, run in docker




type App struct {
	// grpc        grpc.Server
	authUseCase auth.UseCase
}

// interface {Signup, Signin}; stuct Grpc - own realize; struct http - own realize, grpcServer
//singletone - prepare app, connect layers with interface; init app

func NewApp() *App {
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
	// 	return nil
	// }
	// repo := mongoRepo.NewUserRepository(db.(*mongo.Database), viper.GetString("mongo.user_collection"))

	sqlObject := db.NewDbObject("postgresql", viper.GetString("postgres.username"), viper.GetString("postgres.password"), viper.GetString("postgres.host"), viper.GetString("postgres.port"), viper.GetString("postgres.tableName"), viper.GetString("postgres.dbName"))
	db, err := sqlObject.InitDb()
	if err != nil {
		log.Println(err)
		return nil
	}
	repo := psql.NewUserRepository(db.(*sql.DB))

	log.Println("init db")

	// log.Print(repoMongo, "init repo")
	return &App{
		authUseCase: usecase.NewAuthUseCase(repo, []byte(viper.GetString("auth.hash_salt")), []byte(viper.GetString("auth.secret_key")), viper.GetDuration("auth.token_ttl")),
		// httpServer:  server.(http.Server),
	}
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
  
func (app *App) Run(port string) error {
	//grpc || http create server

	factory := transport.GetFactory("http")
	transportProtocol := factory.GetTransport()
	server := transportProtocol.InitTransport(viper.GetString("port")).(http.Server)
	log.Println("init transport")

	var frontend fs.FS = os.DirFS("../../client/public")
	httpFS := http.FS(frontend)
	fileServer := http.FileServer(httpFS)
	serveIndex := serveFileContents("index.html", httpFS)
	http.Handle("/", intercept404(fileServer, serveIndex))

	authHttp.InitRoutes(app.authUseCase)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Print("run server port: ", viper.GetString("port"))

	//refactor logger go func()
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer file.Close()
	// log.SetOutput(file)
	// log.Print("logger start")

	//gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}

//func NewServer(){}
