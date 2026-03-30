import { Injectable } from '@angular/core';
import { MsalService } from '@azure/msal-angular';

@Injectable({ providedIn: 'root' })
export class AuthService {
  constructor(private msal: MsalService) {}

  login(): void {
    this.msal.loginRedirect();
  }

  logout(): void {
    this.msal.logoutRedirect();
  }

  getToken(): Promise<string> {
    return this.msal
      .acquireTokenSilent({
        scopes: ['api://entrax-api/access_as_user']
      })
      .then((result) => result.accessToken);
  }
}
