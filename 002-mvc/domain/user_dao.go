package domain

var(
	users = map[int64]*User{
		123: &User{
			Id:1,
			FirstName:"Sashank",
			LastName:"Bhogu",
			Email:"myemail@gmail.com",
		}
	}
)


func GetUser(userId int64) (*User, error) {
  var user User = users[userId]

  if user == nil {
	  return User{}, errors.New(fmt.Println("user %v was not found",userId))
  }
}


func Another(){
	user, integer, err := GetUser(1);

	if err != nil{

	}
}