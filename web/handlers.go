package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	auth "bitbucket.org/envirovisionsolutions/showandtell/auth"
	database "bitbucket.org/envirovisionsolutions/showandtell/database"
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
//	@Param        request 			body      types.LoginRequest  true  "Login request"
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
	keycloak := auth.NewKeycloak()
	decodedToken, _, err := keycloak.Gocloak.DecodeAccessToken(context.Background(), jwt.AccessToken, keycloak.Realm)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.ApiResponse{Data: fmt.Sprintf("Invalid or malformed token: %s", err.Error())})
	}
	decodedAccessToken := types.DecodedAccessToken{}
	tokenData, _ := json.Marshal(decodedToken)
	_ = json.Unmarshal(tokenData, &decodedAccessToken)
	ls := &types.LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
		Decoded:      decodedAccessToken,
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: ls})
}

// ApiAuthLogout Logout from Keycloak
//
//	@Summary      Logout from Keycloak
//	@Description  Logout from Keycloak
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param        request 			body      types.LogoutRequest  true  "Logout request"
//	@Success      200						{object}  types.ApiResponse
//	@Failure      400						{object}  error
//	@Failure      403						{object}  error
//	@Router       /auth/logout 			[post]
func (c *Controller) ApiAuthLogout(ctx *fiber.Ctx) error {
	lr := types.LogoutRequest{}
	if err := ctx.BodyParser(&lr); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ApiResponse{Data: "Bad request"})
	}

	err := c.Keycloak.Gocloak.Logout(context.Background(),
		c.Keycloak.ClientId,
		c.Keycloak.ClientSecret,
		c.Keycloak.Realm,
		lr.RefreshToken,
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: "logged out"})
}

// ApiGetUsers lists a single or a list of user records
//
//	@Summary      List users
//	@Description  get users
//	@Tags         users
//	@Accept       json
//	@Produce      json
//	@Param        Authorization 						header      	string  true  "Authorization header"  Format(string)
//	@Param        id 					 		path      int 					false  "find by id"  Format(integer)
//	@Param        b  					 		query     int 					false  "list record position begin"  Format(integer)
//	@Param        e  					 		query     int 					false  "list record position end"  Format(integer)
//	@Success      200					 		{array}   database.User
//	@Failure      400					 		{object}  error
//	@Failure      404					 		{object}  error
//	@Failure      500					 		{object}  error
//	@Router       /users/ 		 		[get]
//	@Router       /users/{id}  		[get]
func (c *Controller) ApiGetUsers(ctx *fiber.Ctx) error {
	begin := ctx.QueryInt("b", 0)
	end := ctx.QueryInt("e", 10)
	id := ctx.Params("id")

	if id != "" {
		user := database.User{}
		database.DB.Db.Where("id = ?", id).First(&user)
		return ctx.Status(200).JSON(types.ApiResponse{Data: user})
	} else {
		users := []database.User{}
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
//	@Success      200						{array}   database.Schedule
//	@Failure      400						{object}  error
//	@Failure      404						{object}  error
//	@Failure      500						{object}  error
//	@Router       /schedules/ 			[get]
//	@Router       /schedules/{id}  [get]
func (c *Controller) ApiGetSchedules(ctx *fiber.Ctx) error {
	currentTime := time.Now()
	sixMonthsFromCurrent := currentTime.AddDate(0, 6, 0)
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
		schedule := database.Schedule{}
		err = database.DB.Db.Model(&database.Schedule{}).Preload("Speakers", func(db *gorm.DB) *gorm.DB {
			return db.Order("`speakers`.`order` ASC")
		}).Where("id = ?", id).First(&schedule).Error
		if err != nil {
			log.Println(err)
		}
		return ctx.Status(200).JSON(types.ApiResponse{Data: schedule})
	} else {
		schedules := []database.Schedule{}
		err = database.DB.Db.Model(&database.Schedule{}).Preload("Speakers", func(db *gorm.DB) *gorm.DB {
			return db.Order("`speakers`.`order` ASC")
		}).Where("start_date_time >= ? AND start_date_time <= ?", from, to).Find(&schedules).Error
		if err != nil {
			log.Println(err)
		}
		return ctx.Status(200).JSON(types.ApiResponse{Data: schedules})
	}
}

// ApiPostSchedules add a new schedule
//
//	@Summary      Add schedule
//	@Description  add schedule
//	@Tags         schedules
//	@Accept       json
//	@Produce      json
//	@Param        request 			body      types.SchedulePostRequest  true  "Schedule"
//	@Success      200						{object}  database.Schedule
//	@Failure      400						{object}  error
//	@Failure      404						{object}  error
//	@Failure      500						{object}  error
//	@Router       /schedules 		[post]
func (c *Controller) ApiPostSchedules(ctx *fiber.Ctx) error {
	sr := types.SchedulePostRequest{}
	if err := ctx.BodyParser(&sr); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ApiResponse{Data: "Bad request"})
	}

	schedule := database.Schedule{
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     nil,
		Title:         sr.Title,
		Description:   sr.Description,
		StartDateTime: sr.StartDateTime,
		EndDateTime:   sr.EndDateTime,
	}

	result := database.DB.Db.Create(&schedule)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: schedule})
}

