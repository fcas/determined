import { BaseComponent, NamedComponent } from 'e2e/models/BaseComponent';

/**
 * Returns a representation of the Modal component from Ant.
 * This constructor represents the contents in antd/es/modal/index.d.ts.
 * @param {object} obj
 * @param {CanBeParent} obj.parent - The parent used to locate this Modal
 * @param {string} obj.selector - Used instead of `defaultSelector`
 */
export class Modal extends NamedComponent {
  readonly defaultSelector = '.ant-modal-content';
  readonly header = new ModalHeader({ parent: this, selector: '.ant-modal-header' });
  readonly body = new BaseComponent({ parent: this, selector: '.ant-modal-body' });
  readonly footer = new ModalFooter({ parent: this, selector: '.ant-modal-footer' });
}

/**
 * Returns a representation of the Modal's Footer component from Ant.
 * This constructor represents the footer in antd/es/modal/index.d.ts..
 * @param {object} obj
 * @param {CanBeParent} obj.parent - The parent used to locate this Modal
 * @param {string} obj.selector - Used instead of `defaultSelector`
 */
class ModalHeader extends BaseComponent {
  readonly title = new BaseComponent({ parent: this, selector: '.ant-modal-title' });
}

/**
 * Returns a representation of the Modal's Footer component from Ant.
 * This constructor represents the footer in antd/es/modal/index.d.ts..
 * @param {object} obj
 * @param {implementsGetLocator} obj.parent - The parent used to locate this Modal
 * @param {string} obj.selector - Used instead of `defaultSelector`
 */
class ModalFooter extends BaseComponent {
  readonly submit = new BaseComponent({ parent: this, selector: '[type="submit"]' });
  readonly cancel = new BaseComponent({ parent: this, selector: '[type="cancel"]' });
}
