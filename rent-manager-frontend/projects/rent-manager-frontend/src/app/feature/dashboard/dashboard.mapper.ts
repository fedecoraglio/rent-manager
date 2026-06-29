import { PropertySummary } from '@core/property/property.model';
import { PropertySummaryCardData } from '@ui/property-summary-card/property-summary-card.model';

export class DashboardMapper {
  static toCardData(property: PropertySummary): PropertySummaryCardData {
    if (!property || !property.summary) {
      console.error('No property summary found');
    }
    return {
      id: property.id,
      rentalContractId: property.summary.rental_contract_id,
      title: property.title,
      totalPayments: property.summary.total_payments,
      paidPayments: property.summary.paid_payments,
      remainingPayments: property.summary.remaining_payments,
      currentAmount: property.summary.current_amount,
      nextSuggestedAmount: property.summary.next_suggested_amount,
      nextPendingPeriod: property.summary.next_pending_period,
      nextAdjustmentPeriod: property.summary.next_adjustment_period,
    };
  }
}
