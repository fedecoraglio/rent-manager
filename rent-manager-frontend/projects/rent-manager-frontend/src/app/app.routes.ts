import { Routes } from '@angular/router';

import { authGuard } from './core/auth/auth.guard';
import { AuthLayoutComponent } from './layout/auth/auth-layout.component';
import { MainLayoutComponent } from './layout/main/main-layout.component';

export const routes: Routes = [
  {
    path: '',
    component: AuthLayoutComponent,
    children: [
      {
        path: 'login',
        loadChildren: () => import('./feature/login/login.routes'),
      },
      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'login',
      },
    ],
  },
  {
    path: 'app',
    component: MainLayoutComponent,
    canActivate: [authGuard],
    children: [
      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'dashboard',
      },
      {
        path: 'dashboard',
        loadComponent: () =>
          import('./feature/dashboard/dashboard.component').then((m) => m.DashboardComponent),
      },
      {
        path: 'users',
        loadChildren: () => import('@feature/user/user.routes'),
      },
      {
        path: 'owners',
        loadChildren: () => import('@feature/owner/owner.routes'),
      },
      {
        path: 'tenants',
        loadChildren: () => import('@feature/tenant/tenant.routes'),
      },
      {
        path: 'properties',
        loadChildren: () => import('./feature/property/property.routes'),
      },
      {
        path: 'rent-contracts',
        loadChildren: () => import('./feature/rent-contract/rent-contract.routes'),
      },
      {
        path: 'rent-payments',
        loadChildren: () => import('./feature/rent-payment/rent-payment.routes'),
      },
      {
        path: 'rent-contracts/overview/:id',
        loadComponent: () =>
          import('./feature/rent-contract/overview/rent-contract-overview.component').then(
            (m) => m.RentContractOverviewComponent,
          ),
      },
      {
        path: 'inflation-indexes',
        loadChildren: () =>
          import('@feature/inflation-index/inflation-index.routes').then(
            (m) => m.inflationIndexRoutes,
          ),
      },
    ],
  },
  {
    path: '**',
    redirectTo: 'login',
  },
];
