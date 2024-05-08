import { NamedComponent } from 'e2e/models/BaseComponent';
import { FilterForm } from 'e2e/models/components/FilterForm';
import { Dropdown } from 'e2e/models/hew/Dropdown';

/**
 * Returns a representation of the ColumnPickerMenu component.
 * This constructor represents the contents in src/components/ColumnPickerMenu.tsx.
 * @param {object} obj
 * @param {CanBeParent} obj.parent - The parent used to locate this ColumnPickerMenu
 * @param {string} obj.selector - Used instead of `defaultSelector`
 */
export class ColumnPickerMenu extends NamedComponent {
  readonly defaultSelector = '[data-test-component="columnPickerMenu"]';
  readonly dropdown = new Dropdown({
    parent: this._parent,
    selector: 'Button' + this.defaultSelector,
  });
  readonly filterForm = new FilterForm({ parent: this.dropdown._menu });
}
