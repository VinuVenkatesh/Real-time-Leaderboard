import { Component, inject } from '@angular/core';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { MatDialogRef, MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { UserAuth } from '../services/userauth';

@Component({
  selector: 'app-login-modal',
  imports: [    
    ReactiveFormsModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule
  ],
  templateUrl: './login-modal.html',
  styleUrl: './login-modal.css',
})
export class LoginModal {
  private dialogRef = inject(MatDialogRef<LoginModal>);
  private userAuth = inject(UserAuth);

  loginForm = new FormGroup({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
  });

  onLogin() {
    if (this.loginForm.invalid) {
      this.loginForm.markAllAsTouched(); // show errors
      return;
    }
    this.userAuth.login(
      this.loginForm.value.username!,
      this.loginForm.value.password!
    ).subscribe({
      next: (response) => {
        console.log('Login successful', response);
        this.dialogRef.close(this.loginForm.value); // close dialog with form data
      }
    });
  }
}
