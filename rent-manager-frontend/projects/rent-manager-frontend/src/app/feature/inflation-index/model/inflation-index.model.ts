export interface InflationIndexFormValue {
  period: string;
  percentage: number;
  source: string;
  notes: string;
}

export interface InflationIndexCreateRequest {
  period: string;
  percentage: number;
  source: string;
  notes: string;
}

export interface InflationIndexUpdateRequest {
  period: string;
  percentage: number;
  source: string;
  notes: string;
}
