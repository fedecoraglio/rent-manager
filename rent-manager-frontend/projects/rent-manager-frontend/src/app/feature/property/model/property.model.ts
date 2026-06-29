export interface CreatePropertyRequest {
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
}

export interface UpdatePropertyRequest {
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
}

export interface PropertyFormValue {
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
}
