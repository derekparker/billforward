package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
PricingComponent pricing component

swagger:model PricingComponent
*/
type PricingComponent struct {

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "The charge model of the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ChargeModel string `json:"chargeModel,omitempty"`

	/* { "description" : "The charge type of the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ChargeType string `json:"chargeType,omitempty"`

	/* ComponentValue component value
	 */
	ComponentValue int32 `json:"componentValue,omitempty"`

	/* { "description" : "ID of the pricing-component. This ID does not change with new versions of the pricing component.", "verbs":["POST","PUT","GET"] } When associating a pricing component with a product rate plan, this ID should be used.

	Required: true
	*/
	ConsistentID string `json:"consistentID,omitempty" xml:"id"`

	/* Cost cost
	 */
	Cost float64 `json:"cost,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "The default quantity of the pricing-component. If no value is supplied on a subscription this value will be used. This is useful for setting an expected purchase level of for introducing new pricing components to existing subscriptions and not having to back-fill values", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	DefaultQuantity int32 `json:"defaultQuantity,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	Description string `json:"description,omitempty"`

	/* {"default":"<span class=\"label label-default\">delayed</span>","description":"Default behaviour when a value is downgraded using this pricing component, this behaviour can be overridden when changing the value.<br><span class=\"label label-default\">immediate</span> &mdash; Downgrade will apply at the time the request is made.<br><span class=\"label label-default\">delayed</span> &mdash; Downgrade will apply at the end of the current billing cycle.","verbs":["POST","GET"]}
	 */
	DowngradeMode string `json:"downgradeMode,omitempty"`

	/* { "description" : "", "default" : "", "verbs":["POST","GET"] }

	Required: true
	*/
	DummyField string `json:"dummyField,omitempty" xml:"@type"`

	/* { "description" : "Version ID of the pricing-component. Unique ID for each version of a pricing-component.", "verbs":["POST","PUT","GET"] }
	 */
	ID string `json:"id,omitempty" xml:"versionID"`

	/* { "default" : "Aggregated",  "description" : "For <span class=\"label label-default\">setup</span> pricing components <span class=\"label label-default\">Immediate</span> invoicing will result in an invoice being issued on subscription being set to the AwaitingPayment state, irrespective of the subscription start date. <span class=\"label label-default\">Aggregated</span> invoicing will add a charge to the first invoice of the subscription.", "verbs":["POST","PUT","GET"] }
	 */
	InvoicingType string `json:"invoicingType,omitempty"`

	/* { "description" : "The maximum quantity of the pricing-component.", "verbs":[] }
	 */
	MaxQuantity int32 `json:"maxQuantity,omitempty"`

	/* { "default" : "0", "description" : "The minimum quantity of the pricing-component.", "verbs":[] }
	 */
	MinQuantity int32 `json:"minQuantity,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* { "description" : "", "verbs":[] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* PriceExplanation price explanation
	 */
	PriceExplanation []string `json:"priceExplanation,omitempty"`

	/* PriceExplanationString price explanation string
	 */
	PriceExplanationString string `json:"priceExplanationString,omitempty"`

	/* { "description" : "The product-rate-plan associated with the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ProductRatePlan *ProductRatePlan `json:"productRatePlan,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ProductRatePlanID string `json:"productRatePlanID,omitempty"`

	/* {  "default" : "[]", "description" : "The pricing-component-tiers associated with the pricing-component.", "verbs":["POST","PUT","GET"] }
	 */
	Tiers []*PricingComponentTier `json:"tiers,omitempty"`

	/* UnitOfMeasure unit of measure
	 */
	UnitOfMeasure *UnitOfMeasure `json:"unitOfMeasure,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	UnitOfMeasureID string `json:"unitOfMeasureID,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* {"default":"<span class=\"label label-default\">immediate</span>","description":"Default behaviour when a value is upgraded using this pricing component, this behaviour can be overridden when changing the value.<br><span class=\"label label-default\">immediate</span> &mdash; Upgrade will apply at the time the request is made.<br><span class=\"label label-default\">delayed</span> &mdash; Upgrade will apply at the end of the current billing cycle.","verbs":["POST","GET"]}
	 */
	UpgradeMode string `json:"upgradeMode,omitempty"`

	/* { "default" : "current time", "description" : "The DateTime specifying when the pricing-component is valid from.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ValidFrom strfmt.DateTime `json:"validFrom,omitempty"`

	/* {  "default" : "null", "description" : "The UTC DateTime specifying when the pricing-component is valid till.", "verbs":["POST","PUT","GET"] }
	 */
	ValidTill strfmt.DateTime `json:"validTill,omitempty"`
}

