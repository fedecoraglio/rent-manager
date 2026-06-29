export interface PropertySummaryCardData {
  id: number;
  rentalContractId: number;
  title: string;
  totalPayments: number;
  paidPayments: number;
  remainingPayments: number;
  currentAmount: number;
  nextSuggestedAmount: number;
  nextPendingPeriod: string | null;
  nextAdjustmentPeriod: string | null;
}
