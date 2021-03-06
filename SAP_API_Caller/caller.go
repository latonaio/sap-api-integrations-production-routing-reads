package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-production-routing-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
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
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, productionRoutingGroup, productionRouting)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToMaterialAssignment(url string) ([]sap_api_output_formatter.ToMaterialAssignment, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToSequence(url string) ([]sap_api_output_formatter.ToSequence, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSequence(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToOperation(url string) ([]sap_api_output_formatter.ToOperation, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToComponentAllocation(url string) ([]sap_api_output_formatter.ToComponentAllocation, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToComponentAllocation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
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
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithProductPlant(req, plant, product)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BillOfOperationsDesc(plant, billOfOperationsDesc string) {
	data, err := c.callProductionRoutingSrvAPIRequirementBillOfOperationsDesc("ProductionRoutingHeader", plant, billOfOperationsDesc)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementBillOfOperationsDesc(api, plant, billOfOperationsDesc string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBillOfOperationsDesc(req, plant, billOfOperationsDesc)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SequenceText(sequenceText string) {
	data, err := c.callProductionRoutingSrvAPIRequirementSequenceText("ProductionRoutingSequence", sequenceText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementSequenceText(api, sequenceText string) ([]sap_api_output_formatter.Sequence, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSequenceText(req, sequenceText)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSequence(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) OperationText(plant, operationText string) {
	data, err := c.callProductionRoutingSrvAPIRequirementOperationText("ProductionRoutingOperation", plant, operationText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementOperationText(api, plant, operationText string) ([]sap_api_output_formatter.Operation, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithOperationText(req, plant, operationText)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, productionRoutingGroup, productionRouting string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ProductionRoutingGroup eq '%s' and ProductionRouting eq '%s'", productionRoutingGroup, productionRouting))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithProductPlant(req *http.Request, plant, product string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and Product eq '%s'", plant, product))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithBillOfOperationsDesc(req *http.Request, plant, billOfOperationsDesc string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and substringof('%s', BillOfOperationsDesc)", plant, billOfOperationsDesc))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSequenceText(req *http.Request, sequenceText string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("substringof('%s', SequenceText)", sequenceText))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithOperationText(req *http.Request, plant, operationText string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Plant eq '%s' and substringof('%s', OperationText)", plant, operationText))
	req.URL.RawQuery = params.Encode()
}
