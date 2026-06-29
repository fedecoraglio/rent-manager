export interface Country {
  id: number;
  code: string;
  name: string;
}

export interface State {
  id: number;
  country_id: number;
  code: string;
  name: string;
}
