import re
def extract_model_info(django_model):
    model_info = {}
    
    # استخراج نام کلاس
    model_name = re.search(r'class (\w+)', django_model).group(1)
    model_info['class_name'] = model_name
    
    # استخراج فیلدها و نوع داده‌ها
    fields = {}
    for line in django_model.splitlines():
        if 'models.CharField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'str'
        elif 'models.TextField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'str'
        elif 'models.BigIntegerField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'int'
        elif 'models.DateField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'date'
        elif 'models.jDateField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
        elif 'models.PositiveBigIntegerField' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'int'
        elif 'models.ForeignKey' in line:
            field_name = re.search(r'(\w+)\s*=', line).group(1)
            fields[field_name] = 'int'
        # می‌توانی سایر انواع فیلدها رو هم اضافه کنی
    
    model_info['fields'] = fields
    return model_info