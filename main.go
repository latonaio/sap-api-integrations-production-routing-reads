package main

import (
	sap_api_caller "sap-api-integrations-production-routing-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-production-routing-reads/SAP_API_Input_Reader"
	"sap-api-integrations-production-routing-reads/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	gc := sap_api_request_client_header_setup.NewSAPRequestClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		gc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Production_Routing_Product_Plant_sample.json")
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&inputSDC)
	accepter := inputSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "ProductPlant",
			"BillOfOperationsDesc", "SequenceText", "OperationText",
		}
	}

	caller.AsyncGetProductionRouting(
		inputSDC.ProductionRouting.ProductionRoutingGroup,
		inputSDC.ProductionRouting.ProductionRouting,
		inputSDC.ProductionRouting.Plant,
		inputSDC.ProductionRouting.MaterialAssignment.Product,
		inputSDC.ProductionRouting.BillOfOperationsDesc,
		inputSDC.ProductionRouting.Sequence.SequenceText,
		inputSDC.ProductionRouting.Sequence.Operation.OperationText,
		accepter,
	)
}
