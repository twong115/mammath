package user

type User struct {
	name string
	points int
}

func New(username string, points int) *User {
	return &User{name: username, points: points}
}

func (u *User) SetName(newName string) {
	u.name = newName
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetPoints(point int) {
	u.points = point
}

func (u *User) GetPoints() int {
	return u.points
}
