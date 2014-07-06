package resources

import (
	"github.com/emicklei/go-restful"
)

type User struct {
	Id, Name string
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

	ws.Route(ws.GET("/{user-id}").To(u.noop).
		Doc("Get a User").
		Operation("getUser").
		Param(ws.PathParameter("user-id", "User Id").DataType("string")).
		Writes(User{}))

	ws.Route(ws.PUT("/{user-id}").To(u.noop).
		Doc("Update a User").
		Operation("updateUser").
		Param(ws.PathParameter("user-id", "User Id").DataType("string")).
		Reads(User{}))

	ws.Route(ws.POST("").To(u.noop).
		Doc("Create a User").
		Operation("createUser").
		Reads(User{}))

	ws.Route(ws.DELETE("/{user-id}").To(u.noop).
		Doc("Delete a User").
		Operation("deleteUser").
		Param(ws.PathParameter("user-id", "User Id").DataType("string")))

	container.Add(ws)
}

func (u UserResource) noop(request *restful.Request, response *restful.Response) {
}
