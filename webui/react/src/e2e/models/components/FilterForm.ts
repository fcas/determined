import { NamedComponent } from 'e2e/models/BaseComponent';

/**
 * Returns a representation of the FilterForm component.
 * This constructor represents the contents in src/components/FilterForm.tsx.
 * @param {object} obj
 * @param {CanBeParent} obj.parent - The parent used to locate this FilterForm
 * @param {string} [obj.selector] - Used instead of `defaultSelector`
 */
export class FilterForm extends NamedComponent {
  readonly defaultSelector = '[data-test-component="filterForm"]';
  // todo webui/react/src/components/FilterForm/components/FilterForm.tsx
}
