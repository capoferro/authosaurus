package resources

import (
	"strconv"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id string
	Name string
}

type UserResource struct {
	// TODO persistence
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
	id := strconv.Itoa(len(u.Users) + 1)
	log.Printf("Creating user #" + id + " (" + user.Name + ")")
	user.Id = id
	// TODO persistence
	u.Users[user.Id] = *user
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(user)
}
