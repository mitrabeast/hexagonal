package api

import (
	"encoding/json"
	"net/http"

	"idiomatic/user"
	"idiomatic/util"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) MountRoutes(r chi.Router) {
	r.Get("/user", c.GetAll)
	r.Get("/user/{id}", c.GetByID)
	r.Post("/user", c.Register)
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.List(r.Context())
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}
	util.Success(w, http.StatusOK, users.PublicInfo())
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	userID := util.ParseInt64(chi.URLParam(r, "id"))
	user, err := c.userService.Retrieve(r.Context(), userID)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}
	util.Success(w, http.StatusOK, user.PublicInfo())
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user user.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}
	if err := c.userService.Register(r.Context(), &user); err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}
	util.Success(w, http.StatusOK, user.PublicInfo())
}
