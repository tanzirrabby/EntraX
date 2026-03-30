import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { MsalService } from '@azure/msal-angular';

export const authGuard: CanActivateFn = () => {
  const msal = inject(MsalService);
  const account = msal.instance.getActiveAccount();

  if (!account) {
    msal.loginRedirect();
    return false;
  }

  return true;
};
