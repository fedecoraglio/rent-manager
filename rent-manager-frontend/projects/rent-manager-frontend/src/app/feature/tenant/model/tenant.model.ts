export interface CreateTenantRequest {
  country_id: number;
  state_id: number;
  name: string;
  email: string;
  document_number: string;
  phone: string;
  city: string;
  street: string;
  street_number: string;
  floor: string;
  apartment: string;
  postal_code: string;
}

export interface UpdateTenantRequest {
  country_id: number;
  state_id: number;
  name: string;
  email: string;
  document_number: string;
  phone: string;
  city: string;
  street: string;
  street_number: string;
  floor: string;
  apartment: string;
  postal_code: string;
}

export interface TenantFormValue {
  country_id: number;
  state_id: number;
  name: string;
  email: string;
  document_number: string;
  phone: string;
  city: string;
  street: string;
  street_number: string;
  floor: string;
  apartment: string;
  postal_code: string;
}
