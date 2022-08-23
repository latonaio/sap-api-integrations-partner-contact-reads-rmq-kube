package responses

type PartnerContactCollection struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			ObjectID                          string `json:"ObjectID"`
			PartnerContactID                  string `json:"PartnerContactID"`
			PartnerContactUUID                string `json:"PartnerContactUUID"`
			UserID                            string `json:"UserID"`
			IdentityUUID                      string `json:"IdentityUUID"`
			CreateUser                        bool   `json:"CreateUser"`
			BusinessPartnerID                 string `json:"BusinessPartnerID"`
			PartnerID                         string `json:"PartnerID"`
			StatusCode                        string `json:"StatusCode"`
			StatusCodeText                    string `json:"StatusCodeText"`
			TitleCode                         string `json:"TitleCode"`
			TitleCodeText                     string `json:"TitleCodeText"`
			AcademicTitleCode                 string `json:"AcademicTitleCode"`
			AcademicTitleCodeText             string `json:"AcademicTitleCodeText"`
			BusinessPartnerFormattedName      string `json:"BusinessPartnerFormattedName"`
			FirstName                         string `json:"FirstName"`
			LastName                          string `json:"LastName"`
			MiddleName                        string `json:"MiddleName"`
			AdditionalLastName                string `json:"AdditionalLastName"`
			GenderCode                        string `json:"GenderCode"`
			GenderCodeText                    string `json:"GenderCodeText"`
			MaritalStatusCode                 string `json:"MaritalStatusCode"`
			MaritalStatusCodeText             string `json:"MaritalStatusCodeText"`
			LanguageCode                      string `json:"LanguageCode"`
			LanguageCodeText                  string `json:"LanguageCodeText"`
			BirthDate                         string `json:"BirthDate"`
			VIPContactCode                    string `json:"VIPContactCode"`
			VIPContactCodeText                string `json:"VIPContactCodeText"`
			Department                        string `json:"Department"`
			JobTitle                          string `json:"JobTitle"`
			FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
			CountryCode                       string `json:"CountryCode"`
			CountryCodeText                   string `json:"CountryCodeText"`
			StateCode                         string `json:"StateCode"`
			StateCodeText                     string `json:"StateCodeText"`
			AddressLine1                      string `json:"AddressLine1"`
			AddressLine2                      string `json:"AddressLine2"`
			HouseNumber                       string `json:"HouseNumber"`
			Street                            string `json:"Street"`
			AddressLine4                      string `json:"AddressLine4"`
			AddressLine5                      string `json:"AddressLine5"`
			City                              string `json:"City"`
			StreetPostalCode                  string `json:"StreetPostalCode"`
			Phone                             string `json:"Phone"`
			Mobile                            string `json:"Mobile"`
			Fax                               string `json:"Fax"`
			Email                             string `json:"Email"`
			BestReachedByCode                 string `json:"BestReachedByCode"`
			BestReachedByCodeText             string `json:"BestReachedByCodeText"`
			NormalisedPhone                   string `json:"NormalisedPhone"`
			NormalisedMobile                  string `json:"NormalisedMobile"`
			CreatedOn                         string `json:"CreatedOn"`
			CreatedBy                         string `json:"CreatedBy"`
			CreatedByIdentityUUID             string `json:"CreatedByIdentityUUID"`
			ChangedOn                         string `json:"ChangedOn"`
			ChangedBy                         string `json:"ChangedBy"`
			ChangedByIdentityUUID             string `json:"ChangedByIdentityUUID"`
			EntityLastChangedOn               string `json:"EntityLastChangedOn"`
			ETag                              string `json:"ETag"`
			Deferred                          struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"results"`
	} `json:"d"`
}