// ApiPutSchedules update a schedule
//
//	@Summary      Edit schedule
//	@Description  edit schedule
//	@Tags         schedules
//	@Accept       json
//	@Produce      json
//	@Param        id 					 			path      int  true  "request id"  Format(integer)
//	@Param        request 		 			body      types.SchedulePutRequest  true  "Schedule"
//	@Success      200					 			{object}  database.Schedule
//	@Failure      400					 			{object}  error
//	@Failure      404					 			{object}  error
//	@Failure      500					 			{object}  error
//	@Router       /schedules/{id} 	[put]
func (c *Controller) ApiPutSchedules(ctx *fiber.Ctx) error {
	sr := types.SchedulePutRequest{}
	id := ctx.Params("id")
	if err := ctx.BodyParser(&sr); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ApiResponse{Data: "Bad request"})
	}
	scheduleId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Bad request. Schedule ID not valid"})
	}

	schedule := database.Schedule{
		ID:            uint(scheduleId),
		UpdatedAt:     time.Now(),
		DeletedAt:     nil,
		Title:         sr.Title,
		Description:   sr.Description,
		StartDateTime: sr.StartDateTime,
		EndDateTime:   sr.EndDateTime,
	}

	result := database.DB.Db.Save(&schedule)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: schedule})
}

// ApiDeleteSchedules delete a schedule
//
//	@Summary      Delete schedule
//	@Description  delete schedule
//	@Tags         schedules
//	@Accept       json
//	@Produce      json
//	@Param        id 					 			path      int  true  "request id"  Format(integer)
//	@Success      200					 			{object}  database.Schedule
//	@Failure      400					 			{object}  error
//	@Failure      404					 			{object}  error
//	@Failure      500					 			{object}  error
//	@Router       /schedules/{id} 	[delete]
func (c *Controller) ApiDeleteSchedules(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	scheduleId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Bad request. Schedule ID not valid"})
	}

	result := database.DB.Db.Where("id = ?", scheduleId).Delete(&database.Schedule{})

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: "Deleted"})
}

// ApiGetSpeakers lists a single or a list of speaker records
//
//	@Summary      List speakers
//	@Description  get speakers
//	@Tags         speakers
//	@Accept       json
//	@Produce      json
//	@Param        id 					 		path      int  false  "find by id"  Format(integer)
//	@Param        b  					 		query     int  false  "list record position begin"  Format(integer)
//	@Param        e  					 		query     int  false  "list record position end"  Format(integer)
//	@Param        f  					 		query     int  false  "list record date from"  Format(date)
//	@Param        t  					 		query     int  false  "list record date to"  Format(date)
//	@Success      200					 		{array}   database.Speaker
//	@Failure      400					 		{object}  error
//	@Failure      404					 		{object}  error
//	@Failure      500					 		{object}  error
//	@Router       /speakers/ 			[get]
//	@Router       /speakers/{id}  [get]
func (c *Controller) ApiGetSpeakers(ctx *fiber.Ctx) error {
	currentTime := time.Now()
	sixMonthsFromCurrent := currentTime.AddDate(0, 6, 0)
	log.Printf("Six months from current: %s\n", sixMonthsFromCurrent.Format("2006-01-02"))
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
		speaker := database.Speaker{}
		err = database.DB.Db.Model(&database.Speaker{}).Preload("Speakers", func(db *gorm.DB) *gorm.DB {
			return db.Order("`speakers`.`order` ASC")
		}).Where("id = ?", id).First(&speaker).Error
		if err != nil {
			log.Println(err)
		}
		return ctx.Status(200).JSON(types.ApiResponse{Data: speaker})
	} else {
		speakers := []database.Speaker{}
		err = database.DB.Db.Model(&database.Speaker{}).Preload("Speakers", func(db *gorm.DB) *gorm.DB {
			return db.Order("`speakers`.`order` ASC")
		}).Where("start_date_time >= ? AND start_date_time <= ?", from, to).Find(&speakers).Error
		if err != nil {
			log.Println(err)
		}
		return ctx.Status(200).JSON(types.ApiResponse{Data: speakers})
	}
}

