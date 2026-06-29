import {
  AfterViewInit,
  ChangeDetectionStrategy,
  Component,
  inject,
  Input,
  OnDestroy,
} from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { filter, ReplaySubject, Subject, takeUntil } from 'rxjs';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { TranslatePipe } from '@ngx-translate/core';

import { OwnerFormField } from './field/owner-form-field';
import { RoleService } from '@core/role/role.service';
import { Owner } from '@core/owner/owner.model';
import { OwnerFormValue } from '@feature/owner/model/owner.model';

@Component({
  selector: 'rm-owner-form',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    TranslatePipe,
  ],
  templateUrl: './owner-form.component.html',
  styleUrl: './owner-form.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OwnerFormComponent implements OnDestroy, AfterViewInit {
  private readonly destroy$ = new Subject<void>();
  private readonly roleService = inject(RoleService);

  readonly rolesSignal = this.roleService.rolesSignal;
  readonly fields = OwnerFormField;
  readonly formGroup = new FormGroup({
    [OwnerFormField.Name]: new FormControl('', Validators.required),
    [OwnerFormField.Email]: new FormControl('', [Validators.required, Validators.email]),
    [OwnerFormField.DocumentNumber]: new FormControl(''),
    [OwnerFormField.Phone]: new FormControl(''),
  });

  readonly owner$ = new ReplaySubject<Owner | null>(1);

  @Input() isReadonly = false;
  @Input() set owner(owner: Owner | null) {
    this.owner$.next(owner);
  }

  get nameField() {
    return this.formGroup.get(OwnerFormField.Name);
  }

  get emailField() {
    return this.formGroup.get(OwnerFormField.Email);
  }

  get isValid() {
    this.formGroup.markAllAsTouched();
    return this.formGroup.valid;
  }

  get value(): OwnerFormValue {
    return {
      ...(this.formGroup.getRawValue() as OwnerFormValue),
    };
  }

  ngAfterViewInit() {
    this.owner$.pipe(filter(Boolean), takeUntil(this.destroy$)).subscribe((owner: Owner) => {
      this.formGroup.patchValue({
        name: owner.name,
        email: owner.email,
        document_number: owner.document_number,
        phone: owner.phone,
      });

      if (this.isReadonly) {
        this.formGroup.disable();
      }

      this.formGroup.updateValueAndValidity();
    });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
