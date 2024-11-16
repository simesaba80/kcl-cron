package db

import (
	"time"

	"github.com/uptrace/bun"
)

// User is a struct that represents a user.
type User struct {
	bun.BaseModel `bun:"table:users,select:users" json:"-"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	UID           string    `bun:"uid,unique,notnull" json:"uid"`
	FitbitUserID  string    `bun:"fitbit_user_id" json:"fitbit_user_id"`
	Name          string    `bun:"name,notnull" json:"name"`
	Sex           string    `bun:"sex,notnull" json:"sex"`
	Height        int       `bun:"height,nullzero" json:"height"`
	Weight        int       `bun:"weight,nullzero" json:"weight"`
	Age           int       `bun:"age,nullzero" json:"age"`
	Job           string    `bun:"job,nullzero" json:"job"`
	AccessToken   string    `bun:"access_token" json:"access_token"`
	RefreshToken  string    `bun:"refresh_token" json:"refresh_token"`
	CreatedAt     time.Time `bun:"created_at" json:"created_at"`
}

type Momentum struct {
	// Momentum is a struct that represents a user's Momentum data.
	bun.BaseModel `bun:"table:momentum,select:momentum" json:"-"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	UserID        string    `bun:"user_id,notnull" json:"user_id"`
	Steps         int       `bun:"steps,notnull" json:"steps"`
	Calories      int       `bun:"calories,notnull" json:"calories"`
	Distance      int       `bun:"distance,notnull" json:"distance"`
	MaxHeartRate  int       `bun:"max_heart_rate,notnull" json:"max_heart_rate"`
	MinHeartRate  int       `bun:"min_heart_rate,notnull" json:"min_heart_rate"`
	Date          string    `bun:"date,notnull" json:"date"`
	CreatedAt     time.Time `bun:"created_at" json:"created_at"`
}

type Sleep struct {
	// Sleep is a struct that represents a user's sleep data.
	bun.BaseModel `bun:"table:sleep,select:sleep" json:"-"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	UserID        string    `bun:"user_id,notnull" json:"user_id"`
	Minutes       int       `bun:"hours,notnull" json:"minutes"`
	DeepSleep     int       `bun:"deep_sleep,notnull" json:"deep_sleep"`
	LightSleep    int       `bun:"light_sleep,notnull" json:"light_sleep"`
	RemSleep      int       `bun:"rem_sleep,notnull" json:"rem_sleep"`
	Wake          int       `bun:"wake,notnull" json:"wake"`
	Date          string    `bun:"date,notnull" json:"date"`
	CreatedAt     time.Time `bun:"created_at" json:"created_at"`
}

type Meal struct {
	// Meal is a struct that represents a user's meal data.
	bun.BaseModel `bun:"table:meal,select:meal" json:"-"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	UserID        string    `bun:"user_id,notnull" json:"user_id"`
	MealName      string    `bun:"meal_name,notnull" json:"meal_name"`
	Calories      int       `bun:"calories,notnull" json:"calories"`
	Protein       float64   `bun:"protein,notnull" json:"protein"`
	Fat           float64   `bun:"fat,notnull" json:"fat"`
	Carbohydrates float64   `bun:"carbohydrates,notnull" json:"carbohydrates"`
	Salt          float64   `bun:"salt,notnull" json:"salt"`
	Calcium       float64   `bun:"calcium,notnull" json:"calcium"`
	Date          string    `bun:"date,notnull" json:"date"`
	CreatedAt     time.Time `bun:"created_at" json:"created_at"`
}
