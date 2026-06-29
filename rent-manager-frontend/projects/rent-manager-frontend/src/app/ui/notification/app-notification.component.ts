import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';

import { NotificationService } from '@core/notification/notification.service';

@Component({
  selector: 'rm-app-notification',
  standalone: true,
  imports: [MatIconModule],
  templateUrl: './app-notification.component.html',
  styleUrl: './app-notification.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppNotificationComponent {
  readonly notificationService = inject(NotificationService);
}
