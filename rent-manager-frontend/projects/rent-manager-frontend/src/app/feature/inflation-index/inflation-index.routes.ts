import { Routes } from '@angular/router';

export const inflationIndexRoutes: Routes = [
  {
    path: '',
    loadComponent: () =>
      import('./list/inflation-index-list.component').then((m) => m.InflationIndexListComponent),
  },
  {
    path: 'create',
    loadComponent: () =>
      import('./create/inflation-index-create.component').then(
        (m) => m.InflationIndexCreateComponent,
      ),
  },
  {
    path: 'edit/:id',
    loadComponent: () =>
      import('./edit/inflation-index-edit.component').then((m) => m.InflationIndexEditComponent),
  },
];
