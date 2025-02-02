# InlineResponse200129

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number of the sensor that took the reading. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNetwork**](OrganizationsOrganizationIdSensorReadingsHistoryNetwork.md) |  | [optional] 
**Ts** | Pointer to **string** | Time at which the reading occurred, in ISO8601 format. | [optional] 
**Metric** | Pointer to **string** | Type of sensor reading. | [optional] 
**Battery** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryBattery**](OrganizationsOrganizationIdSensorReadingsHistoryBattery.md) |  | [optional] 
**Button** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryButton**](OrganizationsOrganizationIdSensorReadingsHistoryButton.md) |  | [optional] 
**Door** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryDoor**](OrganizationsOrganizationIdSensorReadingsHistoryDoor.md) |  | [optional] 
**Humidity** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryHumidity**](OrganizationsOrganizationIdSensorReadingsHistoryHumidity.md) |  | [optional] 
**IndoorAirQuality** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality**](OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality.md) |  | [optional] 
**Noise** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNoise**](OrganizationsOrganizationIdSensorReadingsHistoryNoise.md) |  | [optional] 
**Pm25** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryPm25**](OrganizationsOrganizationIdSensorReadingsHistoryPm25.md) |  | [optional] 
**Temperature** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryTemperature**](OrganizationsOrganizationIdSensorReadingsHistoryTemperature.md) |  | [optional] 
**Tvoc** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryTvoc**](OrganizationsOrganizationIdSensorReadingsHistoryTvoc.md) |  | [optional] 
**Water** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryWater**](OrganizationsOrganizationIdSensorReadingsHistoryWater.md) |  | [optional] 

## Methods

### NewInlineResponse200129

`func NewInlineResponse200129() *InlineResponse200129`

NewInlineResponse200129 instantiates a new InlineResponse200129 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200129WithDefaults

`func NewInlineResponse200129WithDefaults() *InlineResponse200129`

NewInlineResponse200129WithDefaults instantiates a new InlineResponse200129 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200129) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200129) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200129) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200129) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200129) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200129) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200129) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200129) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTs

`func (o *InlineResponse200129) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200129) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200129) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200129) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMetric

`func (o *InlineResponse200129) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *InlineResponse200129) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *InlineResponse200129) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *InlineResponse200129) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetBattery

`func (o *InlineResponse200129) GetBattery() OrganizationsOrganizationIdSensorReadingsHistoryBattery`

GetBattery returns the Battery field if non-nil, zero value otherwise.

### GetBatteryOk

`func (o *InlineResponse200129) GetBatteryOk() (*OrganizationsOrganizationIdSensorReadingsHistoryBattery, bool)`

GetBatteryOk returns a tuple with the Battery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBattery

`func (o *InlineResponse200129) SetBattery(v OrganizationsOrganizationIdSensorReadingsHistoryBattery)`

SetBattery sets Battery field to given value.

### HasBattery

`func (o *InlineResponse200129) HasBattery() bool`

HasBattery returns a boolean if a field has been set.

### GetButton

