import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'rm-tenant',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './tenant.component.html',
  styleUrl: './tenant.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TenantComponent {}
