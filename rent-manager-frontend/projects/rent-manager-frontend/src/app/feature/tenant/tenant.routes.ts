import { Routes } from '@angular/router';

import { TenantWriteApiService } from './service/tenant-write-api.service';
import { TenantWriteService } from './service/tenant-write.service';
import { TenantComponent } from './tenant.component';
import { TenantCreateComponent } from './create/tenant-create.component';
import { TenantEditComponent } from './edit/tenant-edit.component';
import { TenantListComponent } from './list/tenant-list.component';
import { TenantViewComponent } from './view/tenant-view.component';

export default [
  {
    path: '',
    providers: [TenantWriteService, TenantWriteApiService],
    component: TenantComponent,
    children: [
      {
        path: '',
        component: TenantListComponent,
      },
      {
        path: 'create',
        component: TenantCreateComponent,
      },
      {
        path: 'edit/:id',
        component: TenantEditComponent,
      },
      {
        path: 'view/:id',
        component: TenantViewComponent,
      },
    ],
  },
] as Routes;
