package users

import (
	"net/http"

	"github.com/Oudom/bookstore_users-api/domain/users"
	"github.com/Oudom/bookstore_users-api/services"
	"github.com/Oudom/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// fmt.Println(user)
// bytes, err := ioutil.ReadAll(c.Request.Body)
// if err != nil {
// 	//TODO: Handle error
// 	return
// }
// if err := json.Unmarshal(bytes, &user); err != nil {
// 	fmt.Println(err.Error())
// 	//TODO: Handle json error
// 	return
// }
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	// userId, userErr := strconv.ParseInt(c.Params("user_id"), 10, 64)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
