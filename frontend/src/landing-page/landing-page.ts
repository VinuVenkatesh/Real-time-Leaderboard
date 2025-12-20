import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { RegisterModal } from '../register-modal/register-modal';

@Component({
  selector: 'app-landing-page',
  standalone: true,
  imports: [MatButtonModule, MatDialogModule],
  templateUrl: './landing-page.html',
  styleUrl: './landing-page.css',
})
export class LandingPage {
  private dialog = inject(MatDialog);

  login() {
    console.log('Login button clicked');
  }

  register() {
    this.dialog.open(RegisterModal,{
      width: '400px'
    }).afterClosed().subscribe(result => {
      if (!result) return;
      console.log('Registered with', result);
    });
  }
}
