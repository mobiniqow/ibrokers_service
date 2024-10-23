import re

def pascal_to_snake(pascal_str):
    # اضافه کردن "_" قبل از هر حرف بزرگ که در وسط رشته است و تبدیل به lowercase
    snake_str = re.sub(r'(?<!^)(?=[A-Z])', '_', pascal_str).lower()
    return snake_str