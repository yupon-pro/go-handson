package warehouse

type User struct{
	Id int `json:"id" xml:"id" form:"id" query:"id"`
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Age int	`json:"age" xml:"age" form:"age" query:"age"`
}

type UserWareHouse struct{
	LastId int
	UserList []User
}

var UserWH = UserWareHouse{
	LastId: 3,
	UserList: []User{{
		Id: 1,
		Name: "Joe Biden",
		Age: 82,
	},
	{
		Id: 2,
		Name: "Donald Trump",
		Age: 79,
	},
	{
		Id: 3,
		Name: "Robert Kennedy Junior",
		Age: 70,
	},
	},
}