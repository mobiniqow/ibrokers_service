import os
import re
from utils.pascal_to_snake import pascal_to_snake
from utils.config import LOCATION

def create_repository_file(class_name):
    package_name = pascal_to_snake(class_name)
    
    file_package_name = LOCATION+package_name
    
    if not os.path.exists(file_package_name):
        os.makedirs(file_package_name)
    repository_filename = os.path.join(LOCATION+package_name, "repository.go")

    # ایجاد محتویات فایل
    repository_content = f"""package {package_name}

import (
    "errors"
    "ibrokers_service/pkg/helper"
    "ibrokers_service/pkg/middleware/filter/operators"
    "ibrokers_service/pkg/middleware/pagination"

    "gorm.io/gorm"
)

type Repository struct {{
    DB *gorm.DB
}}

func (r *Repository) Create{class_name}(item {class_name}) ({class_name}, error) {{
    result := r.DB.Create(&item)
    if result.Error != nil {{
        return {class_name}{{}}, result.Error
    }}
    return item, nil
}}

func (r *Repository) GetAll{class_name}s(limit, page int, filters []operators.FilterBlock) (items []{class_name}, count int64) {{
    _query := helper.QueryBuilder({class_name}{{}}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}}

func (r *Repository) Update{class_name}(item {class_name}) error {{
    result := r.DB.Save(&item)
    return result.Error
}}

func (r *Repository) Delete{class_name}(item {class_name}) error {{
    result := r.DB.Delete(&item)
    return result.Error
}}

func (r *Repository) Find{class_name}ById(id int) ({class_name}, error) {{
    var item {class_name}
    result := r.DB.First(&item, id)
    if result.Error != nil {{
        return {class_name}{{}}, errors.New("not found {class_name.lower()}")
    }}
    return item, nil
}}
"""

    # ذخیره مدل در فایل
    with open(repository_filename, 'w') as file:
        file.write(repository_content)
    
    print(f"Repository done.")

 
