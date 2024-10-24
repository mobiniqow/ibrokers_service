import re
def extract_model_info(django_model):
    model_info = {}
    
    model_name = re.search(r'class (\w+)', django_model).group(1)
    model_info['class_name'] = model_name
 
    fields = {}
    for line in django_model.splitlines():
        if 'models.CharField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'string'
        elif 'models.TextField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'string'
        elif 'models.BigIntegerField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'int'
        elif 'models.IntegerField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'int'
        elif 'models.DateField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'time.Time'
        elif 'models.jDateField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'time.Time'
        elif 'models.PositiveBigIntegerField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'int'
        elif 'models.ForeignKey' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'int'
    
    model_info['fields'] = fields
    return model_info
