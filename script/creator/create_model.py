from utils.pascal_to_snake import pascal_to_snake
import os
from utils.config import LOCATION

def create_struct_to_file(model_info):
    model_name = model_info['class_name']
    package_name = pascal_to_snake(model_name)
    
    file_package_name = LOCATION+package_name
    
    if not os.path.exists(file_package_name):
        os.makedirs(file_package_name)
    golang_struct = f"package {package_name}\n\n"
    for field_name, field_type in model_info['fields'].items():
        if "time.Time" in field_type:
            golang_struct += 'import "time";'
            break
    
    golang_struct += f"type {model_name} struct {{\n"
    
    for field_name, field_type in model_info['fields'].items():
        if field_name == "id":
            golang_struct += "    Id            int    `form:\"id\" gorm:\"primary_key\"`\n"
        else:
            golang_struct += f"    {field_name.capitalize()} {field_type} `form:\"{field_name}\"`\n"
    
    golang_struct += "}\n"
    file_path = os.path.join(LOCATION+package_name, f"model.go")
    with open(file_path, 'w') as file:
        file.write(golang_struct)
    
    print(f"Model done.")
