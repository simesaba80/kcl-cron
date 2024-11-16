package db

import (
	"time"

	"github.com/uptrace/bun"
)

// User is a struct that represents a user.
type User struct {
	bun.BaseModel `bun:"table:users,select:users"`
	ID            int       `bun:"id,pk,autoincrement"`
	UID           string    `bun:"uid,unique,notnull"`
	FitbitUserID  string    `bun:"fitbit_user_id"`
	Name          string    `bun:"name,notnull"`
	Sex           string    `bun:"sex,notnull"`
	Height        int       `bun:"height,nullzero"`
	Weight        int       `bun:"weight,nullzero"`
	Age           int       `bun:"age,nullzero"`
	Job           string    `bun:"job,nullzero"`
	AccessToken   string    `bun:"access_token"`
	RefreshToken  string    `bun:"refresh_token"`
	CreatedAt     time.Time `bun:"created_at"`
}

type Activities struct {
	// Momentum is a struct that represents a user's Momentum data.
	bun.BaseModel `bun:"table:activities,select:activities"`
	ID            int       `bun:"id,pk,autoincrement"`
	UserID        string    `bun:"user_id,notnull"`
	Steps         int       `bun:"steps,notnull"`
	Calories      int       `bun:"calories,notnull"`
	Date          string    `bun:"date,notnull"`
	CreatedAt     time.Time `bun:"created_at"`
}

type Sleep struct {
	// Sleep is a struct that represents a user's sleep data.
	bun.BaseModel `bun:"table:sleep,select:sleep"`
	ID            int       `bun:"id,pk,autoincrement"`
	UserID        string    `bun:"user_id,notnull"`
	Minutes       int       `bun:"minutes,notnull"`
	DeepSleep     int       `bun:"deep_sleep,notnull"`
	LightSleep    int       `bun:"light_sleep,notnull"`
	RemSleep      int       `bun:"rem_sleep,notnull"`
	Wake          int       `bun:"wake,notnull"`
	Date          string    `bun:"date,notnull"`
	CreatedAt     time.Time `bun:"created_at"`
}

type Meal struct {
	// Meal is a struct that represents a user's meal data.
	bun.BaseModel `bun:"table:meal,select:meal"`
	ID            int       `bun:"id,pk,autoincrement"`
	UserID        string    `bun:"user_id,notnull"`
	MealName      string    `bun:"meal_name,notnull"`
	Calories      int       `bun:"calories,notnull"`
	Protein       float64   `bun:"protein,notnull"`
	Fat           float64   `bun:"fat,notnull"`
	Carbohydrates float64   `bun:"carbohydrates,notnull"`
	Salt          float64   `bun:"salt,notnull"`
	Calcium       float64   `bun:"calcium,notnull"`
	Date          string    `bun:"date,notnull"`
	CreatedAt     time.Time `bun:"created_at"`
}

type ToGetFitbitData struct {
	bun.BaseModel `bun:"table:users,select:users" json:"-"`
	FitbitUserID  string `bun:"fitbit_user_id" json:"fitbit_user_id"`
	AccessToken   string `bun:"access_token" json:"access_token"`
}
