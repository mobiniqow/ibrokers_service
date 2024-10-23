from creator.create_endpoint import  create_endpoints_file
from creator. create_handler import  create_handler_file
from creator. create_mapper import  create_mapper
from creator. create_model import   create_struct_to_file
from creator. create_repository import  create_repository_file
from creator. create_service import  create_service_file
from creator. creator_reqres import  create_model_file
from utils.class_extracker import extract_model_info

model=""" 
class Offer(models.Model):
    buyMethodId = models.ForeignKey(BuyMethod, on_delete=models.DO_NOTHING)
    brokerId = models.ForeignKey(Broker, on_delete=models.DO_NOTHING)
    commodityId = models.ForeignKey(Commodity, on_delete=models.DO_NOTHING)
    contractTypeId = models.ForeignKey(ContractType, on_delete=models.DO_NOTHING)
    currencyId = models.ForeignKey(CurrencyUnit, on_delete=models.DO_NOTHING)
    deliveryPlaceId = models.ForeignKey(DeliveryPlace, on_delete=models.DO_NOTHING)
    initPrice = models.BigIntegerField(null=True)
    initVolume = models.CharField(default='0', null=True, blank=True, max_length=122)
    lotSize = models.BigIntegerField(null=True)
    manufacturerId = models.ForeignKey(Manufacturers, on_delete=models.DO_NOTHING)
    maxInitPrice = models.BigIntegerField(null=True)
    maxIncOfferVol = models.BigIntegerField(null=True)
    maxOrderVol = models.BigIntegerField(null=True)
    maxOfferPrice = models.BigIntegerField(null=True)
    measureUnitId = models.ForeignKey(MeasureUnit, on_delete=models.DO_NOTHING)
    minAllocationVol = models.BigIntegerField(null=True)
    minOfferVol = models.BigIntegerField(null=True)
    minInitPrice = models.BigIntegerField(null=True)
    minOrderVol = models.BigIntegerField(null=True)
    minOfferPrice = models.BigIntegerField(null=True)
    offerModeId = models.ForeignKey(OfferMod, on_delete=models.DO_NOTHING)
    offerTypeId = models.ForeignKey(OfferType, on_delete=models.DO_NOTHING)
    offerVol = models.BigIntegerField(null=True)
    packagingTypeId = models.ForeignKey(PackagingType, on_delete=models.DO_NOTHING)
    permissibleError = models.BigIntegerField(null=True)
    priceDiscoveryMinOrderVol = models.BigIntegerField(null=True)
    prepaymentPercent = models.BigIntegerField(null=True)
    securityTypeId = models.BigIntegerField(null=True)
    settlementTypeId = models.ForeignKey(Settlement, on_delete=models.DO_NOTHING)
    supplierId = models.ForeignKey(Supplier, on_delete=models.DO_NOTHING)
    tickSize = models.BigIntegerField(null=True)
    tradingHallId = models.ForeignKey(TradingHall, on_delete=models.DO_NOTHING)
    weightFactor = models.BigIntegerField(null=True)
    id = models.PositiveBigIntegerField(primary_key=True)
    deliveryDate = jmodels.jDateField(null=True)
    description = models.TextField(null=True)
    offerDate = jmodels.jDateField(null=True)
    offerRing = models.TextField(null=True)
    offerSymbol = models.TextField(null=True)
    securityTypeNote = models.CharField(
        null=True, blank=True, max_length=90
    )
    tradeStatus = models.TextField(null=True)
"""

example = extract_model_info(model)
create_endpoints_file(example['class_name'])
create_model_file(example)
create_handler_file(example['class_name'])
create_repository_file(example['class_name'])
create_service_file(example['class_name'])
create_mapper(example['class_name'],example['fields'])
create_struct_to_file(example)