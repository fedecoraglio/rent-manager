export interface RentPayment {
  id: number;
  rental_contract_id: number;
  period: string;
  due_date: string;
  payment_date: string | null;
  base_amount: number;
  suggested_adjustment_percentage: number;
  applied_adjustment_percentage: number;
  suggested_interest_amount: number;
  applied_interest_amount: number;
  total_amount: number;
  paid_amount: number;
  is_paid: boolean;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface RentPaymentScheduleItem {
  rental_contract_id: number;
  rent_payment_id: number;
  period: string;
  due_date: string;
  base_amount: number;
  suggested_adjustment_percentage: number;
  suggested_adjustment_status: string;
  suggested_adjustment_message: string;
  suggested_interest_amount: number;
  suggested_total_amount: number;
  applied_adjustment_percentage: number;
  applied_interest_amount: number;
  total_amount: number;
  paid_amount: number;
  payment_date: string | null;
  is_paid: boolean;
}

export interface RentPaymentSuggestion {
  rental_contract_id: number;
  period: string;
  due_date: string;
  payment_date: string;
  base_amount: number;
  suggested_adjustment_percentage: number;
  suggested_adjustment_amount: number;
  suggested_adjustment_status: string;
  suggested_adjustment_message: string;
  suggested_interest_amount: number;
  suggested_total_amount: number;
}

export interface RentalContractSummary {
  rental_contract_id: number;
  total_payments: number;
  paid_payments: number;
  remaining_payments: number;
  current_amount: number;
  next_suggested_amount: number;
  next_pending_period: string | null;
  next_adjustment_period: string | null;
}
