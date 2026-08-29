import {Property} from "@core/property/property.model";
import {Tenant} from "@core/tenant/tenant-core.model";

export interface RentContract {
  id: number;
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
  property?: Property | null;
  tenant?: Tenant | null;
  created_at: string;
  updated_at: string;
}
