import os 
from utils.pascal_to_snake import pascal_to_snake


def create_model_file(model_data):
    class_name = model_data['class_name']
    package_name = pascal_to_snake(class_name)
    model_filename = os.path.join(package_name, f"reqres.go")

    # Create fields based on the provided model_data
    fields = []
    for field_name, field_type in model_data['fields'].items():
        if field_type == 'str':
            fields.append(f'\t{field_name.capitalize()} *string `json:"{field_name}"`')
        elif field_type == 'int':
            fields.append(f'\t{field_name.capitalize()} *int `json:"{field_name}"`')
        elif field_type == 'date':
            fields.append(f'\t{field_name.capitalize()} *time.Time `json:"{field_name}"`')
    
    # Create the content of the model file
    model_content = f"""package {package_name}

import "time"

type Create{class_name}Request struct {{
{"\n".join(fields)}
}}

type Response struct {{
    {class_name} {class_name} `json:"{class_name}"`
}}
"""

    # Create the package directory if it doesn't exist
    if not os.path.exists(package_name):
        os.makedirs(package_name)

    # Save the model in the file
    with open(model_filename, 'w') as file:
        file.write(model_content)

    print(f"File {model_filename} saved.")
 
 
 