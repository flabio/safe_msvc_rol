package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"

	"github.com/msvc_rol/core/interfaces"
	"github.com/msvc_rol/core/repositories"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/infrastructure/entities"
	"github.com/msvc_rol/usecases/dto"
	"github.com/ulule/deepcopier"
)

type rolService struct {
	Irol interfaces.IRol
}

func NewRolService() globalinterfaces.IRolGlobal {
	return &rolService{
		Irol: repositories.GetRolInstance(),
	}
}

func (rolService *rolService) GetFindAll(c *fiber.Ctx) error {
	result, err := rolService.Irol.GetFindAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.STATUS: http.StatusOK,
		constants.DATA:   result,
	})
}
func (rolService *rolService) GetFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	result, err := rolService.Irol.GetFindById(id)
	log.Println(result)
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  http.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
func (rolService *rolService) Create(c *fiber.Ctx) error {
	var rolCreate entities.Rol
	var rol dto.RolDTO

	rolDtos, msgError := validateRol(0, rol, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(rolDtos).To(&rolCreate)
	result, err := rolService.Irol.Create(rolCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.STATUS:  http.StatusOK,
		constants.MESSAGE: constants.CREATED,
		constants.DATA:    result,
	})
}
func (rolService *rolService) Update(c *fiber.Ctx) error {

	var rolEntity entities.Rol
	var rolDto dto.RolDTO

	id, _ := strconv.Atoi(c.Params(constants.ID))
	rol, _ := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  http.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}

	rolDtos, msgError := validateRol(rol.Id, rolDto, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(rolDtos).To(&rolEntity)
	result, err := rolService.Irol.Update(rol.Id, rolEntity)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		constants.STATUS:  http.StatusCreated,
		constants.MESSAGE: constants.UPDATED,
		constants.DATA:    result,
	})
}

func (rolService *rolService) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	rol, _ := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  http.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}
	result, err := rolService.Irol.Delete(rol.Id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		constants.STATUS:  http.StatusCreated,
		constants.MESSAGE: constants.REMOVED,
		constants.DATA:    result,
	})
}
func validateRol(id uint, rolDto dto.RolDTO, rolService *rolService, c *fiber.Ctx) (dto.RolDTO, string) {
	var msg string = ""
	b := c.Body()

	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)

	if errJson != nil {
		msg = errJson.Error()
	}
	msgValid := validateField(dataMap)
	if msgValid != constants.EMPTY {
		return dto.RolDTO{}, msgValid
	}

	MapToStruct(&rolDto, dataMap)
	msgRequired := validateRequired(rolDto)
	if msgRequired != constants.EMPTY {
		return dto.RolDTO{}, msgRequired
	}
	existName, _ := rolService.Irol.GetFindByName(id, rolDto.Name)
	if existName {
		msg = constants.NAME_ALREADY_EXIST
	}
	return rolDto, msg
}

func MapToStruct(dataDto *dto.RolDTO, dataMap map[string]interface{}) {
	rol := dto.RolDTO{
		Name:   dataMap[constants.NAME].(string),
		Active: dataMap[constants.ACTIVE].(bool),
	}
	*dataDto = rol
}
func validateField(value map[string]interface{}) string {
	var msg string = constants.EMPTY
	if value[constants.NAME] == nil {
		msg = constants.NAME_FIELD_IS_REQUIRED
	}
	if value[constants.ACTIVE] == nil {
		msg = constants.ACTIVE_FIELD_IS_REQUIRED
	}
	return msg
}

func validateRequired(field dto.RolDTO) string {
	var msg string = constants.EMPTY
	if field.Name == constants.EMPTY {
		msg = constants.NAME_IS_REQUIRED
	}
	return msg
}
