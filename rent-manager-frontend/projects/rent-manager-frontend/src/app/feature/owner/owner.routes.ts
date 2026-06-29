import { Routes } from '@angular/router';

import { OwnerReadApiService } from '@core/owner/owner-read-api.service';
import { OwnerReadService } from '@core/owner/owner-read.service';
import { OwnerComponent } from './owner.component';
import { OwnerCreateComponent } from './create/owner-create.component';
import { OwnerEditComponent } from './edit/owner-edit.component';
import { OwnerListComponent } from './list/owner-list.component';
import { OwnerViewComponent } from './view/owner-view.component';

export default [
  {
    path: '',
    providers: [OwnerReadService, OwnerReadApiService],
    component: OwnerComponent,
    children: [
      {
        path: '',
        component: OwnerListComponent,
      },
      {
        path: 'create',
        component: OwnerCreateComponent,
      },
      {
        path: 'edit/:id',
        component: OwnerEditComponent,
      },
      {
        path: 'view/:id',
        component: OwnerViewComponent,
      },
    ],
  },
] as Routes;
