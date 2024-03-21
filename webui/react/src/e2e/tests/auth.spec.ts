import { expect } from '@playwright/test';
import { test } from 'e2e/fixtures/global-fixtures';
import { SignIn } from 'e2e/models/pages/SignIn';

test.describe('Authentication', () => {
  test.beforeEach(async ({ dev }) => {
    await dev.setServerAddress();
  });
  test.afterEach(async ({ page, auth }) => {
    if (await page.title() != SignIn.title) {
      await auth.logout()
    }
  });

  test('Login and Logout', async ({ page, auth }) => {
    await test.step('Login', async () => {
      await auth.login();
      await expect(page).toHaveTitle('Home - Determined');
      await expect(page).toHaveURL(/dashboard/);
    });

    await test.step('Logout', async () => {
      await auth.logout();
      await expect(page).toHaveTitle(SignIn.title);
      await expect(page).toHaveURL(/login/);
    });
  });
  
  test('Redirect to the target URL after login', async ({ page, auth }) => {
    await test.step('Visit a page and expect redirect back to login', async () => {
      await page.goto('./models');
      await expect(page).toHaveURL(/login/);
    });
    
    await test.step('Login and expect redirect to previous page', async () => {
      await auth.login(/models/);
      await expect(page).toHaveTitle('Model Registry - Determined');
    });
  });
  
  test('Bad Credentials should throw an error', async ({ page, auth }) => {
    await auth.login(/login/, {username: 'jcom', password: 'superstar'});
    await expect(page).toHaveTitle(SignIn.title);
    await expect(page).toHaveURL(/login/);
    const signInPage = new SignIn(page)
    await expect(signInPage.detAuth.error.pwLocator).toBeVisible();
    // text assertion not working yet, but we can see the error locator
    await expect(signInPage.detAuth.error.message.pwLocator).toContainText('Login Failed');
    await expect(signInPage.detAuth.error.description.pwLocator).toHaveText('invalid credentials');
  });
});
