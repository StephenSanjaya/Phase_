package models

type Recipe struct {
	RecipeID    int     `json:"recipe_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CookingTime string  `json:"cooking_time"`
	Rating      float64 `json:"rating"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	UserID     int    `required:"true" json:"user_id"`
	Email      string `required:"true" json:"email" regex:"true"`
	Password   string `required:"true" json:"password" minLen:"8"`
	FullName   string `required:"true" json:"fullname" minLen:"6" maxLen:"15"`
	Age        int    `required:"true" json:"age" min:"17"`
	Occupation string `required:"true" json:"occupation"`
	Role       string `required:"true" json:"role"`
}

type SuccessMessage struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Datas   interface{} `json:"datas,omitempty"`
}

type ErrorMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
