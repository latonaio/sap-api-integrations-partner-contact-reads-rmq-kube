package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-partner-contact-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToPartnerContactCollection(raw []byte, l *logger.Logger) ([]PartnerContactCollection, error) {
	pm := &responses.PartnerContactCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerContactCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerContactCollection := make([]PartnerContactCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerContactCollection = append(partnerContactCollection, PartnerContactCollection{
			ObjectID:                          data.ObjectID,
			PartnerContactID:                  data.PartnerContactID,
			PartnerContactUUID:                data.PartnerContactUUID,
			UserID:                            data.UserID,
			IdentityUUID:                      data.IdentityUUID,
			CreateUser:                        data.CreateUser,
			BusinessPartnerID:                 data.BusinessPartnerID,
			PartnerID:                         data.PartnerID,
			StatusCode:                        data.StatusCode,
			StatusCodeText:                    data.StatusCodeText,
			TitleCode:                         data.TitleCode,
			TitleCodeText:                     data.TitleCodeText,
			AcademicTitleCode:                 data.AcademicTitleCode,
			AcademicTitleCodeText:             data.AcademicTitleCodeText,
			BusinessPartnerFormattedName:      data.BusinessPartnerFormattedName,
			FirstName:                         data.FirstName,
			LastName:                          data.LastName,
			MiddleName:                        data.MiddleName,
			AdditionalLastName:                data.AdditionalLastName,
			GenderCode:                        data.GenderCode,
			GenderCodeText:                    data.GenderCodeText,
			MaritalStatusCode:                 data.MaritalStatusCode,
			MaritalStatusCodeText:             data.MaritalStatusCodeText,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			BirthDate:                         data.BirthDate,
			VIPContactCode:                    data.VIPContactCode,
			VIPContactCodeText:                data.VIPContactCodeText,
			Department:                        data.Department,
			JobTitle:                          data.JobTitle,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			StateCode:                         data.StateCode,
			StateCodeText:                     data.StateCodeText,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			City:                              data.City,
			StreetPostalCode:                  data.StreetPostalCode,
			Phone:                             data.Phone,
			Mobile:                            data.Mobile,
			Fax:                               data.Fax,
			Email:                             data.Email,
			BestReachedByCode:                 data.BestReachedByCode,
			BestReachedByCodeText:             data.BestReachedByCodeText,
			NormalisedPhone:                   data.NormalisedPhone,
			NormalisedMobile:                  data.NormalisedMobile,
			CreatedOn:                         data.CreatedOn,
			CreatedBy:                         data.CreatedBy,
			CreatedByIdentityUUID:             data.CreatedByIdentityUUID,
			ChangedOn:                         data.ChangedOn,
			ChangedBy:                         data.ChangedBy,
			ChangedByIdentityUUID:             data.ChangedByIdentityUUID,
			EntityLastChangedOn:               data.EntityLastChangedOn,
			ETag:                              data.ETag,
		})
	}

	return partnerContactCollection, nil
}
