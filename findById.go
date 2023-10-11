package dadata

type SuggestionDataFIO struct {
	Surname *string `json:"surname"`
	Name *string `json:"name"`
	Patronymic *string `json:"patronymic"`
}

type SuggestionDataOPF struct {
	Code *string `json:"code"` // код ОКОПФ
	Full *string `json:"full"` // полное название ОПФ
	Short *string `json:"short"` // краткое название ОПФ
	Type *string `json:"type"` // версия справочника ОКОПФ (99, 2012 или 2014)
}

type SuggestionDataManagement struct {
	Name *string `json:"name"`
	Post *string `json:"post"` //должность руководителя
	Disqualified *string `json:"disqualified"`
}

type SuggestionDataAddressData struct {
	Source *string `json:"source"` //адрес одной строкой как в ЕГРЮЛ
	Qc *string  `json:"qc"` //код проверки адреса 0        — адрес распознан уверенно, 1 или 3  — требуется ручная проверка
}

type SuggestionDataAddressState struct {
	Actuality_date int64 `json:"Actuality_date"` // дата последних изменений
	Registration_date int64 `json:"registration_date"` // дата регистрации
	Liquidation_date int64 `json:"liquidation_date"` // дата ликвидации
	Status *string `json:"status"` //статус организации 
	/*
		ACTIVE       — действующая
		LIQUIDATING  — ликвидируется
		LIQUIDATED   — ликвидирована
		BANKRUPT     — банкротство
		REORGANIZING — в процессе присоединения к другому юрлицу, с последующей ликвидацией	
	*/
	Code *string `json:"code"` // детальный статус
	
}

type SuggestionDataAddress struct {
	Value *string `json:"value"` //адрес одной строкой
	Unrestricted_value *string `json:"unrestricted_value"` //адрес одной строкой (полный, с индексом) стандартизован, поэтому может отличаться от записанного в ЕГРЮЛ
	Data *SuggestionDataAddressData `json:"data"` //гранулярный адрес
	State *SuggestionDataAddressState `json:"state"` //Состояние
}

type SuggestionDataName struct {
	Full_with_opf *string `json:"full_with_opf"`
	Short_with_opf *string `json:"short_with_opf"`
	Full *string `json:"full"`
	Short *string `json:"short"`
}

type Email struct {
	Source *string `json:"source"`
	Local *string `json:"local"`
	Domain *string `json:"domain"`
}

type Phone struct {
	Source *string `json:"source"`
	Type *string `json:"type"`
	Country_code *string `json:"country_code"`
	Number *string `json:"number"`
	Provider *string `json:"provider"`
	Region *string `json:"region"`
	City *string `json:"city"`
	Timezone *string `json:"timezone"`
}

type SuggestionData struct {
	Inn *string `json:"inn"`
	Kpp *string `json:"kpp"`
	Ogrn *string `json:"ogrn"`
	Ogrn_date int64 `json:"ogrn_date"`
	Type *string `json:"type"`
	Name *SuggestionDataName `json:"name"`
	Fio *SuggestionDataFIO `json:"fio"`
	
	Okato *string `json:"okato"`// Код ОКАТО
	Oktmo *string `json:"oktmo"`// Код ОКТМО
	Okpo *string `json:"okpo"`// Код ОКПО
	Okogu *string `json:"okogu"`// Код ОКОГУ
	Okfs *string `json:"okfs"`// Код ОКФС
	Okved *string `json:"okved"`// Код ОКВЭД
	Okved_type *string `json:"okved_type"`// Версия справочника ОКВЭД (2001 или 2014)
	
	Opf *SuggestionDataOPF `json:"opf"`	
	Management *SuggestionDataManagement `json:"management"`
	Branch_count int `json:"branch_count"` // Количество филиалов
	
	Address *SuggestionDataAddress `json:"address"`	
	
	Capital *string `json:"capital"`
	
	//+расширенный тариф
	Phones []*Phone `json:"phones"`
	Emails []*Email `json:"emails"`
}

type Suggestion struct {
	Value *string `json:"value"`
	Data *SuggestionData `json:"data"`
}

type FindByIdResponse struct {
	Suggestions []*Suggestion `json:"suggestions"`
}

