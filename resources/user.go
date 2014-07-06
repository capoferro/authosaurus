package resources

import (
	"strconv"
	"log"
	"time"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./authosaurus.db")
	if err != nil {
		log.Printf("Error connecting to the database: " + err.Error())
	}
}

type User struct {
	Id int64 `sql:not null`
	Name string `sql:not null`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserResource struct {
	Users map[string]User
}

func (u UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Doc("Manage Users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// ws.Route(ws.GET("/{user-id}").To(u.noop).
	// 	Doc("Get a User").
	// 	Operation("getUser").
	// 	Param(ws.PathParameter("user-id", "User Id").DataType("string")).
	// 	Writes(User{}))

	// ws.Route(ws.PUT("/{user-id}").To(u.noop).
	// 	Doc("Update a User").
	// 	Operation("updateUser").
	// 	Param(ws.PathParameter("user-id", "User Id").DataType("string")).
	// 	Reads(User{}))

	ws.Route(ws.POST("").To(u.createUser).
		Doc("Create a User").
		Operation("createUser").
		Reads(User{}))

	// ws.Route(ws.DELETE("/{user-id}").To(u.noop).
	// 	Doc("Delete a User").
	// 	Operation("deleteUser").
	// 	Param(ws.PathParameter("user-id", "User Id").DataType("string")))

	container.Add(ws)
}

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	user := &User{}
	err := request.ReadEntity(user)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, "Error parsing user JSON: " + err.Error())
		return
	}
	
	db.Save(user)

	log.Printf("Created user #" + strconv.Itoa(int(user.Id)) + " (" + user.Name + ")")
	
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(user)
}
