package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-production-routing-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

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

func (c *SAPAPICaller) Header(productionRoutingGroup, productionRouting string) {
	headerData, err := c.callProductionRoutingSrvAPIRequirementHeader("ProductionRoutingHeader", productionRoutingGroup, productionRouting)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	materialAssignmentData, err := c.callToMaterialAssignment(headerData[0].ToMaterialAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialAssignmentData)

	sequenceData, err := c.callToSequence(headerData[0].ToSequence)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(sequenceData)

	operationData, err := c.callToOperation(sequenceData[0].ToOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationData)

	componentAllocationData, err := c.callToComponentAllocation(operationData[0].ToComponentAllocation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(componentAllocationData)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementHeader(api, productionRoutingGroup, productionRouting string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, productionRoutingGroup, productionRouting)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToMaterialAssignment(url string) ([]sap_api_output_formatter.ToMaterialAssignment, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToSequence(url string) ([]sap_api_output_formatter.ToSequence, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSequence(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToOperation(url string) ([]sap_api_output_formatter.ToOperation, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToOperation(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToComponentAllocation(url string) ([]sap_api_output_formatter.ToComponentAllocation, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToComponentAllocation(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProductPlant(plant, product string) {
	data, err := c.callProductionRoutingSrvAPIRequirementProductPlant("ProductionRoutingMatlAssgmt", plant, product)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

	sequenceData, err := c.callToSequence(
		fmt.Sprintf("%s/API_PRODUCTION_ROUTING/ProductionRoutingHeader(ProductionRoutingGroup='%s',ProductionRouting='%s',ProductionRoutingInternalVers='%s')/to_Sequence",
			c.baseURL, data[0].ProductionRoutingGroup, data[0].ProductionRouting, data[0].ProductionRtgMatlAssgmtIntVers,
		),
	)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(sequenceData)

	operationData, err := c.callToOperation(sequenceData[0].ToOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationData)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementProductPlant(api, plant, product string) ([]sap_api_output_formatter.MaterialAssignment, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")

	param := c.getQueryWithProductPlant(map[string]string{}, plant, product)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BillOfOperationsDesc(plant, billOfOperationsDesc string) {
	billOfOperationsDescdata, err := c.callProductionRoutingSrvAPIRequirementBillOfOperationsDesc("ProductionRoutingHeader", plant, billOfOperationsDesc)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(billOfOperationsDescdata)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementBillOfOperationsDesc(api, plant, billOfOperationsDesc string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")

	param := c.getQueryWithBillOfOperationsDesc(map[string]string{}, plant, billOfOperationsDesc)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SequenceText(sequenceText string) {
	sequenceTextdata, err := c.callProductionRoutingSrvAPIRequirementSequenceText("ProductionRoutingSequence", sequenceText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(sequenceTextdata)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementSequenceText(api, sequenceText string) ([]sap_api_output_formatter.Sequence, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")

	param := c.getQueryWithSequenceText(map[string]string{}, sequenceText)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSequence(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) OperationText(plant, operationText string) {
	operationTextdata, err := c.callProductionRoutingSrvAPIRequirementOperationText("ProductionRoutingOperation", plant, operationText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationTextdata)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementOperationText(api, plant, operationText string) ([]sap_api_output_formatter.Operation, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")

	param := c.getQueryWithOperationText(map[string]string{}, plant, operationText)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToOperation(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, productionRoutingGroup, productionRouting string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ProductionRoutingGroup eq '%s' and ProductionRouting eq '%s'", productionRoutingGroup, productionRouting)
	return params
}

func (c *SAPAPICaller) getQueryWithProductPlant(params map[string]string, plant, product string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Plant eq '%s' and Product eq '%s'", plant, product)
	return params
}

func (c *SAPAPICaller) getQueryWithBillOfOperationsDesc(params map[string]string, plant, billOfOperationsDesc string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Plant eq '%s' and substringof('%s', BillOfOperationsDesc)", plant, billOfOperationsDesc)
	return params
}

func (c *SAPAPICaller) getQueryWithSequenceText(params map[string]string, sequenceText string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("substringof('%s', SequenceText)", sequenceText)
	return params
}

func (c *SAPAPICaller) getQueryWithOperationText(params map[string]string, plant, operationText string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Plant eq '%s' and substringof('%s', OperationText)", plant, operationText)
	return params
}
