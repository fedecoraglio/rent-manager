import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'rm-property',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './property.component.html',
  styleUrl: './property.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PropertyComponent {}
