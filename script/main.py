from creator.create_endpoint import  create_endpoints_file
from creator. create_handler import  create_handler_file
from creator. create_mapper import  create_mapper
from creator. create_model import   create_struct_to_file
from creator. create_repository import  create_repository_file
from creator. create_service import  create_service_file
from creator. creator_reqres import  create_model_file
from utils.class_extracker import extract_model_info

model=""" 
class GroupHall(models.Model):
    id = models.IntegerField(unique=True, primary_key=True)
    group = models.ForeignKey(MenuSubGroup, on_delete=CASCADE)
    hall = models.ForeignKey(TradingHall, on_delete=CASCADE)
"""

example = extract_model_info(model)
create_endpoints_file(example['class_name'])
create_model_file(example)
create_handler_file(example['class_name'])
create_repository_file(example['class_name'])
create_service_file(example['class_name'])
create_mapper(example['class_name'],example['fields'])
create_struct_to_file(example)