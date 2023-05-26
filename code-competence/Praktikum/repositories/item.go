package repositories

import (
	"code_competence/database"
	"code_competence/models"
	"code_competence/models/input"
)

type ItemRepositoryImp struct {
}

func InitItemRepository() ItemRepository {
	return &ItemRepositoryImp{}
}

func (ur *ItemRepositoryImp) Create(itemInput input.ItemInput) (models.Item, error) {
	var item models.Item = models.Item{
		Name: itemInput.Name,
		Price: itemInput.Price,
		Description: itemInput.Description,
		Stock: itemInput.Stock,
		CategoryID: itemInput.CategoryID,
	}

	if err := database.ConnectDB().Create(&item).Error; err != nil {
		return models.Item{}, err
	}

	if err := database.ConnectDB().Last(&item).Error; err != nil {
		return models.Item{}, err
	}

    return item, nil
}

func (ur *ItemRepositoryImp) GetAll() ([]models.Item, error) {
	var items []models.Item

	if err := database.ConnectDB().Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (ur *ItemRepositoryImp) GetById(id string) (models.Item, error) {
	var item models.Item

	if err := database.ConnectDB().First(&item, id).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (ur *ItemRepositoryImp) GetItemsByCategoryId(id string) ([]models.Item, error) {
	var items []models.Item

	if err := database.ConnectDB().Find(&items, "`category_id` = ?", id).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (ur *ItemRepositoryImp) GetItemsByName(name string) ([]models.Item, error) {
	var items []models.Item

	if err := database.ConnectDB().Find(&items, "name = ?", name).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (ur *ItemRepositoryImp) Update(itemInput input.ItemInput, id string) (models.Item, error) {
	item, err := ur.GetById(id)

	if err != nil {
		return models.Item{}, err
	}

	item.Name = itemInput.Name
	item.Price = itemInput.Price
	item.Description = itemInput.Description
	item.Stock = itemInput.Stock
	item.CategoryID = itemInput.CategoryID

	if err := database.ConnectDB().Save(&item).Error; err != nil {
		return models.Item{}, err
	}

    return item, nil
}

func (ur *ItemRepositoryImp) Delete(id string) error {
	item, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.ConnectDB().Delete(&item).Error; err != nil {
		return err
	}

    return nil
}