// Validate validates this pricing component
func (m *PricingComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDowngradeMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDummyField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoicingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductRatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductRatePlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pricingComponentChargeModelEnum []interface{}

func (m *PricingComponent) validateChargeModelEnum(path, location string, value string) error {
	if pricingComponentChargeModelEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["flat","tiered","tiered_volume"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			pricingComponentChargeModelEnum = append(pricingComponentChargeModelEnum, v)
		}
	}
	return validate.Enum(path, location, value, pricingComponentChargeModelEnum)
}

func (m *PricingComponent) validateChargeModel(formats strfmt.Registry) error {

	if err := validate.Required("chargeModel", "body", string(m.ChargeModel)); err != nil {
		return err
	}

	if err := m.validateChargeModelEnum("chargeModel", "body", m.ChargeModel); err != nil {
		return err
	}

	return nil
}

var pricingComponentChargeTypeEnum []interface{}

func (m *PricingComponent) validateChargeTypeEnum(path, location string, value string) error {
	if pricingComponentChargeTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["setup","subscription","arrears","usage"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			pricingComponentChargeTypeEnum = append(pricingComponentChargeTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, pricingComponentChargeTypeEnum)
}

func (m *PricingComponent) validateChargeType(formats strfmt.Registry) error {

	if err := validate.Required("chargeType", "body", string(m.ChargeType)); err != nil {
		return err
	}

	if err := m.validateChargeTypeEnum("chargeType", "body", m.ChargeType); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponent) validateConsistentID(formats strfmt.Registry) error {

	if err := validate.Required("consistentID", "body", string(m.ConsistentID)); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponent) validateDefaultQuantity(formats strfmt.Registry) error {

	if err := validate.Required("defaultQuantity", "body", int32(m.DefaultQuantity)); err != nil {
		return err
	}

	return nil
}

var pricingComponentDowngradeModeEnum []interface{}

func (m *PricingComponent) validateDowngradeModeEnum(path, location string, value string) error {
	if pricingComponentDowngradeModeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			pricingComponentDowngradeModeEnum = append(pricingComponentDowngradeModeEnum, v)
		}
	}
	return validate.Enum(path, location, value, pricingComponentDowngradeModeEnum)
}

func (m *PricingComponent) validateDowngradeMode(formats strfmt.Registry) error {

	if err := m.validateDowngradeModeEnum("downgradeMode", "body", m.DowngradeMode); err != nil {
		return err
	}

	return nil
}

var pricingComponentDummyFieldEnum []interface{}

func (m *PricingComponent) validateDummyFieldEnum(path, location string, value string) error {
	if pricingComponentDummyFieldEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["TieredPricingComponent","FlatPricingComponent","TieredVolumePricingComponent"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			pricingComponentDummyFieldEnum = append(pricingComponentDummyFieldEnum, v)
		}
	}
	return validate.Enum(path, location, value, pricingComponentDummyFieldEnum)
}

func (m *PricingComponent) validateDummyField(formats strfmt.Registry) error {

	if err := validate.Required("dummyField", "body", string(m.DummyField)); err != nil {
		return err
	}

	if err := m.validateDummyFieldEnum("dummyField", "body", m.DummyField); err != nil {
		return err
	}

	return nil
}

var pricingComponentInvoicingTypeEnum []interface{}

func (m *PricingComponent) validateInvoicingTypeEnum(path, location string, value string) error {
	if pricingComponentInvoicingTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Immediate","Aggregated"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			pricingComponentInvoicingTypeEnum = append(pricingComponentInvoicingTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, pricingComponentInvoicingTypeEnum)
}

func (m *PricingComponent) validateInvoicingType(formats strfmt.Registry) error {

	if err := m.validateInvoicingTypeEnum("invoicingType", "body", m.InvoicingType); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponent) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponent) validateProductRatePlan(formats strfmt.Registry) error {

	if m.ProductRatePlan != nil {

		if err := m.ProductRatePlan.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PricingComponent) validateProductRatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("productRatePlanID", "body", string(m.ProductRatePlanID)); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponent) validateUnitOfMeasureID(formats strfmt.Registry) error {

	if err := validate.Required("unitOfMeasureID", "body", string(m.UnitOfMeasureID)); err != nil {
		return err
	}

	return nil
}

var pricingComponentUpgradeModeEnum []interface{}

func (m *PricingComponent) validateUpgradeModeEnum(path, location string, value string) error {
	if pricingComponentUpgradeModeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["immediate","delayed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			pricingComponentUpgradeModeEnum = append(pricingComponentUpgradeModeEnum, v)
		}
	}
	return validate.Enum(path, location, value, pricingComponentUpgradeModeEnum)
}

func (m *PricingComponent) validateUpgradeMode(formats strfmt.Registry) error {

	if err := m.validateUpgradeModeEnum("upgradeMode", "body", m.UpgradeMode); err != nil {
		return err
	}

	return nil
}

func (m *PricingComponent) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", strfmt.DateTime(m.ValidFrom)); err != nil {
		return err
	}

	return nil
}
