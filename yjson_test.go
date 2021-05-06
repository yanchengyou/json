package json

import (
	"fmt"
	"testing"
)

func TestParseToMap(t *testing.T) {
	a := `{
	   "p": {},
	   "test": "111123asd",
	   "a": 23,
	   "b": [{
		   "id": "20210403564",
		   "inquiryId": "20210401174",
		   "$.inquiryno": "20210401418_01",
		   "$.inquiryname": "20210423鸭蛋牛肉_01",
		   "purOrgId": "101118017",
		   "purOrgName": "炼化司务处",
		   "$.purmethodid": "2021",
		   "$.quotationmethod": "onlinequotation",
		   "$.dealmethod": "itemizeddeal",
		   "$.contracttype": "purchaseorder",
		   "$.quotationstarttime": "2021-04-23 17:30:51",
		   "$.quotationendtime": "2021-04-23 22:00:55",
		   "$.quotationrevealtime": "2021-04-23 22:01:55",
		   "$.purusername": "于景伟",
		   "$.puruserphone": "",
		   "$.purusermobile": "",
		   "$.puruseremail": "",
		   "vendorId": "20200700465",
		   "vendorName": "东营鑫昊学堂蔬菜批发摊",
		   "durationDesc": "1天",
		   "payMode": "",
		   "vendorUserName": "闫学堂",
		   "vendorUserPhone": "",
		   "vendorUserFax": "",
		   "vendorUserEmail": "251965388@qq.com",
		   "version": "0"
	   }],
	   "bObjectCount": 1,
	   "b1": [{
		   "applierName": "于景伟",
		   "applierOrgName": "司务处",
		   "deliveryTime": "2021-04-24",
		   "ext$": {
			   "usedbyorgname": "司务处",
			   "mtldescription": "鸭蛋  ",
			   "mtlitemid": "37623195",
			   "mtlitemcode": "000005031701000013",
			   "planid": "20210401112",
			   "projectid": "20200700008",
			   "pricecap": "0",
			   "projectname": "利华益炼化有限公司-司务处",
			   "rn_": "1",
			   "planname": "20210423鸭蛋",
			   "mtlunitid": "jin",
			   "usedbyorgid": "101118017",
			   "projectno": "SWC_LH",
			   "planno": "20210401112",
			   "latestbargainprice": "",
			   "matchagreement": "no"
		   },
		   "freightCharge": "0.00",
		   "goodsAmountTaxExcl": 1200,
		   "goodsAmountTaxIncl": 1200,
		   "goodsPriceTaxExcl": 7.5,
		   "goodsPriceTaxIncl": "7.5",
		   "handlingCharge": "0.00",
		   "id": 42963206,
		   "inquiryLineId": 42961860,
		   "inquiryQty": 160,
		   "inquiryQuotationId": "20210403564",
		   "planLineId": 42959107,
		   "readonly": "0",
		   "schemeLineId": 42960162,
		   "taxAmount": "0",
		   "taxRate": "0.00",
		   "totalAmount": 1200,
		   "brandModel": "鸭蛋",
		   "remark": ""
	   }, {
		   "applierName": "于景伟",
		   "applierOrgName": "司务处",
		   "deliveryTime": "2021-04-24",
		   "ext$": {
			   "usedbyorgname": "司务处",
			   "mtldescription": "熟牛肉  ",
			   "mtlitemid": "37623229",
			   "mtlitemcode": "000005031701000028",
			   "planid": "20210401112",
			   "projectid": "20200700008",
			   "pricecap": "0",
			   "projectname": "利华益炼化有限公司-司务处",
			   "rn_": "2",
			   "planname": "20210423鸭蛋",
			   "mtlunitid": "jin",
			   "usedbyorgid": "101118017",
			   "projectno": "SWC_LH",
			   "planno": "20210401112",
			   "latestbargainprice": "",
			   "matchagreement": "no"
		   },
		   "freightCharge": "0.00",
		   "goodsAmountTaxExcl": 2580,
		   "goodsAmountTaxIncl": 2580,
		   "goodsPriceTaxExcl": 43,
		   "goodsPriceTaxIncl": "43",
		   "handlingCharge": "0.00",
		   "id": 42963207,
		   "inquiryLineId": 42961861,
		   "inquiryQty": 60,
		   "inquiryQuotationId": "20210403564",
		   "planLineId": 42959054,
		   "readonly": "0",
		   "schemeLineId": 42960163,
		   "taxAmount": "0",
		   "taxRate": "0.00",
		   "totalAmount": 2580,
		   "brandModel": "熟牛肉",
		   "remark": ""
	   }]
	}`
	x:= Parse(a).Get("test")
	fmt.Printf("%v\n%+v\n",x.Type,x.Raw)
	fmt.Println(x.Type)

}