// ApiPostSpeakers add a new speaker
//
//	@Summary      Add speaker
//	@Description  add speaker
//	@Tags         speakers
//	@Accept       json
//	@Produce      json
//	@Param        request 			body      types.SpeakerPostRequest  true  "Speaker"
//	@Success      200						{object}  database.Speaker
//	@Failure      400						{object}  error
//	@Failure      404						{object}  error
//	@Failure      500						{object}  error
//	@Router       /speakers 		[post]
func (c *Controller) ApiPostSpeakers(ctx *fiber.Ctx) error {
	sr := types.SpeakerPostRequest{}
	if err := ctx.BodyParser(&sr); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ApiResponse{Data: "Bad request"})
	}

	schedule := database.Schedule{}
	err := database.DB.Db.Model(&database.Schedule{}).Preload("Speakers").Where("id = ?", sr.ScheduleId).First(&schedule).Error
	if err != nil {
		log.Println(err)
	}
	order := len(schedule.Speakers)
	if order == 0 {
		order = 1
	} else {
		order++
	}
	speaker := database.Speaker{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Name:        sr.Name,
		Topic:       sr.Topic,
		Description: sr.Description,
		Order:       order,
		Minutes:     sr.Minutes,
		ScheduleId:  sr.ScheduleId,
		Private:     sr.Private,
	}

	result := database.DB.Db.Create(&speaker)

	if result.Error != nil {
		log.Printf("ERROR: %v\n", result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: speaker})
}

// ApiPutSpeakers update a speaker
//
//	@Summary      Edit speaker
//	@Description  edit speaker
//	@Tags         speakers
//	@Accept       json
//	@Produce      json
//	@Param        id 					 		path      int  true  "request id"  Format(integer)
//	@Param        request 		 		body      types.SpeakerPutRequest  true  "Speaker"
//	@Success      200					 		{object}  database.Speaker
//	@Failure      400					 		{object}  error
//	@Failure      404					 		{object}  error
//	@Failure      500					 		{object}  error
//	@Router       /speakers/{id} 	[put]
func (c *Controller) ApiPutSpeakers(ctx *fiber.Ctx) error {
	sr := types.SpeakerPutRequest{}
	id := ctx.Params("id")
	if err := ctx.BodyParser(&sr); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.ApiResponse{Data: "Bad request"})
	}
	speakerId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Bad request. Speaker ID not valid"})
	}

	speaker := database.Speaker{
		ID:          uint(speakerId),
		UpdatedAt:   time.Now(),
		Name:        sr.Name,
		Topic:       sr.Topic,
		Description: sr.Description,
		Order:       sr.Order,
		Minutes:     sr.Minutes,
		ScheduleId:  sr.ScheduleId,
		Private:     sr.Private,
	}

	result := database.DB.Db.Save(&speaker)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: speaker})
}

// ApiDeleteSpeakers delete a speaker
//
//	@Summary      Delete speaker
//	@Description  delete speaker
//	@Tags         speakers
//	@Accept       json
//	@Produce      json
//	@Param        id 					 		path      int  true  "request id"  Format(integer)
//	@Success      200					 		{object}  database.Speaker
//	@Failure      400					 		{object}  error
//	@Failure      404					 		{object}  error
//	@Failure      500					 		{object}  error
//	@Router       /speakers/{id} 	[delete]
func (c *Controller) ApiDeleteSpeakers(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	speakerId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Bad request. Speaker ID not valid"})
	}

	result := database.DB.Db.Where("id = ?", speakerId).Delete(&database.Speaker{})

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.ApiResponse{Data: "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(types.ApiResponse{Data: "Deleted"})
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
