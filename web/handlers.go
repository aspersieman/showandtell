package web

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	auth "bitbucket.org/envirovisionsolutions/showandtell/auth"
	database "bitbucket.org/envirovisionsolutions/showandtell/database"
	models "bitbucket.org/envirovisionsolutions/showandtell/models"
	types "bitbucket.org/envirovisionsolutions/showandtell/types"
	utils "bitbucket.org/envirovisionsolutions/showandtell/utils"
)

type Controller struct {
	Keycloak *auth.Keycloak
}

func NewController(keycloak *auth.Keycloak) *Controller {
	return &Controller{
		Keycloak: keycloak,
	}
}

func (c *Controller) Home(ctx *fiber.Ctx) error {
	return ctx.Render("web/static/src/dist/index", fiber.Map{
		"Version": utils.GetAppVersion(),
	})
}

// ApiAuthLogin Login with Keycloak
//
//	@Summary      Login with Keycloak
//	@Description  Login to Keycloak with username and password
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param        request 						body      types.LoginRequest  true  "Login request"
//	@Success      200						{object}  types.LoginResponse
//	@Failure      400						{object}  error
//	@Failure      403						{object}  error
//	@Router       /auth/login 			[post]
func (c *Controller) ApiAuthLogin(ctx *fiber.Ctx) error {
	lr := types.LoginRequest{}
	if err := ctx.BodyParser(&lr); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ApiResponse{Data: "Bad request"})
	}

	jwt, err := c.Keycloak.Gocloak.Login(context.Background(),
		c.Keycloak.ClientId,
		c.Keycloak.ClientSecret,
		c.Keycloak.Realm,
		lr.Username,
		lr.Password,
	)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(types.ApiResponse{Data: err.Error()})
	}
	ls := &types.LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: ls})
}

// ApiGetUsers lists a single or a list of user records
//
//	@Summary      List users
//	@Description  get users
//	@Tags         users
//	@Accept       json
//	@Produce      json
//	@Param        Authorization 						header      string  true  "Authorization header"  Format(string)
//	@Param        id 						path      int  false  "find by id"  Format(integer)
//	@Param        b  						query     int  false  "list record position begin"  Format(integer)
//	@Param        e  						query     int  false  "list record position end"  Format(integer)
//	@Success      200						{array}   models.User
//	@Failure      400						{object}  error
//	@Failure      404						{object}  error
//	@Failure      500						{object}  error
//	@Router       /users/ 			[get]
//	@Router       /users/{id}  [get]
func (c *Controller) ApiGetUsers(ctx *fiber.Ctx) error {
	begin := ctx.QueryInt("b", 0)
	end := ctx.QueryInt("e", 10)
	id := ctx.Params("id")

	if id != "" {
		user := models.User{}
		database.DB.Db.Where("id = ?", id).First(&user)
		return ctx.Status(200).JSON(types.ApiResponse{Data: user})
	} else {
		users := []models.User{}
		database.DB.Db.Offset(begin).Limit(end).Find(&users)
		return ctx.Status(200).JSON(types.ApiResponse{Data: users})
	}
}

// ApiGetSchedules lists a single or a list of schedule records
//
//	@Summary      List schedules
//	@Description  get schedules
//	@Tags         schedules
//	@Accept       json
//	@Produce      json
//	@Param        id 						path      int  false  "find by id"  Format(integer)
//	@Param        b  						query     int  false  "list record position begin"  Format(integer)
//	@Param        e  						query     int  false  "list record position end"  Format(integer)
//	@Param        f  						query     int  false  "list record date from"  Format(date)
//	@Param        t  						query     int  false  "list record date to"  Format(date)
//	@Success      200						{array}   models.Schedule
//	@Failure      400						{object}  error
//	@Failure      404						{object}  error
//	@Failure      500						{object}  error
//	@Router       /schedules/ 			[get]
//	@Router       /schedules/{id}  [get]
func (c *Controller) ApiGetSchedules(ctx *fiber.Ctx) error {
	currentTime := time.Now()
	sixMonthsFromCurrent := currentTime.AddDate(0, 6, 0)
	begin := ctx.QueryInt("b", 0)
	end := ctx.QueryInt("e", 10)
	fromQuery := ctx.Query("f", currentTime.Format("2006-01-02"))
	from, err := time.Parse("2006-01-02", fromQuery)
	if err != nil {
		from = currentTime
	}
	toQuery := ctx.Query("t", sixMonthsFromCurrent.Format("2006-01-02"))
	to, err := time.Parse("2006-01-02", toQuery)
	if err != nil {
		to = sixMonthsFromCurrent
	}
	id := ctx.Params("id")

	if id != "" {
		schedule := models.Schedule{}
		database.DB.Db.Where("id = ?", id).First(&schedule)
		return ctx.Status(200).JSON(types.ApiResponse{Data: schedule})
	} else {
		schedules := []models.Schedule{}
		database.DB.Db.Where("start_date_time >= ? AND end_date_time <= ?", from, to).Offset(begin).Limit(end).Find(&schedules)
		return ctx.Status(200).JSON(types.ApiResponse{Data: schedules})
	}
}

// ApiGetVersion lists current and supported versions
//
//	@Summary      List versions
//	@Description  get versions
//	@Tags         versions
//	@Accept       json
//	@Produce      json
//	@Success      200						{array}   types.Version
//	@Failure      400						{object}  error
//	@Failure      404						{object}  error
//	@Failure      500						{object}  error
//	@Router       /version 			[get]
func (c *Controller) ApiGetVersion(ctx *fiber.Ctx) error {
	version := utils.GetAppVersion()
	return ctx.Status(200).JSON(version)
}
