import { ChangeDetectionStrategy, Component, input, output } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { TranslatePipe } from '@ngx-translate/core';

import { Property } from '@core/property/property.model';

@Component({
  selector: 'rm-property-list-table',
  standalone: true,
  imports: [MatTableModule, MatIconModule, MatButtonModule, TranslatePipe],
  templateUrl: './property-list-table.component.html',
  styleUrl: './property-list-table.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PropertyListTableComponent {
  readonly properties = input<Property[]>([]);

  readonly view = output<number>();
  readonly edit = output<number>();

  readonly displayedColumns = ['id', 'code', 'title', 'city', 'street', 'actions'];
}
