package entity

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}
