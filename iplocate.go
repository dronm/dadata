package dadata

type IPLocateData struct {
	Postal_code  *string `json:"postal_code"`
	Country *string `json:"country"`
	Country_iso_code *string `json:"country_iso_code"`
	Federal_district *string `json:"federal_district"`
	Region_fias_id *string `json:"region_fias_id"`
	Region_kladr_id *string `json:"region_kladr_id"`
	Region_iso_code *string `json:"region_iso_code"`
	Region_with_type *string `json:"region_with_type"`
	Region_type *string `json:"region_type"`
	Region_type_full *string `json:"region_type_full"`
	Region *string `json:"region"`
	Area_fias_id *string `json:"area_fias_id"`
	Area_kladr_id *string `json:"area_kladr_id"`
	Area_with_type *string `json:"area_with_type"`
	Area_type *string `json:"area_type"`
	Area_type_full *string `json:"area_type_full"`
	Crea *string `json:"area"`
	City_fias_id *string `json:"city_fias_id"`
	City_kladr_id *string `json:"city_kladr_id"`
	City_with_type *string `json:"city_with_type"`
	City_type *string `json:"city_type"`
	City_type_full *string `json:"city_type_full"`
	City *string `json:"city"`
	City_area *string `json:"city_area"`
	City_district_fias_id *string `json:"city_district_fias_id"`
	City_district_kladr_id *string `json:"city_district_kladr_id"`
	City_district_with_type *string `json:"city_district_with_type"`
	City_district_type *string `json:"city_district_type"`
	City_district_type_full *string `json:"city_district_type_full"`
	City_district *string `json:"city_district"`
	Settlement_fias_id *string `json:"settlement_fias_id"`
	Settlement_kladr_id *string `json:"settlement_kladr_id"`
	Settlement_with_type *string `json:"settlement_with_type"`
	Settlement_type *string `json:"settlement_type"`
	Settlement_type_full *string `json:"settlement_type_full"`
	Settlement *string `json:"settlement"`
	Street_fias_id *string `json:"street_fias_id"`
	Street_kladr_id *string `json:"street_kladr_id"`
	Street_with_type *string `json:"street_with_type"`
	Street_type *string `json:"street_type"`
	Street_type_full *string `json:"street_type_full"`
	Street *string `json:"street"`
	Stead_fias_id *string `json:"stead_fias_id"`
	Stead_cadnum *string `json:"stead_cadnum"`
	Stead_type *string `json:"stead_type"`
	Stead_type_full *string `json:"stead_type_full"`
	Stead *string `json:"stead"`
	House_fias_id *string `json:"house_fias_id"`
	House_kladr_id *string `json:"house_kladr_id"`
	House_cadnum *string `json:"house_cadnum"`
	House_type *string `json:"house_type"`
	House_type_full *string `json:"house_type_full"`
	House *string `json:"house"`
	Block_type *string `json:"block_type"`
	Block_type_full *string `json:"block_type_full"`
	Block *string `json:"block"`
	Entrance *string `json:"entrance"`
	Floor *string `json:"floor"`
	Flat_fias_id *string `json:"flat_fias_id"`
	Flat_cadnum *string `json:"flat_cadnum"`
	Flat_type *string `json:"flat_type"`
	Flat_type_full *string `json:"flat_type_full"`
	Flat *string `json:"flat"`
	Flat_area *string `json:"flat_area"`
	Square_meter_price *string `json:"square_meter_price"`
	Flat_price *string `json:"flat_price"`	
	Postal_box *string `json:"postal_box"`
	Fias_id *string `json:"fias_id"`
	Fias_code *string `json:"fias_code"`
	Fias_level *string `json:"fias_level"`
	Fias_actuality_state *string `json:"fias_actuality_state"`
	Kladr_id *string `json:"kladr_id"`
	Geoname_id *string `json:"geoname_id"`
	Capital_marker *string `json:"capital_marker"`
	Okato *string `json:"okato"`
	Oktmo *string `json:"oktmo"`
	Tax_office *string `json:"tax_office"`
	Tax_office_legal *string `json:"tax_office_legal"`
	Timezone *string `json:"timezone"`
	Geo_lat *string `json:"geo_lat"`
	Geo_lon *string `json:"geo_lon"`
	Beltway_hit *string `json:"beltway_hit"`
	Beltway_distance *string `json:"beltway_distance"`
	Metro *string `json:"metro"`
	Divisions *string `json:"divisions"`
	Qc_geo *string `json:"qc_geo"`
	Qc_complete *string `json:"qc_complete"`
	Qc_house *string `json:"qc_house"`
	History_values *string `json:"history_values"`
	Unparsed_parts *string `json:"unparsed_parts"`
	Source *string `json:"source"`
	Qc *string `json:"qc"`	
}

type IPLocateLocation struct {
	Value string `json:"value"`
	Unrestricted_value string `json:"unrestricted_value"`
	Data *IPLocateData `json:"data"`
}

type IPLocateResponse struct {
	Location *IPLocateLocation `json:"location"`
}


