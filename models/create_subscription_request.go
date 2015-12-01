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
Entity for requesting that a subscription be created.

swagger:model CreateSubscriptionRequest
*/
type CreateSubscriptionRequest struct {

	/* {"description":"ID of the BillForward Account who will own this subscription. You should ensure beforehand that the customer has had a BillForward Account created for them.","verbs":["POST"]}

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* {"default":false,"description":"Whether this subscription should become an 'aggregating subscription', collecting charges (starting now) from all other subscriptions (current and future) belonging to this BillForward Account.","verbs":["POST"]}
	 */
	AggregateAllSubscriptionsOnAccount bool `json:"aggregateAllSubscriptionsOnAccount,omitempty"`

	/* {"default":true,"description":"Whether to override the `end` date to line up with the current period end of the 'aggregating subscription' to which this subscription belongs.","verbs":["POST"]}
	 */
	AlignPeriodWithAggregatingSubscription bool `json:"alignPeriodWithAggregatingSubscription,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* {"default":"(null)","description":"Description of the created subscription. This is primarily for your benefit &mdash; for example, you could write here the mechanism through which you obtained this customer. (e.g. 'Customer obtained through Lazy Wednesdays promotion').","verbs":["POST"]}
	 */
	Description string `json:"description,omitempty"`

	/* {"default":"(1 period ahead of the `start` time)","description":"ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should finish its first service period.","verbs":["POST"]}
	 */
	End *strfmt.DateTime `json:"end,omitempty"`

	/* {"default":"(Subscription will be named after the rate plan to which the subscription subscribes)","description":"Name of the created subscription. This is primarily for your benefit &mdash; for example, to enable you to identify subscriptions at a glance in the BillForward web interface (e.g. 'Customer 1425, guy@mail.com, Premium membership').","verbs":["POST"]}
	 */
	Name string `json:"name,omitempty"`

	/* {"default":"(Auto-populated using your authentication credentials)","description":"ID of the BillForward Organization within which the requested Subscription should be created. If omitted, this will be auto-populated using your authentication credentials.","verbs":["POST"]}
	 */
	OrganizationID string `json:"organizationID,omitempty"`

	/* {"default":"(If a subscription exists which 'aggregates all subscriptions belonging to this BillForward Account', refer to the ID of that subscription. Otherwise: null)","description":"ID of a parent subscription which will collect the charges raised by this subscription. The parent becomes responsible for paying those charges. If a subscription exists which 'aggregates all subscriptions belonging to this BillForward Account', then that parent will override any parent specified here.","verbs":["POST"]}
	 */
	ParentID string `json:"parentID,omitempty"`

	/* {"default":"(empty list)","description":"Quantities that this subscription possesses (upon beginning service), of pricing components upon the subscription's rate plan. For example: you can set the subscription to begin its service with '5 widgets' consumed. Otherwise the 'default quantity' will be observed instead, for each pricing component upon the rate plan.","verbs":["POST"]}
	 */
	PricingComponentQuantities []*BillingEntityBase `json:"pricingComponentQuantities,omitempty"`

	/* {"description":"ID of the product to which the subscription will be subscribing. If omitted: the product's ID will be inferred from the rate plan -- if and only if the rate plan is specified by ID.","verbs":["POST"]}
	 */
	ProductID string `json:"productID,omitempty"`

	/* {"description":"ID or name of the rate plan to which the subscription will be subscribing. Lookup by name is only possible if a `productID` is specified.","verbs":["POST"]}

	Required: true
	*/
	ProductRatePlan string `json:"productRatePlan,omitempty"`

	/* ProductRatePlanID product rate plan ID
	 */
	ProductRatePlanID string `json:"productRatePlanID,omitempty"`

	/* {"default":"(ServerNow upon receiving request)","description":"ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should enter its first service period.","verbs":["POST"]}
	 */
	Start *strfmt.DateTime `json:"start,omitempty"`

	/* {"default":"Provisioned","description":"The state in which the created subscription will begin.<br><span class=\"label label-default\">Provisioned</span> &mdash; The subscription will wait (without raising any invoices or beginning its service) until explicit action is taken to change its state.<br><span class=\"label label-default\">AwaitingPayment</span> &mdash; The subscription is activated. After `start` time is surpassed, it will begin service and raise its first invoice.","verbs":["POST"]}
	 */
	State string `json:"state,omitempty"`

	/* Type type
	 */
	Type string `json:"type,omitempty"`
}

// Validate validates this create subscription request
func (m *CreateSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductRatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSubscriptionRequest) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

var createSubscriptionRequestBillingEntityEnum []interface{}

func (m *CreateSubscriptionRequest) validateBillingEntityEnum(path, location string, value string) error {
	if createSubscriptionRequestBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createSubscriptionRequestBillingEntityEnum = append(createSubscriptionRequestBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, createSubscriptionRequestBillingEntityEnum)
}

func (m *CreateSubscriptionRequest) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubscriptionRequest) validateProductRatePlan(formats strfmt.Registry) error {

	if err := validate.Required("productRatePlan", "body", string(m.ProductRatePlan)); err != nil {
		return err
	}

	return nil
}

var createSubscriptionRequestStateEnum []interface{}

func (m *CreateSubscriptionRequest) validateStateEnum(path, location string, value string) error {
	if createSubscriptionRequestStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Trial","Provisioned","Paid","AwaitingPayment","Cancelled","Failed","Expired"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createSubscriptionRequestStateEnum = append(createSubscriptionRequestStateEnum, v)
		}
	}
	return validate.Enum(path, location, value, createSubscriptionRequestStateEnum)
}

func (m *CreateSubscriptionRequest) validateState(formats strfmt.Registry) error {

	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var createSubscriptionRequestTypeEnum []interface{}

func (m *CreateSubscriptionRequest) validateTypeEnum(path, location string, value string) error {
	if createSubscriptionRequestTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Subscription","FixedTerm","Trial"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createSubscriptionRequestTypeEnum = append(createSubscriptionRequestTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, createSubscriptionRequestTypeEnum)
}

func (m *CreateSubscriptionRequest) validateType(formats strfmt.Registry) error {

	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
