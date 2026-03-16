package dadata

type AddressSuggestionsResponse struct {
	Suggestions []AddressSuggestion `json:"suggestions"`
}

type AddressSuggestion struct {
	Value             string                `json:"value"`
	UnrestrictedValue string                `json:"unrestricted_value"`
	Data              AddressSuggestionData `json:"data"`
}

type AddressSuggestionData struct {
	PostalCode      *string `json:"postal_code"`
	Country         *string `json:"country"`
	CountryISOCode  *string `json:"country_iso_code"`
	FederalDistrict *string `json:"federal_district"`
	RegionFiasID    *string `json:"region_fias_id"`
	RegionKladrID   *string `json:"region_kladr_id"`
	RegionISOCode   *string `json:"region_iso_code"`
	RegionWithType  *string `json:"region_with_type"`
	RegionType      *string `json:"region_type"`
	RegionTypeFull  *string `json:"region_type_full"`
	Region          *string `json:"region"`

	AreaFiasID   *string `json:"area_fias_id"`
	AreaKladrID  *string `json:"area_kladr_id"`
	AreaWithType *string `json:"area_with_type"`
	AreaType     *string `json:"area_type"`
	AreaTypeFull *string `json:"area_type_full"`
	Area         *string `json:"area"`

	CityFiasID   *string `json:"city_fias_id"`
	CityKladrID  *string `json:"city_kladr_id"`
	CityWithType *string `json:"city_with_type"`
	CityType     *string `json:"city_type"`
	CityTypeFull *string `json:"city_type_full"`
	City         *string `json:"city"`
	CityArea     *string `json:"city_area"`

	CityDistrictFiasID   *string `json:"city_district_fias_id"`
	CityDistrictKladrID  *string `json:"city_district_kladr_id"`
	CityDistrictWithType *string `json:"city_district_with_type"`
	CityDistrictType     *string `json:"city_district_type"`
	CityDistrictTypeFull *string `json:"city_district_type_full"`
	CityDistrict         *string `json:"city_district"`

	SettlementFiasID   *string `json:"settlement_fias_id"`
	SettlementKladrID  *string `json:"settlement_kladr_id"`
	SettlementWithType *string `json:"settlement_with_type"`
	SettlementType     *string `json:"settlement_type"`
	SettlementTypeFull *string `json:"settlement_type_full"`
	Settlement         *string `json:"settlement"`

	StreetFiasID   *string `json:"street_fias_id"`
	StreetKladrID  *string `json:"street_kladr_id"`
	StreetWithType *string `json:"street_with_type"`
	StreetType     *string `json:"street_type"`
	StreetTypeFull *string `json:"street_type_full"`
	Street         *string `json:"street"`

	SteadFiasID   *string `json:"stead_fias_id"`
	SteadCadnum   *string `json:"stead_cadnum"`
	SteadType     *string `json:"stead_type"`
	SteadTypeFull *string `json:"stead_type_full"`
	Stead         *string `json:"stead"`

	HouseFiasID    *string `json:"house_fias_id"`
	HouseKladrID   *string `json:"house_kladr_id"`
	HouseCadnum    *string `json:"house_cadnum"`
	HouseFlatCount *int    `json:"house_flat_count"`
	HouseType      *string `json:"house_type"`
	HouseTypeFull  *string `json:"house_type_full"`
	House          *string `json:"house"`

	BlockType     *string `json:"block_type"`
	BlockTypeFull *string `json:"block_type_full"`
	Block         *string `json:"block"`

	Entrance *string `json:"entrance"`
	Floor    *string `json:"floor"`

	FlatFiasID       *string `json:"flat_fias_id"`
	FlatCadnum       *string `json:"flat_cadnum"`
	FlatType         *string `json:"flat_type"`
	FlatTypeFull     *string `json:"flat_type_full"`
	Flat             *string `json:"flat"`
	FlatArea         *string `json:"flat_area"`
	SquareMeterPrice *string `json:"square_meter_price"`
	FlatPrice        *string `json:"flat_price"`

	PostalBox *string `json:"postal_box"`

	FiasID             *string     `json:"fias_id"`
	FiasCode           *string     `json:"fias_code"`
	FiasLevel          *string     `json:"fias_level"`
	FiasActualityState *string     `json:"fias_actuality_state"`
	KladrID            *string     `json:"kladr_id"`
	GeonameID          *string     `json:"geoname_id"`
	CapitalMarker      *string     `json:"capital_marker"`
	Okato              *string     `json:"okato"`
	Oktmo              *string     `json:"oktmo"`
	TaxOffice          *string     `json:"tax_office"`
	TaxOfficeLegal     *string     `json:"tax_office_legal"`
	Timezone           *string     `json:"timezone"`
	GeoLat             *string     `json:"geo_lat"`
	GeoLon             *string     `json:"geo_lon"`
	BeltwayHit         *string     `json:"beltway_hit"`
	BeltwayDistance    *string     `json:"beltway_distance"`
	Metro              *string     `json:"metro"`
	Divisions          interface{} `json:"divisions"`
	QcGeo              *string     `json:"qc_geo"`
	QcComplete         *string     `json:"qc_complete"`
	QcHouse            *string     `json:"qc_house"`
	HistoryValues      []string    `json:"history_values"`
	UnparsedParts      *string     `json:"unparsed_parts"`
	Source             *string     `json:"source"`
	Qc                 *string     `json:"qc"`
}
