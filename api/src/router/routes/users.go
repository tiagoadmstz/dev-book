package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		AuthenticationRequire: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.FindAllUsers,
		AuthenticationRequire: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.FindUserById,
		AuthenticationRequire: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		AuthenticationRequire: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		AuthenticationRequire: false,
	},
}
