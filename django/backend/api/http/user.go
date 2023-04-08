package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"django-hexagonal/hexagon/user"
	"django-hexagonal/util"
)

type UserController struct {
	usecase user.UserUsecase
}

func NewUserController(usecase user.UserUsecase) *UserController {
	return &UserController{
		usecase: usecase,
	}
}

func (c *UserController) MountRoutes(r chi.Router) {
	r.Get("/user/{id}", c.GetByID)
	r.Post("/user", c.Register)
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	userID := util.ParseInt64(chi.URLParam(r, "id"))
	user, err := c.usecase.Retrieve(r.Context(), userID)
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
	if err := c.usecase.Register(r.Context(), &user); err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}
	util.Success(w, http.StatusOK, user.PublicInfo())
}
