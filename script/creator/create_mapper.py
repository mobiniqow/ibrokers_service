import os
import re
from utils.pascal_to_snake import pascal_to_snake
from utils.config import LOCATION

def create_mapper(class_name, fields):
    package_name = pascal_to_snake(class_name)
    
    file_package_name = LOCATION+package_name
    
    if not os.path.exists(file_package_name):
        os.makedirs(file_package_name)
    function_filename = os.path.join(LOCATION+package_name, f"mapper.go")

    # Create the fields for the response
    response_fields = []
    time=''
    for field_name, field_type in fields.items():
        if field_type == 'string':
            response_fields.append(f'\t\t{field_name.capitalize()}: &buyMethod.{field_name.capitalize()},')
        elif field_type == 'int':
            response_fields.append(f'\t\t{field_name.capitalize()}: &buyMethod.{field_name.capitalize()},')
        elif field_type == 'date':
            time='import "time"'
            response_fields.append(f'\t\t{field_name.capitalize()}: &buyMethod.{field_name.capitalize()},')

    # Create the content of the ToResponse function
    
    function_content = f"""package {package_name}

{time}


func To{class_name}Response(buyMethod {class_name}) {class_name}Response {{
    return {class_name}Response {{
{"\n".join(response_fields)}
    }}
}}
"""

    # Save the function in the file
    with open(function_filename, 'w') as file:
        file.write(function_content)

    print(f"Mapper done.")
