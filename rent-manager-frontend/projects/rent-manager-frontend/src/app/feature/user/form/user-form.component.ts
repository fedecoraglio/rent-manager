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

import { User, UserFormValue } from '../model/user.model';
import { UserFormField } from './field/user-form-field';
import { RoleService } from '@core/role/role.service';

@Component({
  selector: 'rm-user-form',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    TranslatePipe,
  ],
  templateUrl: './user-form.component.html',
  styleUrl: './user-form.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserFormComponent implements OnDestroy, AfterViewInit {
  private readonly destroy$ = new Subject<void>();
  private readonly roleService = inject(RoleService);

  readonly rolesSignal = this.roleService.rolesSignal;
  readonly fields = UserFormField;
  readonly formGroup = new FormGroup({
    [UserFormField.Name]: new FormControl('', Validators.required),
    [UserFormField.Email]: new FormControl('', [Validators.required, Validators.email]),
    [UserFormField.Password]: new FormControl(''),
    [UserFormField.RoleId]: new FormControl(2, Validators.required),
  });

  readonly user$ = new ReplaySubject<User | null>(1);

  @Input() isReadonly = false;
  @Input() isEdit = false;

  @Input() set user(user: User | null) {
    this.user$.next(user);
  }

  get nameField() {
    return this.formGroup.get(UserFormField.Name);
  }

  get emailField() {
    return this.formGroup.get(UserFormField.Email);
  }

  get passwordField() {
    return this.formGroup.get(UserFormField.Password);
  }

  get roleIdField() {
    return this.formGroup.get(UserFormField.RoleId);
  }

  get isValid() {
    this.formGroup.markAllAsTouched();
    return this.formGroup.valid;
  }

  get value(): UserFormValue {
    return {
      ...(this.formGroup.getRawValue() as UserFormValue),
    };
  }

  ngAfterViewInit() {
    this.configureForm();

    this.user$.pipe(filter(Boolean), takeUntil(this.destroy$)).subscribe((user: User) => {
      this.formGroup.patchValue({
        name: user.name,
        email: user.email,
        password: '',
        role_id: user.role_id,
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

  private configureForm(): void {
    if (this.isEdit) {
      this.passwordField?.clearValidators();
    } else {
      this.passwordField?.setValidators([Validators.required]);
    }

    this.passwordField?.updateValueAndValidity();
  }
}
