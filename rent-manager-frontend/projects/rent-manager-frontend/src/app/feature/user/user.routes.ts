import { Routes } from '@angular/router';

import { UserApiService } from './service/user-api.service';
import { UserService } from './service/user.service';
import { UserComponent } from './user.component';
import { UserCreateComponent } from './create/user-create.component';
import { UserEditComponent } from './edit/user-edit.component';
import { UserListComponent } from './list/user-list.component';
import { UserViewComponent } from './view/user-view.component';

export default [
  {
    path: '',
    providers: [UserService, UserApiService],
    component: UserComponent,
    children: [
      {
        path: '',
        component: UserListComponent,
      },
      {
        path: 'create',
        component: UserCreateComponent,
      },
      {
        path: 'edit/:id',
        component: UserEditComponent,
      },
      {
        path: 'view/:id',
        component: UserViewComponent,
      },
    ],
  },
] as Routes;
