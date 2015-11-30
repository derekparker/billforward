package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*Billing entities are models in the BillForward system, of objects in the real-world or otherwise.

swagger:model BillingEntityBase
*/
type BillingEntityBase struct {

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* {"description":"The discount percentage applied to this pricing component.","verbs":["POST","GET"]}
	 */
	PercentageDiscount float64 `json:"percentageDiscount,omitempty"`

	/* {"description":"Name or ID of the pricing component purchased.","verbs":["POST","GET"]}

	Required: true
	*/
	PricingComponent string `json:"pricingComponent,omitempty"`

	/* {"description":"Quantity consumed of the pricing component.","verbs":["POST","GET"]}

	Required: true
	*/
	Quantity int32 `json:"quantity,omitempty"`
}

// Validate validates this billing entity base
func (m *BillingEntityBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingEntity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var billingEntityBaseBillingEntityEnum []interface{}

func (m *BillingEntityBase) validateBillingEntityEnum(path, location string, value string) error {
	if billingEntityBaseBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			billingEntityBaseBillingEntityEnum = append(billingEntityBaseBillingEntityEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, billingEntityBaseBillingEntityEnum); err != nil {
		return err
	}
	return nil
}

func (m *BillingEntityBase) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *BillingEntityBase) validatePricingComponent(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponent", "body", string(m.PricingComponent)); err != nil {
		return err
	}

	return nil
}

func (m *BillingEntityBase) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", int32(m.Quantity)); err != nil {
		return err
	}

	return nil
}
