import { Component, inject } from '@angular/core';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { MatDialogRef, MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { UserAuth } from '../services/userauth';

@Component({
  selector: 'app-register-modal',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule
  ],
  templateUrl: './register-modal.html',
  styleUrls: ['./register-modal.css']
})
export class RegisterModal {
  private dialogRef = inject(MatDialogRef<RegisterModal>);
  private userAuthService = inject(UserAuth);

  // Reactive form
  registerForm = new FormGroup({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  });

  onRegister() {
    if (this.registerForm.invalid) {
      this.registerForm.markAllAsTouched(); // show errors
      return;
    }
    this.userAuthService.register(
      this.registerForm.value.username!,
      this.registerForm.value.password!
    ).subscribe({
      next: (response) => {
        console.log('Registration successful', response);
        this.dialogRef.close(this.registerForm.value); // close dialog with form data
      }
    // close dialog with form data
  });
  }
}