`func (o *InlineResponse200129) GetButton() OrganizationsOrganizationIdSensorReadingsHistoryButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *InlineResponse200129) GetButtonOk() (*OrganizationsOrganizationIdSensorReadingsHistoryButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *InlineResponse200129) SetButton(v OrganizationsOrganizationIdSensorReadingsHistoryButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *InlineResponse200129) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetDoor

`func (o *InlineResponse200129) GetDoor() OrganizationsOrganizationIdSensorReadingsHistoryDoor`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *InlineResponse200129) GetDoorOk() (*OrganizationsOrganizationIdSensorReadingsHistoryDoor, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *InlineResponse200129) SetDoor(v OrganizationsOrganizationIdSensorReadingsHistoryDoor)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *InlineResponse200129) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetHumidity

`func (o *InlineResponse200129) GetHumidity() OrganizationsOrganizationIdSensorReadingsHistoryHumidity`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *InlineResponse200129) GetHumidityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryHumidity, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *InlineResponse200129) SetHumidity(v OrganizationsOrganizationIdSensorReadingsHistoryHumidity)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *InlineResponse200129) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIndoorAirQuality

`func (o *InlineResponse200129) GetIndoorAirQuality() OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality`

GetIndoorAirQuality returns the IndoorAirQuality field if non-nil, zero value otherwise.

### GetIndoorAirQualityOk

`func (o *InlineResponse200129) GetIndoorAirQualityOk() (*OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality, bool)`

GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndoorAirQuality

`func (o *InlineResponse200129) SetIndoorAirQuality(v OrganizationsOrganizationIdSensorReadingsHistoryIndoorAirQuality)`

SetIndoorAirQuality sets IndoorAirQuality field to given value.

### HasIndoorAirQuality

`func (o *InlineResponse200129) HasIndoorAirQuality() bool`

HasIndoorAirQuality returns a boolean if a field has been set.

### GetNoise

`func (o *InlineResponse200129) GetNoise() OrganizationsOrganizationIdSensorReadingsHistoryNoise`

GetNoise returns the Noise field if non-nil, zero value otherwise.

### GetNoiseOk

`func (o *InlineResponse200129) GetNoiseOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNoise, bool)`

GetNoiseOk returns a tuple with the Noise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoise

`func (o *InlineResponse200129) SetNoise(v OrganizationsOrganizationIdSensorReadingsHistoryNoise)`

SetNoise sets Noise field to given value.

### HasNoise

`func (o *InlineResponse200129) HasNoise() bool`

HasNoise returns a boolean if a field has been set.

### GetPm25

`func (o *InlineResponse200129) GetPm25() OrganizationsOrganizationIdSensorReadingsHistoryPm25`

GetPm25 returns the Pm25 field if non-nil, zero value otherwise.

### GetPm25Ok

`func (o *InlineResponse200129) GetPm25Ok() (*OrganizationsOrganizationIdSensorReadingsHistoryPm25, bool)`

GetPm25Ok returns a tuple with the Pm25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm25

`func (o *InlineResponse200129) SetPm25(v OrganizationsOrganizationIdSensorReadingsHistoryPm25)`

SetPm25 sets Pm25 field to given value.

### HasPm25

`func (o *InlineResponse200129) HasPm25() bool`

HasPm25 returns a boolean if a field has been set.

### GetTemperature

`func (o *InlineResponse200129) GetTemperature() OrganizationsOrganizationIdSensorReadingsHistoryTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *InlineResponse200129) GetTemperatureOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *InlineResponse200129) SetTemperature(v OrganizationsOrganizationIdSensorReadingsHistoryTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *InlineResponse200129) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTvoc

`func (o *InlineResponse200129) GetTvoc() OrganizationsOrganizationIdSensorReadingsHistoryTvoc`

GetTvoc returns the Tvoc field if non-nil, zero value otherwise.

### GetTvocOk

`func (o *InlineResponse200129) GetTvocOk() (*OrganizationsOrganizationIdSensorReadingsHistoryTvoc, bool)`

GetTvocOk returns a tuple with the Tvoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvoc

`func (o *InlineResponse200129) SetTvoc(v OrganizationsOrganizationIdSensorReadingsHistoryTvoc)`

SetTvoc sets Tvoc field to given value.

### HasTvoc

`func (o *InlineResponse200129) HasTvoc() bool`

HasTvoc returns a boolean if a field has been set.

### GetWater

`func (o *InlineResponse200129) GetWater() OrganizationsOrganizationIdSensorReadingsHistoryWater`

GetWater returns the Water field if non-nil, zero value otherwise.

### GetWaterOk

`func (o *InlineResponse200129) GetWaterOk() (*OrganizationsOrganizationIdSensorReadingsHistoryWater, bool)`

GetWaterOk returns a tuple with the Water field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWater

`func (o *InlineResponse200129) SetWater(v OrganizationsOrganizationIdSensorReadingsHistoryWater)`

SetWater sets Water field to given value.

### HasWater

`func (o *InlineResponse200129) HasWater() bool`

HasWater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


