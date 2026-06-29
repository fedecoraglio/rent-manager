import { Routes } from '@angular/router';

import { PropertyWriteApiService } from './service/property-write-api.service';
import { PropertyWriteService } from './service/property-write.service';
import { PropertyComponent } from './property.component';
import { PropertyCreateComponent } from './create/property-create.component';
import { PropertyEditComponent } from './edit/property-edit.component';
import { PropertyListComponent } from './list/property-list.component';
import { PropertyViewComponent } from './view/property-view.component';

export default [
  {
    path: '',
    providers: [PropertyWriteService, PropertyWriteApiService],
    component: PropertyComponent,
    children: [
      {
        path: '',
        component: PropertyListComponent,
      },
      {
        path: 'create',
        component: PropertyCreateComponent,
      },
      {
        path: 'edit/:id',
        component: PropertyEditComponent,
      },
      {
        path: 'view/:id',
        component: PropertyViewComponent,
      },
    ],
  },
] as Routes;
