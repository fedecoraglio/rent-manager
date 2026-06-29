import { Routes } from '@angular/router';

import { RentPaymentWriteApiService } from './service/rent-payment-write-api.service';
import { RentPaymentWriteService } from './service/rent-payment-write.service';
import { RentPaymentComponent } from './rent-payment.component';
import { RentPaymentCreateComponent } from './create/rent-payment-create.component';
import { RentPaymentEditComponent } from './edit/rent-payment-edit.component';

export default [
  {
    path: '',
    providers: [RentPaymentWriteService, RentPaymentWriteApiService],
    component: RentPaymentComponent,
    children: [
      {
        path: 'create',
        component: RentPaymentCreateComponent,
      },
      {
        path: 'edit/:id',
        component: RentPaymentEditComponent,
      },
    ],
  },
] as Routes;
