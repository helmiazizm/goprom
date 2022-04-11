package delivery

import (
	"fmt"
	"goprom/model"

	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	publicRoute *gin.RouterGroup
}

func NewStudentApi(publicRoute *gin.RouterGroup) *StudentApi {
	studentApi := new(StudentApi)
	studentApi.publicRoute = publicRoute
	studentApi.initRoute()
	return studentApi
}

func (api *StudentApi) initRoute() {
	stdRoute := api.publicRoute.Group("/students")
	stdRoute.POST("", api.createStudent)
	stdRoute.GET("/:idcard", api.getStudentById)
}

func (api *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Created student: %v", student),
	})
}

func (api *StudentApi) getStudentById(c *gin.Context) {
	name := c.Param("idcard")
	c.JSON(200, gin.H{
		"message": name,
	})
}