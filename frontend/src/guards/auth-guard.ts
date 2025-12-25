import { J } from '@angular/cdk/keycodes';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, GuardResult, MaybeAsync, Router, RouterStateSnapshot } from '@angular/router';

@Injectable({
  providedIn: 'root',
})
export class AuthGuard implements CanActivate {
  
  constructor(private router: Router) {}
  canActivate(): boolean {
    const token = localStorage.getItem('jwtToken');
    if (!token){
      this.router.navigate(['']);
      return false;
    }
    const tokenExpiry = JSON.parse(atob(token.split('.')[1])).exp * 1000;
    if (Date.now() >= tokenExpiry){
      this.router.navigate(['']);
      localStorage.removeItem('jwtToken');
      return false;
    }
    return true
  }
}
