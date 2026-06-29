import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'rm-owners',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './owner.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OwnerComponent {}
