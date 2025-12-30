import { Component, inject } from '@angular/core';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { MatDialogRef, MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { UserAuth } from '../services/userauth';
import { User } from '../models/user';

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
  private userModel: User = new User();
  loginForm = new FormGroup({
    username: new FormControl(this.userModel.username, Validators.required),
    password: new FormControl(this.userModel.password, Validators.required),
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
        this.dialogRef.close(true); // close dialog with form data
      }
    });
  }
}
