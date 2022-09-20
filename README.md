# sap-api-integrations-production-routing-reads  
sap-api-integrations-production-routing-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で作業手順データを取得するマイクロサービスです。  
sap-api-integrations-production-routing-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-production-routing-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PRODUCTION_ROUTING_0001/overview  

## 動作環境

sap-api-integrations-production-routing-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）  
・ AION のリソース （推奨)  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用

sap-api-integrations-production-routing-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-production-routing-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PRODUCTION_ROUTING_0001/overview  
* APIサービス名(=baseURL): API_PRODUCTION_ROUTING

## 本レポジトリ に 含まれる API名
sap-api-integrations-production-routing-reads には、次の API をコールするためのリソースが含まれています。  

* ProductionRoutingHeader（作業手順 - ヘッダ）※作業手順の詳細データを取得するために、ToMaterialAssignment、ToSequence、ToOperation、ToComponentAllocation、と合わせて利用されます。
* ToMaterialAssignment（作業手順 - 品目 ※To）
* ToSequence（作業手順 - 順序 ※To）
* ToOperation（作業手順 - 作業 ※To）
* ToComponentAllocation（作業手順 - 構成品目割当 ※To）
* ProductionRoutingMatlAssgmt（作業手順 - 品目）※作業手順の詳細データを取得するために、ToSequence、ToOperation、と合わせて利用されます。
* ToSequence（作業手順 - 順序 ※To）
* ToOperation（作業手順 - 作業 ※To）
* ToComponentAllocation（作業手順 - 構成品目割当 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-production-routing-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ProductionRouting.ProductionRoutingGroup（作業手順グループ）
* inoutSDC.ProductionRouting.ProductionRouting（作業手順）
* inoutSDC.ProductionRouting.Plant（プラント）
* inoutSDC.ProductionRouting.MaterialAssignment.Product（品目）
* inoutSDC.ProductionRouting.BillOfOperationsDesc（作業手順説明）
* inoutSDC.ProductionRouting.Sequence.SequenceText（順序テキスト）
* inoutSDC.ProductionRouting.Sequence.Operation.OperationText（作業テキスト）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ProductPlant" が指定されています。    
  
```
	"api_schema": "ProductionRoutingReads",
	"accepter": ["Header"],
	"material_code": "",
	"plant": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "ProductionRoutingReads",
	"accepter": ["All"],
	"material_code": "",
	"plant": "",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetProductionRouting(productionRoutingGroup, productionRouting, plant, product, billOfOperationsDesc, sequenceText, operationText string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(productionRoutingGroup, productionRouting)
				wg.Done()
			}()
		case "ProductPlant":
			func() {
				c.ProductPlant(plant, product)
				wg.Done()
			}()
		case "BillOfOperationsDesc":
			func() {
				c.BillOfOperationsDesc(plant, billOfOperationsDesc)
				wg.Done()
			}()
		case "SequenceText":
			func() {
				c.SequenceText(sequenceText)
				wg.Done()
			}()
		case "OperationText":
			func() {
				c.OperationText(plant, operationText)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。   
以下の sample.json の例は、SAP 作業手順 ヘッダ が取得された結果の JSON の例です。  
以下の項目のうち、"ProductionRoutingGroup" ～ "ProductionRoutingInternalVers" は、/SAP_API_Output_Formatter/type.go 内 の type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。    

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-production-routing-reads/SAP_API_Caller/caller.go#L73",
	"function": "sap-api-integrations-production-routing-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"ProductionRoutingGroup": "40000060",
			"ProductionRouting": "1",
			"ProductionRoutingInternalVers": "1",
			"IsMarkedForDeletion": false,
			"BillOfOperationsDesc": "MTS - ELECTRIC FAN",
			"Plant": "1010",
			"BillOfOperationsUsage": "1",
			"BillOfOperationsStatus": "4",
			"ResponsiblePlannerGroup": "",
			"MinimumLotSizeQuantity": "1",
			"MaximumLotSizeQuantity": "99999999",
			"BillOfOperationsUnit": "PC",
			"CreationDate": "2021-02-01T09:00:00+09:00",
			"CreatedByUser": "SAP_SYSTEM",
			"LastChangeDate": "",
			"ValidityStartDate": "2021-02-01T09:00:00+09:00",
			"ValidityEndDate": "9999-12-31T09:00:00+09:00",
			"ChangeNumber": "",
			"PlainLongText": "",
			"to_MatlAssgmt": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCTION_ROUTING/ProductionRoutingHeader(ProductionRoutingGroup='40000060',ProductionRouting='1',ProductionRoutingInternalVers='1')/to_MatlAssgmt",
			"to_Sequence": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCTION_ROUTING/ProductionRoutingHeader(ProductionRoutingGroup='40000060',ProductionRouting='1',ProductionRoutingInternalVers='1')/to_Sequence"
		}
	],
	"time": "2022-01-28T14:08:56+09:00"
}
```

