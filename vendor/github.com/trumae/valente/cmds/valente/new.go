package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	path "path/filepath"
	"strings"
)

const tplMainGo = `package main

import (
   "log"
	 "net/http"
   "runtime"
	 "sync"
	 "time"

	 "github.com/satori/go.uuid"
   "github.com/trumae/valente"
   "github.com/trumae/valente/status"
	 "{{ . }}"

	 "github.com/gorilla/websocket"
)

//App is a Web Application representation
type App struct {
    valente.App
}

const timeout = 300
const gctime = 1

var (
	sessions map[string]*App
	mutex    sync.Mutex

	upgrader = websocket.Upgrader{}
)

//addSession include a new app on sessions
func addSession(key string, app *App) {
	mutex.Lock()
	defer mutex.Unlock()

	sessions[key] = app
}

//getSession return the app by key
func getSession(key string) *App {
	return sessions[key]
}

func gcStepSession() {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now().Unix()
	for key, app := range sessions {
		if now-app.LastAccess.Unix() > timeout {
			log.Println("Collecting", key)
			delete(sessions, key)
		}
	}
}

//Initialize inits the App
func (app *App) Initialize() {
    log.Println("App Initialize")

    app.AddForm("login", forms.FormLogin{})
    app.AddForm("home", forms.FormHome{})

    app.GoTo("login", nil)
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    log.Println("Init sessions")
    sessions = make(map[string]*App)
    mutex = sync.Mutex{}
	  
	  go func() {
		  for {
		    time.Sleep(gctime * time.Second)
		    gcStepSession()
		  }
		}()

		fs := http.FileServer(http.Dir("public"))
		http.Handle("/", fs)

    http.HandleFunc("/status", status.ValenteStatusHandler)
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        ws, err := upgrader.Upgrade(w, r, nil)
	      if err != nil {
					      log.Println(err)
	              return 
	      }
	      defer ws.Close()

	      err = ws.WriteMessage(websocket.TextMessage, []byte("__GETSESSION__"))
        if err != nil {
					      log.Println(err)
	              return 
        }

        _, bid, err := ws.ReadMessage()
       if err != nil {
					      log.Println(err)
	              return 
       }
       idSession := string(bid)

       var app *App
       app = getSession(idSession)
       if app == nil {
               su1, err := uuid.NewV4()
	       if err != nil {
			 log.Println(err)
	                 return 
	       }
	       u1 := su1.String()
               log.Println("New session", u1)
               app = &App{}
               app.LastAccess = time.Now()
               addSession(u1, app)
               err = ws.WriteMessage(websocket.TextMessage, []byte(u1))
               if err != nil {
			 log.Println(err)
	                 return 
               }
               app.WebSocket(ws)
               app.Initialize()
       } else {
               app.LastAccess = time.Now()
               log.Println("Reusing session", idSession)
               err := ws.WriteMessage(websocket.TextMessage, []byte(idSession))
               if err != nil {
					         log.Println(err)
	                 return 
               }
               app.WebSocket(ws)
       }
       app.Run()
    })

		log.Println("Server running ...")
    http.ListenAndServe(":8000", nil)
}

`

func createApp(name string) {
	curpath, _ := os.Getwd()
	log.Println("current path:", curpath)
	gopath := os.Getenv("GOPATH")
	log.Println("gopath:", gopath)
	if gopath == "" {
		log.Println("[ERRO] $GOPATH not found")
		log.Println("[HINT] Set $GOPATH in your environment variables")
		os.Exit(2)
	}

	haspath := false
	appsrcpath := ""

	wgopath := path.SplitList(gopath)
	for _, wg := range wgopath {

		wg = path.Join(wg, "src")

		if strings.HasPrefix(strings.ToLower(curpath), strings.ToLower(wg)) {
			haspath = true
			appsrcpath = wg
			break
		}

		wg, _ = path.EvalSymlinks(wg)

		if strings.HasPrefix(strings.ToLower(curpath), strings.ToLower(wg)) {
			haspath = true
			appsrcpath = wg
			break
		}

	}

	if !haspath {
		log.Printf("[ERRO] Unable to create an application outside of $GOPATH%ssrc(%s%ssrc)\n", string(path.Separator), gopath, string(path.Separator))
		log.Printf("[HINT] Change your work directory by `cd ($GOPATH%ssrc)`\n", string(path.Separator))
		os.Exit(2)
	}

	apppath := path.Join(curpath, name)

	if isExist(apppath) {
		log.Printf("[ERRO] Path (%s) already exists\n", apppath)
		log.Printf("[WARN] Do you want to overwrite it? [yes|no]]")
		if !askForConfirmation() {
			os.Exit(2)
		}
	}
	pathToValente := path.Join(appsrcpath, "github.com", "trumae", "valente")
	pathToValenteData := path.Join(pathToValente, "data")
	pathToValenteDataPublic := path.Join(pathToValenteData, "public")
	pathToValenteDataForms := path.Join(pathToValenteData, "forms")

	log.Println("Creating application ...")

	os.MkdirAll(apppath, 0755)
	fmt.Println(apppath + string(path.Separator))

	copyDir(pathToValenteDataPublic, path.Join(apppath, "public"))
	copyDir(pathToValenteDataForms, path.Join(apppath, "forms"))

	packageForms := path.Join(apppath[len(appsrcpath)+1:], "forms")

	tmpl, err := template.New("forms").Parse(tplMainGo)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, packageForms)
	if err != nil {
		panic(err)
	}
	writetofile(path.Join(apppath, "main.go"), buf.String())

}
