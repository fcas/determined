import { BasePage } from 'e2e/models/BasePage';
import { DeterminedAuth } from 'e2e/models/components/DeterminedAuth';
import { PageComponent } from 'e2e/models/components/Page';

/**
 * Returns a representation of the SignIn page.
 *
 * @remarks
 * This constructor represents the contents in src/pages/SignIn.tsx.
 *
 * @param {Page} page - The '@playwright/test' Page being used by a test
 */
export class SignIn extends BasePage {
  readonly url: string = 'login'
  readonly pageComponent: PageComponent = new PageComponent({parent: this});
  readonly detAuth: DeterminedAuth = new DeterminedAuth({parent: this.pageComponent})
  static title: string | RegExp = 'Sign In - Determined';
  // TODO add SSO page model as well
}
