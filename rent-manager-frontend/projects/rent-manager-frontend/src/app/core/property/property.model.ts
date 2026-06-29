import { RentalContractSummary } from '@core/rent-payment/rent-payment.model';

export interface Property {
  id: number;
  owner_id: number;
  type_id: number;
  status_id: number;
  country_id: number;
  state_id: number;
  code: string;
  title: string;
  description: string;
  street: string;
  street_number: string;
  floor: string;
  apartment: string;
  city: string;
  postal_code: string;
  created_at: string;
  updated_at: string;
}

export interface PropertySummary {
  id: number;
  title: string;
  summary: RentalContractSummary;
}
