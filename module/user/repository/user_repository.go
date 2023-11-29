package repository

import (
	"errors"
	"fmt"
	"go-template-project/constant"
	"go-template-project/dto"
	user_entity "go-template-project/module/user/entity"
	"go-template-project/util"
	"gorm.io/gorm"
	"strings"
)

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) UserList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError) {
	var ListEntityUser []*user_entity.User

	db := r.db.Debug().Model(&user_entity.User{})

	if input.SearchText != "" {
		search := "%" + strings.ToLower(input.SearchText) + "%"
		db = db.Where(fmt.Sprintf("LOWER(%v.name) LIKE ? OR LOWER(%v.email) LIKE ? OR LOWER(%v.phone) LIKE ? ", constant.TABLE_MS_USER, constant.TABLE_MS_USER, constant.TABLE_MS_USER),
			search, search, search)
	}

	if input.IsActive != nil {
		db = db.Where(fmt.Sprintf("%v.is_active = ?", constant.TABLE_MS_USER), input.IsActive)
	}

	var count int64

	dbCount := db.Table(constant.TABLE_MS_USER).Model(user_entity.User{}).Count(&count)
	if dbCount.Error != nil && !errors.Is(dbCount.Error, gorm.ErrRecordNotFound) {
		return nil, 0, dto.ResponseError{Error: dbCount.Error, Code: 500}
	}

	if count < 1 {
		return []*dto.UsersResponse{}, 0, dto.ResponseError{}
	}
	orderByQuery := ""
	if input.OrderField != "" {
		orderColumn, orderType := util.SplitOrderQuery(input.OrderField)
		switch orderColumn {
		case "id":
			orderByQuery += fmt.Sprintf(" %v.id %v", constant.TABLE_MS_USER, orderType)
		case "name":
			orderByQuery += fmt.Sprintf(" %v.name %v", constant.TABLE_MS_USER, orderType)
		default:
			orderByQuery += fmt.Sprintf(" %v.id DESC", constant.TABLE_MS_USER)
		}
	} else {
		orderByQuery += "id DESC"
	}
	db = db.Order(orderByQuery)

	if input.Limit > 0 && input.Page > 0 {
		offset := input.Page*input.Limit - input.Limit
		db = db.Limit(input.Limit).Offset(offset)
	}

	Find := db.Debug().Find(&ListEntityUser)
	if Find.Error != nil {
		if errors.Is(Find.Error, gorm.ErrRecordNotFound) {
			return nil, 0, dto.ResponseError{Error: Find.Error, Code: 401}
		}
		return nil, 0, dto.ResponseError{Error: Find.Error, Code: 500}
	}

	var ListSchemaUser []*dto.UsersResponse
	for _, user := range ListEntityUser {
		DataUser := dto.UsersResponse{
			Id:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			AvatarPath: user.AvatarPath,
		}
		if user.IsActive != nil {
			DataUser.IsActive = *user.IsActive
		}

		ListSchemaUser = append(ListSchemaUser, &DataUser)
	}

	return ListSchemaUser, count, dto.ResponseError{}
}

func (r *UserRepository) GetById(id uint64) (*user_entity.User, dto.ResponseError) {
	EntityUser := user_entity.User{}

	db := r.db.Where("id = ?", id)

	Find := db.First(&EntityUser)

	if Find.Error != nil {
		if errors.Is(Find.Error, gorm.ErrRecordNotFound) {
			return nil, dto.ResponseError{Error: Find.Error, Code: 401}
		}
		return nil, dto.ResponseError{Error: Find.Error, Code: 500}
	}

	return &EntityUser, dto.ResponseError{}
}

func (r *UserRepository) UserInsert(input user_entity.User) dto.ResponseError {

	IsActive := true
	input.IsActive = &IsActive
	Create := r.db.Create(&input)
	if Create.Error != nil {
		return dto.ResponseError{Error: Create.Error, Code: 500}
	}
	if Create.RowsAffected < 1 {
		return dto.ResponseError{Error: errors.New("user Create Error"), Code: 500}
	}

	return dto.ResponseError{}

}

func (r *UserRepository) CheckEmail(email string) (*user_entity.User, dto.ResponseError) {

	var user user_entity.User
	CheckEmail := r.db.Debug().Where("email = ?", email).First(&user)
	if CheckEmail.Error != nil {
		return nil, dto.ResponseError{Error: CheckEmail.Error, Code: 404}
	}

	return &user, dto.ResponseError{}
}

func (r *UserRepository) CheckUsername(username string) (*user_entity.User, dto.ResponseError) {
	var user user_entity.User
	CheckUsername := r.db.Where("username = ?", username).First(&user)
	if CheckUsername.Error != nil {
		return nil, dto.ResponseError{Error: CheckUsername.Error, Code: 404}
	}

	return &user, dto.ResponseError{}
}

func (r *UserRepository) UserUpdate(input user_entity.User) (*user_entity.User, dto.ResponseError) {

	var entity user_entity.User
	find := r.db.Where("id = ?", input.ID).First(&entity)
	if find.Error != nil {
		if errors.Is(find.Error, gorm.ErrRecordNotFound) {
			return nil, dto.ResponseError{Error: find.Error, Code: 401}
		}
		return nil, dto.ResponseError{Error: find.Error, Code: 500}
	}

	entity.Name = input.Name
	entity.Email = input.Email
	entity.Phone = input.Phone

	entity.IsActive = input.IsActive

	update := r.db.Save(&entity)
	if update.Error != nil {
		return nil, dto.ResponseError{Error: update.Error, Code: 500}
	}

	return &entity, dto.ResponseError{}
}

func (r *UserRepository) ChangePassword(user_id uint64, new_password string) dto.ResponseError {

	entity := user_entity.User{
		ID:       user_id,
		Password: new_password,
	}

	Update := r.db.Updates(&entity)
	if Update.Error != nil {
		return dto.ResponseError{Error: Update.Error, Code: 500}
	}

	return dto.ResponseError{}
}
