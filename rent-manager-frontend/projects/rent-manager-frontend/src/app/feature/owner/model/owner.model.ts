export interface CreateOwnerRequest {
  name: string;
  email: string;
  document_number: string;
  phone: string;
}

export interface UpdateOwnerRequest {
  name: string;
  email: string;
  document_number: string;
  phone: string;
}

export interface OwnerFormValue {
  name: string;
  email: string;
  document_number: string;
  phone: string;
}
