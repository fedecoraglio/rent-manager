import { Routes } from '@angular/router';

import { RentContractWriteApiService } from './service/rent-contract-write-api.service';
import { RentContractWriteService } from './service/rent-contract-write.service';
import { RentContractComponent } from './rent-contract.component';
import { RentContractCreateComponent } from './create/rent-contract-create.component';
import { RentContractEditComponent } from './edit/rent-contract-edit.component';
import { RentContractListComponent } from './list/rent-contract-list.component';
import { RentContractViewComponent } from '@feature/rent-contract/view/rent-contract-view.component';

export default [
  {
    path: '',
    providers: [RentContractWriteService, RentContractWriteApiService],
    component: RentContractComponent,
    children: [
      {
        path: '',
        component: RentContractListComponent,
      },
      {
        path: 'create',
        component: RentContractCreateComponent,
      },
      {
        path: 'edit/:id',
        component: RentContractEditComponent,
      },
      {
        path: 'view/:id',
        component: RentContractViewComponent,
      },
    ],
  },
] as Routes;
