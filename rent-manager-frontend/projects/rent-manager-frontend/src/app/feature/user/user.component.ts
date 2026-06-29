import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'rm-users',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './user.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserComponent {}
