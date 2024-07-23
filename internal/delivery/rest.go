package delivery

import (
	"encoding/json"
	"myapp/internal/entity/user"
	"myapp/internal/service/userservice"
	"net/http"
)

type RESTController struct {
	userService userservice.UserService
}

func NewRESTController(s *userservice.Service) *RESTController {
	return &RESTController{userService: s}
}

func (c *RESTController) Run(addr string) {
	http.HandleFunc("/register", c.Register)
	http.HandleFunc("/activate", c.Activate)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/check-balance", c.CheckBalance)

	http.ListenAndServe(addr, nil)
}

func (c *RESTController) Register(w http.ResponseWriter, r *http.Request) {
	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.userService.Register(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (c *RESTController) Activate(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	if err := c.userService.Activate(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "activated"})
}

func (c *RESTController) Login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	u, err := c.userService.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (c *RESTController) CheckBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	balance, err := c.userService.CheckBalance(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}
