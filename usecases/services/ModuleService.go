package services

import (
	"encoding/json"
	"strconv"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/core/repositories"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/infrastructure/entities"
	"github.com/msvc_rol/infrastructure/ui/uicore"
	"github.com/msvc_rol/usecases/dto"

	"github.com/ulule/deepcopier"
)

type moduleService struct {
	uiModule uicore.UIModuleCore
}

func ModuleInstance() globalinterfaces.UIModuleGlobal {
	return &moduleService{
		uiModule: repositories.GetModuleInstance(),
	}
}

func (s *moduleService) GetModuleFindAll(c *fiber.Ctx) error {
	query, err := s.uiModule.GetModuleFindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS: fiber.StatusBadRequest,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS: fiber.StatusOK,
		constants.DATA:   query,
	})
}
func (s *moduleService) GetModuleFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	query, err := s.uiModule.GetModuleFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			constants.STATUS: fiber.StatusNotFound,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	if query.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			constants.STATUS: fiber.StatusNotFound,
			constants.DATA:   constants.ID_NO_EXIST,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS: fiber.StatusOK,
		constants.DATA:   query,
	})
}
func (s *moduleService) CreateModule(c *fiber.Ctx) error {
	var module entities.Module
	var moduleDto dto.ModuleDTO
	data, msgError := validateModule(0, moduleDto, s, c)
	if msgError != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusBadRequest,
			constants.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(data).To(&module)
	query, err := s.uiModule.CreateModule(module)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			constants.STATUS: fiber.StatusInternalServerError,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		constants.STATUS:  fiber.StatusCreated,
		constants.MESSAGE: constants.CREATED,
		constants.DATA:    query,
	})
}
func (s *moduleService) UpdateModule(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	var module entities.Module
	var moduleDto dto.ModuleDTO

	existId, err := s.uiModule.GetModuleFindById(uint(id))
	if existId.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusInternalServerError,
			constants.MESSAGE: constants.ERROR_QUERY,
		})
	}
	data, msgError := validateModule(uint(existId.Id), moduleDto, s, c)
	if msgError != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusBadRequest,
			constants.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(data).To(&module)
	query, err := s.uiModule.UpdateModule(existId.Id, module)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			constants.STATUS: fiber.StatusInternalServerError,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS:  fiber.StatusOK,
		constants.MESSAGE: constants.UPDATED,
		constants.DATA:    query,
	})
}
func (s *moduleService) DeleteModule(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	existId, err := s.uiModule.GetModuleFindById(uint(id))
	if existId.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusInternalServerError,
			constants.MESSAGE: constants.ERROR_QUERY,
		})
	}
	query, err := s.uiModule.DeleteModule(existId.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusInternalServerError,
			constants.MESSAGE: constants.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS:  fiber.StatusOK,
		constants.MESSAGE: constants.REMOVED,
		constants.DATA:    query,
	})
}

// TODO-> module with role
func (s *moduleService) GetModuleRoleFindAll(c *fiber.Ctx) error {
	result, err := s.uiModule.GetModuleRoleFindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS: fiber.StatusOK,
		constants.DATA:   result,
	})
}
func (s *moduleService) CreateModuleRole(c *fiber.Ctx) error {
	var module entities.ModuleRole
	var moduleDto dto.ModuleRoleDTO
	data, msgError := validateModuleRole(moduleDto, s, c)
	if msgError != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusBadRequest,
			constants.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(data).To(&module)
	query, err := s.uiModule.CreateModuleRole(module)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			constants.STATUS: fiber.StatusInternalServerError,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		constants.STATUS:  fiber.StatusCreated,
		constants.MESSAGE: constants.CREATED,
		constants.DATA:    query,
	})
}
func (s *moduleService) DeleteModuleRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	query, err := s.uiModule.DeleteModuleRole(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusInternalServerError,
			constants.MESSAGE: constants.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS:  fiber.StatusOK,
		constants.MESSAGE: constants.REMOVED,
		constants.DATA:    query,
	})
}

func validateModule(id uint, moduleDto dto.ModuleDTO, s *moduleService, c *fiber.Ctx) (dto.ModuleDTO, string) {
	var msg string
	body := c.Body()
	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
		return moduleDto, msg
	}
	msg = validateFieldModule(dataMap)
	if msg != "" {
		return dto.ModuleDTO{}, msg
	}
	MapToStructModule(&moduleDto, dataMap)
	msg = validateRequiredModule(moduleDto)
	if msg != "" {
		return dto.ModuleDTO{}, msg
	}
	isExistName, _ := s.uiModule.GetModuleFindByName(id, moduleDto.Name)
	if isExistName {
		msg = constants.NAME_ALREADY_EXIST
		return moduleDto, msg
	}
	return moduleDto, msg
}
func MapToStructModule(dataDto *dto.ModuleDTO, dataMap map[string]interface{}) {
	rol := dto.ModuleDTO{
		Name:   dataMap[constants.NAME].(string),
		Icon:   dataMap["icon"].(string),
		Order:  int(dataMap["order"].(float64)),
		Active: dataMap[constants.ACTIVE].(bool),
	}
	*dataDto = rol
}
func validateFieldModule(value map[string]interface{}) string {
	var msg string = constants.EMPTY
	if value[constants.NAME] == nil {
		msg = constants.NAME_FIELD_IS_REQUIRED
	}
	if value["icon"] == nil {
		msg = "The icon field is required "
	}
	if value["order"] == nil {
		msg = "The order field is required "
	}
	if value[constants.ACTIVE] == nil {
		msg = constants.ACTIVE_FIELD_IS_REQUIRED
	}
	return msg
}
func validateRequiredModule(field dto.ModuleDTO) string {
	var msg string = constants.EMPTY
	if field.Name == constants.EMPTY {
		msg = constants.NAME_IS_REQUIRED
	}
	if field.Icon == constants.EMPTY {
		msg = "The icon is required "
	}
	if field.Order == 0 {
		msg = "The order is required "
	}
	return msg
}

// module with role
func validateModuleRole(moduleDto dto.ModuleRoleDTO, s *moduleService, c *fiber.Ctx) (dto.ModuleRoleDTO, string) {
	var msg string
	body := c.Body()
	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
		return moduleDto, msg
	}
	msg = validateFieldModuleRole(dataMap)
	if msg != "" {
		return dto.ModuleRoleDTO{}, msg
	}
	MapToStructModuleRole(&moduleDto, dataMap)
	msg = validateRequiredModuleRole(moduleDto)
	if msg != "" {
		return dto.ModuleRoleDTO{}, msg
	}
	return moduleDto, msg
}

func MapToStructModuleRole(dataDto *dto.ModuleRoleDTO, dataMap map[string]interface{}) {
	fileds := dto.ModuleRoleDTO{
		RoleId:   uint(dataMap["role_id"].(float64)),
		ModuleId: uint(dataMap["module_id"].(float64)),
		Active:   dataMap[constants.ACTIVE].(bool),
	}
	*dataDto = fileds
}
func validateFieldModuleRole(value map[string]interface{}) string {
	var msg string = constants.EMPTY
	if value["role_id"] == nil {
		msg = "The field role_id is required"
	}
	if value["module_id"] == nil {
		msg = "The field module_id is required "
	}
	if value[constants.ACTIVE] == nil {
		msg = constants.ACTIVE_FIELD_IS_REQUIRED
	}
	return msg
}
func validateRequiredModuleRole(field dto.ModuleRoleDTO) string {
	var msg string = constants.EMPTY

	if field.RoleId == 0 {
		msg = "The role is required "
	}
	if field.ModuleId == 0 {
		msg = "The role is required "
	}
	return msg
}
