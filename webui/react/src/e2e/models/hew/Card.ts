import { CanBeParent, NamedComponent, NamedComponentArgs } from 'e2e/models/BaseComponent';
import { WorkspaceActionDropdown } from 'e2e/models/components/WorkspaceActionDropdown';

import { Dropdown } from './Dropdown';

/**
 * Returns a representation of the card component from Hew.
 * This constructor represents the contents in hew/src/kit/Card.tsx.
 * @param {object} obj
 * @param {implementsGetLocator} obj.parent - The parent used to locate this card
 * @param {string} obj.selector - Used instead of `defaultSelector`
 */
export class Card extends NamedComponent {
  override defaultSelector: string = ''; // must be provided
  static actionMenuSelector = '[aria-label="Action menu"]';

  readonly actionMenu: Dropdown = new WorkspaceActionDropdown({
    parent: this,
    selector: Card.actionMenuSelector,
  });

  static withName<T extends Card>(
    props: { name: string; parent: CanBeParent },
    cardType: new (args: NamedComponentArgs) => T,
  ): T {
    return new cardType({
      parent: props.parent,
      selector: `[data-testid="card-${props.name}"]`,
    });
  }
}
