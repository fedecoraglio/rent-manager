import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { User } from '@feature/user/model/user.model';

@Component({
  selector: 'rm-user-list-table',
  standalone: true,
  imports: [MatTableModule, MatIconModule, MatButtonModule, TranslatePipe],
  templateUrl: './user-list-table.component.html',
  styleUrl: './user-list-table.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserListTableComponent {
  readonly users = input<User[]>([]);

  readonly view = output<number>();
  readonly edit = output<number>();
  readonly delete = output<number>();

  readonly displayedColumns = ['id', 'name', 'email', 'role_name', 'actions'];
}
