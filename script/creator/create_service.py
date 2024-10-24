import os
import re
from utils.pascal_to_snake import pascal_to_snake
from utils.config import LOCATION

def create_service_file(class_name):
    package_name = pascal_to_snake(class_name)
    
    file_package_name = LOCATION+package_name
    
    if not os.path.exists(file_package_name):
        os.makedirs(file_package_name)
        
    service_filename = os.path.join(LOCATION+package_name, "service.go")
 
    service_content = f"""package {package_name}

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {{
    Repository Repository
}}

func (s *Service) Create{class_name}(item {class_name}) ({class_name}, error) {{
    return s.Repository.Create{class_name}(item)
}}

func (s *Service) Update{class_name}(item {class_name}) error {{
    _, err := s.Repository.Find{class_name}ById(item.Id)
    if err != nil {{
        return err
    }}
    return s.Repository.Update{class_name}(item)
}}

func (s *Service) Delete{class_name}(item {class_name}) error {{
    return s.Repository.Delete{class_name}(item)
}}

func (s *Service) GetAll{class_name}s(limit, page int, filters []operators.FilterBlock) ([]{class_name}, int64) {{
    return s.Repository.GetAll{class_name}s(limit, page, filters)
}}
"""

 
    with open(service_filename, 'w') as file:
        file.write(service_content)
    
    print(f"Service done.")
 
 