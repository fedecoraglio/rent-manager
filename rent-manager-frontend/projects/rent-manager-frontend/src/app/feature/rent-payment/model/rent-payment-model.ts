export interface RentPaymentFormValue {
  rental_contract_id: number;
  period: string;
  due_date: string;
  payment_date: string;
  base_amount: number;
  suggested_adjustment_percentage: number;
  applied_adjustment_percentage: number;
  suggested_interest_amount: number;
  applied_interest_amount: number;
  total_amount: number;
  paid_amount: number;
  is_paid: boolean;
  notes: string;
}

export interface CreateRentPaymentRequest extends RentPaymentFormValue {}

export interface UpdateRentPaymentRequest extends RentPaymentFormValue {}
