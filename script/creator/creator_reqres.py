import os 
from utils.pascal_to_snake import pascal_to_snake
from utils.config import LOCATION

def create_model_file(model_data):
    class_name = model_data['class_name']
    package_name = pascal_to_snake(class_name)
    
    file_package_name = LOCATION+package_name
    
    if not os.path.exists(file_package_name):
        os.makedirs(file_package_name)
        
    model_filename = os.path.join(LOCATION+package_name, f"reqres.go")
    time=''
    # Create fields based on the provided model_data
    fields = []
    for field_name, field_type in model_data['fields'].items():
        if field_type == 'string':
            fields.append(f'\t{field_name.capitalize()} *string `form:"{field_name}"`')
        elif field_type == 'int':
            fields.append(f'\t{field_name.capitalize()} *int `form:"{field_name}"`')
        elif field_type == 'date':
            time='import "time";'
            fields.append(f'\t{field_name.capitalize()} *time.Time `form:"{field_name}"`')
    
    # Create the content of the model file
    model_content = f"""package {package_name}

{time}

type Create{class_name}Request struct {{
{"\n".join(fields)}
}}

type {class_name}Response struct {{
    {"\n".join(fields)}
}}
"""

    # Save the model in the file
    with open(model_filename, 'w') as file:
        file.write(model_content)

    print(f"Reqres done.")
 
 
 