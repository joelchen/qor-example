package admin

import (
	"github.com/qor/exchange"
	"github.com/qor/i18n"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
)

var ProductExchange *exchange.Resource
var TranslationExchange *exchange.Resource

func init() {
	ProductExchange = exchange.NewResource(&models.Product{}, exchange.Config{PrimaryField: "Code"})
	ProductExchange.Meta(&exchange.Meta{Name: "Code"})
	ProductExchange.Meta(&exchange.Meta{Name: "Name"})
	ProductExchange.Meta(&exchange.Meta{Name: "Price"})

	ProductExchange.AddValidator(func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
		if utils.ToInt(metaValues.Get("Price").Value) < 100 {
			return validations.NewError(record, "Price", "price can't less than 100")
		}
		return nil
	})

	TranslationExchange = exchange.NewResource(&i18n.Translation{}, exchange.Config{PrimaryField: "Key"})
	TranslationExchange.Meta(&exchange.Meta{Name: "Locale"})
	TranslationExchange.Meta(&exchange.Meta{Name: "Key"})
	TranslationExchange.Meta(&exchange.Meta{Name: "Value"})
}
