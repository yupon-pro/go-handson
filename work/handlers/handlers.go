package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yupon-pro/go-handson/work/warehouse"
)

func GetUsers(c echo.Context) error {
	aboveAgeStr := c.QueryParam("age_over")
	belowAgeStr := c.QueryParam("age_under")

	if aboveAgeStr == "" && belowAgeStr == "" {
			return c.JSON(http.StatusOK, warehouse.UserWH.UserList)
	}

	var aboveAge, belowAge int
	var err error

	if aboveAgeStr != "" {
			aboveAge, err = strconv.Atoi(aboveAgeStr)
			if err != nil {
					return c.String(http.StatusBadRequest, "The type of age_over is wrong.")
			}
	}
	
	if belowAgeStr != "" {
			belowAge, err = strconv.Atoi(belowAgeStr)
			if err != nil {
					return c.String(http.StatusBadRequest, "The type of age_under is wrong.")
			}
	}

	var uLis []warehouse.User
	for _, v := range warehouse.UserWH.UserList {
		if aboveAgeStr != "" && belowAgeStr != "" && belowAge >= aboveAge{
			if  aboveAge <= v.Age && belowAge >= v.Age{
				uLis = append(uLis, v)
			}
		}else{
			if aboveAgeStr != "" && v.Age >= aboveAge {
				uLis = append(uLis, v)
			}
			if belowAgeStr != "" && v.Age <= belowAge {
				uLis = append(uLis, v)
			}
		}
	}

	return c.JSON(http.StatusOK, uLis)
}

func GetUser(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.String(http.StatusBadRequest, "The id is wrong.")
	}

	uLis := warehouse.UserWH.UserList
	for _,v := range uLis{
		if v.Id == id{
			return c.JSON(http.StatusOK, v)
		}
	}
	
	return c.String(http.StatusBadRequest, "There was no user information you specified.")
}

func PostUser(c echo.Context) error{
	u := new(warehouse.User)
	if err := c.Bind(u); err != nil{
		return c.String(http.StatusBadRequest, "Bad Request!")
	}
	warehouse.UserWH.LastId++
	u.Id = warehouse.UserWH.LastId
	warehouse.UserWH.UserList = append(warehouse.UserWH.UserList, *u)
	return c.JSON(http.StatusCreated, u)
}

func PatchUser(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.String(http.StatusBadRequest, "The id is wrong.")
	}

	age := c.FormValue("age")
	name := c.FormValue("name")

	if age == "" && name == "" {
		return c.String(http.StatusBadRequest, "Please provide either name or age.")
	}

	for i, v := range warehouse.UserWH.UserList {
		if v.Id == id {
			if age != "" {
				intAge, err := strconv.Atoi(age)
				if err != nil {
					return c.String(http.StatusBadRequest, "The type of age is wrong.")
				}
				warehouse.UserWH.UserList[i].Age = intAge
			}
			if name != "" {
				warehouse.UserWH.UserList[i].Name = name
			}
			return c.JSON(http.StatusOK, warehouse.UserWH.UserList[i])
		}
	}

	return c.String(http.StatusBadRequest, "There was no user information you specified.")
	
}

func DeleteUser(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.String(http.StatusBadRequest, "The id is wrong.")
	}
	for i,v := range warehouse.UserWH.UserList{
		if v.Id == id{
			warehouse.UserWH.UserList = append(warehouse.UserWH.UserList[:i],warehouse.UserWH.UserList[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.String(http.StatusBadRequest, "There was no user information you specified.")
}