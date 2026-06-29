export interface RentContractFormValue {
  property_id: number;
  tenant_id: number;
  status_id: number;
  interest_calculation_type_id: number;
  adjustment_type_id: number;
  start_date: string;
  end_date: string;
  monthly_amount: number;
  deposit_amount: number;
  currency: string;
  due_day: number;
  daily_interest_percentage: number;
  adjustment_frequency_months: number;
  notes: string;
}

export interface CreateRentContractRequest extends RentContractFormValue {}

export interface UpdateRentContractRequest extends RentContractFormValue {}
