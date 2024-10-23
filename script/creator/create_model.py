from utils.pascal_to_snake import pascal_to_snake
import os
def create_struct_to_file(model_info):
    model_name = model_info['class_name']
    package_name = pascal_to_snake(model_name)
    
    golang_struct = f"package {package_name}\n\n"
    golang_struct += f"type {model_name} struct {{\n"
    
    for field_name, field_type in model_info['fields'].items():
        golang_struct += f"    {field_name.capitalize()} {field_type} `json:\"{field_name}\"`\n"
    
    golang_struct += "}\n"
     
    if not os.path.exists(package_name):
        os.makedirs(package_name) 
    file_path = os.path.join(package_name, f"model.go")
    with open(file_path, 'w') as file:
        file.write(golang_struct)
    
    print(f"model {file_path} saved.")
